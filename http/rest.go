/*

   Copyright (C) 2017 Contributors as noted in the AUTHORS file

   This file is part of vetbbedit, veterinabb.sk page editor.

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.

*/

package http

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jkusniar/vetbbedit"
	"github.com/pkg/errors"
)

func logError(err error) {
	log.Printf("ERROR: request failed: %+v\n", err)
}

func renderError(w http.ResponseWriter, err error) {
	logError(err)
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func renderBadJSONError(w http.ResponseWriter, err error) {
	e := errors.Wrap(err, "json decode error")
	logError(e)

	http.Error(w, e.Error(), http.StatusBadRequest)
}

func renderNotAllowed(w http.ResponseWriter, m string) {
	err := errors.Errorf("bad method %s", m)
	logError(err)
	http.Error(w, err.Error(), http.StatusMethodNotAllowed)
}

func renderJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(true)
	if err := enc.Encode(v); err != nil {
		renderError(w, errors.Wrap(err, "error rendering JSON message"))
	}
}

func decodeJSON(r io.Reader, v interface{}) error {
	defer io.Copy(ioutil.Discard, r)
	return json.NewDecoder(r).Decode(v)
}

type pageData struct {
	News     vetbbedit.ItemData     `json:"news"`
	Services vetbbedit.ItemData     `json:"services"`
	Hours    vetbbedit.OpeningHours `json:"hours"`
}

// serveData is http handler for data REST API.
// GET request returns JSON encoded pageData from server.
// PUT request updates pageData on server from JSON passed in request.
// Other methods are not allowed
func (s *Server) serveData(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h, err := s.OpeningHours.Load()
		if err != nil {
			renderError(w, err)
			return
		}

		n, err := s.News.Load()
		if err != nil {
			renderError(w, err)
			return
		}

		ss, err := s.Services.Load()
		if err != nil {
			renderError(w, err)
			return
		}

		renderJSON(w, &pageData{Hours: *h, News: *n, Services: *ss})

	case "PUT":
		var d pageData
		if err := decodeJSON(r.Body, &d); err != nil {
			renderBadJSONError(w, err)
			return
		}

		// save hours
		if err := s.OpeningHours.Save(&d.Hours); err != nil {
			renderError(w, err)
			return
		}

		// save news
		if err := s.News.Save(&d.News); err != nil {
			renderError(w, err)
			return
		}

		// save services
		if err := s.Services.Save(&d.Services); err != nil {
			renderError(w, err)
			return
		}

		// generate page
		genDir, err := s.PageGen.Generate()
		if err != nil {
			renderError(w, err)
			return
		}

		// upload page
		if err := s.Ftp.Upload(genDir); err != nil {
			renderError(w, err)
			return
		}

		// git commit & push
		if err := s.Repo.Push(); err != nil {
			renderError(w, err)
		}

	default:
		renderNotAllowed(w, r.Method)
	}
}

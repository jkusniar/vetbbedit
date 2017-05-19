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

// serveNews is http handler for news REST API.
// GET request returns JSON encoded news data from server.
// PUT request updates data on server from JSON passed in request.
// Other methods are not allowed
func (s *Server) serveNews(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		n, err := s.NewsService.Load()
		if err != nil {
			renderError(w, err)
			return
		}
		renderJSON(w, n)

	case "PUT":
		var n vetbbedit.ItemData
		if err := decodeJSON(r.Body, &n); err != nil {
			renderBadJSONError(w, err)
			return
		}

		if err := s.NewsService.Save(&n); err != nil {
			renderError(w, err)
		}

	default:
		renderNotAllowed(w, r.Method)
	}
}

// serveServices is http handler for services REST API.
// GET request returns JSON encoded services data from server.
// PUT request updates data on server from JSON passed in request.
// Other methods are not allowed
func (s *Server) serveServices(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		n, err := s.ServicesService.Load()
		if err != nil {
			renderError(w, err)
			return
		}
		renderJSON(w, n)

	case "PUT":
		var n vetbbedit.ItemData
		if err := decodeJSON(r.Body, &n); err != nil {
			renderBadJSONError(w, err)
			return
		}

		if err := s.ServicesService.Save(&n); err != nil {
			renderError(w, err)
		}

	default:
		renderNotAllowed(w, r.Method)
	}
}

// serveOpeningHours is http handler for services REST API.
// GET request returns JSON encoded opening hours data from server.
// PUT request updates data on server from JSON passed in request.
// Other methods are not allowed
func (s *Server) serveOpeningHours(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		n, err := s.OpeningHoursService.Load()
		if err != nil {
			renderError(w, err)
			return
		}
		renderJSON(w, n)

	case "PUT":
		var o vetbbedit.OpeningHours
		if err := decodeJSON(r.Body, &o); err != nil {
			renderBadJSONError(w, err)
			return
		}

		if err := s.OpeningHoursService.Save(&o); err != nil {
			renderError(w, err)
		}

	default:
		renderNotAllowed(w, r.Method)
	}
}

// generate page on backend using hugo
func (s *Server) generate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		renderNotAllowed(w, r.Method)
		return
	}
	// TODO do work and return 201 if OK
}

// upload generated page to server
func (s *Server) upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		renderNotAllowed(w, r.Method)
		return
	}

	// TODO do work and return 201 if OK
}

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
	"context"
	"fmt"
	"log"
	"net/http"
	"path"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/jkusniar/vetbbedit"
)

// generate browser client. Rerun every time client directory contents are changed!
//go:generate go-bindata -pkg http -o client.go client/...

// Server serves REST API and client web application
type Server struct {
	srv *http.Server

	News         vetbbedit.NewsService
	Services     vetbbedit.ServicesService
	OpeningHours vetbbedit.OpeningHoursService
	PageGen      vetbbedit.GeneratorService
}

// Serve starts HTTP server.
// Method blocks until server stopped from another goroutine or error occurs.
func (s *Server) Serve(port uint, devMode bool) error {
	s.srv = &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: s.createServeMux(devMode),
	}

	log.Printf("serving at http://localhost:%d", port)

	if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

// create http server URL multiplexer
func (s *Server) createServeMux(devMode bool) *http.ServeMux {
	mux := http.NewServeMux()

	if devMode {
		mux.Handle("/", http.FileServer(http.Dir(path.Join("http", "client"))))
	} else {
		mux.Handle("/",
			http.FileServer(
				&assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo,
					Prefix: "client"}))
	}

	mux.HandleFunc("/data", s.serveData)

	return mux
}

// Shutdown stops server gracefully
func (s *Server) Shutdown() error {
	if s.srv == nil {
		return nil
	}
	log.Println("Server going down...")
	return s.srv.Shutdown(context.Background())
}

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
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/jkusniar/vetbbedit"
)

type newsMock struct{}

// Save is mock implementation
func (s *newsMock) Save(i *vetbbedit.ItemData) error {
	return nil
}

// Load is mock implementation
func (s *newsMock) Load() (i *vetbbedit.ItemData, err error) {
	i = &vetbbedit.ItemData{Items: []string{"a", "b", "c"}}
	return
}

type servicesMock struct{}

// Save is mock implementation
func (s *servicesMock) Save(i *vetbbedit.ItemData) error {
	return nil
}

// Load is mock implementation
func (s *servicesMock) Load() (i *vetbbedit.ItemData, err error) {
	i = &vetbbedit.ItemData{Items: []string{"x", "y", "z"}}
	return
}

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard) //calm down logger in tests
	os.Exit(m.Run())
}

func TestAllHttpHandlers(t *testing.T) {
	var tests = []struct {
		name            string
		reqMethod       string
		reqURL          string
		reqBody         io.Reader
		expCode         int
		expBody         string
		useBodyContains bool // when true, use strings.Contains to compare actual/expected resp body
	}{
		{"serveNews_GET_OK",
			"GET", "/news", nil, 200,
			`{"items":["a","b","c"]}` + "\n",
			false},
		{"serveNews_PUT_OK",
			"PUT", "/news", strings.NewReader(`{"items":["a","b","c"]}`), 200,
			"", false},
		{"serveNews_PUT_BadJSON",
			"PUT", "/news", strings.NewReader(`:-)`), 400,
			"json decode error", true},
		{"serveNews_POST",
			"POST", "/news", nil, 405,
			"bad method POST\n", false},
		{"serveNews_DELETE",
			"DELETE", "/news", nil, 405,
			"bad method DELETE\n", false},

		{"serveServices_GET_OK",
			"GET", "/services", nil, 200,
			`{"items":["x","y","z"]}` + "\n",
			false},
		{"serveServices_PUT_OK",
			"PUT", "/services", strings.NewReader(`{"items":["x","y","z"]}`), 200,
			"", false},
		{"serveServices_PUT_BadJSON",
			"PUT", "/services", strings.NewReader(`:-)`), 400,
			"json decode error", true},
		{"serveServices_POST",
			"POST", "/services", nil, 405,
			"bad method POST\n", false},
		{"serveServices_DELETE",
			"DELETE", "/services", nil, 405,
			"bad method DELETE\n", false},
	}

	handler := (&Server{
		NewsService:     &newsMock{},
		ServicesService: &servicesMock{},
	}).getServeMux(true)

	for _, tt := range tests {
		req, err := http.NewRequest(tt.reqMethod, tt.reqURL, tt.reqBody)
		if err != nil {
			t.Fatal(err)
		}
		resp := httptest.NewRecorder()
		handler.ServeHTTP(resp, req)

		if tt.expCode != resp.Code {
			t.Fatalf("%s failed. Expected return code %d but was %d",
				tt.name, tt.expCode, resp.Code)
		}

		actBody := resp.Body.String()
		var equals bool

		if tt.useBodyContains {
			equals = strings.Contains(actBody, tt.expBody)
		} else {
			equals = tt.expBody == actBody
		}

		if !equals {
			t.Fatalf("%s failed. Expected \n--%s--\n but was \n--%s--\n",
				tt.name, tt.expBody, actBody)
		}
	}
}

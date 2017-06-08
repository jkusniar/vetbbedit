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

type hoursMock struct{}

// Save is mock implementation
func (s *hoursMock) Save(i *vetbbedit.OpeningHours) error {
	return nil
}

// Load is mock implementation
func (s *hoursMock) Load() (i *vetbbedit.OpeningHours, err error) {
	i = &vetbbedit.OpeningHours{Days: []vetbbedit.DayDef{{"m", "11", "22"}},
		Footnotes: []string{"1", "2"}}
	return
}

type generatorMock struct{}

func (*generatorMock) Generate() error {
	return nil
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
		{"serveData_GET_OK",
			"GET", "/data", nil, 200,
			`{"news":{"items":["a","b","c"]},"services":{"items":["x","y","z"]},"hours":{"days":[{"day":"m","am":"11","pm":"22"}],"footnotes":["1","2"]}}` + "\n",
			false},
		{"serveData_PUT_OK",
			"PUT", "/data", strings.NewReader(`{"news":{"items":["a","b","c"]},"services":{"items":["x","y","z"]},"hours":{"days":[{"day":"m","am":"11","pm":"22"}],"footnotes":["1","2"]}}`), 200,
			"", false},
		{"serveData_PUT_BadJSON",
			"PUT", "/data", strings.NewReader(`:-)`), 400,
			"json decode error", true},
		{"serveData_PUT_Empty",
			"PUT", "/data", strings.NewReader(""), 400,
			"json decode error", true},
		{"serveData_POST",
			"POST", "/data", nil, 405,
			"bad method POST\n", false},
		{"serveData_DELETE",
			"DELETE", "/data", nil, 405,
			"bad method DELETE\n", false},
	}

	handler := (&Server{
		News:         &newsMock{},
		Services:     &servicesMock{},
		OpeningHours: &hoursMock{},
		PageGen:      &generatorMock{},
	}).createServeMux(true)

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

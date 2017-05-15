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

package store

import (
	"testing"
	"time"
)

func TestCopyrightYearsTo(t *testing.T) {
	const tf = "2006-Jan-02"
	cases := []struct {
		name     string
		input    string
		expected string
	}{
		{"past", "2015-Dec-31", "2015"},
		{"current_1", "2016-Jan-01", "2016"},
		{"current_2", "2016-Dec-31", "2016"},
		{"future_1", "2017-Jan-01", "2016 - 2017"},
		{"future_2", "2117-Jan-01", "2016 - 2117"},
	}

	for _, c := range cases {
		now, err := time.Parse(tf, c.input)
		if err != nil {
			t.Fatalf("Error parsing test data %s: %s", c.input, err.Error())
		}

		actual := copyrightYearsTo(now)
		if actual != c.expected {
			t.Errorf("Test '%s' expects '%s' but was '%s'",
				c.name, c.expected, actual)
		}
	}
}

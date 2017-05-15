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
	"encoding/json"
	"os"

	"github.com/pkg/errors"
)

func jsonSave(data interface{}, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return errors.Wrapf(err, "error opening JSON file %s for writing", filePath)
	}
	defer file.Close()

	b, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return errors.Wrap(err, "error encoding JSON")
	}

	_, err = file.Write(b)

	return errors.Wrap(err, "file write error")
}

func jsonLoad(data interface{}, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return errors.Wrapf(err, "error loading JSON file %s", filePath)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(data)

	return errors.Wrap(err, "error decoding JSON")
}

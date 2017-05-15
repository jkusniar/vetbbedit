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
	"path"

	"github.com/jkusniar/vetbbedit"
)

const dataSubDir = "data"

// NewsService is filesystem store implementation of vetbbedit.NewsService
type NewsService struct {
	fileName string
}

// NewNewsService creates new NewsService operating on files in pageDir.
func NewNewsService(pageDir string) *NewsService {
	return &NewsService{path.Join(pageDir, dataSubDir, "news.json")}
}

// Save stores ItemData in json file
func (s *NewsService) Save(i *vetbbedit.ItemData) error {
	return jsonSave(i, s.fileName)
}

// Load returns ItemData from json file
func (s *NewsService) Load() (i *vetbbedit.ItemData, err error) {
	err = jsonLoad(&i, s.fileName)
	return
}

// ServicesService is filesystem store implementation of vetbbedit.ServicesService
type ServicesService struct {
	fileName string
}

// NewServicesService creates new ServicesService operating on files in pageDir.
func NewServicesService(pageDir string) *ServicesService {
	return &ServicesService{path.Join(pageDir, dataSubDir, "services.json")}
}

// Save stores ItemData in json file
func (s *ServicesService) Save(i *vetbbedit.ItemData) error {
	return jsonSave(i, s.fileName)
}

// Load returns ItemData from json file
func (s *ServicesService) Load() (i *vetbbedit.ItemData, err error) {
	err = jsonLoad(&i, s.fileName)
	return
}

// OpeningHoursService is filesystem store implementation of vetbbedit.OpeningHoursService
type OpeningHoursService struct {
	fileName string
}

// NewOpeningHoursService creates new OpeningHoursService operating on files in pageDir.
func NewOpeningHoursService(pageDir string) *OpeningHoursService {
	return &OpeningHoursService{path.Join(pageDir, dataSubDir, "hours.json")}
}

// Save stores OpeningHours in json file
func (s *OpeningHoursService) Save(o *vetbbedit.OpeningHours) error {
	return jsonSave(o, s.fileName)
}

// Load returns OpeningHours from json file
func (s *OpeningHoursService) Load() (o *vetbbedit.OpeningHours, err error) {
	err = jsonLoad(&o, s.fileName)
	return
}

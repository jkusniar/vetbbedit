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

package vetbbedit

// ItemData JSON encoded list of strings used for various data lists (news, services, ...)
type ItemData struct {
	Items []string `json:"items"`
}

// Day is JSON encoded day on opening hours
type Day struct {
	Day string `json:"day"`
	AM  string `json:"am"`
	PM  string `json:"pm"`
}

// OpeningHours is JSON encoded opening hours data
type OpeningHours struct {
	Days      []Day    `json:"days"`
	Footnotes []string `json:"footnotes"`
}

// NewsService manges news data (custom data file, default: data/news.json)
type NewsService interface {
	Save(i *ItemData) error
	Load() (*ItemData, error)
}

// ServicesService manges services data (custom data file, default: data/services.json)
type ServicesService interface {
	Save(i *ItemData) error
	Load() (*ItemData, error)
}

// OpeningHoursService manges opening hours data (custom data file, default: data/hours.json)
type OpeningHoursService interface {
	Save(o *OpeningHours) error
	Load() (*OpeningHours, error)
}

// ConfigService updates hugo configuration data (config.json)
type ConfigService interface {
	UpdateCopyrightYears() error
}

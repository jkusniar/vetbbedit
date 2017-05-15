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
	"strconv"
	"time"
)

// Config parameter from hugo config.json
type configParam struct {
	Description    string `json:"description"`
	Author         string `json:"author"`
	CopyrightYears string `json:"copyrightYearsTo"`
}

// Hugo config.json
type config struct {
	BaseURL        string      `json:"baseurl"`
	LanguageCode   string      `json:"languageCode"`
	Title          string      `json:"title"`
	DisableRSS     string      `json:"disableRSS"`
	DisableSitemap string      `json:"disableSitemap"`
	MetadataFormat string      `json:"metaDataFormat"`
	Params         configParam `json:"params"`
}

// ConfigService is filesystem store implementation of vetbbedit.ConfigService
type ConfigService struct {
	fileName string
}

// NewConfigService creates new ConfigService operating on files in pageDir.
func NewConfigService(pageDir string) *ConfigService {
	return &ConfigService{path.Join(pageDir, "config.json")}
}

// UpdateCopyrightYears updates params.CopyrightYears in config.json using current date
func (s *ConfigService) UpdateCopyrightYears() error {
	var c config
	err := jsonLoad(&c, s.fileName)
	if err != nil {
		return err
	}

	c.Params.CopyrightYears = copyrightYearsTo(time.Now())

	return jsonSave(c, s.fileName)
}

// copyrightYearsTo generates string containing copyright years from launch date (Jan 2016) till now.
func copyrightYearsTo(now time.Time) (copyright string) {
	const shortForm = "2006-Jan-02"
	tFrom, _ := time.Parse(shortForm, "2016-Jan-01")

	yFrom := tFrom.Year()
	yNow := now.Year()

	copyright = strconv.Itoa(yNow)
	if yFrom < yNow {
		copyright = strconv.Itoa(yFrom) + " - " + strconv.Itoa(yNow)
	}

	return
}

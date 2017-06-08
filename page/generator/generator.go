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

package generator

import (
	"os"
	"path"

	"github.com/jkusniar/vetbbedit"
	"github.com/pkg/errors"
	"github.com/spf13/hugo/commands"
)

// Hugo is vetbbedit.GeneratorService implementation
type Hugo struct {
	srcDir     string
	cfgService vetbbedit.ConfigService
}

// New creates initialized Hugo structure
func New(srcDir string, cfgServ vetbbedit.ConfigService) *Hugo {
	return &Hugo{srcDir, cfgServ}
}

// Generate generates page from configuration data in g.SrcDir using hugo. Returns generated directory and error if any.
func (g *Hugo) Generate() (string, error) {
	generatedDir := path.Join(g.srcDir, "public")
	if err := g.cfgService.UpdateCopyrightYears(); err != nil {
		return generatedDir, err
	}

	if err := os.RemoveAll(generatedDir); err != nil {
		return generatedDir, errors.Wrap(err, "error deleting generated page")
	}

	commands.HugoCmd.SetArgs([]string{"--source", g.srcDir})
	return generatedDir, errors.Wrap(commands.HugoCmd.Execute(), "error generating page")
}

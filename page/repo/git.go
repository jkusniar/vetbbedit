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

package repo

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"

	"github.com/pkg/errors"
)

var (
	vcs = "git"
)

type PageGitRepo struct {
	repoDir, repoURL, repoBranch string
}

// NewPageGitRepo returns new initialized instance of PageGitRepo
func NewPageGitRepo(repoDir, repoURL, repoBranch string) *PageGitRepo {
	return &PageGitRepo{repoDir: repoDir, repoURL: repoURL, repoBranch: repoBranch}
}

// Update fetches latest changes from repoURL.
// If repo doesn't exist in repoDir, function clones it from repoURL using
// repoBranch.
func (p *PageGitRepo) Update() error {
	exists := true
	if _, err := os.Stat(p.repoDir); os.IsNotExist(err) {
		exists = false
	}

	if !exists {
		log.Printf("Page not existing yet, cloning from %s to %s\n", p.repoURL, p.repoDir)
		if out, err := runVcs(path.Dir(p.repoDir), []string{"clone", p.repoURL, p.repoDir}); err != nil {
			log.Print(out)
			return err
		}
	}

	if out, err := runVcs(p.repoDir, []string{"checkout", p.repoBranch}); err != nil {
		log.Print(out)
		return err

	}

	if exists {
		out, err := runVcs(p.repoDir, []string{"pull", "--ff-only"})
		if err != nil {
			log.Print(out)
			return err
		}
	}

	return nil

}

// Push commits and pushes changes to repoURL.
func (p *PageGitRepo) Push() error {
	if out, err := runVcs(p.repoDir, []string{"diff"}); err != nil {
		log.Print(out)
		return err
	} else if len(out) != 0 {
		if out, err := runVcs(p.repoDir,
			[]string{"commit", "-a", "-m", "vetbbedit " + time.Now().Format(time.RFC3339)}); err != nil {
			log.Print(out)
			return err
		}
		if out, err := runVcs(p.repoDir, []string{"push", "origin", p.repoBranch}); err != nil {
			log.Print(out)
			return err
		}
	} else {
		log.Println("No changes detected, skipping commit & push")
	}

	return nil
}

// Stolen from golang.org/x/tools/go/vcs
func envForDir(dir string) []string {
	env := os.Environ()
	return mergeEnvLists([]string{"PWD=" + dir}, env)
}

func mergeEnvLists(in, out []string) []string {
NextVar:
	for _, inkv := range in {
		k := strings.SplitAfterN(inkv, "=", 2)[0]
		for i, outkv := range out {
			if strings.HasPrefix(outkv, k) {
				out[i] = inkv
				continue NextVar
			}
		}
		out = append(out, inkv)
	}
	return out
}

func runVcs(dir string, args []string) (string, error) {
	_, err := exec.LookPath(vcs)
	if err != nil {
		return "", errors.Wrap(err, "git command not present")
	}

	cmd := exec.Command(vcs, args...)
	cmd.Dir = dir
	cmd.Env = envForDir(cmd.Dir)
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	err = cmd.Run()
	out := buf.String()
	return out, errors.Wrap(err, "error running git")
}

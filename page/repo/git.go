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
	"archive/tar"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/pkg/errors"
)

const (
	repoSubDir            = "repo"
	currentCommitFileName = "current"
)

// PageGitRepo is remote GIT repo
type PageGitRepo struct {
	localDir, repoURL, repoBranch, repoToken string
	projectID                                uint
	tempArchiveFile                          string
	currentCommitFile                        string
	RepoDir                                  string
}

// NewPageGitRepo returns new initialized instance of PageGitRepo
func NewPageGitRepo(localDir, repoURL, repoBranch, repoToken string, projectID uint) *PageGitRepo {
	return &PageGitRepo{
		localDir:          localDir,
		repoURL:           repoURL,
		repoBranch:        repoBranch,
		projectID:         projectID,
		repoToken:         repoToken,
		tempArchiveFile:   path.Join(os.TempDir(), "vetbbedit.tar.gz"),
		currentCommitFile: path.Join(localDir, currentCommitFileName),
		RepoDir:           path.Join(localDir, repoSubDir)}
}

// Pull fetches latest project version from repository using REST API and extracts it to localDir
func (p *PageGitRepo) Pull() error {

	// get current commit SHA
	currentSHA, err := ioutil.ReadFile(p.currentCommitFile)
	if err != nil && !os.IsNotExist(err) {
		errors.Wrap(err, "error reading current commit SHA")
	}

	// get latest commit SHA
	serverSHA, err := p.getLatestCommit()
	if err != nil {
		return err
	}

	// check if repo update is needed
	if strings.Compare(string(currentSHA), serverSHA) == 0 {
		log.Println("Nothing to download, current repo up to date")
		return nil
	}

	// get archive from server
	if err := p.getRepoArchive(serverSHA); err != nil {
		return err
	}

	// delete previous repo version
	if err := os.RemoveAll(p.RepoDir); err != nil {
		return errors.Wrap(err, "error deleting previous repo version")
	}

	// extract archive to work directory,
	// filter out directory "web-xxxxx-xxxxx" from all paths while extracting
	dirFilter := strings.Join([]string{"web", serverSHA, serverSHA}, "-")
	if err := untar(p.RepoDir, p.tempArchiveFile, dirFilter); err != nil {
		return err
	}

	// store current commit SHA
	if err := p.storeCommitSHA(serverSHA); err != nil {
		return err
	}

	// remove archive
	return errors.Wrap(os.RemoveAll(p.tempArchiveFile), "error removing archive")
}

// Push pushes changes to repo using REST API
func (p *PageGitRepo) Push() error {

	type action struct {
		Action   string `json:"action"`
		FilePath string `json:"file_path"`
		Content  string `json:"content"`
	}

	type payload struct {
		Branch        string   `json:"branch"`
		CommitMessage string   `json:"commit_message"`
		Actions       []action `json:"actions"`
		AuthorEmail   string   `json:"author_email"`
		AuthorName    string   `json:"author_name"`
	}

	// load data files and create commit payload
	configData, err := loadFromFile(path.Join(p.RepoDir, "config.json"))
	if err != nil {
		return err
	}

	newsData, err := loadFromFile(path.Join(p.RepoDir, "data", "news.json"))
	if err != nil {
		return err
	}

	servicesData, err := loadFromFile(path.Join(p.RepoDir, "data", "services.json"))
	if err != nil {
		return err
	}

	hoursData, err := loadFromFile(path.Join(p.RepoDir, "data", "hours.json"))
	if err != nil {
		return err
	}

	pl := payload{
		Branch:        p.repoBranch,
		CommitMessage: strings.Join([]string{"vetbbedit", time.Now().Format(time.RFC3339)}, " "),
		AuthorEmail:   "vetbbedit@veterinabb.sk",
		AuthorName:    "vetbbedit",
		Actions: []action{
			action{Action: "update", FilePath: "config.json", Content: configData},
			action{Action: "update", FilePath: "data/news.json", Content: newsData},
			action{Action: "update", FilePath: "data/services.json", Content: servicesData},
			action{Action: "update", FilePath: "data/hours.json", Content: hoursData},
		},
	}

	// encode json to bytes
	var b bytes.Buffer
	if err := json.NewEncoder(&b).Encode(&pl); err != nil {
		return errors.Wrap(err, "error encoding json")
	}

	// send commit to server
	URL := fmt.Sprintf("%s/projects/%d/repository/commits", p.repoURL, p.projectID)
	resp, err := authenticatedPost(URL, p.repoToken, &b)
	if err != nil {
		return errors.Wrap(err, "error committing")
	}

	// store commit to file
	var c commit
	if err := decodeJSONFromHttpResp(resp, &c); err != nil {
		return err
	}

	return p.storeCommitSHA(c.ID)
}

func loadFromFile(fp string) (string, error) {
	data, err := ioutil.ReadFile(fp)
	if err != nil {
		return "", errors.Wrapf(err, "error reading %s", fp)
	}
	return string(data), nil
}

func (p *PageGitRepo) storeCommitSHA(sha string) error {
	if err := ioutil.WriteFile(p.currentCommitFile, []byte(sha), 0664); err != nil {
		return errors.Wrap(err, "error storing commit sha")
	}
	return nil
}

type commit struct {
	ID string `json:"id"`
}

func (p *PageGitRepo) getLatestCommit() (string, error) {
	URL := fmt.Sprintf("%s/projects/%d/repository/branches/%s", p.repoURL, p.projectID, p.repoBranch)
	resp, err := authenticatedGet(URL, p.repoToken)
	if err != nil {
		return "", err
	}

	type branch struct {
		Name   string `json:"name"`
		Commit commit `json:"commit"`
	}

	var b branch
	if err := decodeJSONFromHttpResp(resp, &b); err != nil {
		return "", err
	}

	return b.Commit.ID, nil
}

func (p *PageGitRepo) getRepoArchive(commitID string) error {
	URL := fmt.Sprintf("%s/projects/%d/repository/archive?sha=%s", p.repoURL, p.projectID, commitID)
	resp, err := authenticatedGet(URL, p.repoToken)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(p.tempArchiveFile)
	if err != nil {
		return errors.Wrap(err, "error creating archive file")
	}
	defer out.Close()

	if _, err := io.Copy(out, resp.Body); err != nil {
		return errors.Wrap(err, "error copying data to archive file")
	}

	return nil
}

func untar(dst, tarFile, filterDir string) error {
	r, err := os.Open(tarFile)
	if err != nil {
		return errors.Wrap(err, "error opening file for reading")
	}
	defer r.Close()

	gzr, err := gzip.NewReader(r)
	defer gzr.Close()
	if err != nil {
		return errors.Wrap(err, "error opening gzip")
	}

	tr := tar.NewReader(gzr)

	for {
		header, err := tr.Next()

		switch {

		// if no more files are found return
		case err == io.EOF:
			return nil

		// return any other error
		case err != nil:
			return errors.Wrap(err, "error decompressing file")

		// if the header is nil, just skip it (not sure how this happens)
		case header == nil:
			continue
		}

		// custom target path: remove directory "web-<commitSHA>-<commitSHA>" from target path
		rel, e := filepath.Rel(filterDir, header.Name)
		if e != nil {
			return errors.Wrap(err, "error computing new file path")
		}
		target := filepath.Join(dst, rel)

		// check the file type
		switch header.Typeflag {

		// if its a dir and it doesn't exist create it
		case tar.TypeDir:
			if err := createDir(target); err != nil {
				return err
			}

		// if it's a file create it
		case tar.TypeReg:
			f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return errors.Wrap(err, "error creating file")
			}
			defer f.Close()

			// copy over contents
			if _, err := io.Copy(f, tr); err != nil {
				return errors.Wrap(err, "error writing file contents")
			}
		}
	}
}

// decodes JSON from http.Response.Body to v
func decodeJSONFromHttpResp(resp *http.Response, v interface{}) error {
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
		return errors.Wrap(err, "error parsing JSON")
	}
	return nil
}

// create dir if not exists
func createDir(dirName string) error {
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		return errors.Wrap(os.MkdirAll(dirName, 0755), "error creating directory")
	}
	return nil
}

// authenticatedGet performs HTTP GET request with supplied auth token
func authenticatedGet(URL string, token string) (*http.Response, error) {
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error while creating http request")
	}

	req.Header.Add("PRIVATE-TOKEN", token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "error calling HTTP API")
	}

	if resp.StatusCode != 200 {
		return nil, errors.Errorf("bad HTTP status %d [%s]",
			resp.StatusCode, resp.Status)
	}

	return resp, nil
}

// authenticatedPost performs HTTP POST request with supplied auth token and content.
// Returns http.Response if successful, error otherwise.
func authenticatedPost(URL, token string, content io.Reader) (*http.Response, error) {
	req, err := http.NewRequest("POST", URL, content)
	if err != nil {
		return nil, errors.Wrap(err, "error while creating http request")
	}

	req.Header.Add("PRIVATE-TOKEN", token)
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "error calling HTTP API")
	}

	fmt.Printf("statuscode : %d\n", resp.StatusCode)

	if (resp.StatusCode < 200) || (resp.StatusCode > 299) {
		// dump response & return error
		defer resp.Body.Close()
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return nil, errors.Wrap(err, "error dumping response")
		}
		log.Printf("RESP: %q\n", dump)

		return nil, errors.Errorf("bad HTTP status %d [%s]",
			resp.StatusCode, resp.Status)
	}

	return resp, nil
}

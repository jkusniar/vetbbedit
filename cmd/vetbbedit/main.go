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

package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"os/user"
	"path"
	"strconv"
	"syscall"

	"github.com/jkusniar/vetbbedit/http"
	"github.com/jkusniar/vetbbedit/page/generator"
	"github.com/jkusniar/vetbbedit/page/repo"
	"github.com/jkusniar/vetbbedit/page/store"
	"github.com/jkusniar/vetbbedit/page/uploader"
	"github.com/zserge/webview"
)

var (
	printVersion = flag.Bool("v", false, "print version and exit")
	devMode      = flag.Bool("devMode", false,
		"development mode. Serve static files from http/client dir")

	// Embedded HTTP server settings
	port     = flag.Uint("port", uint(8080), "http server port [env VETBB_PORT]")
	localDir *string

	// GIT repo parameters
	repoURL = flag.String("repoURL", "https://gitlab.com/api/v4",
		"remote repo REST API url [env VETBB_REPO_URL]")
	repoBranch = flag.String("repoBranch", "master",
		"remote repo branch [env VETBB_REPO_BRANCH]")
	repoToken     = flag.String("repoToken", "xxSECRETxx", "remote repo access token [env VETBB_REPO_TOKEN]")
	repoProjectID = flag.Uint("repoProjectID", uint(979247), "remote repo project id [env VETBB_REPO_PRJID]")

	// SSH upload parameters
	sshHost = flag.String("sshHost", "localhost",
		"remote server SSH hostname [env VETBB_SSH_HOST]")
	sshPort = flag.Uint("sshPort", uint(22), "remote server SSH port [env VETBB_SSH_PORT]")
	sshUser *string
	sshPass = flag.String("sshPass", os.Getenv("SOCKSIE_SSH_PASSWORD"),
		"remote server SSH password [env VETBB_SSH_PASS]")
	sshDir = flag.String("sshDir", "public_html", "remote server SSH directory [env VETBB_SSH_DIR]")
)

const version = "v1.1.0"

func init() {
	u, _ := user.Current()
	sshUser = flag.String("sshUser", u.Username, "remote server SSH username [env VETBB_SSH_USER]")
	localDir = flag.String("localDir", path.Join(u.HomeDir, ".vetbbedit"),
		"web page local directiory [env VETBB_LOCAL_DIR]")
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// flags
	flag.Usage = usage
	envVars()
	flag.Parse()

	if *printVersion {
		fmt.Println(version)
		return
	}

	checkPortNum(*port, "port")
	checkPortNum(*sshPort, "sshPort")

	// git clone/pull page repo
	pr := repo.NewPageGitRepo(*localDir, *repoURL, *repoBranch, *repoToken, *repoProjectID)

	// server
	srv := &http.Server{
		News:         store.NewNewsService(pr.RepoDir),
		Services:     store.NewServicesService(pr.RepoDir),
		OpeningHours: store.NewOpeningHoursService(pr.RepoDir),
		PageGen:      generator.New(pr.RepoDir, store.NewConfigService(pr.RepoDir)),
		Ftp:          uploader.New(*sshHost, *sshPort, *sshUser, *sshPass, *sshDir),
		Repo:         pr,
	}

	// shutdown signal handler
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, os.Kill, syscall.SIGTERM)
	go func() {
		<-quit
		if err := srv.Shutdown(); err != nil {
			log.Fatalf("FATAL: shutdown server failed: %+v\n", err)
		}

		log.Println("Server gracefully stopped")
		os.Exit(0)
	}()

	// start server
	go func() {
		if err := srv.Serve(*port, *devMode); err != nil {
			log.Fatalf("FATAL: starting server failed: %+v\n", err)
		}
	}()

	// start webview
	const startHTML = `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <title>title</title>
  </head>
  <body>
    <h1>LOADING</h1>
  </body>
</html>`
	w := webview.New(webview.Settings{
		Width:     1024,
		Height:    768,
		Title:     "vetbbedit",
		URL:       `data:text/html,` + url.PathEscape(startHTML),
		Resizable: true,
	})
	defer w.Exit()
	w.Dispatch(func() {
		log.Println("started loading data")
		if err := pr.Pull(); err != nil {
			log.Fatalf("FATAL: repo download failed: %+v\n", err)
		}
		log.Println("data loaded")
		w.Eval(fmt.Sprintf("window.location.replace(\"http://localhost:%d\");", *port))
	})
	w.Run()

}

func usage() {
	fmt.Fprintln(os.Stderr, "usage: vetbbedit [flags]")
	fmt.Fprintln(os.Stderr, "flags:")
	flag.PrintDefaults()
	os.Exit(2)
}

func envVars() {
	uintVar(port, "VETBB_PORT")
	stringVar(localDir, "VETBB_LOCAL_DIR")
	stringVar(repoURL, "VETBB_REPO_URL")
	stringVar(repoBranch, "VETBB_REPO_BRANCH")
	stringVar(repoToken, "VETBB_REPO_TOKEN")
	uintVar(repoProjectID, "VETBB_REPO_PRJID")
	stringVar(sshHost, "VETBB_SSH_HOST")
	uintVar(sshPort, "VETBB_SSH_PORT")
	stringVar(sshUser, "VETBB_SSH_USER")
	stringVar(sshPass, "VETBB_SSH_PASS")
	stringVar(sshDir, "VETBB_SSH_DIR")
}

func uintVar(v *uint, env string) {
	s := os.Getenv(env)
	if len(s) != 0 {
		i, err := strconv.Atoi(s)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Env variable %s conversion error: %s. Ignoring.\n",
				env, err)
		} else if i < 0 {
			fmt.Fprintf(os.Stderr, "Env variable %s is negative. Ignoring.\n", env)
		} else {
			*v = uint(i)
		}
	}
}

func stringVar(v *string, env string) {
	s := os.Getenv(env)
	if len(s) != 0 {
		*v = s
	}
}

func checkPortNum(port uint, name string) {
	if port == 0 || port > 65535 {
		fmt.Fprintf(os.Stderr, "%s value %d is out of range [1-65535]\n",
			name, port)
		os.Exit(2)
	}
}

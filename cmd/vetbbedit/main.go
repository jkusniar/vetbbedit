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
	"os"
	"os/signal"
	"os/user"
	"path"
	"strconv"
	"syscall"

	"github.com/jkusniar/vetbbedit/http"
	"github.com/jkusniar/vetbbedit/page/store"
)

var (
	printVersion = flag.Bool("v", false, "print version and exit")
	port         = flag.Uint("port", uint(8080), "http server port [env VETBB_PORT]")
	localDir     *string

	// GIT repo parameters
	repoURL = flag.String("repoURL", "git@github.com:jkusniar/veterinabb.sk.git",
		"web page remote repository [env VETBB_REPO_URL]")
	repoBranch = flag.String("repoBranch", "master",
		"repository branch [env VETBB_REPO_BRANCH]")

	// SSH upload parameters
	sshHost = flag.String("sshHost", "localhost",
		"remote server SSH hostname [env VETBB_SSH_HOST]")
	sshPort = flag.Uint("sshPort", uint(22), "remote server SSH port [env VETBB_SSH_PORT]")
	sshUser *string
	sshPass = flag.String("sshPass", os.Getenv("SOCKSIE_SSH_PASSWORD"),
		"remote server SSH password [env VETBB_SSH_PASS]")
	sshDir = flag.String("sshDir", "", "remote server SSH directory [env VETBB_SSH_DIR]")
)

const version = "v1.0.0"

func init() {
	u, _ := user.Current()
	sshUser = flag.String("sshUser", u.Username, "remote server SSH username [env VETBB_SSH_USER]")
	localDir = flag.String("localDir", path.Join(u.HomeDir, "web"),
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
	checkFileExists(*localDir) // TODO: `git clone` if not existing. `git pull` if existing

	// server
	srv := &http.Server{
		NewsService:         store.NewNewsService(*localDir),
		ServicesService:     store.NewServicesService(*localDir),
		OpeningHoursService: store.NewOpeningHoursService(*localDir),
		ConfigService:       store.NewConfigService(*localDir),
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

	// showtime
	if err := srv.Serve(*port); err != nil {
		log.Fatalf("FATAL: starting server failed: %+v\n", err)
	}
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

func checkFileExists(name string) {
	if _, err := os.Stat(name); err != nil {
		fmt.Fprintf(os.Stderr, "Required file %v not found: %v\n", name, err)
		os.Exit(2)
	}
}

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

package uploader

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/pkg/errors"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// SFtpUploader is sftp implementation of vetbbedit.UploaderService
type SFtpUploader struct {
	host                   string
	port                   uint
	user, password, tgtDir string
}

const (
	indexHTML = "index.html"
)

// New creates initialized SFtpUploader structure
func New(host string, port uint, user, password, tgtDir string) *SFtpUploader {
	return &SFtpUploader{host: host, port: port, user: user, password: password, tgtDir: tgtDir}
}

// Upload uploads generated page to server using ssh protocol
// TODO Upload all generated files, not just index.html
func (u *SFtpUploader) Upload(fromDir string) error {
	cfg := &ssh.ClientConfig{
		User: u.user,
		Auth: []ssh.AuthMethod{
			ssh.Password(u.password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", u.host, u.port), cfg)
	if err != nil {
		return errors.Wrap(err, "error opening connection")
	}
	defer client.Close()

	ftp, err := sftp.NewClient(client)
	if err != nil {
		return errors.Wrap(err, "error creating sftp client")
	}
	defer ftp.Close()

	// copy only index.html (for now)
	targetFilePath := strings.TrimSpace(u.tgtDir)
	if len(targetFilePath) != 0 {
		targetFilePath += "/"
	}
	targetFilePath += indexHTML

	target, err := ftp.Create(targetFilePath)
	if err != nil {
		return errors.Wrapf(err, "error creating target file %s", targetFilePath)
	}
	defer target.Close()

	src, err := os.Open(path.Join(fromDir, indexHTML))
	if err != nil {
		return errors.Wrap(err, "error opening local file")
	}
	defer src.Close()

	if _, err := io.Copy(target, src); err != nil {
		return errors.Wrap(err, "error copying data")
	}

	return nil
}

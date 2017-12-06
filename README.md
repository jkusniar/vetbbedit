# Vetbbedit

[![Build Status](https://travis-ci.org/jkusniar/vetbbedit.svg?branch=master)](https://travis-ci.org/jkusniar/vetbbedit)
[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](https://raw.githubusercontent.com/jkusniar/lara/master/LICENSE)

Web based page editor for [veterinabb.sk](http://veterinabb.sk) using [hugo](https://gohugo.io/).
Runs in embedded webview.

## Install

Tested on go 1.9+

```sh
go install github.com/jkusniar/vetbbedit/cmd/vetbbedit
```

* On linux, requires libwebkitgtk-3.0-dev installed. Webkit has problems rendering supplied JavaScript
* On osx, XCode needs to be installed.

## Run

```sh
$ vetbbedit [flags]
```

### MacOS Application bundle

* navigate to `cmd/vetbbedit` and run `macpack build` (development dependency)
* modify generated `vetbbedit.app/Contents/Info.plist` if needed (e.g. add custom environment variables):
```
	<key>LSEnvironment</key>
	<dict>
		<key>VETBB_REPO_BRANCH</key>
		<string>test</string>
		<key>VETBB_REPO_TOKEN</key>
		<string>MY_TOKEN</string>
	</dict>
```
* copy vetbbedit.app to `~/Applications/`
* run from launchpad

## Development

Requires `make` utility. Install build dependencies if building first time:

* `make build-deps`
* `make`

## Tests

```sh
make test
```

# License

The project license is specified in LICENSE

Vetbbedit is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

Lara is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.
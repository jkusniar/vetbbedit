# Vetbbedit

[![Build Status](https://travis-ci.org/jkusniar/vetbbedit.svg?branch=master)](https://travis-ci.org/jkusniar/vetbbedit)
[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](https://raw.githubusercontent.com/jkusniar/lara/master/LICENSE)

Web based page editor for [veterinabb.sk](http://veterinabb.sk) using [hugo](https://gohugo.io/).

## Install

Tested on go 1.9+

```sh
go install github.com/jkusniar/vetbbedit/cmd/vetbbedit
```

## Run

* install `git`,
* configure git (user.name/user.email),
* setup ssh access to web page repo and run: 

```sh
$ vetbbedit [flags]
```

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
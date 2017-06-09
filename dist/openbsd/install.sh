#!/bin/sh
#
#   Copyright (C) 2016-2017 Contributors as noted in the AUTHORS file
#
#   This file is part of vetbbedit, veterinabb.sk page editor.
#
#   This program is free software: you can redistribute it and/or modify
#   it under the terms of the GNU General Public License as published by
#   the Free Software Foundation, either version 3 of the License, or
#   (at your option) any later version.
#
#   This program is distributed in the hope that it will be useful,
#   but WITHOUT ANY WARRANTY; without even the implied warranty of
#   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
#   GNU General Public License for more details.
#
#   You should have received a copy of the GNU General Public License
#   along with this program.  If not, see <http://www.gnu.org/licenses/>.
#

# create user _vetbbedit
useradd -L daemon -d /var/vetbbedit -s /sbin/nologin -g =uid _vetbbedit

# create dist dir
mkdir /var/vetbbedit
chown _vetbbedit:_vetbbedit /var/vetbbedit

# copy files to dist dir
cp vetbbedit /var/vetbbedit/vetbbedit
chmod +x /var/vetbbedit/vetbbedit
chown _vetbbedit:_vetbbedit /var/vetbbedit/vetbbedit

# copy rc.d script
cp vetbbedit.rc /etc/rc.d/vetbbedit
chmod +x /etc/rc.d/vetbbedit

# .gitconfig
cp gitconfig /var/vetbbedit/.gitconfig
chown _vetbbedit:_vetbbedit /var/vetbbedit/.gitconfig

# ssh keys
mkdir /var/vetbbedit/.ssh
chown _vetbbedit:_vetbbedit /var/vetbbedit/.ssh
ssh-keygen -C "vetbbedit@veterinabb.sk" -f /var/vetbbedit/.ssh/id_rsa -N ''
chown _vetbbedit:_vetbbedit /var/vetbbedit/.ssh/id_rsa
chown _vetbbedit:_vetbbedit /var/vetbbedit/.ssh/id_rsa.pub
ssh-keyscan github.com 2>/dev/null > /var/vetbbedit/.ssh/known_hosts
ssh-keyscan gitlab.com 2>/dev/null >> /var/vetbbedit/.ssh/known_hosts
chown _vetbbedit:_vetbbedit /var/vetbbedit/.ssh/known_hosts

echo "IMPORTANT MANUAL STEPS BEFORE FIRST START:"
echo "\tInstall SSH key /var/vetbbedit/.ssh/id_rsa.pub to git server"
echo "\tFix command line arguments in /etc/rc.d/vetbbedit"
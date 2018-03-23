# === This file is part of octl <http://github.com/teo/octl> ===
#
#  Copyright 2017 CERN and copyright holders of ALICE O².
#  Author: Teo Mrnjavac <teo.mrnjavac@cern.ch>
#
#  Based on Torus project Makefile <https://github.com/coreos/torus>
#
#  This program is free software: you can redistribute it and/or modify
#  it under the terms of the GNU General Public License as published by
#  the Free Software Foundation, either version 3 of the License, or
#  (at your option) any later version.
#
#  This program is distributed in the hope that it will be useful,
#  but WITHOUT ANY WARRANTY; without even the implied warranty of
#  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
#  GNU General Public License for more details.
#
#  You should have received a copy of the GNU General Public License
#  along with this program.  If not, see <http://www.gnu.org/licenses/>.
#
#  In applying this license CERN does not waive the privileges and
#  immunities granted to it by virtue of its status as an
#  Intergovernmental Organization or submit itself to any jurisdiction.

VERSION := 0.1
BUILD := `git rev-parse --short HEAD`

HOST_GOOS=$(shell go env GOOS)
HOST_GOARCH=$(shell go env GOARCH)
REPOPATH = github.com/teo/octl

VERBOSE_1 := -v
VERBOSE_2 := -v -x

WHAT := dummy
SRC_DIRS := ./*

# Use linker flags to provide version/build settings to the target
LDFLAGS=-ldflags "-X=$(REPOPATH).Version=$(VERSION) -X=$(REPOPATH).Build=$(BUILD)"

.PHONY: build generate clean help

build:
	@for target in $(WHAT); do \
		echo "Building $$target"; \
		$(BUILD_ENV_FLAGS) go build $(VERBOSE_$(V)) -o bin/$$target $(LDFLAGS) ./cmd/$$target; \
	done

generate:
	go generate .

clean:
	rm -rf ./bin/*

tools/protoc:
	@echo "Installing Go protoc"
	go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
	go get -u google.golang.org/grpc

help:
	@echo "Influential make variables"
	@echo "  V                 - Build verbosity {0,1,2}."
	@echo "  BUILD_ENV_FLAGS   - Environment added to 'go build'."
	@echo "  WHAT              - Command to build. (e.g. WHAT=octld)"

#
PWD=$(shell pwd)
Name=$(shell basename ${PWD})
CommitID=$(shell git rev-parse HEAD)
LastCommitTime=$(shell git show -s --date=iso-strict-local --format=%cd)
Version=$(shell git describe --tags || git rev-parse --abbrev-ref HEAD)

ifneq ($(strip $(shell git status -s)),)
  Version := ${Version}-dirty
endif

LDFLAGS=-ldflags "-X 'git.woa.com/rio/version.Name=${Name}' -X 'git.woa.com/rio/version.Version=${Version}' -X 'git.woa.com/rio/version.CommitID=${CommitID}' -X 'git.woa.com/rio/version.LastCommitTime=${LastCommitTime}'"

build:
	go build ${LDFLAGS} -o ${Name}
test:
	go test -cover ./...

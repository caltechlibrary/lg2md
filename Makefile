#
# Simple Makefile
#
PROJECT = lg2md

VERSION = $(shell grep -m1 'Version = ' $(PROJECT).go | cut -d\"  -f 2)

BRANCH = $(shell git branch | grep '* ' | cut -d\  -f 2)

build: bin/lg2json bin/lg2md

bin/lg2json: lg2md.go cmds/lg2json/lg2json.go
	go build -o bin/lg2json cmds/lg2json/lg2json.go 

bin/lg2md: lg2md.go cmds/lg2md/lg2md.go
	go build -o bin/lg2md cmds/lg2md/lg2md.go 

test:
	go test

status:
	git status

save:
	if [ "$(msg)" != "" ]; then git commit -am "$(msg)"; else git commit -am "Quick Save"; fi
	git push origin $(BRANCH)

refresh:
	git fetch origin
	git pull origin $(BRANCH)

clean: 
	if [ -d bin ]; then /bin/rm -fR bin; fi

install:
	env GOBIN=$(HOME)/bin go install cmds/lg2json/lg2json.go



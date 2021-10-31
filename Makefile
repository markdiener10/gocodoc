SHELL := /usr/bin/env bash
#Requires MacOS or Linux to work

.DEFAULT_GOAL := verify
#COMMIT_SHA := $(shell git rev-parse HEAD)

.PHONY: clean test ci dev git gen1 gen2

#Helper function to clean up the docker environment
clean:
	@echo "GocoDoc Clean ##############"
	go clean -cache -modcache -testcache

#By default run the macos
test: clean
	@echo "GocoDoc Test ##############"
	go mod tidy
	go vet .
	go test .

#Continuous integration in Github Actions
ci: 
	@echo "Make CI"
	sh ./lambdagolangsetup.sh	
	sh ./lambdapythonsetup.sh
	sh ./lambdanodesetup.sh		
	go clean -cache -modcache
	go mod tidy
	go vet ./pkg/
	go test ./pkg/

#Quick iteration of unit tests
dev: 
	go clean -testcache 
	go test ./pkg

gen1: 
	go clean -testcache 
	go test -run GenOnePackageInPath ./pkg
	git add .
	git commit -m "doc formatting tests"
	git push

gen2: 
	go clean -testcache 
	go test -run GenOnePackageInPath ./pkg
	git add .
	git commit -m "doc formatting tests"
	git push

git: 
	git add .
	git commit -m "doc formatting tests"
	git push

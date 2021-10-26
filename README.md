# gocodoc

Utility for private github golang library repos to generate developer documentation generally triggered by a project Makefile

# Demo

[Show README](#gocodoc/pkg/README.md)

# installation

go install github.com/markdiener10/gocodoc

# usage

gocodoc --src=path --dest=path 

--src  Source code root directory 
--dest Destination of generated documtenation

# doc tags

What is different about this document generator is that developers control documentation output by placing tags within the source code.  Just put either a /* */ or // on the line before a chunk of golang source code and it will be picked up

<|publish|> Cause the code generator to pick up the symbol for publishing to the documentation

# motivation

Reviewing was is currently available:

	godoc - http server oriented with html output 
		* Not self contained for github
		* Relies on external resource for display
		* More general use case, not focused like a razor on github (and its markup
	
# future development

	Expand to support bitbucket and gitlab.  Generate html for usage by the external github html server

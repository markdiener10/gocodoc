# gocodoc - Golang Code Documenter

Golang library document generator

## Demo

[Show Fake Lib Demo](./fakelib/README.md)

## installation

go install github.com/markdiener10/gocodoc

## usage

gocodoc --src=path --dest=path 

--src  Source code root directory path
--dest Destination path of generated documentation

## Motivation

The reason for this utility is that there is a missing piece in golang build pipelines for private source code repositories.  While the pkg.dev site is good for documentation of public source code, there is nothing built for private developer consumption.  

## Other options

Reviewing was is currently available:

	godoc - http server oriented with html output 
		* Not self contained for github
		* Relies on external resource for display
		* More general use case, not focused like a razor on github (and its markup)
		* Difficult to properly simulate command line with godoc and wget

	go doc - google supplied documentation generator
		* Not purpose built for the generation of properly formatted output
		* Little control over the desired output

	go101/golds - https server with partial command line support 
		* generates html 
		* file generation option hard to use
		* poorly documented
		* generates code down the external dependency chain. (Bloat)

# comment tags

What is different about this document generator is that developers control documentation output by placing tags within the comments to control display.  Very similar to HTML markup tags.  

The format for tags is:<|tag,tag,tag|>

It must be placed on the line above a given symbol either in // or /* */ comment blocks without any blank spaces between the symbol and the comment block
	
# future development

	Expand to support bitbucket and gitlab.  Generate html.  Currently built to generator github.com oriented output.

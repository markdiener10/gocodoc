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

Reviewing what was available at the time of this utility's creation:

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

# comment blocks and tags

What is different about this document generator is that developers control documentation output by placing intelligence within the source code to control display.

Comments may be placed:
	- using slashes (// or ///)
	- using multi-line blocks (/*  */)

Comments that begin with 2 slashes are visible in the documentation. Use 3 slashes to suppress visibility.

Example:

```
// This is a comment that will appear in documenation
/// With 3 slashes, this comment will not appear in documentation
```

Comments must be place either immediately above a given golang symbol or on the same line after the symbol.  Documentation terminates upon the first blank line encountered.

Example:
```
//This is above a blank line so it does not appear in documentation

//This is included in documentation because it is immediately above code symbol
var digita int  //This is included in documentation 

/* This is above a blank line so it does not appear in documentation */

/* This is included in documentation */
var digitb int  /* This is included in documentation */
```

The format for tags is:<|tag,tag,tag|>

Example:
```
//<|tag,tag,tag|>  These tags are included in document generator output processing
var digita int  
```
	
# future development

	Define what tags should be enabled and their behavior

	Expand to support bitbucket and gitlab.  Generate html.  Currently built to generator github.com oriented output.

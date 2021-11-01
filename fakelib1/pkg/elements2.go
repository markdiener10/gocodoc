/*
These are multiline comments so we can see what it takes to parse it

<|tag|> */
package fakelib1

import (
	"os"
)

type FakeLibType21 int64 //This should show in documents for the FakeLibType21
type FakeLibType22 bool  ///This should not appear in the comments

const (
	FakeLibConst21 = iota
	FakeLibConst22
	FakeLibConst23
	FakeLibConst24
)

//Here is a comment above the FakeVar21
var FakeVar21 int32   //This is a comment for FakeVar1
var FakeVar22 float64 //This is a comment for FakeVar2

type TestInterface2 interface {
	FuncOne(parm1 string, parm2 int) (string, bool)
	FuncTwo(parm1 string, parm2 int) (string, bool)
}

type TestStruct2A struct {
	PublicParm  string
	privateParm string
}

func (g *TestStruct2A) FuncOne(parm1 string, parm2 int) (string, bool) {
	if len(os.Args) > 0 {
		return "False", false
	}
	return "True", true
}

func (g TestStruct2A) funcTwo(parm1 string, parm2 int) bool {
	return false
}

type TestStruct2B struct {
	PublicParm  string
	privateParm float64
}

func (g *TestStruct2B) FuncOne(parm1 string, parm2 int64) (string, bool) {
	if len(os.Args) > 0 {
		return "False", false
	}
	return "True", true
}

func (g TestStruct2B) funcTwo(parm1 string, parm2 complex128) bool {
	return false
}

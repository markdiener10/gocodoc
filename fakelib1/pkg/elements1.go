//More comments at the top of the file
//<|tagA,tagB,tagC|>
package fakelib1 // This is a comment to show in the documentation for fakelib package

import (
	"os"
)

type FakeLibType11 int64 // This is a comment that WILL appear in the documents
type FakeLibType12 bool  /// This is a comment that WILL NOT appear in the documents (3 slashes)

//Here is a comment above the FakeVar1
var FakeVar1 int32   //This is a comment for FakeVar1
var FakeVar2 float64 //This is a comment for FakeVar2

//Comment above the FAKECONST
const FAKECONST = 3 /* Comment next to FAKECONST */

// This is a comment that will appear above the const type
const (
	FakeLibConst11 = iota // Comment
	FakeLibConst12        ///Comment that does not show in documentation
	FakeLibConst13
	FakeLibConst14 //Comment that does show
)

// This is a comment that will appear above the interface type
type TestInterface1 interface {
	FuncOne(parm1 string, parm2 int) (string, bool) //Comment this will show in documentation
	FuncTwo(parm1 string, parm2 int) (string, bool)
}

func FuncOneAlone(parm1 string, parm2 int) (one string, two bool) { return "", true }
func FuncTwoAlone(parm1 string, parm2 int) string                 { return "" }

type TestStruct1A struct {
	PublicParm  string
	privateParm string
}

func (g *TestStruct1A) FuncOne(parm1 string, parm2 int) (one string, two bool) {
	if len(os.Args) > 0 {
		return "False", false
	}
	return "True", true
}

func (g TestStruct1A) funcTwo(parm1 string, parm2 int) bool {
	return false
}

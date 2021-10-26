package gocodoc

import (
	"os"
	"strings"
	"testing"
)

func TestFileExists(t *testing.T) {

	if err := fileexists("../VERSION"); err != nil {
		t.Error("Error in CmdLine", err.Error())
	}

}

func TestCommandLine(t *testing.T) {

	//We are one directory under the root directory
	os.Args[1] = "--src=../fakelib"
	os.Args[2] = "--dest=../fakedocs"

	src, dest, reb := cmdlineprocess()
	if !reb {
		t.Error("Error in CmdLine")
		return
	}

	if !strings.Contains(os.Args[1], src) {
		t.Error("Source not processed properly")
	}
	if !strings.Contains(os.Args[2], dest) {
		t.Error("Dest not processed properly")
	}
	if err := fileexists(src); err != nil {
		t.Error("Source Does not exist:", err.Error())
	}

	//if err := fileexists(dest); err != nil {
	//	t.Error("Destination Does not exist:", err.Error())
	//}

}

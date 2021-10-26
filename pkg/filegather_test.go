package gocodoc

import (
	"os"
	"testing"
)

func TestGen(t *testing.T) {

	var err error

	//We are one directory under the root directory
	os.Args[1] = "--src=../fakelib"
	os.Args[2] = "--dest=../fakedocs"

	packs := &Tpacks{}
	packs.Init()

	src, dest, reb := cmdlineprocess()
	if !reb {
		t.Error("Error in CmdLine")
		return
	}

	if err = fileexists(src); err != nil {
		t.Error("Source Does not exist:", err.Error())
	}
	if err = fileexists(dest); err != nil {
		t.Error("Destination Does not exist:", err.Error())
	} else {
		//Now remove the old stuff
		os.RemoveAll(dest)
	}

	err = filerecurse(src, packs)
	if err != nil {
		t.Error("File recurse error:", err.Error())
	}

	gengitmarkup(dest, packs)

}

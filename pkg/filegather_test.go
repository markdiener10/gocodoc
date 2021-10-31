package gocodoc

import (
	"os"
	"testing"
)

func Gen(t *testing.T) {

	var err error

	packs := &Tpacks{}
	packs.Init()

	src, dest, reb := Cmdlineprocess()
	if !reb {
		t.Error("Error in CmdLine")
		return
	}

	if err = fileexists(src); err != nil {
		t.Error("Source Does not exist:", err.Error())
	}
	if err = fileexists(dest); err == nil {
		os.RemoveAll(dest)
	}

	err = Filerecurse(src, packs)
	if err != nil {
		t.Error("File recurse error:", err.Error())
	}

	err = Gengitmarkup(dest, packs)
	if err != nil {
		t.Error("Generate Github documentation error:", err.Error())
	}

}

func TestGenOnePackageInPath(t *testing.T) {
	//1 directories, 1 packages
	os.Args[1] = "--src=../fakelib1"
	os.Args[2] = "--dest=../fakedocs1"
	Gen(t)

}

func TestGenMultiplePackagesInPath(t *testing.T) {
	//1 directories, 1 packages
	os.Args[1] = "--src=../fakelib2"
	os.Args[2] = "--dest=../fakedocs2"
	Gen(t)

}

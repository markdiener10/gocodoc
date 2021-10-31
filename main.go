package main

import (
	"fmt"
	"strings"

	gocodoc "./pkg/"
)

func main() {

	src, dest, reb := gocodoc.Cmdlineprocess()
	if !reb {
		return
	}

	//We need to make sure the source code is not deleted
	if strings.Contains(src, dest) {
		//Do not allow the destination to delete the source
		fmt.Println("Destination path cannot be upstream of source path.")
		return
	}

	packs := &gocodoc.Tpacks{}
	packs.Init()

	_ = gocodoc.Filerecurse(src, packs)

	_ = packs

	//_ = gocodoc.Gengitmarkup(dest, packs)

	//Now generate documentation based on the syntax tree

}

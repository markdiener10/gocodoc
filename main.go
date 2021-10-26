package main

import (
	gocodoc "./pkg/"
)

func main() {

	src, dest, reb := gocodoc.cmdlineprocess()
	if !reb {
		return
	}
	packs := &gocodoc.Tpacks{}
	packs.Init()

	_ = gocodoc.filerecurse(src, packs)

	_ = packs

	_ = gocodoc.gengitmarkup(dest, packs)

	//Now generate documentation based on the syntax tree

}

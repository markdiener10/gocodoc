package main

import (
	gocodoc "./pkg/"
)

func main() {

	src, dest, reb := gocodoc.Cmdlineprocess()
	if !reb {
		return
	}

	packs := &gocodoc.Tpacks{}
	packs.Init()

	_ = gocodoc.Filerecurse(src, packs)

	_ = packs
	_ = dest

	//_ = gocodoc.Gengitmarkup(dest, packs)

	//Now generate documentation based on the syntax tree

}

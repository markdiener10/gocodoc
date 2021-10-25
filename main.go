package main

import (
	"./src/gocodoc"
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

	//Now generate documentation based on the syntax tree

}

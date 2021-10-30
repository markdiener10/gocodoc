package gocodoc

import (
	"errors"
	"os"
	"strings"
)

//Now we want to generate a README.md as our generated documentation head.
//Then we want to count the number of packages in the parsed code

func Wn(g *os.File, pre int, out string) {
	g.WriteString(out)
}

func Ws(g *os.File, pre int, out string) {
	if pre > 0 {
		g.WriteString(strings.Repeat(" ", pre))
	}
	g.WriteString(out + "\r\n")
}

func W(g *os.File, out string) {
	g.WriteString(out + "\r\n")
}

func Gengitmarkup(dest string, packs *Tpacks) error {

	var err error

	if packs == nil {
		return errors.New("Invalid package memory supplied")
	}
	if err = fileexists(dest); err != nil {
		err = os.MkdirAll(dest, 0777)
		if err != nil {
			return nil
		}
	}

	//Lets just generate a big README.md file for now
	g, err := os.OpenFile(dest+"/README.md", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer g.Close()

	W(g, "# Documentation")
	W(g, "")

	count := packs.CodeCount()
	if count == 0 {
		W(g, "There are no packages with code in the related repository:"+dest)
		return nil
	}

	W(g, "## Packages")

	P := packs.Reset()

	for packs.Next() {
		P = packs.P
		W(g, P.Name+"  ")
	}

	return nil
}

package gocodoc

import (
	"errors"
	"os"
	"strconv"
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
	g.WriteString(out + "  \r\n")
}

func W(g *os.File, out string) {
	g.WriteString(out + "  \r\n")
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

	packcnt := packs.PackCount()
	if packcnt == 0 {
		W(g, "There are no packages with code in the related repository:"+dest)
		return nil
	}

	var P *Tpack

	if packcnt == 1 {
		return gengitpackage(g, P)
	}

	W(g, "## Packages")

	packs.Reset()
	for packs.Next() {
		P = packs.P
		switch P.Codes.Count() {
		case 0:
			continue
		case 1:
			W(g, "### ["+P.Name+"](#"+P.Name+")")
		default:
			W(g, "### ["+P.Name+"](#"+P.Name+") Files:"+strconv.Itoa(P.Codes.Count()))
		}
	}

	packs.Reset()
	for packs.Next() {
		P = packs.P
		if P.Codes.Count() == 0 {
			continue
		}

		gp, err := os.OpenFile(dest+"/"+P.Name+".md", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0777)
		if err != nil {
			return err
		}
		defer gp.Close()
		err = gengitpackage(g, P)
		if err != nil {
			return err
		}
	}
	return nil
}

func gengitpackage(g *os.File, gp *Tpack) error {
	return nil
}

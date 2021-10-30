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

	P := packs.Reset()

	if packcnt == 1 {
		return gengitpackage(g, P)
	}

	W(g, "## Packages")

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

	W(g, "### ["+gp.Name+"](#"+gp.Name+")")

	var gc *Tcode
	var gv *Tvar
	var gf *Tfunc
	var gs *Tstru
	var gi *Tinterface

	gp.Codes.Reset()
	for gp.Codes.Next() {

		gc = gp.Codes.C

		//types
		W(g, "#### Types")
		gc.Types.Reset()
		for gc.Types.Next() {
			gv = gc.Types.V
			W(g, gv.Name)
		}

		//vars
		if gc.Vars.Count() > 0 {
			W(g, "#### Vars")
			gc.Vars.Reset()
			for gc.Vars.Next() {
				gv = gc.Vars.V
				W(g, gv.Name)
			}
		}

		//interfaces
		if gc.Interfaces.Count() > 0 {
			W(g, "#### Interfaces")
			gc.Interfaces.Reset()
			for gc.Interfaces.Next() {
				gi = gc.Interfaces.I
				W(g, gi.Name)
			}
		}

		//structs
		if gc.Structs.Count() > 0 {
			W(g, "#### Structs")
			gc.Structs.Reset()
			for gc.Structs.Next() {
				gs = gc.Structs.S
				W(g, gs.Name)
			}
		}

		//funcs
		if gc.Funcs.Count() > 0 {
			W(g, "#### Funcs")
			gc.Funcs.Reset()
			for gc.Funcs.Next() {
				gf = gc.Funcs.F
				W(g, gf.Name)
			}
		}
	}

	return nil
}

func gengitcode(g *os.File, gp *Tpack) error {
	return nil
}

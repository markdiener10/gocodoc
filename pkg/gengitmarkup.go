package gocodoc

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

//Now we want to generate a README.md as our generated documentation head.
//Then we want to count the number of packages in the parsed code

func Wn(g *os.File, out string) {
	g.WriteString(out + "  ")
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

func Ws(input string) string {
	input = strings.TrimSpace(input)
	if len(input) > 2 {
		if input[0:3] == "///" {
			return "" //Hidden comment
		}
	}
	return strings.ReplaceAll(input, "/", "")
}

func Wpre(g *os.File, gm *Tmarkup) {
	if len(gm.Precomments) == 0 {
		return
	}
	//W(g, "") //Add blank line from previous entry
	for _, line := range gm.Precomments {
		W(g, "###### "+Ws(line))
	}
}

func Wfunc(g *os.File, pre string, gf *Tfunc) {
	W(g, pre+"Func:"+gf.Name)
	//for line := range gm.Precomments {
}

func We(g *os.File, pre string, out string) {
	if len(out) == 0 {
		return
	}
	g.WriteString(pre + out + "  \r\n")
}

func gengitpackage(g *os.File, gp *Tpack) error {

	W(g, "### Package:"+gp.Name)

	var gc *Tcode
	var gv *Tvar
	var gf *Tfunc
	var gs *Tstru
	var gi *Tinterface
	var gm *Tmarkup
	var gco *Tconst
	var idx int

	gp.Codes.Reset()
	for gp.Codes.Next() {

		gc = gp.Codes.C

		W(g, gc.Filename+" "+gc.Path)
		if gc.Cgo {
			W(g, "C Linkage notice (look at source)")
		}

		//Consts
		W(g, "#### Constants")
		gc.Consts.Reset()
		for gc.Consts.Next() {
			gco = gc.Consts.C
			gm = &gco.Markup
			W(g, "") //Blank line
			Wpre(g, gm)
			for idx, _ = range gco.Items {
				if gco.Public[idx] == false {
					continue
				}
				W(g, gco.Items[idx]+" "+Ws(gco.Comments[idx]))
			}
		}

		//types
		W(g, "#### Types")
		gc.Types.Reset()
		for gc.Types.Next() {
			gv = gc.Types.V
			gm = &gv.Markup
			Wpre(g, gm)
			W(g, gv.Name+" "+gv.Type+" "+Ws(gm.Comment))
		}

		//vars
		if gc.Vars.Count() > 0 {
			W(g, "#### Vars")
			gc.Vars.Reset()
			for gc.Vars.Next() {
				gv = gc.Vars.V
				gm = &gv.Markup
				Wpre(g, gm)
				W(g, gv.Name+" "+gv.Type+" "+Ws(gm.Comment))
			}
		}

		//interfaces
		if gc.Interfaces.Count() > 0 {
			We(g, "#### ", "Interfaces")
			gc.Interfaces.Reset()
			for gc.Interfaces.Next() {
				gi = gc.Interfaces.I
				gm = &gi.Markup
				Wpre(g, gm)
				W(g, gi.Name)
				gi.Funcs.Reset()
				for gi.Funcs.Next() {
					gf = gi.Funcs.F
					Wfunc(g, "- ", gf)
				}
			}
		}

		//structs
		if gc.Structs.Count() > 0 {
			W(g, "#### Structs")
			gc.Structs.Reset()
			for gc.Structs.Next() {
				gs = gc.Structs.S
				gm = &gs.Markup
				Wpre(g, gm)
				W(g, gs.Name+" "+Ws(gm.Comment))
			}
		}

		//funcs
		if gc.Funcs.Count() > 0 {
			W(g, "#### Funcs")
			gc.Funcs.Reset()
			W(g, "```golang")
			for gc.Funcs.Next() {
				gf = gc.Funcs.F
				W(g, "Func:"+gf.Name)
			}
			W(g, "```")
		}

	}

	return nil
}

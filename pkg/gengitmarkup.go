package gocodoc

import (
	"errors"
	"os"
	"strconv"
)

//Now we want to generate a README.md as our generated documentation head.
//Then we want to count the number of packages in the parsed code
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

	w(g, "# Documentation")

	packcnt := packs.PackCount()
	if packcnt == 0 {
		w(g, "There are no packages with code in the related repository:"+dest)
		return nil
	}

	P := packs.Reset()

	if packcnt == 1 {
		err = gengitsummary(g, P)
		if err != nil {
			return err
		}
		return gengitdetailed(g, P, false)
	}

	w(g, "## Packages")

	for packs.Next() {
		P = packs.P
		switch P.Codes.Count() {
		case 0:
			continue
		case 1:
			w(g, "### ["+P.Name+"](#"+P.Name+")")
		default:
			w(g, "### ["+P.Name+"](#"+P.Name+") Files:"+strconv.Itoa(P.Codes.Count()))
		}
	}

	//Now, we want to build a summary page
	//and each package gets its own detailed page

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
		err = gengitsummary(gp, P)
		if err != nil {
			return err
		}
	}
	return nil
}

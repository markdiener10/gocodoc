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
			return err
		}
	} else {
		err = cleanout(dest, ".md")
		if err != nil {
			return err
		}
	}

	gm := &Tmarkdown{}
	gm.init()

	//Lets just generate a big README.md file for now
	err = gm.open(dest + "README.md")
	if err != nil {
		return err
	}
	defer gm.close()

	gm.wh(1, "Documentation")

	packcnt := packs.PackCount()
	if packcnt == 0 {
		gm.w("There are no packages with code in the related repository:" + dest)
		return nil
	}

	P := packs.Reset()

	if packcnt == 1 {
		err = gengitsummary(gm, P)
		if err != nil {
			return err
		}
		return nil
		//return gengitdetailed(gm, P, false)
	}

	gm.wh(2, "Packages")

	for packs.Next() {
		P = packs.P
		switch P.Codes.Count() {
		case 0:
			continue
		case 1:
			gm.wh(3, "["+P.Name+"](#"+P.Name+")")
		default:
			gm.wh(3, "["+P.Name+"](#"+P.Name+") Files:"+strconv.Itoa(P.Codes.Count()))
		}
	}

	//Now, we want to build a summary page
	//and each package gets its own detailed page

	gh := &Tmarkdown{}
	gh.init()
	defer gh.close()

	packs.Reset()
	for packs.Next() {
		P = packs.P
		if P.Codes.Count() == 0 {
			continue
		}
		gh.close()
		err = gh.open(dest + P.Name + ".md")
		if err != nil {
			return err
		}
		err = gengitsummary(gh, P)
		if err != nil {
			return err
		}
	}
	return nil
}

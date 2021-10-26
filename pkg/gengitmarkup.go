package gocodoc

import (
	"errors"
	"os"
)

//Now we want to generate a README.md as our generated documentation head.
//Then we want to count the number of packages in the parsed code

func gengitmarkup(dest string, packs *Tpacks) error {

	var err error 
	if packs == nil {
		return errors.New("Invalid package memory supplied")
	}
	if err = fileexists(dest); err != nil {
		err = os.MkdirAll(dest,os.ModeAppend)
		if err != nil {
			return nil
		}
	}



	return nil
}

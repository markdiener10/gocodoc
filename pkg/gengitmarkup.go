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
		err = os.MkdirAll(dest, 0777)
		if err != nil {
			return nil
		}
	}

	//Lets just generate a big README.md file for now
	file, err := os.OpenFile(dest+"/README.md", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	file.WriteString("# Documentation")
	file.WriteString("  ")
	pack := packs.Reset()
	for {
		if pack = packs.Next(); pack == nil {
			break
		}
		file.WriteString(pack.Name + "  ")
	}

	file.Close()
	return nil
}

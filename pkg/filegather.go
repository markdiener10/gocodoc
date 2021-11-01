package gocodoc

import (
	"bufio"
	"os"
	"strings"
)

func Filerecurse(path string, packs *Tpacks) error {

	var s string
	var lines *[]string

	filelist, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, file := range filelist {
		if file.IsDir() {
			err = Filerecurse(path+"/"+file.Name(), packs)
			if err != nil {
				return err
			}
			continue
		}
		s = file.Name()
		if s[len(s)-3:] != ".go" {
			continue
		}
		if strings.Contains(s, "_test.go") {
			continue
		}
		lines, err = fileload(s, path)
		if err != nil {
			return err
		}
		pack, packline, err := parsepackage(s, path, lines, packs)
		if err != nil {
			return err
		}
		code, err := parsecode(s, path, pack)
		if err != nil {
			return err
		}
		//Scan above this like for package
		err = parsemarkup(&code.Markup, packline, false, lines)
		if err != nil {
			return err
		}
		err = parsecgo(code, packline, lines)
		if err != nil {
			return err
		}
		err = parseconst(code, packline, lines)
		if err != nil {
			return err
		}

		err = parseinterface(code, packline, lines)
		if err != nil {
			return err
		}

		err = parsetypes(code, packline, lines)
		if err != nil {
			return err
		}

		err = parsevars(code, packline, lines)
		if err != nil {
			return err
		}

	}
	return nil
}

func fileload(filename string, path string) (*[]string, error) {

	var line string
	var lines []string
	var idx int

	fh, err := os.Open(path + "/" + filename)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	fh.Close()
	if err = scanner.Err(); err != nil {
		return nil, err
	}

	//Remove all leading and trailing space (gofmt already does this)
	for idx, line = range lines {
		lines[idx] = strings.TrimSpace(line)
	}
	return &lines, nil
}

func cleanout(path string, ext string) error {

	var s string

	ext = strings.TrimSpace(ext)
	extlen := len(ext)

	filelist, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, file := range filelist {
		if file.IsDir() {
			err = cleanout(path+"/"+file.Name(), ext)
			if err != nil {
				return err
			}
			continue
		}
		s = file.Name()
		if len(s) <= extlen {
			continue
		}
		if extlen > 0 {
			if s[len(s)-extlen:] != ext {
				continue
			}
		}
		err = os.Remove(path + "/" + s)
		if err != nil {
			return err
		}
	}
	return nil
}

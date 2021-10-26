package gocodoc /*  what is this */ //More comments

import (
	"errors"
	"strconv"
	"strings"
)

func parsepackage(filename string, path string, lines *[]string, packs *Tpacks) (*Tpack, int, error) {

	var packname string
	var packline int = -1
	var idx int
	var line string

	for idx, line = range *lines {
		if len(line) < 7 {
			continue
		}
		if line[0:7] != "package" {
			continue
		}
		packline = idx
		break
	}
	if packline == -1 {
		return nil, 0, errors.New("Found source file without package declaration:" + filename)
	}
	line = (*lines)[packline]

	//package NameOfPack //
	firstspace := Iat(line, " ", 0, -1)
	if firstspace == -1 {
		return nil, 0, errors.New("Found source file with invalid package declaration:" + filename)
	}

	line = stripcomment(line, uint(firstspace+1))

	nextspace := Iat(line, " ", uint(firstspace+1), -1)
	if nextspace == -1 {
		packname = line[firstspace+1:]
	} else {
		packname = line[firstspace+1 : nextspace]
	}
	pack := packs.Get(packname)
	if pack == nil {
		return nil, 0, errors.New("Unable to load a new package in memory:" + filename)
	}
	return pack, packline, nil
}

func parsecode(filename string, path string, pack *Tpack) (*Tcode, error) {

	code := pack.Codes.Get(path, filename)
	if code == nil {
		return nil, errors.New("Unable to load a new code in memory:" + path + " " + filename)
	}
	return code, nil
}

/*
 */
//package packname /// Comment
func parsemarkup(markup *Tmarkup, lineRef int, noscan bool, lines *[]string) error {

	line := (*lines)[lineRef]
	idx := Iat(line, "//", 0, -1)
	if idx > -1 {
		markup.Comment = line[idx:]
	}
	if lineRef == 0 {
		return nil
	}

	if noscan == true {
		return nil
	}

	//Now move backwards and find the beginning of the comments
	//Either a /* CODEBLOCK  */
	//Or an unbroken sequence of //
	commentTop := -1
	commentBottom := lineRef - 1

	//Now we need to analyze the line right above the reference line
	line = (*lines)[commentBottom]

	if len(line) < 2 {
		return nil
	}

	//Find the end of the comment block on the first line above our reference line
	idx = Iat(line, "*/", 0, -1)
	if idx == -1 {
		//No comment block used, now just have a series of // line comments
		for lineIdx := commentBottom; lineIdx >= 0; lineIdx-- {
			line = (*lines)[lineIdx]
			if len(line) < 2 {
				break
			}
			if line[0:2] != "//" {
				break
			}
			commentTop = lineIdx
			continue
		}
		if commentTop == -1 {
			return nil
		}

	} else {

		//Now go searching for the beginning of the comment block
		for lineIdx := commentBottom; lineIdx >= 0; lineIdx-- {
			line = (*lines)[lineIdx]
			if len(line) < 2 {
				continue
			}
			if line[0:1] != "/*" {
				continue
			}
			commentTop = lineIdx
			break
		}
		if commentTop == -1 {
			return nil
		}
	}

	//Now roll up the comments
	for lineIdx := commentTop; lineIdx <= commentBottom; lineIdx++ {
		line = (*lines)[lineIdx]
		markup.Precomments = append(markup.Precomments, line)
	}

	//Now Fish out the tags of the first line above our reference line
	//<|tagA,tagB,tagC|>
	line = (*lines)[commentBottom]
	idx = Iat(line, "<|", 0, -1)
	if idx == -1 {
		return nil
	}
	idxb := Iat(line, "|>", uint(idx+2), -1)
	if idxb == -1 {
		return nil
	}
	markup.Tags = strings.Split(line[idx+1:idxb], ",")
	return nil
}

func parsecgo(code *Tcode, packline int, lines *[]string) error {
	for idx, line := range *lines {
		if idx <= packline {
			continue
		}
		if len(line) < 6 {
			continue
		}
		if line[0:6] != "import" {
			continue
		}

		line = stripcomment(line, uint(5))

		if strings.Contains(line, `"C"`) {
			code.Cgo = true
			return nil
		}
	}
	return nil
}

func parseconst(code *Tcode, packline int, lines *[]string) error {

	var comment string
	var idxComment int

	for idx, line := range *lines {

		if idx <= packline {
			continue
		}
		if len(line) < 5 {
			continue
		}
		if line[0:5] != "const" {
			continue
		}

		line = stripcomment(line, uint(5))

		sconst := code.Consts.Add()
		sconst.Line = idx

		err := parsemarkup(&sconst.Markup, idx, false, lines)
		if err != nil {
			return err
		}

		//Now determine if we have a multi-line const or a single line const
		idxTest := Iat(line, "(", uint(4), -1)
		if idxTest == -1 {

			//const WHAT = 3
			idxTestb := Iat(line, "=", uint(4), -1)
			if idxTestb == -1 {
				continue
			}
			sconst.Add(strings.TrimSpace(line[5:idxTestb]), "")
			continue
		}

		for idxSub := idx + 1; idxSub < len(*lines); idxSub++ {
			line = (*lines)[idxSub]
			if len(line) == 0 {
				continue
			}
			if line[0:1] == ")" {
				break
			}

			idxComment = Iat(line, "//", 0, -1)
			if idxComment == -1 {
				idxComment = Iat(line, "/*", 0, -1)
			}
			comment = ""
			if idxComment > -1 {
				comment = strings.TrimSpace(line[idxComment:])
				line = strings.TrimSpace(line[0:idxComment])
			}
			idxTest = Iat(line, "=", 0, -1)
			if idxTest > -1 {
				line = strings.TrimSpace(line[0:idxTest])
			}
			sconst.Add(line, comment)
		}
	}
	return nil
}

func parsefunc(code *Tcode, stru *Tstru, line string, idx int) (*Tfunc, error) {

	var parms []string
	var chunks []string
	var parm string
	var parmblock string
	var sfunc *Tfunc
	var svar *Tvar

	line = stripcomment(line, uint(3))

	//Cannot be a function without having a bracket to open the function
	//Regardless of whether the closing bracking is on the same line or later
	idxBracket := Iat(line, "{", uint(0), -1)
	if idxBracket == -1 {
		return nil, errors.New("Unable to find function bracket:" + code.Filename + " line:" + strconv.Itoa(idx))
	}

	cntLeft := strings.Count(line, "(")
	if cntLeft == 0 {
		return nil, errors.New("Unable to find left paren:" + code.Filename + " line:" + strconv.Itoa(idx))
	}
	cntRight := strings.Count(line, ")")
	if cntRight == 0 {
		return nil, errors.New("Unable to find right paren:" + code.Filename + " line:" + strconv.Itoa(idx))
	}
	if cntLeft != cntRight {
		return nil, errors.New("Mismatched paren count:" + code.Filename + " line:" + strconv.Itoa(idx))
	}

	idxStruRight := -1
	idxStruLeft := -1
	idxFuncRight := -1
	idxFuncLeft := -1
	idxRetRight := -1
	idxRetLeft := -1

	switch cntLeft {
	case 1: //Function, no structure or return
		if stru != nil {
			return nil, errors.New("Insufficient paren count:" + code.Filename + " line:" + strconv.Itoa(idx))
		}
		idxFuncLeft = Iat(line, "(", 0, -1)
		idxFuncRight = Iat(line, ")", uint(idxFuncLeft+1), -1)
		idxRetLeft = idxFuncRight
		idxRetRight = idxBracket
	case 2: //Structure and Function, No return
		idxStruLeft = Iat(line, "(", uint(0), idxBracket-1)
		idxStruRight = Iat(line, ")", uint(idxStruLeft+1), idxBracket-1)
		idxFuncLeft = Iat(line, "(", uint(idxStruRight+1), idxBracket-1)
		idxFuncRight = Iat(line, ")", uint(idxFuncLeft+1), idxBracket-1)
		idxRetLeft = idxFuncRight
		idxRetRight = idxBracket
	default: //Structure, Function, and Return parentheses
		idxStruLeft = Iat(line, "(", uint(0), idxBracket-1)
		idxStruRight = Iat(line, "(", uint(idxStruLeft+1), idxBracket-1)
		idxFuncLeft = Iat(line, "(", uint(idxStruRight+1), idxBracket-1)
		idxFuncRight = Iat(line, ")", uint(idxFuncLeft+1), idxBracket-1)
		idxRetLeft = Iat(line, "(", uint(idxFuncRight+1), idxBracket-1)
		idxRetRight = Iat(line, ")", uint(idxRetLeft+1), idxBracket-1)
	}

	//func myfunc(parms)(returns)
	//func myfunc(parms)returns
	//func (stru) myfunc(parms)(returns)
	//func (stru) myfunc(parms) returns
	funcname := strings.TrimSpace(line[idxStruRight+1 : idxFuncLeft-1])
	if stru != nil {
		sfunc = stru.Funcs.Get(funcname)
	} else {
		sfunc = code.Funcs.Get(funcname)
	}
	if sfunc == nil {
		return nil, errors.New("Unable to allocate function memory:" + code.Filename)
	}

	parmblock = strings.TrimSpace(line[idxFuncLeft+1 : idxFuncRight-1])
	parms = strings.Split(parmblock, ",")
	for _, parm = range parms {
		parm = strings.TrimSpace(parm)
		chunks = strings.Split(parm, " ")
		if len(chunks) != 2 {
			continue
		}
		svar = sfunc.Parms.Get(chunks[0])
		svar.Type = chunks[1]
		svar.Line = idx
	}

	parmblock = strings.TrimSpace(line[idxRetLeft+1 : idxRetRight-1])
	parms = strings.Split(parmblock, ",")
	for _, parm = range parms {
		parm = strings.TrimSpace(parm)
		chunks = strings.Split(parm, " ")
		switch len(chunks) {
		case 0:
			continue
		case 1:
			svar = sfunc.Returns.Add()
			svar.Type = chunks[0]
			svar.Line = idx
		default: //2
			svar = sfunc.Returns.Get(chunks[0])
			svar.Type = chunks[1]
			svar.Line = idx
		}
	}
	return sfunc, nil
}

func parsestruct(code *Tcode, packline int, lines *[]string) error {

	var idxType int
	var idxSub int
	var idxSpace int
	var name string
	var stype string
	var err error
	var svar *Tvar

	for idx, line := range *lines {

		if idx <= packline {
			continue
		}
		if len(line) < 4 {
			continue
		}
		if line[0:4] != "type" {
			continue
		}

		line = stripcomment(line, uint(3))

		idxType = Iat(line, "struct", 0, -1)
		if idxType == -1 {
			continue
		}
		name = strings.TrimSpace(line[4 : idxType-1])

		stru := code.Structs.Get(name)
		stru.Line = idx
		err = parsemarkup(&stru.Markup, idx, false, lines)
		if err != nil {
			return err
		}

		//Now load up the member variables on following lines
		for idxSub = idx + 1; idxSub < len(*lines); idxSub++ {
			line = (*lines)[idxSub]
			if len(line) == 0 {
				continue
			}
			//End of struct definition
			if line[0:1] == "}" {
				break
			}

			//Could be a composition struct member with no name, just type
			idxSpace = Iat(line, " ", 0, -1)
			if idxSpace == -1 {
				name = line
				stype = ""
			} else {
				name = strings.TrimSpace(line[0:idxSpace])
				stype = strings.TrimSpace(line[idxSpace+1 : 0])
			}
			svar = stru.Vars.Get(name)
			svar.Line = idx
			svar.Type = stype
		}

		//Now go down and find the member functions
		for idxSub = idx + 1; idxSub < len(*lines); idxSub++ {

			line = (*lines)[idxSub]
			if len(line) < 4 {
				continue
			}
			//End of struct definition
			if line[0:4] != "func" {
				continue
			}
			//Is the function one of ours
			if !strings.Contains(line, stru.Name) {
				continue
			}
			//Trim off "func" from front of line
			sfunc, err := parsefunc(code, stru, line[3:], idx)
			if err != nil {
				return err
			}
			err = parsemarkup(&sfunc.Markup, idx, false, lines)
			if err != nil {
				return err
			}

		}
	}
	return nil
}

func parseinterface(code *Tcode, packline int, lines *[]string) error {

	var idxType int
	var idxSub int
	var name string
	var err error
	var svar *Tvar

	for idx, line := range *lines {

		if idx <= packline {
			continue
		}
		if len(line) < 4 {
			continue
		}
		if line[0:4] != "type" {
			continue
		}

		line = stripcomment(line, uint(3))

		idxType = Iat(line, "interface", 0, -1)
		if idxType == -1 {
			continue
		}
		name = strings.TrimSpace(line[4:idxType])

		sinterface := code.Interfaces.Get(name)
		sinterface.Line = idx

		err = parsemarkup(&sinterface.Markup, idx, false, lines)
		if err != nil {
			return err
		}

		//Now load up the member function on following lines
		for idxSub = idx + 1; idxSub < len(*lines); idxSub++ {
			line = (*lines)[idxSub]
			if len(line) == 0 {
				continue
			}
			//End of interface definition
			if line[0:1] == "}" {
				break
			}

			line = stripcomment(line, 0)

			cntLeft := strings.Count(line, "(")
			if cntLeft == 0 {
				return errors.New("Unable to find interface left paren:" + code.Filename + " line:" + strconv.Itoa(idx))
			}
			cntRight := strings.Count(line, ")")
			if cntRight == 0 {
				return errors.New("Unable to find interface right paren:" + code.Filename + " line:" + strconv.Itoa(idx))
			}
			if cntLeft != cntRight {
				return errors.New("interface paren count mismatch:" + code.Filename + " line:" + strconv.Itoa(idx))
			}

			idxFuncLeft := Iat(line, "(", 0, -1)
			idxFuncRight := Iat(line, ")", uint(idxFuncLeft+1), -1)
			idxRetLeft := idxFuncRight
			idxRetRight := len(line)
			if cntLeft == 2 {
				idxRetLeft = Iat(line, "(", uint(idxFuncRight+1), -1)
				idxRetRight = Iat(line, ")", uint(idxRetLeft+1), -1)
			}
			//myfunc(parms)(returns)
			//myfunc(parms)returns
			funcname := strings.TrimSpace(line[0:idxFuncLeft])
			sfunc := sinterface.Funcs.Get(funcname)
			if sfunc == nil {
				return errors.New("Unable to allocate function memory:" + code.Filename)
			}

			err = parsemarkup(&sfunc.Markup, idx, false, lines)
			if err != nil {
				return err
			}

			parmblock := strings.TrimSpace(line[idxFuncLeft+1 : idxFuncRight])
			parms := strings.Split(parmblock, ",")
			for _, parm := range parms {
				parm = strings.TrimSpace(parm)
				chunks := strings.Split(parm, " ")
				if len(chunks) != 2 {
					continue
				}
				svar = sfunc.Parms.Get(chunks[0])
				svar.Type = chunks[1]
				svar.Line = idx
			}

			parmblock = strings.TrimSpace(line[idxRetLeft+1 : idxRetRight])
			parms = strings.Split(parmblock, ",")
			for _, parm := range parms {
				parm = strings.TrimSpace(parm)
				chunks := strings.Split(parm, " ")
				switch len(chunks) {
				case 0:
					continue
				case 1:
					svar = sfunc.Returns.Add()
					svar.Type = chunks[0]
					svar.Line = idx
				case 2:
					svar = sfunc.Returns.Get(chunks[0])
					svar.Type = chunks[1]
					svar.Line = idx
				}
			}
		}
	}
	return nil
}

func parsetypes(code *Tcode, packline int, lines *[]string) error {

	var idxType int
	var idxFirst int
	var idxLast int
	var name string
	var err error

	for idx, line := range *lines {

		if idx <= packline {
			continue
		}
		if len(line) < 4 {
			continue
		}
		if line[0:4] != "type" {
			continue
		}

		line = stripcomment(line, uint(3))

		idxType = Iat(line, "interface", 0, -1)
		if idxType > -1 {
			continue
		}
		idxType = Iat(line, "struct", 0, -1)
		if idxType > -1 {
			continue
		}
		idxFirst = Iat(line, " ", 4, -1)
		idxLast = RevIat(line, " ", 0, -1)

		if idxFirst >= idxLast {
			return errors.New("type parens mis-formatted:" + code.Filename + " line:" + strconv.Itoa(idx))
		}
		name = strings.TrimSpace(line[idxFirst+1 : idxLast])
		svar := code.Types.Get(name)
		svar.Type = strings.TrimSpace(line[idxLast+1:])
		err = parsemarkup(&svar.Markup, idx, false, lines)
		if err != nil {
			return err
		}
	}
	return nil
}

func parsevars(code *Tcode, packline int, lines *[]string) error {

	var idxFirst int
	var idxLast int
	var name string
	var err error

	for idx, line := range *lines {

		if idx <= packline {
			continue
		}
		if len(line) < 3 {
			continue
		}
		if line[0:3] != "var" {
			continue
		}

		line = stripcomment(line, uint(3))

		idxFirst = Iat(line, " ", 2, -1)
		idxLast = RevIat(line, " ", 0, -1)
		if idxFirst >= idxLast {
			return errors.New("var parens mis-formatted:" + code.Filename + " line:" + strconv.Itoa(idx))
		}
		name = strings.TrimSpace(line[idxFirst+1 : idxLast])
		svar := code.Vars.Get(name)
		svar.Type = strings.TrimSpace(line[idxLast+1:])
		err = parsemarkup(&svar.Markup, idx, false, lines)
		if err != nil {
			return err
		}
	}
	return nil
}

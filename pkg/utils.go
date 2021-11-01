package gocodoc

import (
	"strings"
)

func Iat(haystack string, needle string, offset uint, max int) int {

	noffset := int(offset)

	hayLen := len(haystack)
	if noffset >= hayLen {
		return -1
	}
	if max > -1 {
		if max < hayLen {
			hayLen = max
		}
	}

	needLen := len(needle)
	if needLen == 0 {
		return -1
	}
	if hayLen < needLen {
		return -1
	}

	idxa := noffset
	idxb := 0

	for {

	Nextloop:

		if idxa >= hayLen {
			return -1
		}

		for idxb = 0; idxb < needLen; idxb++ {
			if idxa+idxb >= hayLen {
				return -1
			}
			if needle[idxb] == haystack[idxa+idxb] {
				continue
			}
			idxa++
			goto Nextloop
		}
		return idxa
	}
	return -1
}

//Optimized - look backwards
func RevIat(haystack string, needle string, offset uint, max int) int {

	noffset := int(offset)

	hayLen := len(haystack)
	if max > -1 {
		if max+noffset >= hayLen {
			return -1
		}
		hayLen = max
	} else {
		if noffset >= hayLen {
			return -1
		}
	}

	needLen := len(needle)
	if needLen == 0 {
		return -1
	}
	if hayLen < needLen {
		return -1
	}

	//Reverse the needle for easier searching
	var revneedle string
	for _, v := range needle {
		revneedle = string(v) + revneedle
	}

	idxa := hayLen - (1 + noffset)
	idxb := 0

	for {

	Nextloop:

		if idxa < 0 {
			return -1
		}

		for idxb = 0; idxb < needLen; idxb++ {

			if revneedle[idxb] == haystack[idxa-idxb] {
				continue
			}
			idxa--
			goto Nextloop
		}
		return (idxa - (idxb - 1))
	}
	return -1

}

func stripcomment(line string, offset uint) string {

	idxSlash := Iat(line, "//", offset, -1)
	idxBlock := Iat(line, "/*", offset, -1)
	if idxSlash == -1 {
		if idxBlock == -1 {
			return line
		}
	}
	if idxSlash == -1 {
		return strings.TrimSpace(line[0:idxBlock])
	}
	if idxBlock == -1 {
		return strings.TrimSpace(line[0:idxSlash])
	}
	if idxBlock < idxSlash {
		return strings.TrimSpace(line[0:idxBlock])
	}
	return strings.TrimSpace(line[0:idxSlash])
}

func bubblesort(sz int, fn func(i int, j int) bool) {
	var n int
	var swapped bool
	for n = sz; n > 1; n-- {
		swapped = false
		for i := 1; i < n; i++ {
			if fn(i-1, i) {
				swapped = true
			}
		}
		if !swapped {
			return
		}
	}
}

//Allow us to compress the Tcodes based on file/path
//into Tcodes based on path only (all source codes within directory are merged code-wise)
func SplitByPath(g *Tpack) []*Tcode {

	glist := []*Tcodes{}
	var gcodes *Tcodes
	var codea *Tcode
	var codeb *Tcode
	var found bool

	for _, codea = range g.Codes.List {

		found = false
		for _, gcodes = range glist {
			if len(gcodes.List) == 0 {
				continue
			}
			codeb = gcodes.List[0]
			if codea.Path != codeb.Path {
				break
			}
			gcodes.List = append(gcodes.List, codea)
			found = true
			break
		}
		if found == false {
			gcodes := &Tcodes{}
			gcodes.Init()
			gcodes.List = append(gcodes.List, codea)
			glist = append(glist, gcodes)
		}
	}

	//Now our codes are organized by path
	//We need to squash down the list and sort the funcs, structs, and interfaces by name
	//And stuff them into a single code

	gclist := []*Tcode{}

	for _, gcodes = range glist {
		if len(gcodes.List) == 0 {
			continue
		}
		codea := &Tcode{}
		codea.Init()
		codea.Cgo = false
		gcodes.Reset()
		for gcodes.Next() {
			codeb = gcodes.C
			if codeb.Cgo == true {
				codea.Cgo = true
			}
			codea.Path = codeb.Path
			codea.Funcs.List = append(codea.Funcs.List, codeb.Funcs.List...)
			codea.Structs.List = append(codea.Structs.List, codeb.Structs.List...)
			codea.Interfaces.List = append(codea.Interfaces.List, codeb.Interfaces.List...)
			codea.Types.List = append(codea.Types.List, codeb.Types.List...)
			codea.Vars.List = append(codea.Vars.List, codeb.Vars.List...)
		}
		gclist = append(gclist, codea)
	}

	for _, codea = range gclist {

		bubblesort(codea.Funcs.Count(),
			func(i, j int) bool {
				gfa := codea.Funcs.I(i)
				gfb := codea.Funcs.I(j)
				if gfa.Name <= gfb.Name {
					return false
				}
				codea.Funcs.Pi(i, gfb)
				codea.Funcs.Pi(j, gfa)
				return true
			})

		bubblesort(codea.Structs.Count(),
			func(i, j int) bool {
				gfa := codea.Structs.I(i)
				gfb := codea.Structs.I(j)
				if gfa.Name <= gfb.Name {
					return false
				}
				codea.Structs.Pi(i, gfb)
				codea.Structs.Pi(j, gfa)
				return true
			})

		bubblesort(codea.Interfaces.Count(),
			func(i, j int) bool {
				gfa := codea.Interfaces.Idx(i)
				gfb := codea.Interfaces.Idx(j)
				if gfa.Name <= gfb.Name {
					return false
				}
				codea.Interfaces.Pidx(i, gfb)
				codea.Interfaces.Pidx(j, gfa)
				return true
			})

		bubblesort(codea.Types.Count(),
			func(i, j int) bool {
				gfa := codea.Types.I(i)
				gfb := codea.Types.I(j)
				if gfa.Name <= gfb.Name {
					return false
				}
				codea.Types.Pi(i, gfb)
				codea.Types.Pi(j, gfa)
				return true
			})

	}
	return gclist
}

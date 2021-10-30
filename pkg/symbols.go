package gocodoc

import (
	"strings"
)

//package
//var
//struct
//interface
//func
//const
//import
//import "C"

type Tpack struct {
	Name  string
	Codes Tcodes
}

func (g *Tpack) Init() {
	g.Codes = Tcodes{}
	g.Codes.Init()
}

type Tpacks struct {
	idx  int
	P    *Tpack
	list []*Tpack
}

type Ppack *Tpack

func (g *Tpacks) Init() {
	g.list = []*Tpack{}
	g.idx = -1
	g.P = nil
}

func (g *Tpacks) Find(name string) *Tpack {
	var item *Tpack
	if len(g.list) == 0 {
		return nil
	}
	for _, item = range g.list {
		if item.Name != name {
			continue
		}
		return item
	}
	return nil
}

func (g *Tpacks) Get(name string) *Tpack {
	item := g.Find(name)
	if item != nil {
		return item
	}
	item = &Tpack{Name: name}
	item.Init()
	g.list = append(g.list, item)
	return item
}

func (g *Tpacks) Reset() *Tpack {
	g.idx = -1
	if len(g.list) == 0 {
		return nil
	}
	return g.list[0]
}

func (g *Tpacks) Next() bool {
	if g.idx < -1 {
		g.idx = -1
	}
	g.idx++
	if g.idx >= len(g.list) {
		return false
	}
	g.P = g.list[g.idx]
	return true
}

func (g *Tpacks) PackCount() int {
	var packcount int = 0
	_ = g.Reset()
	for g.Next() {
		if g.P.Codes.Count() == 0 {
			continue
		}
		packcount++
	}
	return packcount
}

type Tmarkup struct {
	Comment     string
	Precomments []string
	Tags        []string
}

func (g *Tmarkup) Init() {
	g.Precomments = []string{}
	g.Tags = []string{}
	g.Comment = ""
}

type Tcode struct {
	Filename   string
	Path       string //We need to see if there are packages that cross paths
	Cgo        bool
	Vars       Tvars
	Types      Tvars
	Funcs      Tfuncs
	Structs    Tstructures
	Interfaces Tinterfaces
	Consts     Tconsts
	Markup     Tmarkup
}

func (g *Tcode) Init() {
	g.Vars = Tvars{}
	g.Vars.Init()
	g.Types = Tvars{}
	g.Types.Init()
	g.Funcs = Tfuncs{}
	g.Funcs.Init()
	g.Structs = Tstructures{}
	g.Structs.Init()
	g.Interfaces = Tinterfaces{}
	g.Interfaces.Init()
	g.Markup = Tmarkup{}
	g.Markup.Init()
}

type Tcodes struct {
	idx  int
	list []*Tcode
}

func (g *Tcodes) Init() {
	g.list = []*Tcode{}
}

func (g *Tcodes) Count() int {
	return len(g.list)
}

func (g *Tcodes) Find(path string, filename string) *Tcode {
	var item *Tcode
	if len(g.list) == 0 {
		return nil
	}
	for _, item = range g.list {
		if item.Path != path {
			continue
		}
		if item.Filename != filename {
			continue
		}
		return item
	}
	return nil
}

func (g *Tcodes) Get(path string, filename string) *Tcode {

	item := g.Find(path, filename)
	if item != nil {
		return item
	}
	item = &Tcode{Filename: filename, Path: path}
	item.Init()
	g.list = append(g.list, item)
	return item
}

func (g *Tcodes) Reset() *Tcode {
	g.idx = -1
	if len(g.list) == 0 {
		return nil
	}
	return g.list[0]
}

func (g *Tcodes) Next() *Tcode {
	if g.idx < -1 {
		g.idx = -1
	}
	g.idx++
	if len(g.list) <= g.idx {
		return nil
	}
	return g.list[g.idx]
}

type Tvar struct {
	Public bool
	Name   string
	Line   int
	Type   string
	Markup Tmarkup
}

func (g *Tvar) Init() {
	g.Markup = Tmarkup{}
	g.Markup.Init()
}

type Tvars struct {
	idx  int
	list []*Tvar
}

func (g *Tvars) Init() {
	g.list = []*Tvar{}
}

func (g *Tvars) Find(name string) *Tvar {
	var item *Tvar
	if len(g.list) == 0 {
		return nil
	}
	for _, item = range g.list {
		if item.Name == "" {
			//Return values usually do not have names
			continue
		}
		if item.Name != name {
			continue
		}
		return item
	}
	return nil
}

func (g *Tvars) Get(name string) *Tvar {
	item := g.Find(name)
	if item != nil {
		return item
	}
	item = &Tvar{Name: name}
	item.Init()
	if name[0:0] == strings.ToUpper(name[0:0]) {
		item.Public = true
	}
	g.list = append(g.list, item)
	return item
}

func (g *Tvars) Add() *Tvar {
	item := &Tvar{}
	item.Init()
	item.Public = true
	g.list = append(g.list, item)
	return item
}

func (g *Tvars) Reset() *Tvar {
	g.idx = -1
	if len(g.list) == 0 {
		return nil
	}
	return g.list[0]
}

func (g *Tvars) Next() *Tvar {
	if g.idx < -1 {
		g.idx = -1
	}
	(g.idx)++
	if len(g.list) <= g.idx {
		return nil
	}
	return g.list[g.idx]
}

type Tfunc struct {
	Public  bool
	Name    string
	Line    int
	Parms   Tvars
	Returns Tvars
	Markup  Tmarkup
}

func (g *Tfunc) Init() {
	g.Parms = Tvars{}
	g.Parms.Init()
	g.Markup = Tmarkup{}
	g.Markup.Init()
}

type Tfuncs struct {
	idx  int
	list []*Tfunc
}

func (g *Tfuncs) Init() {
	g.list = []*Tfunc{}
}

func (g *Tfuncs) Find(name string) *Tfunc {
	var item *Tfunc
	if len(g.list) == 0 {
		return nil
	}
	for _, item = range g.list {
		if item.Name != name {
			continue
		}
		return item
	}
	return nil
}

func (g *Tfuncs) Get(name string) *Tfunc {
	item := g.Find(name)
	if item != nil {
		return item
	}
	item = &Tfunc{Name: name}
	item.Init()
	if name[0:0] == strings.ToUpper(name[0:0]) {
		item.Public = true
	}

	g.list = append(g.list, item)
	return item
}

func (g *Tfuncs) Next() *Tfunc {
	if g.idx < -1 {
		g.idx = -1
	}
	(g.idx)++
	if len(g.list) <= g.idx {
		return nil
	}
	return g.list[g.idx]
}

type Tstru struct {
	Public bool
	Name   string
	Line   int
	Vars   Tvars
	Funcs  Tfuncs
	Markup Tmarkup
}

func (g *Tstru) Init() {
	g.Vars = Tvars{}
	g.Vars.Init()
	g.Funcs = Tfuncs{}
	g.Funcs.Init()
	g.Markup = Tmarkup{}
	g.Markup.Init()
}

type Tstructures struct {
	idx  int
	list []*Tstru
}

func (g *Tstructures) Init() {
	g.list = []*Tstru{}
}

func (g *Tstructures) Find(name string) *Tstru {
	var item *Tstru
	if len(g.list) == 0 {
		return nil
	}
	for _, item = range g.list {
		if item.Name != name {
			continue
		}
		return item
	}
	return nil
}

func (g *Tstructures) Get(name string) *Tstru {
	item := g.Find(name)
	if item != nil {
		return item
	}
	item = &Tstru{Name: name}
	item.Init()
	g.list = append(g.list, item)
	return item
}

func (g *Tstructures) Next() *Tstru {
	if g.idx < -1 {
		g.idx = -1
	}
	(g.idx)++
	if len(g.list) <= g.idx {
		return nil
	}
	return g.list[g.idx]
}

type Tinterface struct {
	Public bool
	Name   string
	Line   int
	Funcs  Tfuncs
	Markup Tmarkup
}

func (g *Tinterface) Init() {
	g.Funcs = Tfuncs{}
	g.Funcs.Init()
	g.Markup = Tmarkup{}
	g.Markup.Init()
}

type Tinterfaces struct {
	idx  int
	list []*Tinterface
}

func (g *Tinterfaces) Init() {
	g.list = []*Tinterface{}
}

func (g *Tinterfaces) Find(name string) *Tinterface {
	var item *Tinterface
	if len(g.list) == 0 {
		return nil
	}
	for _, item = range g.list {
		if item.Name != name {
			continue
		}
		return item
	}
	return nil
}

func (g *Tinterfaces) Get(name string) *Tinterface {
	item := g.Find(name)
	if item != nil {
		return item
	}
	item = &Tinterface{Name: name}
	item.Init()
	if name[0:0] == strings.ToUpper(name[0:0]) {
		item.Public = true
	}

	g.list = append(g.list, item)
	return item
}

func (g *Tinterfaces) Next() *Tinterface {
	if g.idx < -1 {
		g.idx = -1
	}
	(g.idx)++
	if len(g.list) <= g.idx {
		return nil
	}
	return g.list[g.idx]
}

type Tconst struct {
	Line     int
	Items    []string
	Comments []string
	Public   []bool
	Markup   Tmarkup
}

func (g *Tconst) Init() {
	g.Items = []string{}
	g.Comments = []string{}
	g.Public = []bool{}
	g.Markup = Tmarkup{}
	g.Markup.Init()
}

func (g *Tconst) Add(item string, comment string) {
	if len(item) == 0 {
		return
	}
	g.Items = append(g.Items, item)
	g.Comments = append(g.Comments, comment)
	if item[0:1] == strings.ToUpper(item[0:1]) {
		g.Public = append(g.Public, true)
	} else {
		g.Public = append(g.Public, false)
	}
}

type Tconsts struct {
	idx  int
	list []*Tconst
}

func (g *Tconsts) Init() {
	g.list = []*Tconst{}
}

func (g *Tconsts) Add() *Tconst {
	gc := &Tconst{}
	gc.Init()
	g.list = append(g.list, gc)
	return gc
}

func (g *Tconsts) Next() *Tconst {
	if g.idx < -1 {
		g.idx = -1
	}
	(g.idx)++
	if len(g.list) <= g.idx {
		return nil
	}
	return g.list[g.idx]
}

func checktype(stype string) bool {
	switch stype {
	case "struct":
		break
	case "interface":
		break
	case "interface{}":
		break
	case "int":
		break
	case "int6":
		break
	case "int16":
		break
	case "int32":
		break
	case "int64":
		break
	case "uint":
		break
	case "uint6":
		break
	case "uint16":
		break
	case "uint32":
		break
	case "uint64":
		break
	case "byte":
		break
	case "rune":
		break
	case "float32":
		break
	case "float64":
		break
	case "bool":
		break
	case "complex64":
		break
	case "complex128":
		break
	default:
		return false
	}
	return true
}

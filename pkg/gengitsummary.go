package gocodoc

import (
	"os"
)

func gengitsummary(g *os.File, gp *Tpack) error {

	w(g, "### Package:"+gp.Name)

	var gc *Tcode
	var gcs *Tcodes
	gcodesarr := gp.SplitByPath()
	for _, gcs = range gcodesarr {
		for _, gc = range gcs.List {
			gengitsumcode(g, gp, gc)
		}
	}
	return nil
}

func gengitsumcode(g *os.File, gp *Tpack, gc *Tcode) error {

	var gv *Tvar
	var gf *Tfunc
	var gs *Tstru
	var gi *Tinterface
	var gm *Tmarkup
	var gco *Tconst
	var idx int

	w(g, gc.Path)
	if gc.Cgo {
		w(g, "C Linkage notice (look at source)")
	}

	//Consts
	w(g, "#### Constants")
	gc.Consts.Reset()
	for gc.Consts.Next() {
		gco = gc.Consts.C
		gm = &gco.Markup
		wpre(g, gm)
		w(g, "<pre><code>")
		for idx, _ = range gco.Items {
			if gco.Public[idx] == false {
				continue
			}
			w(g, gco.Items[idx]+"      "+ws(gco.Comments[idx]))
		}
		w(g, "</code></pre>")
	}

	//types
	w(g, "#### Types")
	gc.Types.Reset()
	for gc.Types.Next() {
		gv = gc.Types.V
		gm = &gv.Markup
		wpre(g, gm)
		w(g, "<pre><code>")
		w(g, gv.Name+" "+gv.Type+" "+ws(gm.Comment))
		w(g, "</code></pre>")
	}

	//vars
	if gc.Vars.Count() > 0 {
		w(g, "#### Vars")
		gc.Vars.Reset()
		for gc.Vars.Next() {
			gv = gc.Vars.V
			gm = &gv.Markup
			wpre(g, gm)
			w(g, "<pre><code>")
			w(g, gv.Name+" "+gv.Type+" "+ws(gm.Comment))
		}
	}

	//interfaces
	if gc.Interfaces.Count() > 0 {
		we(g, "#### ", "Interfaces")
		gc.Interfaces.Reset()
		for gc.Interfaces.Next() {
			gi = gc.Interfaces.I
			gm = &gi.Markup
			wpre(g, gm)
			w(g, "<pre><code>")
			w(g, gi.Name)
			gi.Funcs.Reset()
			for gi.Funcs.Next() {
				gf = gi.Funcs.F
				wfunc(g, "- ", gf)
			}
			w(g, "</code></pre>")
		}
	}

	//structs
	if gc.Structs.Count() > 0 {
		w(g, "#### Structs")
		gc.Structs.Reset()
		for gc.Structs.Next() {
			gs = gc.Structs.S
			gm = &gs.Markup
			wpre(g, gm)
			w(g, "<pre><code>")
			w(g, gs.Name+" "+ws(gm.Comment))
			w(g, "</code></pre>")
		}
	}

	//funcs
	if gc.Funcs.Count() > 0 {
		w(g, "#### Funcs")
		gc.Funcs.Reset()
		w(g, "<pre><code>")
		for gc.Funcs.Next() {
			gf = gc.Funcs.F
			w(g, "Func:"+gf.Name)
		}
		w(g, "</code></pre>")
	}

	return nil

}

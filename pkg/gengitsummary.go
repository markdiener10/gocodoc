package gocodoc

func gengitsummary(gh *Tmarkdown, gp *Tpack) error {

	gh.wh(3, "Package:"+gp.Name)

	var gc *Tcode
	var gcs *Tcodes
	var cgo bool
	gcodesarr := gp.SplitByPath()

	for _, gcs = range gcodesarr {
		cgo = false
		for _, gc = range gcs.List {
			if gc.Cgo != true {
				continue
			}
			cgo = true
			break
		}
		for _, gc = range gcs.List {
			gengitsumcode(gh, gp, gc, cgo)
		}
	}
	return nil
}

func gengitsumcode(gh *Tmarkdown, gp *Tpack, gc *Tcode, cgo bool) error {

	var gv *Tvar
	var gf *Tfunc
	var gs *Tstru
	var gi *Tinterface
	var gm *Tmarkup
	var gco *Tconst
	var idx int

	gh.w(gc.Path)
	if cgo {
		gh.w("C Linkage notice (look at source)")
	}

	//Consts
	gh.wh(4, "Constants")
	gc.Consts.Reset()
	for gc.Consts.Next() {
		gco = gc.Consts.C
		gm = &gco.Markup
		gh.wpre(gm)
		for idx, _ = range gco.Items {
			if gco.Public[idx] == false {
				continue
			}
			gh.wcode(gco.Items[idx] + gh.wcomment("   ", gco.Comments[idx]))
		}
	}

	//types
	if gc.Types.Count() > 0 {
		gh.wh(4, "Types")
		gc.Types.Reset()
		for gc.Types.Next() {
			gv = gc.Types.V
			gm = &gv.Markup
			gh.wpre(gm)
			gh.wcode(gv.Name + gh.we("    ", gv.Type) + gh.wcomment("   ", gm.Comment))
		}
	}

	//vars
	if gc.Vars.Count() > 0 {
		gh.wh(4, "Variables")
		gc.Vars.Reset()
		for gc.Vars.Next() {
			gv = gc.Vars.V
			gm = &gv.Markup
			gh.wpre(gm)
			gh.wcode(gv.Name + gh.we("    ", gv.Type) + gh.wcomment("   ", gm.Comment))
		}
	}

	//interfaces
	if gc.Interfaces.Count() > 0 {
		gh.wh(4, "Interfaces")
		gc.Interfaces.Reset()
		for gc.Interfaces.Next() {
			gi = gc.Interfaces.I
			gm = &gi.Markup
			gh.wpre(gm)
			gh.w(gi.Name)
			gh.wcode(gi.Name + gh.wcomment("   ", gm.Comment))
			gi.Funcs.Reset()
			for gi.Funcs.Next() {
				gf = gi.Funcs.F
				gh.wfunc("    ", gf)
			}
		}
	}

	//structs
	if gc.Structs.Count() > 0 {
		gh.wh(4, "Structs")
		gc.Structs.Reset()
		for gc.Structs.Next() {
			gs = gc.Structs.S
			gm = &gs.Markup
			gh.wpre(gm)
			gh.wcode(gs.Name + gh.wcomment("    ", gm.Comment))
		}
	}

	//funcs
	if gc.Funcs.Count() > 0 {
		gh.wh(4, "Functions")
		gc.Funcs.Reset()
		for gc.Funcs.Next() {
			gf = gc.Funcs.F
			gm = &gf.Markup
			gh.wpre(gm)
			gh.wfunc(" ", gf)
		}
	}

	return nil

}

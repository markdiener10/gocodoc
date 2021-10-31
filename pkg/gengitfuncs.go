package gocodoc

import (
	"os"
	"strings"
)

//Now we want to generate a README.md as our generated documentation head.
//Then we want to count the number of packages in the parsed code

func w(g *os.File, out string) {
	g.WriteString(out + "  \r\n")
}

func wn(g *os.File, out string) {
	g.WriteString(out)
}

func ws(input string) string {
	input = strings.TrimSpace(input)
	if len(input) > 2 {
		if input[0:3] == "///" {
			return "" //Hidden comment
		}
	}
	return strings.ReplaceAll(input, "/", "")
}

func wpre(g *os.File, gm *Tmarkup) {
	if len(gm.Precomments) == 0 {
		return
	}
	//w(g, "") //Add blank line from previous entry
	for _, line := range gm.Precomments {
		w(g, "###### "+ws(line))
	}
}

func wfunc(g *os.File, pre string, gf *Tfunc) {
	w(g, pre+"Func:"+gf.Name)
	//for line := range gm.Precomments {
}

func we(g *os.File, pre string, out string) {
	if len(out) == 0 {
		return
	}
	g.WriteString(pre + out + "  \r\n")
}

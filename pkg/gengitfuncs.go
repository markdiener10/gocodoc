package gocodoc

import (
	"os"
	"strings"
)

//Now we want to generate a README.md as our generated documentation head.
//Then we want to count the number of packages in the parsed code
type Tmarkdown struct {
	file *os.File
}

func (g *Tmarkdown) init() {
	g.file = nil
}

func (g *Tmarkdown) open(outfile string) error {
	var err error
	g.file, err = os.OpenFile(outfile, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	return nil
}

func (g *Tmarkdown) close() {
	if g.file == nil {
		return
	}
	g.file.Close()
	g.file = nil
	return
}

func (g *Tmarkdown) w(out string) {
	g.file.WriteString(out + "\r\n")
}

func (g *Tmarkdown) wh(depth int, out string) {
	if depth > 0 {
		out = strings.Repeat("#", depth) + " " + out
	}
	g.w(out)
}

func (g *Tmarkdown) wline() {
	g.w("-------------")
}

func (g *Tmarkdown) we(pre string, out string) string {
	if len(out) == 0 {
		return ""
	}
	return pre + out
}

func (g *Tmarkdown) wcode(input string) {
	g.w("    " + input)
}

func (g *Tmarkdown) wcomment(pre string, input string) string {
	input = strings.TrimSpace(input)
	if len(input) > 2 {
		switch input[0:2] {
		case "/*":
			input = input[2:]
			input = strings.Replace(input, "*/", "", 1)
		case "//":
			if input[0:3] == "///" {
				return "" //Hidden comment
			}
			input = strings.ReplaceAll(input, "/", "")
		}
	}
	return pre + input
}

func (g *Tmarkdown) wpre(gm *Tmarkup) {
	if len(gm.Precomments) == 0 {
		return
	}
	for _, line := range gm.Precomments {
		if len(line) > 2 {
			if line[0:3] == "///" {
				continue
			}
		}
		g.w("> " + g.wcomment("", line))
		g.w("")
	}
}

func (g *Tmarkdown) wfunc(pre string, gf *Tfunc) {
	g.w(pre + "Func:" + gf.Name)
	//for line := range gf.parms, returns {
}

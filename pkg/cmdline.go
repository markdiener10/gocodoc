package gocodoc

import (
	"fmt"
	"os"
	"strings"
)

func fileexists(fname string) error {

	if _, err := os.Stat(fname); err != nil {
		return err
	}
	return nil
}

func Cmdlineprocess() (string, string, bool) {
	//Search for --src --dest
	fmt.Print("Gocodoc Documentation Utility")
	if len(os.Args) < 3 {
		fmt.Print("Missing Source And Destination Parameter")
		fmt.Print("Usage: gocodoc --src=PathtoGoCode --dest=PathToDocOutput")
		fmt.Print("")
		return "", "", false
	}

	var src string
	var dest string
	var pieces []string

	arg1 := strings.Trim(os.Args[1], " ")
	arg2 := strings.Trim(os.Args[2], " ")

	if !strings.Contains(arg1, "--src=") {
		if !strings.Contains(arg2, "--src=") {
			fmt.Print("Missing Source Parameter: --src=PathToGoCode")
			return "", "", false
		}
	}
	if !strings.Contains(arg1, "--dest=") {
		if !strings.Contains(arg2, "--dest=") {
			fmt.Print("Missing Destination Parameter: --dest=PathToDocOutput")
			return "", "", false
		}
	}
	if strings.Contains(arg1, "src=") {
		pieces = strings.Split(arg1, "=")
		src = pieces[1]
	}
	if strings.Contains(arg2, "src=") {
		pieces = strings.Split(arg2, "=")
		src = pieces[1]
	}
	if strings.Contains(arg1, "dest=") {
		pieces = strings.Split(arg1, "=")
		dest = pieces[1]
	}
	if strings.Contains(arg2, "dest=") {
		pieces = strings.Split(arg2, "=")
		dest = pieces[1]
	}
	src = strings.TrimSpace(src)
	dest = strings.TrimSpace(dest)
	if src[len(src)-1:] != "/" {
		src += "/"
	}
	if dest[len(dest)-1:] != "/" {
		dest += "/"
	}
	return src, dest, true
}

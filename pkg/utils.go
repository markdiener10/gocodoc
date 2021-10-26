package gocodoc

import "strings"

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

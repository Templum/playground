package text

import (
	"golang.org/x/text/language"
)

func Testcase() {
	value := language.MustParse("de-De")
	println(value.String())
}

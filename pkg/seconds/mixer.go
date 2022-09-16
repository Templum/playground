package seconds

import (
	"github.com/tidwall/gjson"
	"golang.org/x/text/language"
)

const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

func Testcase() {

	value := language.MustParse("de-De")
	println(value.String())

	val := gjson.Get(json, "name.last")
	println(val.String())
}

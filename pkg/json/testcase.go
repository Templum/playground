package json

import (
	"github.com/tidwall/gjson"
)

const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`


func Testcase() {
	value := gjson.Get(json, "name.last")
	println(value.String())
}
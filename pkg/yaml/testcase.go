package yaml

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

var data = `
a: Easy!
b:
  c: 2
  d: [3, 4]
`

// Note: struct fields must be public in order for unmarshal to
// correctly populate the data.
type T struct {
	A string
	B struct {
		RenamedC int   `yaml:"c"`
		D        []int `yaml:",flow"`
	}
}

func Testcase() {
	t := T{}
	_ = yaml.Unmarshal([]byte(data), &t)
	fmt.Printf("--- t:\n%v\n\n", t)
}

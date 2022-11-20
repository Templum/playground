package main

import (
	"github.com/Templum/playground/pkg/json"
	"github.com/Templum/playground/pkg/seconds"
	"github.com/Templum/playground/pkg/text"
	"github.com/Templum/playground/pkg/yaml"
	privatelib "github.com/Templum/private-lib"
	"golang.org/x/text/language"
)

func main() {
	json.Testcase()
	text.Testcase()
	yaml.Testcase()
	seconds.Testcase()

	supportedLanguages := []language.Tag{language.Albanian, language.Afrikaans, language.German}
	println(supportedLanguages)

	println(privatelib.NeedToSetupPrivate())
}

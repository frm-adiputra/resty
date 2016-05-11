package main

import (
	"html/template"
	"regexp"
	"strings"

	"github.com/frm-adiputra/resty/assets"
)

var (
	reDencoPath = regexp.MustCompile(`\{(.*?)\}`)
	tmpl        = template.New("root")
)

func init() {
	tmpl.Funcs(template.FuncMap{
		"dencoPath":            tmplDencoPath,
		"upperCase":            strings.ToUpper,
		"upperCaseFirstLetter": tmplUpperCaseFirstLetter,
	})

	a, err := assets.AssetDir("templates")
	if err != nil {
		panic(err)
	}

	for _, n := range a {
		t := tmpl.New(n)
		_, err = t.Parse(string(assets.MustAsset("templates/" + n)))
		if err != nil {
			panic(err)
		}
	}
}

func tmplDencoPath(s string) string {
	return reDencoPath.ReplaceAllString(s, ":${1}")
}

func tmplUpperCaseFirstLetter(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[0:1]) + s[1:]
}

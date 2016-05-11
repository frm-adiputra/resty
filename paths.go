package main

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/go-openapi/spec"
)

type pathData struct {
	Path string
	spec spec.PathItem
}

func allPathData(spec *spec.Swagger) []pathData {
	m := spec.Paths.Paths
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var a []pathData
	for _, k := range keys {
		a = append(a, newPathDataFromSpec(k, m[k]))
	}
	return a
}

func newPathDataFromSpec(path string, spec spec.PathItem) pathData {
	p := pathData{path, spec}
	return p
}

func (p pathData) OperationID(method string) (string, error) {
	v := reflect.ValueOf(p.spec).FieldByName(method)
	if !v.IsValid() {
		return "", fmt.Errorf("invalid method name: %s", method)
	}

	if v.IsNil() {
		return "", nil
	}

	op := v.Interface().(*spec.Operation)
	if op == nil {
		return "", nil
	}

	// if op.ID

	return op.ID, nil
}

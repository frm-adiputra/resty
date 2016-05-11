package main

import (
	"strings"
	"testing"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/loads/fmts"
	"github.com/go-openapi/spec"
	"github.com/stretchr/testify/assert"
)

var _swagDocs = make(map[string]*loads.Document)

func swagDocs(f string) *loads.Document {
	d, ok := _swagDocs[f]
	if ok {
		return d
	}

	r, err := fmts.YAMLDoc(f)
	if err != nil {
		panic(err)
	}

	d, err = loads.Analyzed(r, "")
	if err != nil {
		panic(err)
	}

	s := d.Spec()
	err = spec.ExpandSpec(s)
	if err != nil {
		panic(err)
	}

	_swagDocs[f] = d
	return d
}

func TestUniqueAndRequiredOperationIDDuplicate(t *testing.T) {
	a := assert.New(t)

	d := swagDocs("_fixtures/validate_operationID_duplicate.yaml")
	err := doc(*d).UniqueAndRequiredOperationID()
	a.Error(err)
	a.True(strings.HasPrefix(err.Error(), "duplicate operationId:"), "expected error: %s", err)
}

func TestUniqueAndRequiredOperationIDRequired(t *testing.T) {
	a := assert.New(t)

	d := swagDocs("_fixtures/validate_operationID_required.yaml")
	err := doc(*d).UniqueAndRequiredOperationID()
	a.Error(err)
	a.True(strings.HasPrefix(err.Error(), "required operationId:"), "expected error: %s", err)
}

func TestUniqueAndRequiredOperationIDValid(t *testing.T) {
	a := assert.New(t)

	d := swagDocs("_fixtures/validate_operationID_valid.yaml")
	err := doc(*d).UniqueAndRequiredOperationID()
	if !a.NoError(err) {
		t.FailNow()
	}
}

func TestPathOperationID(t *testing.T) {
	a := assert.New(t)

	d := swagDocs("_fixtures/validate_operationID_valid.yaml")
	p := "/pets"
	pd := newPathDataFromSpec(p, d.Spec().Paths.Paths[p])

	s, err := pd.OperationID("Post")
	if !a.NoError(err) {
		t.FailNow()
	}
	a.Equal("addPet", s)
}

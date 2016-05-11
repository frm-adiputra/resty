package main

import (
	"testing"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/loads/fmts"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	a := assert.New(t)

	r, err := fmts.YAMLDoc("swagger.yaml")
	if !a.NoError(err) {
		t.FailNow()
	}

	d, err := loads.Analyzed(r, "")
	if !a.NoError(err) {
		t.FailNow()
	}

	// Validate spec
	v := strfmt.NewFormats()
	err = validate.Spec(d, v)
	if !a.NoError(err) {
		t.FailNow()
	}

	s := d.Spec()
	err = spec.ExpandSpec(s)
	if !a.NoError(err) {
		t.FailNow()
	}

	a.Equal("getProducts", s.Paths.Paths["/products"].Get.ID)
	a.Equal("double", s.Paths.Paths["/products"].Get.Parameters[0].TypeName())
	a.Equal("getProducts", s.Paths.Paths["/products"].Get.ID)
	a.Nil(s.Paths.Paths["/products"].Post)
}

// spec.PathItem{
//         Refable:spec.Refable{
//                 Ref:spec.Ref{
//                         Ref:jsonreference.Ref{
//                                 referenceURL:(*url.URL)(nil),
//                                 referencePointer:jsonpointer.Pointer{
//                                         referenceTokens:[]string(nil)
//                                 },
//                                 HasFullURL:false,
//                                 HasURLPathOnly:false,
//                                 HasFragmentOnly:false,
//                                 HasFileScheme:false,
//                                 HasFullFilePath:false
//                         }
//                 }
//         },
//         VendorExtensible:spec.VendorExtensible{
//                 Extensions:spec.Extensions(nil)
//         },
//         pathItemProps:spec.pathItemProps{
//                 Get:(*spec.Operation)(0xc8202760e0),
//                 Put:(*spec.Operation)(nil),
//                 Post:(*spec.Operation)(nil),
//                 Delete:(*spec.Operation)(nil),
//                 Options:(*spec.Operation)(nil),
//                 Head:(*spec.Operation)(nil),
//                 Patch:(*spec.Operation)(nil), Parameters:[]spec.Parameter(nil)}}

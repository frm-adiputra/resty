package main

import (
	"bytes"
	"strings"
	"testing"

	"golang.org/x/tools/imports"

	"github.com/stretchr/testify/assert"
)

const _expectedDeclareHandler = `// Handlers variable contains all handler for this API.
// Fields int the Handlers sruct format are generated from ` + "`" + `operationId` + "`" + `.
var Handlers = struct {

	// pet
	AddPet            mwchain.Handler
	DeletePet         mwchain.Handler
	FindPetsByStatus  mwchain.Handler
	FindPetsByTags    mwchain.Handler
	GetPetById        mwchain.Handler
	UpdatePet         mwchain.Handler
	UpdatePetWithForm mwchain.Handler

	// store
	DeleteOrder  mwchain.Handler
	GetOrderById mwchain.Handler
	PlaceOrder   mwchain.Handler

	// user
	CreateUser                mwchain.Handler
	CreateUsersWithArrayInput mwchain.Handler
	CreateUsersWithListInput  mwchain.Handler
	DeleteUser                mwchain.Handler
	GetUserByName             mwchain.Handler
	LoginUser                 mwchain.Handler
	LogoutUser                mwchain.Handler
	UpdateUser                mwchain.Handler
}{}`

const _expectedHandlerFuncDecl = `// Handler returns the http handler for this API.
func Handler(prefix string) (http.Handler, error) {
	mux := denco.NewMux()
	handler, err := mux.Build([]denco.Handler{
		
		// pet
		mux.PUT(prefix+"/pets", toDenco(Handlers.UpdatePet)),
		mux.POST(prefix+"/pets", toDenco(Handlers.AddPet)),
		mux.GET(prefix+"/pets/findByStatus", toDenco(Handlers.FindPetsByStatus)),
		mux.GET(prefix+"/pets/findByTags", toDenco(Handlers.FindPetsByTags)),
		mux.GET(prefix+"/pets/:petId", toDenco(Handlers.GetPetById)),
		mux.POST(prefix+"/pets/:petId", toDenco(Handlers.UpdatePetWithForm)),
		mux.DELETE(prefix+"/pets/:petId", toDenco(Handlers.DeletePet)),

		// store
		mux.POST(prefix+"/stores/order", toDenco(Handlers.PlaceOrder)),
		mux.GET(prefix+"/stores/order/:orderId", toDenco(Handlers.GetOrderById)),
		mux.DELETE(prefix+"/stores/order/:orderId", toDenco(Handlers.DeleteOrder)),

		// user
		mux.POST(prefix+"/users", toDenco(Handlers.CreateUser)),
		mux.POST(prefix+"/users/createWithArray", toDenco(Handlers.CreateUsersWithArrayInput)),
		mux.POST(prefix+"/users/createWithList", toDenco(Handlers.CreateUsersWithListInput)),
		mux.GET(prefix+"/users/login", toDenco(Handlers.LoginUser)),
		mux.GET(prefix+"/users/logout", toDenco(Handlers.LogoutUser)),
		mux.GET(prefix+"/users/:username", toDenco(Handlers.GetUserByName)),
		mux.PUT(prefix+"/users/:username", toDenco(Handlers.UpdateUser)),
		mux.DELETE(prefix+"/users/:username", toDenco(Handlers.DeleteUser)),
	})
	if err != nil {
		return nil, err
	}

	return handler, nil
}`

func _execTmpl(name, expected string, sd *specData) (string, string, error) {
	var b bytes.Buffer
	err := tmpl.ExecuteTemplate(&b, name, sd)
	if err != nil {
		return "", "", err
	}

	opt := imports.Options{Fragment: true, FormatOnly: true, Comments: true, TabIndent: true}
	_actual, err := imports.Process("", b.Bytes(), &opt)
	if err != nil {
		return "", "", err
	}

	_expected, err := imports.Process("", []byte(expected), &opt)
	if err != nil {
		return "", "", err
	}

	return strings.TrimSpace(string(_expected)), strings.TrimSpace(string(_actual)), nil
}

func TestTmplFuncDencoPath(t *testing.T) {
	a := assert.New(t)

	var fixtures = []struct {
		Input    string
		Expected string
	}{
		{"/{name}", "/:name"},
		{"/{name}/{id}", "/:name/:id"},
		{"{name}/{id}", ":name/:id"},
	}

	for _, fx := range fixtures {
		actual := tmplDencoPath(fx.Input)
		a.Equal(fx.Expected, actual)
	}
}

func TestTmplFuncUpperCaseFirstLetter(t *testing.T) {
	a := assert.New(t)

	var fixtures = []struct {
		Input    string
		Expected string
	}{
		{"", ""},
		{".", "."},
		{"addPet", "AddPet"},
		{"12level", "12level"},
		{"v1.2.3", "V1.2.3"},
	}

	for _, fx := range fixtures {
		actual := tmplUpperCaseFirstLetter(fx.Input)
		a.Equal(fx.Expected, actual)
	}
}

func TestTmpl_DECLARE_HANDLER(t *testing.T) {
	a := assert.New(t)

	d := swagDocs("_fixtures/petstore.yaml")
	sd, err := newSpecData(d)
	if !a.NoError(err) {
		t.FailNow()
	}

	expected, actual, err := _execTmpl("DECLARE_HANDLER", _expectedDeclareHandler, sd)
	if !a.NoError(err) {
		t.FailNow()
	}

	if !a.Equal(expected, actual) {
		t.Logf("expected:\n%s\n\nactual:\n%s", expected, actual)
	}
}

func TestTmpl_HANDLER_FUNC_DECL(t *testing.T) {
	a := assert.New(t)

	d := swagDocs("_fixtures/petstore.yaml")
	sd, err := newSpecData(d)
	if !a.NoError(err) {
		t.FailNow()
	}

	expected, actual, err := _execTmpl("HANDLER_FUNC_DECL", _expectedHandlerFuncDecl, sd)
	if !a.NoError(err) {
		t.FailNow()
	}

	if !a.Equal(expected, actual) {
		t.Logf("expected:\n%s\n\nactual:\n%s", expected, actual)
	}
}

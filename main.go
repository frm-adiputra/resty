package main

//go:generate go-bindata -o assets/assets.go -pkg assets -prefix _assets _assets/...

func main() {
	// r, err := fmts.YAMLDoc("swagger.yaml")
	// if err != nil {
	// 	panic(err)
	// }
	//
	// d, err := loads.Analyzed(r, "")
	// if err != nil {
	// 	panic(err)
	// }
}

package main

import (
	"github.com/kalifun/protoc-gen-goutput/internal/generator"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	g := generator.Goutput{}
	protogen.Options{}.Run(g.Generate)
}

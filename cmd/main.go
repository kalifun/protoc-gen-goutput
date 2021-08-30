package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kalifun/protoc-gen-goutput/generator"
	_ "github.com/kalifun/protoc-gen-goutput/plugin"
	"google.golang.org/protobuf/proto"
)

func main() {
	gen := generator.New()

	// data, err := ioutil.ReadAll(os.Stdin)
	filename := "hello.proto"
	f, _ := os.Open(filename)
	data, err := ioutil.ReadAll(f)
	if err != nil {
		gen.Error(err, "read file fail")
	}
	fmt.Println(string(data))
	if err := proto.Unmarshal(data, gen.Request); err != nil {
		gen.Error(err, "unmarshal failed")
	}

	if len(gen.Request.FileToGenerate) == 0 {
		gen.Fail("no files to generate")
	}

	gen.CommandLineParameters(gen.Request.GetParameter())

	gen.WrapTypes()
	gen.SetPackageNames()
	gen.BuildTypeNameMap()
	gen.GenerateAllFiles()

	data, err = proto.Marshal(gen.Response)
	if err != nil {
		gen.Error(err, "marshal failed")
	}

	_, err = os.Stdout.Write(data)
	if err != nil {
		gen.Error(err, "failed to write")
	}

}

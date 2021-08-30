package plugin

import "github.com/kalifun/protoc-gen-goutput/generator"

type plugin struct {
	*generator.Generator
}

func NewPlugin() generator.Plugin {
	return &plugin{}
}

// Name identifies the plugin.
func (p *plugin) Name() string {
	return "goutput"
}

// Init is called once after data structures are built but before
// code generation begins.
func (p *plugin) Init(g *generator.Generator) {
	p.Generator = g
}

// Generate produces the code generated by the plugin for this file,
// except for the imports, by calling the generator's methods P, In, and Out.
func (p *plugin) Generate(file *generator.FileDescriptor) {
	panic("not implemented") // TODO: Implement
}

// GenerateImports produces the import declarations for this file.
// It is called after Generate.
func (p *plugin) GenerateImports(file *generator.FileDescriptor) {
	panic("not implemented") // TODO: Implement
}

func init() {
	generator.RegisterPlugin(NewPlugin())
}

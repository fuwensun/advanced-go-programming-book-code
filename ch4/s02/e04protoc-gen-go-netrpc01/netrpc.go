/*
1,名字改protoc-gen-go-netrpc01
2,移动到～/bin下
*/

package main

import (
	"github.com/golang/protobuf/protoc-gen-go/generator"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

type netrpcPlugin struct{ *generator.Generator }

func (p *netrpcPlugin) Name() string                { return "netrpc" }
func (p *netrpcPlugin) Init(g *generator.Generator) { p.Generator = g }

func (p *netrpcPlugin) GenerateImports(file *generator.FileDescriptor) {
	if len(file.Service) > 0 {
		p.genImportCode(file)
	}
}

func (p *netrpcPlugin) Generate(file *generator.FileDescriptor) {
	for _, svc := range file.Service {
		p.genServiceCode(svc)
	}
}


func (p *netrpcPlugin) genImportCode(file *generator.FileDescriptor) {
	p.P("// TODO: import code")
}

func (p *netrpcPlugin) genServiceCode(svc *descriptor.ServiceDescriptorProto) {
	p.P("// TODO: service code, Name = " + svc.GetName())
}

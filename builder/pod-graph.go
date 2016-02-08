package builder

import (
	"bytes"
	"github.com/n0rad/go-erlog/logs"
	"io/ioutil"
	"os"
)

func (p *Pod) Graph() {
	logs.WithF(p.fields).Info("Graphing")
	os.MkdirAll(p.target, 0777)

	var buffer bytes.Buffer
	buffer.WriteString("digraph {\n")
	buffer.WriteString("  {\n")
	buffer.WriteString("  ")
	buffer.WriteString("\"")
	buffer.WriteString(p.manifest.Name.ShortNameId())
	buffer.WriteString("\"")
	buffer.WriteString(" [style=filled, fillcolor=yellow, shape=box]\n")
	buffer.WriteString("  }\n")

	for _, e := range p.manifest.Pod.Apps {
		for _, d := range e.Dependencies {
			buffer.WriteString("  ")
			buffer.WriteString("\"")
			buffer.WriteString(d.ShortNameId())
			buffer.WriteString("\"")
			buffer.WriteString(" -> ")
			buffer.WriteString("\"")
			buffer.WriteString(p.manifest.Name.ShortNameId())
			buffer.WriteString("\"")
			buffer.WriteString("\n")
		}
	}

	buffer.WriteString("}\n")

	ioutil.WriteFile(p.target+"/graph.dot", buffer.Bytes(), 0644)

}
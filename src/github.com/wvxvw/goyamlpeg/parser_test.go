package goyamlpeg

import "testing"

func TestMinimalParser(t *testing.T) {
	yaml := &Yaml{
		Buffer: `%YAML 1.2
---
body
...`,
	}
	yaml.Init()
	if err := yaml.Parse(); err != nil {
		t.Fatalf("Parsing failed: %+v", err)
	}
	yaml.Execute()
	// yaml.PrintSyntax()
	yaml.PrintSyntaxTree()
}

func TestTwoMessages(t *testing.T) {
	yaml := &Yaml{
		Buffer: `
%YAML 1.2
---
body
...
%TAG ! foo
---
morebody
...
`,
	}
	yaml.Init()
	if err := yaml.Parse(); err != nil {
		t.Fatalf("Parsing failed: %+v", err)
	}
}

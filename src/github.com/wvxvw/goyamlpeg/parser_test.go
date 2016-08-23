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

func TestQuotedLiteral(t *testing.T) {
	yaml := &Yaml{
		Buffer: `
"simple quote"
`,
	}
	yaml.Init()
	if err := yaml.Parse(); err != nil {
		t.Fatalf("Parsing failed: %+v", err)
	}
	yaml.PrintSyntaxTree()
}

func TestAposLiteral(t *testing.T) {
	yaml := &Yaml{
		Buffer: `
'simple apostrophe'
`,
	}
	yaml.Init()
	if err := yaml.Parse(); err != nil {
		t.Fatalf("Parsing failed: %+v", err)
	}
	yaml.PrintSyntaxTree()
}

func TestQuotedCodepoints(t *testing.T) {
	yaml := &Yaml{
		Buffer: `
"simple \x36 quote \u123f"
`,
	}
	yaml.Init()
	if err := yaml.Parse(); err != nil {
		t.Fatalf("Parsing failed: %+v", err)
	}
	yaml.PrintSyntaxTree()
}

func TestQuotedMultiline(t *testing.T) {
	yaml := &Yaml{
		Buffer: `
"
first line
second line

third line
"
`,
	}
	yaml.Init()
	if err := yaml.Parse(); err != nil {
		t.Fatalf("Parsing failed: %+v", err)
	}
	yaml.PrintSyntaxTree()
}

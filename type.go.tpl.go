package main

var typeTpl = `import (
	"fmt"
)

{{ range $_, $enum := .Enums }}
{{ $TypeName := ExportName $enum.Id }}
type {{ $TypeName }} int32

func (v {{ $TypeName }}) String() string {
	switch v {
	{{ range $_, $value := $enum.Values }}
	case {{ $value.Value }}:
		return "{{ $value.Nick }}"
	{{ end }}
	}

	panic(fmt.Sprintf("should not reach here, unknown value %v", int(v)))
}

const (
	{{ range $_, $value := $enum.Values }}
	{{ $TypeName }}{{ ExportName $value.Nick }} {{ $TypeName }} = {{ $value.Value }}
	{{ end }}
)
{{ end }}

{{ range $_, $flags := .Flags }}
{{ $TypeName := ExportName $flags.Id }}
type {{ $TypeName }} uint32

func (v {{ $TypeName }}) String() string {
	switch v {
	{{ range $_, $value := $flags.Values }}
	case {{ $value.Value }}:
		return "{{ $value.Nick }}"
	{{ end }}
	}

	panic(fmt.Sprintf("should not reach here, unknown value %v", int(v)))
}

const {
	{{ range $_, $value := $flags.Values }}
	{{ $TypeName }}Flags{{ExportName $value.Nick}} {{ $TypeName }} = {{ $value.Value }}
	{{ end }}
}
{{ end }}
`

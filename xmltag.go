package main

// TODO: support choice, aliases.

type Value struct {
	Nick  string `xml:"nick,attr"`
	Value int    `xml:"value,attr"`
}

type Enum struct {
	Id     string  `xml:"id,attr"`
	Values []Value `xml:"value"`
}

func (e Enum) IsValid() bool {
	return e.Id != ""
}

type Flags struct {
	Id     string  `xml:"id,attr"`
	Values []Value `xml:"value"`
}

func (f Flags) IsValid() bool {
	return f.Id != ""
}

type Range struct {
	Min string `xml:"min,attr"`
	Max string `xml:"max,attr"`
}

type Key struct {
	Name        string `xml:"name,attr"`
	Type        string `xml:"type,attr"`
	Enum        string `xml:"enum,attr"`
	Flags       string `xml:"flags,attr"`
	Default     string `xml:"default"`
	Summary     string `xml:"summary"`
	Description string `xml:"description"`
	Range       Range
}

func (key Key) IsEnum() bool {
	return key.Enum != ""
}

func (key Key) IsFlags() bool {
	return key.Flags != ""
}

type Child struct {
	Name   string `xml:"name:attr"`
	Schema string `xml:"schema:attr"`
}

type Schema struct {
	Path     string  `xml:"path,attr"`
	Id       string  `xml:"id,attr"`
	Children []Child `xml:"child"`
	Keys     []Key   `xml:"key"`
}

type SchemaList struct {
	Schemas []Schema `xml:"schema"`
	Enums   []Enum   `xml:"enum"`
	Flags   []Flags  `xml:"flags"`
}

func (l SchemaList) FindEnum(id string) Enum {
	for _, enum := range l.Enums {
		if enum.Id == id {
			return enum
		}
	}
	return Enum{}
}

func (l *SchemaList) FindFlags(id string) Flags {
	for _, flags := range l.Flags {
		if flags.Id == id {
			return flags
		}
	}
	return Flags{}
}

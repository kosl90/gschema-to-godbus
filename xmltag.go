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

// FIXME: support different type range.
type Range struct {
	Min string `xml:"min,attr"`
	Max string `xml:"max,attr"`
}

type Choice struct {
	Value string `xml:"value,attr"`
}

type Choices struct {
	Choice []Choice `xml:"choice"`
}

type Alias struct {
	Value string `xml:"value,attr"`
}

type Aliases struct {
	Values []Alias `xml:"alias"`
}

type Default struct {
	L10n    string `xml:"l10n,attr"`
	Context string `xml:"context,attr"`
	Value   string `xml:",chardata"`
}

type Key struct {
	Name string `xml:"name,attr"` // can only contain lowercase letters, number and '-'

	// exactly one of type, enum or flags must be given
	Type  string `xml:"type,attr"`  // GVariant type string
	Enum  string `xml:"enum,attr"`  // id of an enum type that has been defined earlier
	Flags string `xml:"flags,attr"` // id of an flags type that has been defined earlier

	Default     Default `xml:"default"`
	Summary     string  `xml:"summary"`
	Description string  `xml:"description"`
	Range       Range   `xml:"range"`
	Choices     Choices `xml:"choices"`
	Aliases     Aliases `xml:"aliases"`
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

type Override struct {
	Name    string `xml:"name,attr"`
	L10n    string `xml:"l10n,attr"`
	Context string `xml:"context,attr"`
}

type Schema struct {
	Path          string `xml:"path,attr"`
	Id            string `xml:"id,attr"`
	GettextDomain string `xml:"gettext-domain,attr"`
	Extends       string `xml:"extends,attr"`
	ListOf        string `xml:list-of,attr"`

	Children  []Child    `xml:"child"`
	Keys      []Key      `xml:"key"`
	Overrides []Override `xml:"override"`
}

type SchemaList struct {
	Schemas []Schema `xml:"schema"`
	Enums   []Enum   `xml:"enum"`
	Flags   []Flags  `xml:"flags"` // cannot find flags in dtd for schemalist, but it works.

	GettextDomain string `xml:"gettext-domain,attr"`
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

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"text/template"
)

var funcMap = template.FuncMap{
	"ExportName":             exportName,
	"PkgName":                pkgName,
	"ToolVersion":            toolVersion,
	"MapType":                mapType,
	"MapTypeSetter":          mapTypeSetter,
	"MapTypeGetter":          mapTypeGetter,
	"TrimQuote":              trimQuote,
	"DBusName":               dbusName,
	"DBusPath":               dbusPath,
	"ConvertToDBusInterface": convertToDBusInterface,
	"GetKeyType":             getKeyType,
	"GetDefaultValue":        getDefaultValue,
	"Title":                  strings.Title,
	"Predefined":             predefined,
	"GetRangeType":           getRangeType,
	"String":                 toString,
}

func predefined() []string {
	return []string{"int16", "uint16", "int32", "uint32", "int64", "uint64", "double"}
}

func mustParse(tpl *template.Template, cont string) *template.Template {
	return template.Must(tpl.Parse(cont))
}

func normalize(id string) string {
	names := strings.Split(id, "-")
	l := len(names)
	normalizedNames := make([]string, l)
	for i, name := range names {
		normalizedNames[i] = strings.Title(name)
	}
	normalizedID := strings.Join(normalizedNames, "")
	return normalizedID
}

func exportName(id string) string {
	names := strings.Split(id, ".")
	return normalize(names[len(names)-1])
}

var (
	__VERSION__ = "v0.1.0"
)

func pkgName() string {
	return *_pkgName
}

func toolVersion() string {
	return __VERSION__
}

var TypeMap = map[string]string{
	"s":  "string",
	"o":  "dbus.ObjectPath",
	"g":  "string", // dbus.Signature
	"b":  "bool",
	"y":  "byte",
	"n":  "int16",
	"q":  "uint16",
	"i":  "int32",
	"u":  "uint32",
	"x":  "int64",
	"t":  "uint64",
	"h":  "int32",
	"d":  "double",
	"as": "[]string",
	"ao": "[]dbus.ObjectPath",
	"ag": "[]string", // dbus.Signature
	"ab": "[]bool",
	"ay": "[]byte",
	"an": "[]int16",
	"aq": "[]uint16",
	"ai": "[]int32",
	"au": "[]uint32",
	"ax": "[]int64",
	"at": "[]uint64",
	"ah": "[]int32",
	"ad": "[]double",
}

// TODO: support tuple and struct
var VariantMap = map[string]string{
	"s":  "String",
	"o":  "ObjectPath",
	"g":  "Signature", // dbus.Signature
	"b":  "Boolean",
	"y":  "Byte",
	"n":  "Int16",
	"q":  "Uint16",
	"i":  "Int32",
	"u":  "Uint32",
	"x":  "Int64",
	"t":  "Uint64",
	"h":  "Int32",
	"d":  "Double",
	"as": "Strv",
	// TODO: NOT support now.
	// "ao": "[]dbus.ObjectPath",
	// "ag": "[]string", // dbus.Signature
	// "ab": "[]bool",
	// "ay": "[]byte",
	// "an": "[]int16",
	// "aq": "[]uint16",
	// "ai": "[]int32",
	// "au": "[]uint32",
	// "ax": "[]int64",
	// "at": "[]uint64",
	// "ah": "[]int32",
	// "ad": "[]double",
}

func mapType(ty string) string {
	v, ok := TypeMap[ty]
	if ok {
		return v
	}
	return ""
}

func mapTypeGetter(ty string) string {
	v, ok := VariantMap[ty]
	if ok {
		return v
	} else {
		panic(fmt.Sprintf("not support %q now", ty))
	}
	return ""
}

func mapTypeSetter(ty string) string {
	v, ok := VariantMap[ty]
	if ok {
		return v
	} else {
		panic(fmt.Sprintf("not support %q now", ty))
	}
	return ""
}

func getKeyType(k Key) string {
	if k.IsEnum() {
		return "int32"
	} else if k.IsFlags() {
		return "uint32"
	}
	return mapType(k.Type)
}

func ReadFile(filename string) string {
	c, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(c)
}

func trimQuote(a string) string {
	return strings.Trim(a, "'")
}

func convertToDBusInterface(schemaID string) string {
	names := strings.Split(schemaID, "-")
	exportName := strings.Title(names[len(names)-1])
	names[len(names)-1] = exportName
	return strings.Join(names, "")
}

func dbusName() string {
	return *_dbusName
}

func dbusPath() string {
	return *_dbusPath
}

type Result struct {
	Value string
	Err   error
}

func getDefaultValue(schemas SchemaList, key Key) Result {
	if key.IsEnum() {
		enum := schemas.FindEnum(key.Enum)
		if enum.IsValid() {
			for _, value := range enum.Values {
				if value.Nick == trimQuote(key.Default.Value) {
					return Result{Value: exportName(enum.Id) + exportName(value.Nick)}
				}
			}
			return Result{Err: fmt.Errorf("invalid value %v", key.Default.Value)}
		}
		return Result{Err: fmt.Errorf("invalid enum id %q", key.Enum)}
	}
	return Result{Value: key.Default.Value}
}

func getRangeType(key Key) string {
	if key.IsEnum() {
		return "int32"
	} else if key.IsFlags() {
		return "uint32"
	}

	return "Range" + VariantMap[key.Type]
}

func toString(s []byte) string {
	fmt.Println(string(s))
	return string(s)
}

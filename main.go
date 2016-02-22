package main

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"fmt"
	"go/format"
	"os"
	"path"
	"sync"
	"text/template"
)

var (
	funcMap = template.FuncMap{
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
	}

	pref             = template.New("prefix header").Funcs(funcMap)
	typeTemplate     = template.New("types").Funcs(funcMap)
	settingsTemplate = template.New("settings").Funcs(funcMap)
)

const (
	typeFile     = "autogen_settings_types.go"
	settingsFile = "autogen_settings.go"
)

func writeToFileWithFormat(bytes []byte, f *os.File, formatSrc bool) {
	output := bufio.NewWriter(f)
	if formatSrc {
		src, err := format.Source(bytes)
		if err != nil {
			fmt.Println(err)
		}
		output.Write(src)
	} else {
		output.Write(bytes)
	}
	output.Flush()
}

func main() {
	parseCMD()

	var (
		typeOutputFile    *os.File
		settingOutputFile *os.File
	)

	err := Resolve(func() error {
		f, err := os.OpenFile(path.Join(*_outputDir, typeFile), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
		typeOutputFile = f
		return err
	}).Resolve(func() error {
		f, err := os.OpenFile(path.Join(*_outputDir, settingsFile), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
		settingOutputFile = f
		return err
	}).GetError()

	if err != nil {
		fmt.Println(err)
		return
	}

	defer typeOutputFile.Close()
	defer settingOutputFile.Close()

	typeOutput := bytes.NewBufferString("")
	settingOutput := bytes.NewBufferString("")
	{
		t := mustParse(pref, Prefix)
		t.Execute(typeOutput, "")
		t.Execute(settingOutput, "")
	}

	var v SchemaList
	xml.Unmarshal([]byte(ReadFile(*_schema)), &v)

	var WorkGroup sync.WaitGroup
	WorkGroup.Add(2)
	go func() {
		defer WorkGroup.Done()
		if err := mustParse(typeTemplate, typeTpl).Execute(typeOutput, v); err != nil {
			fmt.Println(err)
			return
		}
		writeToFileWithFormat(typeOutput.Bytes(), typeOutputFile, true)
	}()

	go func() {
		defer WorkGroup.Done()
		if err := mustParse(settingsTemplate, settingsTpl).Execute(settingOutput, v); err != nil {
			fmt.Println(err)
			return
		}
		writeToFileWithFormat(settingOutput.Bytes(), settingOutputFile, true)
	}()
	WorkGroup.Wait()
}

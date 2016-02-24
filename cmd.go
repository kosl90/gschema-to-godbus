package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	_schema    *string = kingpin.Flag("schema", "schema which is going to be translated.").Required().Short('s').String()
	_pkgName   *string = kingpin.Flag("pkg_name", "package name used for golang package, default is main.").Default("main").Short('n').String()
	_dbusName  *string = kingpin.Flag("dest", "dbus dest used for dbus.").Required().Short('d').String()
	_dbusPath  *string = kingpin.Flag("path", "dbus path used for dbus.").Required().Short('p').String()
	_outputDir *string = kingpin.Flag("output_dir", "output directory to save generated files, default is current directory.").Default(".").Short('o').ExistingDir()
)

func parseCMD() {
	kingpin.Parse()
}

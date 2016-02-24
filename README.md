# Intro

Generate dbus code for golang according to gschema file. A little convenient tool for working in [Deepin](https://github.com/linuxdeepin).

# Request

1. [kingpin](https://github.com/alecthomas/kingpin) for commandline parse.
2. [pkg.deepin.io/lib/dbus](https://github.com/linuxdeepin/go-lib/tree/develop/dbus)
3. [gir](https://github.com/linuxdeepin/go-gir-generator)

# Usage
```bash
$ gschema-to-godbus --help
usage: gschema-to-godbus --schema=SCHEMA --dest=DEST --path=PATH [<flags>]

Flags:
      --help               Show context-sensitive help (also try --help-long and --help-man).
  -s, --schema=SCHEMA    schema which is going to be translated.
  -n, --pkg_name="main"  package name used for golang package, default is main.
  -d, --dest=DEST        dbus dest used for dbus.
  -p, --path=PATH        dbus path used for dbus.
  -o, --output_dir=.     output directory to save generated files, default is current directory.
```

`autogen_gsettings.go` and `autogen_gsettings_type.go` will be generated, a setter, a getter and a changed signal will be generated for each key. If the **range** tag is specific, a `GetRangeOfX` method will be generated too, like `GetRangeOfZoomLevel`.

A function named `NewX` and a function named `NewXWithHook`, like `NewDesktopPreferences` and `NewDesktopPreferencesWithHook`, will be generated for each schema. The hook is nullable which should follow [SettingHook](settings.go.tpl.go#L11) interface, and if you want to implement your own hook, [DefaultSettingHook](settings.go.tpl.go#L18) can be embeded for convenient.

# References

1. [gschema.dtd](https://git.gnome.org/browse/glib/tree/gio/gschema.dtd).
2. [GVariant doc](https://developer.gnome.org/glib/stable/glib-GVariant.html)
3. [GVariant format strings](https://developer.gnome.org/glib/stable/gvariant-format-strings.html)
4. [GSettings doc](https://developer.gnome.org/gio/stable/GSettings.html)

# License
[MIT License](http://opensource.org/licenses/MIT)


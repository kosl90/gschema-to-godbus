package main

var settingsTpl = `import (
	"sync"

	"gir/gio-2.0"
	"gir/glib-2.0"
	"pkg.deepin.io/lib/dbus"
)

type SettingHook interface {
	finializeOnce sync.Once
	BeforeSettingFilter(key string, oldValue interface{}, newValue interface{}) bool
	OnChanged(key string)
}

type _DefaultHook struct {
}

func (_DefaultHook) BeforeSettingFilter(key string, oldValue interface{}, newValue interface{}) bool {
	return true
}

func (_DefaultHook) OnChanged(string) {
}

var _defaultHook = _DefaultHook{}

{{ $schemas := . }}
{{ range $_, $schema := .Schemas }}{{ if $schema.Keys }}{{ $TypeName := ExportName $schema.Id }}{{ if $schema.Keys }}
const (
{{ range $_, $key := $schema.Keys }}
	// {{ $key.Summary }}
	{{$result :=  GetDefaultValue $schemas $key }}{{ if $result.Err }}panic({{ $result.Err }}){{ else }}// default: {{ $result.Value }}{{ end }}
	{{$TypeName}}{{ ExportName $key.Name }} string = "{{ $key.Name }}"
{{ end }}
){{ end }}
{{/* generate setting structure */}}
type {{ $TypeName }} struct {
	settings *gio.Settings
	hook SettingHook

{{ range $_, $key := $schema.Keys }}{{ $PropName :=  ExportName $key.Name }}
	{{ $PropName }}Changed func({{ if $key.IsEnum }}int32{{ else }}{{ if $key.IsFlags }}uint32{{ else }}{{ MapType $key.Type }}{{ end }}{{ end }})
{{ end }}
}

func (s *{{ $TypeName}}) GetDBusInfo() dbus.DBusInfo {
	return dbus.DBusInfo{
		Dest:       "{{ DBusName }}",
		ObjectPath: "{{ DBusPath }}",
		Interface:  "{{ ConvertToDBusInterface $schema.Id }}",
	}
}

func New{{ $TypeName }}(hook SettingHook) *{{ $TypeName }} {
	if hook == nil {
		hook = _defaultHook
	}
	s := &{{ $TypeName }} {
		hook: hook,
		settings: gio.NewSettings("{{$schema.Id}}"),
	}
	s.listenSignal()
	return s
}

func (s *{{ $TypeName }}) Finialize() {
	s.finializeOnce.Do(func() {
		s.settings.Unref()
	})
}

func (s *{{ $TypeName }}) listenSignal() {
	s.settings.Connect("changed", func(_ *gio.Settings, key string){
		s.hook.OnChanged(key)
		switch key {
		{{ range $_, $key := $schema.Keys }}{{ $PropName :=  ExportName $key.Name }}
		case "{{ $key.Name }}":
			dbus.Emit(s, "{{ $PropName }}Changed", s.{{ $PropName }}())
		{{ end }}
		}
	})
}

{{ range $_, $key := $schema.Keys }}{{ $PropName :=  ExportName $key.Name }}
// {{ $PropName }} gets {{ $PropName }}'s value.
func (s *{{ $TypeName }}) {{$PropName}}() {{ GetKeyType $key }} {
	{{ if $key.IsEnum }}value := s.settings.GetEnum("{{$key.Name}}")
	{{ else }}{{ if $key.IsFlags }}value := s.settings.GetFlags("{{$key.Name}}")
	{{ else }}value := s.settings.GetValue("{{ $key.Name }}").Get{{ MapTypeGetter $key.Type }}()
	{{ end }}{{ end }}
	return value
}

// set{{ $PropName }} used internal.
func (s *{{ $TypeName }}) set{{ $PropName }}(newValue {{ GetKeyType $key }}) bool {
	oldValue := s.{{ $PropName }}()
	if oldValue == newValue {
		return false
	}

	gs := s.settings
	{{ if $key.IsEnum }}return gs.SetEnum("{{ $key.Name }}", newValue)
	{{ else }}{{ if $key.IsFlags }}return gs.SetFlags("{{ $key.Name }}", newValue)
	{{ else }}return gs.SetValue("{{ $key.Name }}", glib.NewVariant{{ MapTypeSetter $key.Type }}(newValue)){{ end }}{{ end }}
}

// Set{{ $PropName }} sets value of {{ $PropName }} and emit {{$PropName}}Changed signal.
func (s *{{ $TypeName }}) Set{{ $PropName }}(newValue {{ GetKeyType $key }}) {
	s.set{{$PropName}}(newValue)
	dbus.Emit(s, "{{ $PropName }}Changed", newValue)
}
{{ end }}
{{ end }}

{{ end }}

`

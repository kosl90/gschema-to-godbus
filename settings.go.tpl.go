package main

var settingsTpl = `import (
	"runtime"
	"sync"

	"gir/gio-2.0"
	"gir/glib-2.0"
	"pkg.deepin.io/lib/dbus"
)

type SettingHook interface {
	WillSet(gs *gio.Settings, key string, oldValue interface{}, newValue interface{}) bool // return true to continue setting.
	DidSet(gs *gio.Settings, key string, oldValue interface{}, newValue interface{})
	WillChange(gs *gio.Settings, key string) bool // return true to continue handleing change.
	DidChange(gs *gio.Settings, key string)
}

type DefaultSettingHook struct {
}

func (DefaultSettingHook) WillSet(gs *gio.Settings, key string, oldValue interface{}, newValue interface{}) bool {
	return true
}

func (DefaultSettingHook) DidSet(gs *gio.Settings, key string, oldValue interface{}, newValue interface{}) {
}

func (DefaultSettingHook) WillChange(*gio.Settings, string) bool {
	return true
}

func (DefaultSettingHook) DidChange(*gio.Settings, string) {
}

var _defaultHook = DefaultSettingHook{}

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
	finalizeOnce sync.Once
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

func New{{ $TypeName }}() *{{ $TypeName }} {
	return New{{ $TypeName }}WithHook(nil)
}

func New{{ $TypeName }}WithHook(hook SettingHook) *{{ $TypeName }} {
	if hook == nil {
		hook = _defaultHook
	}
	s := &{{ $TypeName }} {
		hook: hook,
		settings: gio.NewSettings("{{$schema.Id}}"),
	}
	s.listenSignal()
	runtime.SetFinalizer(s, func(o interface{}) {
		s := o.(*{{ $TypeName }})
		s.finalize()
	})
	return s
}

func (s *{{ $TypeName }}) finalize() {
	s.finalizeOnce.Do(func() {
		s.settings.Unref()
	})
}

func (s *{{ $TypeName }}) listenSignal() {
	s.settings.Connect("changed", func(gs *gio.Settings, key string){
		if !s.hook.WillChange(gs, key) {
			return
		}
		switch key {
		{{ range $_, $key := $schema.Keys }}{{ $PropName :=  ExportName $key.Name }}
		case "{{ $key.Name }}":
			dbus.Emit(s, "{{ $PropName }}Changed", s.{{ $PropName }}())
		{{ end }}
		}
		s.hook.DidChange(gs, key)
	})
	{{ $sample := index $schema.Keys 0 }}
	// make sure signal work
	// detail: https://github.com/GNOME/glib/commit/8ff5668a458344da22d30491e3ce726d861b3619
	s.{{ ExportName $sample.Name }}()
}

{{ range $_, $key := $schema.Keys }}{{ $PropName :=  ExportName $key.Name }}
{{/* not generated GetRangeOfX for "type", "enum", "flags" */}}
{{ if $key.Range.Min }}
{{ $rangeType := GetRangeType $key }}
// GetRangeOf{{ $PropName }} gets the value range of {{ $PropName }}.
func (s *{{ $TypeName }}) GetRangeOf{{ $PropName }}() {{ $rangeType }} {
	return {{ $rangeType }}{Min: {{ $key.Range.Min }}, Max: {{ $key.Range.Max }}}
}
{{ end }}

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
	if !s.hook.WillSet(gs, "{{ $key.Name }}", oldValue, newValue) {
		return false
	}
	defer s.hook.DidSet(gs, "{{ $key.Name }}", oldValue, newValue)

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

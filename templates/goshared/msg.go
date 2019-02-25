package goshared

const msgTpl = `
{{ if disabled . }}
	{{- cmt "ValidateFields is disabled for " (msgTyp .) ". This method will always return nil." -}}
{{ else }}
	{{- cmt "ValidateFields checks the field values on " (msgTyp .) " with the rules defined in the proto definition for this message. If any rules are violated, an error is returned." -}}
{{ end -}}
func (m {{ (msgTyp .).Pointer }}) ValidateFields(paths ...string) error {
	{{ if disabled . -}}
		return nil
	{{- else -}}
		{{ if .Fields -}}
			if m == nil { return nil }

			if len(paths) == 0 {
				paths = {{ msgTyp . }}FieldPathsNested
			}

			for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
				_ = subs
				switch name {
			{{ range .NonOneOfFields -}}
				case "{{ .Name }}":
					{{ render (context .) }}
			{{ end -}}

			{{ range .OneOfs -}}
				case "{{ .Name }}":
					if len(subs) == 0 {
						subs = []string{
							{{ range .Fields -}}
								"{{ .Name }}",
							{{- end }}
						}
					}
					for name, subs := range _processPaths(subs) {
						_ = subs
						switch name {
						{{ range .Fields -}}
							case "{{ .Name }}":
								{{ render (context .) }}
						{{ end -}}
						{{ if required . }}
							default:
								return {{ errname .Message }}{
									field:  {{ .Name }},
									reason: "value is required",
								}
						{{ end -}}
						}		
					}
			{{ end -}}
				default:
					return {{ errname . }} {
						field:  name,
						reason: "invalid field path",
					}
				}
			}
			return nil
		{{ else -}}
			if len(paths) > 0 {
				return fmt.Errorf("message {{ msgTyp . }} has no fields, but paths %s were specified", paths)
			}
			return nil
		{{ end -}}
	{{- end -}}
}

{{ if needs . "hostname" }}{{ template "hostname" . }}{{ end }}

{{ if needs . "email" }}{{ template "email" . }}{{ end }}

{{ cmt (errname .) " is the validation error returned by " (msgTyp .) ".ValidateFields if the designated constraints aren't met." -}}
type {{ errname . }} struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e {{ errname . }}) Field() string { return e.field }

// Reason function returns reason value.
func (e {{ errname . }}) Reason() string { return e.reason }

// Cause function returns cause value.
func (e {{ errname . }}) Cause() error { return e.cause }

// Key function returns key value.
func (e {{ errname . }}) Key() bool { return e.key }

// ErrorName returns error name.
func (e {{ errname . }}) ErrorName() string { return "{{ errname . }}" }

// Error satisfies the builtin error interface
func (e {{ errname . }}) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %s{{ (msgTyp .) }}.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = {{ errname . }}{}

var _ interface{
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = {{ errname . }}{}

{{ range .Fields }}{{ with (context .) }}{{ $f := .Field }}
	{{ if has .Rules "In" }}{{ if .Rules.In }}
		var {{ lookup .Field "InLookup" }} = map[{{ inType .Field .Rules.In }}]struct{}{
			{{- range .Rules.In }}
				{{ inKey $f . }}: {},
			{{- end }}
		}
	{{ end }}{{ end }}

	{{ if has .Rules "NotIn" }}{{ if .Rules.NotIn }}
		var {{ lookup .Field "NotInLookup" }} = map[{{ inType .Field .Rules.In }}]struct{}{
			{{- range .Rules.NotIn }}
				{{ inKey $f . }}: {},
			{{- end }}
		}
	{{ end }}{{ end }}

	{{ if has .Rules "Pattern"}}{{ if .Rules.Pattern }}
		var {{ lookup .Field "Pattern" }} = regexp.MustCompile({{ lit .Rules.GetPattern }})
	{{ end }}{{ end }}

	{{ if has .Rules "Items"}}{{ if .Rules.Items }}
	{{ if has .Rules.Items.GetString_ "Pattern" }} {{ if .Rules.Items.GetString_.Pattern }}
		var {{ lookup .Field "Pattern" }} = regexp.MustCompile({{ lit .Rules.Items.GetString_.GetPattern }})
	{{ end }}{{ end }}
	{{ end }}{{ end }}

	{{ if has .Rules "Items"}}{{ if .Rules.Items }}
	{{ if has .Rules.Items.GetString_ "In" }} {{ if .Rules.Items.GetString_.In }}
		var {{ lookup .Field "InLookup" }} = map[string]struct{}{
			{{- range .Rules.Items.GetString_.In }}
				{{ inKey $f . }}: {},
			{{- end }}
		}
	{{ end }}{{ end }}
	{{ end }}{{ end }}

	{{ if has .Rules "Items"}}{{ if .Rules.Items }}
	{{ if has .Rules.Items.GetString_ "NotIn" }} {{ if .Rules.Items.GetString_.NotIn }}
		var {{ lookup .Field "NotInLookup" }} = map[string]struct{}{
			{{- range .Rules.Items.GetString_.NotIn }}
				{{ inKey $f . }}: {},
			{{- end }}
		}
	{{ end }}{{ end }}
	{{ end }}{{ end }}

	{{ if has .Rules "Keys"}}{{ if .Rules.Keys }}
	{{ if has .Rules.Keys.GetString_ "Pattern" }} {{ if .Rules.Keys.GetString_.Pattern }}
		var {{ lookup .Field "Pattern" }} = regexp.MustCompile({{ lit .Rules.Keys.GetString_.GetPattern }})
	{{ end }}{{ end }}
	{{ end }}{{ end }}

	{{ if has .Rules "Values"}}{{ if .Rules.Values }}
	{{ if has .Rules.Values.GetString_ "Pattern" }} {{ if .Rules.Values.GetString_.Pattern }}
		var {{ lookup .Field "Pattern" }} = regexp.MustCompile({{ lit .Rules.Values.GetString_.GetPattern }})
	{{ end }}{{ end }}
	{{ end }}{{ end }}

{{ end }}{{ end }}
`

data "iosxr_{{snakeCase .Name}}" "example" {
{{- range  .Attributes}}
{{- if or (eq .Id true) (eq .Reference true)}}
  {{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if eq .Type "StringList"}}["{{.Example}}"]{{else if eq .Type "Int64List"}}[{{.Example}}]{{else}}{{.Example}}{{end}}
{{- end}}
{{- end}}
}

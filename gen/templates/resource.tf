resource "iosxr_{{snakeCase .Name}}" "example" {
{{- range  .Attributes}}
{{- if ne .ExcludeTest true}}
  {{.TfName}} = {{if eq .Type "String"}}"{{end}}{{.Example}}{{if eq .Type "String"}}"{{end}}
{{- end}}
{{- end}}
}

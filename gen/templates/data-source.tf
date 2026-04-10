data "iosxr_{{snakeCase .Name}}" "example" {
{{- if ne .RemovedInVersion ""}}
  # NOTE: This data source is not supported from IOS-XR version {{formatVersionDisplay .RemovedInVersion}} and above
  # Only use with versions earlier than {{formatVersionDisplay .RemovedInVersion}}
{{- end}}
{{- range  .Attributes}}
{{- if and (or .Id .Reference) (not .Legacy) (len .Example)}}
{{- if ne .AddedInVersion ""}}
  # Supported from version {{formatVersionDisplay .AddedInVersion}}
{{- end}}
{{- if ne .RemovedInVersion ""}}
  # Not supported from version {{formatVersionDisplay .RemovedInVersion}} and above
{{- end}}
  {{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if eq .Type "StringList"}}["{{.Example}}"]{{else if eq .Type "Int64List"}}[{{.Example}}]{{else}}{{.Example}}{{end}}
{{- end}}
{{- end}}
}

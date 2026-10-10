data "iosxr_{{snakeCase .Name}}" "example" {
{{- if ne .IntroducedInVersion ""}}
  # NOTE: This data source is only supported from IOS-XR version {{formatVersionDisplay .IntroducedInVersion}} and above
{{- end}}
{{- if ne .RemovedInVersion ""}}
  # NOTE: Only use with versions earlier than {{formatVersionDisplay .RemovedInVersion}}
{{- end}}
{{- range sortedAttrs .Attributes}}
{{- if and (or .Id .Reference) (len .Example) (eq .RemovedInVersion "")}}
  {{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if eq .Type "StringList"}}["{{.Example}}"]{{else if eq .Type "Int64List"}}[{{.Example}}]{{else}}{{.Example}}{{end}}
{{- end}}
{{- end}}
}

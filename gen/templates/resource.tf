resource "iosxr_{{snakeCase .Name}}" "example" {
{{- range  .Attributes}}
{{- if and (not .ExcludeTest) (not .ExcludeExample)}}
{{- if eq .Type "List"}}
  {{.TfName}} = [
    {
      {{- range  .Attributes}}
      {{- if and (not .ExcludeTest) (not .ExcludeExample)}}
      {{- if eq .Type "List"}}
        {{.TfName}} = [
          {
            {{- range  .Attributes}}
            {{- if and (not .ExcludeTest) (not .ExcludeExample)}}
            {{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if eq .Type "StringList"}}["{{.Example}}"]{{else if eq .Type "Int64List"}}[{{.Example}}]{{else}}{{.Example}}{{end}}
            {{- end}}
            {{- end}}
          }
        ]
      {{- else}}
      {{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if eq .Type "StringList"}}["{{.Example}}"]{{else if eq .Type "Int64List"}}[{{.Example}}]{{else}}{{.Example}}{{end}}
      {{- end}}
      {{- end}}
      {{- end}}
    }
  ]
{{- else}}
  {{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if eq .Type "StringList"}}["{{.Example}}"]{{else if eq .Type "Int64List"}}[{{.Example}}]{{else}}{{.Example}}{{end}}
{{- end}}
{{- end}}
{{- end}}
}

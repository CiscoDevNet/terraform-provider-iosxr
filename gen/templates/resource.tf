resource "iosxr_{{snakeCase .Name}}" "example" {
{{- if ne .IntroducedInVersion ""}}
  # NOTE: This resource is only supported from IOS-XR version {{formatVersionDisplay .IntroducedInVersion}} and above
{{- end}}
{{- if ne .RemovedInVersion ""}}
  # NOTE: Only use with versions earlier than {{formatVersionDisplay .RemovedInVersion}}
{{- end}}
{{- range sortedAttrs .Attributes}}
{{- if and (not .ExcludeExample) (not .ExcludeTest) (not .ExcludeTestFrom) (eq .RemovedInVersion "") (or (not (len .TestTags)) .IncludeExample)}}
{{- if eq .Type "List"}}
  {{.TfName}} = [
    {
      {{- range sortedAttrs .Attributes}}
      {{- if and (not .ExcludeExample) (not .ExcludeTest) (not .ExcludeTestFrom) (eq .RemovedInVersion "") (or (not (len .TestTags)) .IncludeExample) (or (eq .Type "List") (len .Example))}}
      {{- if eq .Type "List"}}
        {{.TfName}} = [
          {
            {{- range sortedAttrs .Attributes}}
            {{- if and (not .ExcludeExample) (not .ExcludeTest) (not .ExcludeTestFrom) (eq .RemovedInVersion "") (or (not (len .TestTags)) .IncludeExample) (or (eq .Type "List") (len .Example))}}
            {{- if eq .Type "List"}}
              {{.TfName}} = [
                {
                  {{- range sortedAttrs .Attributes}}
                  {{- if and (not .ExcludeExample) (not .ExcludeTest) (not .ExcludeTestFrom) (eq .RemovedInVersion "") (or (not (len .TestTags)) .IncludeExample) (or (eq .Type "List") (len .Example))}}
                  {{- if eq .Type "List"}}
                    {{.TfName}} = [
                      {
                        {{- range sortedAttrs .Attributes}}
                        {{- if and (not .ExcludeExample) (not .ExcludeTest) (not .ExcludeTestFrom) (eq .RemovedInVersion "") (or (not (len .TestTags)) .IncludeExample) (or (eq .Type "List") (len .Example))}}
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
        ]
      {{- else}}
      {{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if eq .Type "StringList"}}["{{.Example}}"]{{else if eq .Type "Int64List"}}[{{.Example}}]{{else}}{{.Example}}{{end}}
      {{- end}}
      {{- end}}
      {{- end}}
    }
  ]
{{- else if len .Example}}
  {{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if eq .Type "StringList"}}["{{.Example}}"]{{else if eq .Type "Int64List"}}[{{.Example}}]{{else}}{{.Example}}{{end}}
{{- end}}
{{- end}}
{{- end}}
}
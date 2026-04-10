resource "iosxr_{{snakeCase .Name}}" "example" {
{{- if ne .RemovedInVersion ""}}
  # NOTE: This resource is not supported from IOS-XR version {{formatVersionDisplay .RemovedInVersion}} and above
  # Only use with versions earlier than {{formatVersionDisplay .RemovedInVersion}}
{{- end}}
{{- range  .Attributes}}
{{- if and (not .ExcludeTest) (not .ExcludeExample) (or (not (len .TestTags)) .IncludeExample) (not .Legacy)}}
{{- if eq .Type "List"}}
{{- if ne .AddedInVersion ""}}
  # Supported from version {{formatVersionDisplay .AddedInVersion}}
{{- end}}
{{- if ne .RemovedInVersion ""}}
  # Not supported from version {{formatVersionDisplay .RemovedInVersion}} and above
{{- end}}
  {{.TfName}} = [
    {
      {{- range  .Attributes}}
      {{- if and (not .ExcludeTest) (not .ExcludeExample) (or (not (len .TestTags)) .IncludeExample) (not .Legacy) (or (eq .Type "List") (len .Example))}}
      {{- if ne .AddedInVersion ""}}
      # Supported from version {{formatVersionDisplay .AddedInVersion}}
      {{- end}}
      {{- if ne .RemovedInVersion ""}}
      # Not supported from version {{formatVersionDisplay .RemovedInVersion}} and above
      {{- end}}
      {{- if eq .Type "List"}}
        {{.TfName}} = [
          {
            {{- range  .Attributes}}
            {{- if and (not .ExcludeTest) (not .ExcludeExample) (or (not (len .TestTags)) .IncludeExample) (not .Legacy) (or (eq .Type "List") (len .Example))}}
            {{- if ne .AddedInVersion ""}}
            # Supported from version {{formatVersionDisplay .AddedInVersion}}
            {{- end}}
            {{- if ne .RemovedInVersion ""}}
            # Not supported from version {{formatVersionDisplay .RemovedInVersion}} and above
            {{- end}}
            {{- if eq .Type "List"}}
              {{.TfName}} = [
                {
                  {{- range  .Attributes}}
                  {{- if and (not .ExcludeTest) (not .ExcludeExample) (or (not (len .TestTags)) .IncludeExample) (not .Legacy) (or (eq .Type "List") (len .Example))}}
                  {{- if ne .AddedInVersion ""}}
                  # Supported from version {{formatVersionDisplay .AddedInVersion}}
                  {{- end}}
                  {{- if ne .RemovedInVersion ""}}
                  # Not supported from version {{formatVersionDisplay .RemovedInVersion}} and above
                  {{- end}}
                  {{- if eq .Type "List"}}
                    {{.TfName}} = [
                      {
                        {{- range  .Attributes}}
                        {{- if and (not .ExcludeTest) (not .ExcludeExample) (or (not (len .TestTags)) .IncludeExample) (not .Legacy) (or (eq .Type "List") (len .Example))}}
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
{{- if ne .AddedInVersion ""}}
  # Supported from version {{formatVersionDisplay .AddedInVersion}}
{{- end}}
{{- if ne .RemovedInVersion ""}}
  # Not supported from version {{formatVersionDisplay .RemovedInVersion}} and above
{{- end}}
  {{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if eq .Type "StringList"}}["{{.Example}}"]{{else if eq .Type "Int64List"}}[{{.Example}}]{{else}}{{.Example}}{{end}}
{{- end}}
{{- end}}
{{- end}}
}

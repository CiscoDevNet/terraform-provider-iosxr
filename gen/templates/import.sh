terraform import iosxr_{{snakeCase .Name}}.example "{{range $index, $attr := .Attributes}}{{if or $attr.Reference $attr.Id}}{{if $index}},{{end}}<{{$attr.TfName}}>{{end}}{{end}}"

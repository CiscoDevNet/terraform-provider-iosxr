terraform import iosxr_{{snakeCase .Name}}.example "{{range $i, $e := (importAttributes .)}}{{if $i}},{{end}}<{{.TfName}}>{{end}}"

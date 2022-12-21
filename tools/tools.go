//go:build tools

package tools

import (
	// Documentation generation
	_ "github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs"
	// Code generation
	_ "github.com/openconfig/goyang/pkg/yang"
	_ "golang.org/x/tools/cmd/goimports"
	_ "gopkg.in/yaml.v3"
)

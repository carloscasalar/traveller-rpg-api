//go:build tools
// +build tools

package tools

import (
	_ "github.com/mgechev/revive"
	_ "github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen"
	_ "golang.org/x/tools/cmd/goimports"
)

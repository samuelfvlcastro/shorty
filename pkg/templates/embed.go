package templates

import "embed"

//go:embed layouts/*.go.html components/*.go.html pages/*.go.html
var FS embed.FS

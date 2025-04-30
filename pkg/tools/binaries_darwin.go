//go:build darwin
// +build darwin

package tools

import _ "embed"

//go:embed darwin/tailwindcss
var tailwindBinaryData []byte
var tailwindBinaryName = "tailwindcss"

//go:build darwin
// +build darwin

package tools

import _ "embed"

//go:embed darwin/air
var airBinaryData []byte
var airBinaryName = "air"

//go:embed darwin/tailwindcss
var tailwindBinaryData []byte
var tailwindBinaryName = "tailwindcss"

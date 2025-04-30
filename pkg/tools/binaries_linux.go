//go:build linux
// +build linux

package tools

import _ "embed"

//go:embed linux/tailwindcss
var tailwindBinaryData []byte
var tailwindBinaryName = "tailwindcss"

//go:build windows
// +build windows

package tools

import _ "embed"

//go:embed windows/tailwindcss.exe
var tailwindBinaryData []byte
var tailwindBinaryName = "tailwindcss.exe"

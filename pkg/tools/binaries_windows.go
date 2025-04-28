//go:build windows
// +build windows

package tools

import _ "embed"

//go:embed windows/air.exe
var airBinaryData []byte
var airBinaryName = "air.exe"

//go:embed windows/tailwindcss.exe
var tailwindBinaryData []byte
var tailwindBinaryName = "tailwindcss.exe"

//go:build linux
// +build linux

package tools

import _ "embed"

//go:embed linux/air
var airBinaryData []byte
var airBinaryName = "air"

//go:embed linux/tailwindcss
var tailwindBinaryData []byte
var tailwindBinaryName = "tailwindcss"

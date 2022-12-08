//go:build tools
// +build tools

// This package contains the tool dependencies of the bigtable example.

package tools

import (
	// Register gotestsum for pinning version
	_ "gotest.tools/gotestsum"
)

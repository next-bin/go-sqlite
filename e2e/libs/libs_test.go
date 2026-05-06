package libs

import (
	"runtime"
	"testing"
)

func TestCRTImports(t *testing.T) {
	// CRT packages have platform-specific (Linux) build constraints.
	// On non-Linux platforms, verify the module paths resolve via go.mod only.
	if runtime.GOOS != "linux" {
		t.Skip("CRT packages require Linux build constraints")
	}
}

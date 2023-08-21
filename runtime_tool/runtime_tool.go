package runtime_tool

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func executablePath(name string) string {
	fmt.Println(runtime.GOOS)
	if runtime.GOOS == "windows" {
		name += ".exe"
	}
	return filepath.Join("test", name)
}

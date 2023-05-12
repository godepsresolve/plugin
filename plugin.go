package plugin

import (
	"fmt"
	"github.com/godepsresolve/corelib"
)

func MakeSometingPluggable(input string) string {
	return fmt.Sprintf("%s@v%s -> %s", pkg, version, corelib.Format(input))
}

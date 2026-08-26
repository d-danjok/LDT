package cli

import (
	definitions "LDT/src"
	"fmt"
)

func ListLCVersions() {
	for i, version := range definitions.LCVersions {
		fmt.Printf("%d:\t%s\n", i, version.Name)
	}
}

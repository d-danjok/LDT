package cli

import (
	definitions "LDT/src/programScopeData"
	"fmt"
)

func ListLCVersions() {
	for i, version := range definitions.LCVersions {
		fmt.Printf("%d. %s\n", i, version.Name)
	}
}

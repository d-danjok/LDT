package cli

import (
	definitions "LCAD/src"
	"fmt"
)

func ListAvailableAssemblies() {
	for i, assembly := range definitions.Assemblies {
		fmt.Printf("%d:\t%s\n", i, assembly.Name)
	}
}

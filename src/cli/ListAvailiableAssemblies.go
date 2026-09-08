package cli

import (
	definitions "LDT/src/appdata"
	"fmt"
)

func ListAvailableAssemblies() {
	for i, assembly := range definitions.Assemblies {
		fmt.Printf("%d. %s\n", i, assembly.Name)
	}
}

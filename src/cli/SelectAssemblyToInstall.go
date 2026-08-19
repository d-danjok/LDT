package cli

import (
	definitions "LCAD/src"
	"LCAD/src/structures"
	"fmt"
	"strconv"
)

func SelectAssemblyToInstall() (structures.Assembly, error) {
	var assemblyNum int
	var tmp string

	fmt.Printf("Select assembly you want to install by entering corresponding number\n\n")
	ListAvailableAssemblies()

	fmt.Printf("\n: ")
	_, err := fmt.Scanln(&tmp)
	if err != nil {
		return structures.Assembly{}, err
	}

	assemblyNum, err = strconv.Atoi(tmp)
	if err != nil {
		return structures.Assembly{}, err
	}

	if assemblyNum >= len(definitions.Assemblies) || assemblyNum < 0 {
		return structures.Assembly{}, fmt.Errorf("invalid assembly number")
	}

	return definitions.Assemblies[assemblyNum], nil
}

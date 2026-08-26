package cli

import (
	definitions "LDT/src"
	"LDT/src/structures"
)

func SelectAssemblyToInstall() (structures.Assembly, error) {
	assemblyNum, err := SelectByNum("assembly you want to install", len(definitions.Assemblies), ListAvailableAssemblies, nil)
	if err != nil {
		return structures.Assembly{}, err
	}

	return definitions.Assemblies[assemblyNum], nil
}

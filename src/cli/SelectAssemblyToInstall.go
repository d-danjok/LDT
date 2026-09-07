package cli

import (
	definitions "LDT/src/programScopeData"
	"LDT/src/structures/definitionHoldStructures"
)

func SelectAssemblyToInstall() (definitionHoldStructures.Assembly, error) {
	assemblyNum, err := SelectByNum("assembly you want to install", len(definitions.Assemblies), ListAvailableAssemblies, nil)
	if err != nil {
		return definitionHoldStructures.Assembly{}, err
	}

	return definitions.Assemblies[assemblyNum], nil
}

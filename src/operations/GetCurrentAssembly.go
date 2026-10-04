package operations

import (
	"LDT/src/appdata/userdata/assemblyData"
	"LDT/src/structures/assemblyDataHoldStructures"
)

func GetCurrentAssembly() assemblyDataHoldStructures.Assembly {
	return assemblyData.CurrentAssembly
}

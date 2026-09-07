package conversion

import (
	"LDT/src/functions/common"
	definitions "LDT/src/programScopeData"
	"LDT/src/structures/definitionHoldStructures"
	"fmt"
)

func StrToLCVersion(version string) (definitionHoldStructures.LCVersion, error) {
	index := common.FindInList(definitionHoldStructures.LCVersion{Name: version}, definitions.LCVersions)
	if index == -1 {
		return definitionHoldStructures.LCVersion{}, fmt.Errorf("version not found")
	}
	return definitions.LCVersions[index], nil
}

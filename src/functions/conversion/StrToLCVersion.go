package conversion

import (
	definitions "LDT/src/appdata"
	"LDT/src/functions/common"
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

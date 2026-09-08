package appdataHandling

import (
	definitions "LDT/src/appdata"
	"LDT/src/structures/definitionHoldStructures"
)

func CreateAssembliesFromLCVersions() {

	//convert plain versions into individual non modded assemblies
	for _, lcVersion := range definitions.LCVersions {
		definitions.Assemblies = append(definitions.Assemblies,
			definitionHoldStructures.Assembly{
				Name:                  lcVersion.Name + " Non-modded",
				BaseVersionManifestID: lcVersion.ManifestID,
			},
		)
	}
}

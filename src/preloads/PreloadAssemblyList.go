package preloads

import (
	definitions "LDT/src"
	"LDT/src/structures"
)

func PreloadAssemblyList() {

	//convert plain versions into individual non modded assemblies
	for _, lcVersion := range definitions.LCVersions {
		definitions.Assemblies = append(definitions.Assemblies,
			structures.Assembly{
				Name:                  lcVersion.Name + " Non-modded",
				BaseVersionManifestID: lcVersion.ManifestID,
			},
		)
	}
}

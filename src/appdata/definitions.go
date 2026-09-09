package definitions

import (
	"LDT/src/structures/definitionHoldStructures"
)

// manifestIDs for versions download
const (
	V45ManifestID string = "7637156099460715726"
	V50ManifestID string = "2961956797830002840"
	V56ManifestID string = "6074226372806880905"
	V69ManifestID string = "1367019593609280205"
	V73ManifestID string = "1749099131234587692"
)

var Assemblies []definitionHoldStructures.Assembly
var LCVersions []definitionHoldStructures.LCVersion

var CleanMode = true

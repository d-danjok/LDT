package definitions

import "LCAD/src/structures"

// manifestIDs for versions download
const (
	V50HotFixManifestID string = "2961956797830002840"
	V73ManifestID       string = "1749099131234587692"
	V56ManifestID       string = "6074226372806880905"
)

var SteamFolder string
var Assemblies []structures.Assembly
var LCVersions = []structures.LCVersion{
	{"v50HotFix", V50HotFixManifestID},
	{"v56", V56ManifestID},
	{"v73", V73ManifestID},
}

var CleanMode = true

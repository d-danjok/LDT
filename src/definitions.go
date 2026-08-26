package definitions

import "LDT/src/structures"

// manifestIDs for versions download
const (
	V50HotFixManifestID string = "2961956797830002840"
	V73ManifestID       string = "1749099131234587692"
	V56ManifestID       string = "6074226372806880905"
)

var SteamFolder string
var Assemblies = []structures.Assembly{
	{
		Name:                  "v73 Wesley's Basic",
		BaseVersionManifestID: V73ManifestID,
		Mods: structures.RemoteSrc{
			RemoteType: "GoogleDrive",
			Url:        "https://drive.google.com/file/d/12fwUhp5xuKSIlF_w6GllVlDHPDHLefGN/view?usp=sharing",
		},
	},
}

var LCVersions = []structures.LCVersion{
	{"v50HotFix", V50HotFixManifestID, "2024-07-06"},
	{"v56", V56ManifestID, "2024-08-17"},
	{"v73", V73ManifestID, "2026-03-29"},
}

var CleanMode = true

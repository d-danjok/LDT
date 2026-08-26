package definitions

import "LDT/src/structures"

// manifestIDs for versions download
const (
	V45ManifestID string = "7637156099460715726"
	V50ManifestID string = "2961956797830002840"
	V56ManifestID string = "6074226372806880905"
	V69ManifestID string = "1367019593609280205"
	V73ManifestID string = "1749099131234587692"
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
	{"v45", V45ManifestID, "2023-12-09"},
	{"v50", V50ManifestID, "2024-06-28"},
	{"v56", V56ManifestID, "2024-08-17"},
	{"v69", V69ManifestID, "2024-12-13"},
	{"v73", V73ManifestID, "2026-03-29"},
}

var CleanMode = true

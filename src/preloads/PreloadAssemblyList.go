package preloads

import (
	definitions "LCAD/src"
	"LCAD/src/structures"
)

func PreloadAssemblyList() {
	//TODO: Replace hardcode with some kind of online actualisation of available versions
	v50HotFixNonModded := structures.Assembly{Name: "v50 hotfix Non-modded"}
	v50HotFixNonModded.BaseVersionManifestID = definitions.V50HotFixManifest
	v50HotFixNonModded.SetMods("", "")
	definitions.Assemblies = append(definitions.Assemblies, v50HotFixNonModded)

	v56NonModded := structures.Assembly{Name: "v56 Non-modded"}
	v56NonModded.BaseVersionManifestID = definitions.V56ManifestID
	v56NonModded.SetMods("", "")
	definitions.Assemblies = append(definitions.Assemblies, v56NonModded)

	v73NonModded := structures.Assembly{Name: "v73 Non-modded"}
	v73NonModded.BaseVersionManifestID = definitions.V73ManifestID
	v73NonModded.SetMods("", "")
	definitions.Assemblies = append(definitions.Assemblies, v73NonModded)

	v73WesleysBasic := structures.Assembly{Name: "v73 Wesley's Basic"}
	v73WesleysBasic.BaseVersionManifestID = definitions.V73ManifestID
	v73WesleysBasic.SetMods("GoogleDrive", "https://drive.google.com/file/d/12fwUhp5xuKSIlF_w6GllVlDHPDHLefGN/view?usp=sharing")
	definitions.Assemblies = append(definitions.Assemblies, v73WesleysBasic)
}

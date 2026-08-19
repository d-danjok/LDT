package preloads

import (
	definitions "LCAD/src"
	"LCAD/src/structures"
)

func PreloadAssemblyList() {
	//TODO: Replace hardcode with some kind of online actualisation of available versions

	v73NonModded := structures.Assembly{Name: "v73 Non modded"}
	v73NonModded.SetBase("GoogleDrive", "https://drive.google.com/file/d/1MaT6dXBp4GBs0CC2pgXqBwO0HVbWSzip/view?usp=sharing")
	v73NonModded.SetMods("", "")
	definitions.Assemblies = append(definitions.Assemblies, v73NonModded)

	v73WesleysBasic := structures.Assembly{Name: "v73 Wesley's Basic"}
	v73WesleysBasic.SetBase("GoogleDrive", "https://drive.google.com/file/d/1MaT6dXBp4GBs0CC2pgXqBwO0HVbWSzip/view?usp=sharing")
	v73WesleysBasic.SetMods("GoogleDrive", "https://drive.google.com/file/d/12fwUhp5xuKSIlF_w6GllVlDHPDHLefGN/view?usp=sharing")
	definitions.Assemblies = append(definitions.Assemblies, v73WesleysBasic)
}

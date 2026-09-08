package userdataHoldStructures

type Package struct {
	Name    string `json:"name"`
	Author  string `json:"author"`
	Version string `json:"version"`
}

type InstalledAssembly struct {
	Name              string    `json:"name"`
	LocationPath      string    `json:"locationPath"`
	LCVersion         string    `json:"lcVersion"`
	InstalledPackages []Package `json:"installedPackages"`
}

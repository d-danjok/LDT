package operations

import (
	"LDT/src/appdata/userdata/assemblyData"
	"LDT/src/functions/conversion"
	"LDT/src/functions/pkgInstallation"
)

func InstallPkgByLink(link string) error {
	destPath := assemblyData.CurrentAssembly.LocationPath

	version, err := conversion.StrToLCVersion(assemblyData.CurrentAssembly.LCVersion)
	if err != nil {
		return err
	}

	err = pkgInstallation.InstallPkgWithDependenciesByLCVersion(link, destPath, version)
	if err != nil {
		return err
	}

	return nil
}

func InstallPkgByName(pkgAuthor string, pkgName string) error {
	destPath := assemblyData.CurrentAssembly.LocationPath

	version, err := conversion.StrToLCVersion(assemblyData.CurrentAssembly.LCVersion)
	if err != nil {
		return err
	}

	maxPkgDate, err := pkgInstallation.GetMaxPkgDateByLCVersion(version)
	if err != nil {
		return err
	}

	deps, err := pkgInstallation.ResolveAndInstallPkg(pkgAuthor, pkgName, destPath, maxPkgDate)
	if err != nil {
		return err
	}

	return pkgInstallation.InstallDependencies(deps, destPath, maxPkgDate)
	return nil
}

func InstallPkgGroupByCode(code string) error {

	return nil
}

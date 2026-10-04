package main

import (
	definitions "LDT/src/appdata"
	"LDT/src/appdataHandling"
	"LDT/src/cli"
	"LDT/src/gui"
	installs2 "LDT/src/installs"
	"embed"
	"fmt"
	"os"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func runCLINew() error {
	return cli.Main()
}

func runCLI() error {
	cli.ClearTerminal()

	err := installs2.InstallWithIndividualMods()
	if err != nil {
		return fmt.Errorf("Error assembly: %v\n", err)
	}
	return nil
}

func runGUI() error {
	// Create an instance of the app structures
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Lethal Download Tools",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	return err
}

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "-a" {
			definitions.CleanMode = false
		}
	}

	//appdataHandling
	err := appdataHandling.LoadConstants()
	if err != nil {
		gui.ShowErrorPopup(fmt.Sprintf("Error loading constants: %v\n", err))
		return
	}
	err = appdataHandling.LoadUserData()
	if err != nil {
		gui.ShowErrorPopup(fmt.Sprintf("Error loading user data: %v\n", err))
		return
	}

	defer func() {
		err = appdataHandling.StoreUserData()
		if err != nil {
			gui.ShowErrorPopup(fmt.Sprintf("Error storing user data: %v\n", err))
		}
	}()

	err = runCLI()

	if err != nil {
		gui.ShowErrorPopup(fmt.Sprintf("Error: %v\n", err))
		return
	}
}

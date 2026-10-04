package cli

import (
	"LDT/src/functions/common"
	"fmt"
	"strings"
)

type Msg struct {
	welcome,
	homeFmt,
	help,
	selectHome,
	selectSuccessFmt,
	selectFailFmt,
	launch,
	listHome,
	listAssembliesFmt,
	listPackagesFmt,
	installHome,
	installSuccessFmt,
	installFailFmt,
	newHome,
	newSuccessFmt,
	newFailFmt string
}

func Main() error {
	var err error

	var currCmd string
	var currArgs []string

	//TODO: implement this shit
	for {
		//debug
		fmt.Printf("%s %#v\n  : ", currCmd, currArgs)

		switch currCmd {
		case "": //welcome page

		case "home": //home page

		case "help": //list of all basic commands

		case "select": //assembly selection

		case "launch": //launch currently selected assembly

		case "list": //list smth (assemblies, packages installed for currently selected assemblies)

		case "install": //install new package

		case "new":
		//create new assembly (from code or version num)

		case "exit": //exit program
			break

		default: //non-existing program

		}

		inputLine := common.ReadLineWithSpaces()
		if err != nil {
			return err
		}
		lineParts := strings.Split(inputLine, " ")
		currCmd = lineParts[0]
		currArgs = lineParts[1:]

	}
	return nil
}

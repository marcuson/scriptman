package config

import (
	"fmt"

	"github.com/adrg/xdg"
)

func PrintConfig() error {
	absScriptmanHome, err := xdg.DataFile(SCRIPTMAN_HOME_DEFAULT)
	if err != nil {
		return err
	}

	absScriptHome, err := xdg.DataFile(SCRIPT_HOME_DEFAULT)
	if err != nil {
		return err
	}

	linkedBinDir, err := xdg.DataFile(BIN_HOME_DEFAULT)
	if err != nil {
		return err
	}

	fmt.Println("Scriptman home is: " + absScriptmanHome)
	fmt.Println("Script home is: " + absScriptHome)
	fmt.Println("Linked bin dir is: " + linkedBinDir)
	return nil
}

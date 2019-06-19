package main

import (
	"fmt"
	"os"
	"path"
)

type RestoreCommand struct {
	Help           bool   `gli:"^help,h"`
	ConfigPath     string `gli:"config,c" description:"Location of the config yaml file to use" default:"./config.yml"`
	PullConfig     bool   `gli:"pull,p" description:"Pull from remote before restore" default:"true"`
	Overwrite      bool   `gli:"overrwite,o" description:"overwrite existing files during restore" default:"true"`
	IgnoreExisting bool   `gli:"ignore-existing,i" description:"dont overwrite any files that exist on the system"`
	Verbose        bool   `gli:"verbose,v" description:"display operations as they happen"`
}

func (cmd *RestoreCommand) Run() int {
	loggerEnabled = cmd.Verbose

	config, err := config(cmd.ConfigPath)
	if nil != err {
		fmt.Printf("Failed to load config %s", cmd.ConfigPath)
		return CODE_CONFIG_ERROR
	}

	if cmd.PullConfig {
		if err := pullFromRenote(config, cmd.Verbose); nil != err {
			fmt.Println("Failed to pull from remote server")
			return CODE_REMOTE_ERROR
		}
	}

	if err := restoreFiles(config, cmd.Overwrite, cmd.IgnoreExisting, cmd.Verbose); nil != err {
		fmt.Println("Failed to restore files")
		return CODE_STORE_ERROR
	}

	return CODE_SUCCESS
}

func (cmd *RestoreCommand) NeedHelp() bool {
	return cmd.Help
}

func restoreFiles(config *AppConfig, overwrite, ignoreExisting, verbose bool) error {
	if 0 == len(config.Files.Home) {
		logf("no files in the home directory")
	} else {
		for _, file := range config.Files.Home {
			source := path.Join(config.Path.Store, file)
			destination := path.Join(config.Path.Home, file)
			if _, err := os.Stat(source); os.IsExist(err) {
				logf("not found in store: %s", source)
				continue
			}

			logf("%s -> %s", source, destination)

			if err := copyR(source, destination, overwrite, ignoreExisting); nil != err {
				return err
			}
		}
	}

	if 0 == len(config.Files.Absolute) {
		if verbose {
			logf("no absolute file paths")
		}
	} else {
		for _, file := range config.Files.Absolute {
			source := file
			destination := path.Join(config.Path.Store, file)
			if _, err := os.Stat(source); os.IsExist(err) {
				logf("not found in store: %s", source)
				continue
			}

			logf("%s -> %s", source, destination)

			if err := copyR(source, destination, overwrite, ignoreExisting); nil != err {
				return err
			}
		}
	}

	return nil
}

func pullFromRenote(config *AppConfig, verbose bool) error {
	// TODO: implement
	return nil
}

package main

import "fmt"

type RestoreCommand struct {
	Help           bool   `gli:"^help,h"`
	ConfigPath     string `gli:"config,c" description:"Location of the config yaml file to use" default:"./config.yml"`
	PullConfig     bool   `gli:"pull,p" description:"Pull from remote before restore" default:"true"`
	Overwrite      bool   `gli:"overrwite,o" description:"overwrite existing files during restore" default:"true"`
	IgnoreExisting bool   `gli:"ignore-existing,i" description:"dont overwrite any files that exist on the system"`
	Verbose        bool   `gli:"verbose,v" description:"display operations as they happen"`
}

func (cmd *RestoreCommand) Run() int {
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
	// TODO: implement
	return nil
}

func pullFromRenote(config *AppConfig, verbose bool) error {
	// TODO: implement
	return nil
}

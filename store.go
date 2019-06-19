package main

import "fmt"

type StoreCommand struct {
	Help           bool   `gli:"^help,h"`
	ConfigPath     string `gli:"config,c" description:"Location of the config yaml file to use" default:"./config.yml"`
	PushConfig     bool   `gli:"push,p" description:"Push config files to remote store after collection" default:"true"`
	Overwrite      bool   `gli:"overrwite,o" description:"overwrite existing files during restore" default:"true"`
	IgnoreExisting bool   `gli:"ignore-existing,i" description:"dont overwrite any files that exist on the system"`
	Verbose        bool   `gli:"verbose,v" description:"display operations as they happen"`
}

func (cmd *StoreCommand) Run() int {
	config, err := config(cmd.ConfigPath)
	if nil != err {
		fmt.Printf("Failed to load config %s", cmd.ConfigPath)
		return CODE_CONFIG_ERROR
	}

	if err := storeFiles(config, cmd.Overwrite, cmd.IgnoreExisting, cmd.Verbose); nil != err {
		fmt.Println("Failed to store files")
		return CODE_STORE_ERROR
	}

	if cmd.PushConfig {
		if err := pushToRemote(config, cmd.Verbose); nil != err {
			fmt.Println("Failed to push to remote server")
			return CODE_REMOTE_ERROR
		}
	}

	return CODE_SUCCESS
}

func (cmd *StoreCommand) NeedHelp() bool {
	return cmd.Help
}

func storeFiles(config *AppConfig, overwrite, ignoreExisting, verbose bool) error {
	// TODO: implement
	return nil
}

func pushToRemote(config *AppConfig, verbose bool) error {
	// TODO: implement
	return nil
}

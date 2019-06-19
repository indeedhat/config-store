package main

type StoreCommand struct {
	Help       bool   `gli:"^help,h"`
	ConfigPath string `gli:"config,c" description:"Location of the config yaml file to use" default:"./config.yml"`
	PushConfig bool   `gli:"push,p" description:"Push config files to remote store after collection" default:"true"`
}

func (cmd *StoreCommand) Run() int {
	return 0
}

func (cmd *StoreCommand) NeedHelp() bool {
	return cmd.Help
}

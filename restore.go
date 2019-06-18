package configstore

type RestoreCommand struct {
	Help       bool   `gli:"^help,h"`
	ConfigPath string `gli:"config,c" description:"Location of the config yaml file to use" default:"./config.yml"`
	Overwrite  bool   `gli:"overrwite,o" description:"overwrite existing files during restore" default:"true"`
	PullConfig bool   `gli:"pull,p" description:"Pull from remote before restore" default:"true"`
}

func (cmd *RestoreCommand) Run() int {
	return 0
}

func (cmd *RestoreCommand) NeedHelp() bool {
	return cmd.Help
}

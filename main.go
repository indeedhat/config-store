package configstore

import "github.com/indeedhat/gli"

type ConfigClone struct {
	Store   StoreCommand   `gli:"store" description:"Collect config files and store them remotely"`
	Restore RestoreCommand `gli:"restore" description:"restore config files from the store"`
}

func (c *ConfigClone) Run() int {
	return 0
}

func main() {
	app := gli.NewApplication(&ConfigClone{}, "Config Clone")
	app.Run()
}

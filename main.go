package main

import "github.com/indeedhat/gli"

var app *gli.App

type ConfigClone struct {
	Store   StoreCommand   `gli:"store,s" description:"Collect config files and store them remotely"`
	Restore RestoreCommand `gli:"restore,r" description:"restore config files from the store"`
}

func (c *ConfigClone) Run() int {
	app.ShowHelp(true)
	return 0
}

func main() {
	app = gli.NewApplication(&ConfigClone{}, "Config Clone")
	app.Run()
}

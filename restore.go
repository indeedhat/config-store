package main

import (
	"fmt"
	"os"
	"path"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
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
	if cmd.Verbose {
		loggingLevel = loggingLevel | LOG_VERBOSE
	}

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
		logInfo("no files in the home directory")
	} else {
		for _, file := range config.Files.Home {
			source := path.Join(config.Path.Store, file)
			destination := path.Join(config.Path.Home, file)
			if _, err := os.Stat(source); os.IsExist(err) {
				logVerbose("not found in store: %s", source)
				continue
			}

			logVerbose("%s -> %s", source, destination)

			if err := copyR(source, destination, overwrite, ignoreExisting); nil != err {
				return err
			}
		}
	}

	if 0 == len(config.Files.Absolute) {
		if verbose {
			logInfo("no absolute file paths")
		}
	} else {
		for _, file := range config.Files.Absolute {
			source := file
			destination := path.Join(config.Path.Store, file)
			if _, err := os.Stat(source); os.IsExist(err) {
				logVerbose("not found in store: %s", source)
				continue
			}

			logVerbose("%s -> %s", source, destination)

			if err := copyR(source, destination, overwrite, ignoreExisting); nil != err {
				return err
			}
		}
	}

	return nil
}

func pullFromRenote(config *AppConfig, verbose bool) error {
	_, workTree, err := openRepo(config)
	if nil != err {
		return err
	}

	if err := checkoutBranch(workTree, config); nil != err {
		return err
	}

	status, err := workTree.Status()
	if nil != err {
		logErrorf("Failed to get status %s", err)
		return err
	}

	if !status.IsClean() {
		logVerbose("cannot pull, repo has unstaged changes")
		return nil
	}

	err = workTree.Pull(&git.PullOptions{
		Auth: &http.BasicAuth{
			Username: config.Remote.User,
			Password: config.Remote.Token,
		},
	})
	if nil != err {
		logErrorf("failed to pull from remote %s", err)
		return err
	}

	return nil
}

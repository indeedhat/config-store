package main

import (
	"os"
	"os/exec"
	"path"
	"time"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

type StoreCommand struct {
	Help           bool   `gli:"^help,h"`
	ConfigPath     string `gli:"config,c" description:"Location of the config yaml file to use" default:"./config.yml"`
	PushConfig     bool   `gli:"push,p" description:"Push config files to remote store after collection" default:"true"`
	Overwrite      bool   `gli:"overrwite,o" description:"overwrite existing files during restore" default:"true"`
	IgnoreExisting bool   `gli:"ignore-existing,i" description:"dont overwrite any files that exist on the system"`
	Verbose        bool   `gli:"verbose,v" description:"display operations as they happen"`
}

func (cmd *StoreCommand) Run() int {
	if cmd.Verbose {
		loggingLevel = loggingLevel | LOG_VERBOSE
	}

	config, err := config(cmd.ConfigPath)
	if nil != err {
		logErrorf("Failed to load config %s", cmd.ConfigPath)
		return CODE_CONFIG_ERROR
	}

	repo, workTree, err := openRepoForWrite(config)
	if nil != err {
		return CODE_REMOTE_ERROR
	}

	if err := storeFiles(config, cmd.Overwrite, cmd.IgnoreExisting); nil != err {
		logErrorf("Failed to store files")
		return CODE_STORE_ERROR
	}

	if cmd.PushConfig {
		if err := pushToRemote(repo, workTree, config); nil != err {
			logErrorf("Failed to push to remote server")
			return CODE_REMOTE_ERROR
		}
	}

	return CODE_SUCCESS
}

func (cmd *StoreCommand) NeedHelp() bool {
	return cmd.Help
}

func storeFiles(config *AppConfig, overwrite, ignoreExisting bool) error {
	if 0 == len(config.Files.Home) {
		logInfo("no files in the home directory")
	} else {
		for _, file := range config.Files.Home {
			source := path.Join(config.Path.Home, file)
			destination := path.Join(config.Path.Store, file)

			logVerbose("%s -> %s", source, destination)

			if err := copyR(source, destination, overwrite, ignoreExisting); nil != err {
				return err
			}
		}
	}

	if 0 == len(config.Files.Absolute) {
		logInfo("no absolute file paths")
	} else {
		for _, file := range config.Files.Absolute {
			source := file
			destination := path.Join(config.Path.Store, file)

			logVerbose("%s -> %s", source, destination)

			if err := copyR(source, destination, overwrite, ignoreExisting); nil != err {
				return err
			}
		}
	}

	return nil
}

func openRepoForWrite(config *AppConfig) (*git.Repository, *git.Worktree, error) {
	repo, workTree, err := openRepo(config)
	if nil != err {
		return nil, nil, err
	}

	if err := checkoutBranch(workTree, config); nil != err {
		return nil, nil, err
	}

	return repo, workTree, nil
}

func pushToRemote(repo *git.Repository, workTree *git.Worktree, config *AppConfig) error {
	status, err := workTree.Status()
	if nil != err {
		logErrorf("Failed to get status %s", err)
		return err
	}

	if status.IsClean() {
		logVerbose("nothing to commit, exiting")
		return nil
	}

	// TODO: fix this when they fix the Add method
	// _, err = workTree.Add(".")
	cmd := exec.Command("git", "add", ".")
	cmd.Dir = workTree.Filesystem.Root()
	err = cmd.Run()
	if err != nil {
		logErrorf("Failed to add files to repo %s", err)
	}

	_, err = workTree.Commit("Config Update", &git.CommitOptions{
		Author: &object.Signature{
			Name:  config.Remote.User,
			Email: config.Remote.Email,
			When:  time.Now(),
		},
	})

	if nil != err {
		logErrorf("failed to commit %s", err)
		return nil
	}

	options := &git.PushOptions{
		Auth: &http.BasicAuth{
			Username: config.Remote.User,
			Password: config.Remote.Token,
		},
		Progress: os.Stdout,
	}
	if err := repo.Push(options); nil != err {
		logErrorf("Failed to push to remote: %s", err)
		return err
	}

	return nil
}

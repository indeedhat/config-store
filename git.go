package main

import (
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

func openRepo(config *AppConfig) (*git.Repository, *git.Worktree, error) {
	repo, err := git.PlainOpen(config.Path.Store)
	if nil != err {
		logErrorf("Failed to ope repo %s", err)
		return nil, nil, err
	}

	workTree, err := repo.Worktree()
	if nil != err {
		logErrorf("Failed to open worktree %s", err)
		return nil, nil, err
	}

	return repo, workTree, nil
}

func checkoutBranch(workTree *git.Worktree, config *AppConfig) error {
	branch := plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", config.Remote.Branch))
	// this is properly lazy and hacky but meh
	err := workTree.Checkout(&git.CheckoutOptions{Branch: branch})
	if nil == err {
		return nil
	} else {
		logVerbose("failed to checkout to branch %s", err)
	}

	// create new branch on checkout if normal checkout failed
	err = workTree.Checkout(&git.CheckoutOptions{
		Branch: branch,
		Create: true,
		Force:  true,
	})
	if nil != err {
		logErrorf("Failed to checkout: %s", err)
		return err
	}

	return nil
}

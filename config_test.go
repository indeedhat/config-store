package main

import (
	"testing"
)

func TestConfigLoaded(t *testing.T) {
	if nil == t_config {
		t.Error("Failed to load config from file")
	}
}

func TestPathStruct(t *testing.T) {
	if "/home/indeedhat" != t_config.Path.Home {
		t.Error("Home path is invalid")
	}

	if "./store" != t_config.Path.Store {
		t.Errorf("Store path is invalid")
	}
}

func TestRemoteCredentials(t *testing.T) {
	if "https://github.com/indeedhat/config-store" != t_config.Remote.URL {
		t.Errorf("Remote url is invalid")
	}

	if "laptop" != t_config.Remote.Branch {
		t.Error("Remote branch is invalid")
	}

	if "indeedhat" != t_config.Remote.User {
		t.Error("Remote user is invalid")
	}

	if "noreply@phpmatt.com" != t_config.Remote.Email {
		t.Error("Remote email is invalid")
	}

	if "somefaketoken" != t_config.Remote.Token {
		t.Error("Remote token is invalid")
	}
}

func TestHomeFileList(t *testing.T) {
	expected := []string{
		".i3",
		".bashrc",
		".bash_profile",
	}

	if !t_stringSlicesMatch(expected, t_config.Files.Home) {
		t.Errorf("Home file list does not match the expected one")
	}
}

func TestAbsoluteFileList(t *testing.T) {
	expected := []string{
		"/etc/crontab",
		"/etc/php.ini",
	}

	if !t_stringSlicesMatch(expected, t_config.Files.Absolute) {
		t.Error("Absolute file list does not match the expected one")
	}
}

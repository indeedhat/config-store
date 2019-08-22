package main

import (
	"errors"
	"testing"
)

func TestCanDetectFileExistsError(t *testing.T) {
	err := FileExistsError{}

	if !isFileExistsError(err) {
		t.Error("Did not detect FileExistsError")
	}
}

func TestDoesNotMistakeErrorForFileExistsError(t *testing.T) {
	err := errors.New("some error")

	if isFileExistsError(err) {
		t.Error("Incorrectly identified standard error as FileExistsError")
	}
}

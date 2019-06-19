package main

type FileExistsError struct {
	err  string
	Path string
}

func (e FileExistsError) Error() string {
	return e.err
}

func isFileExistsError(err error) bool {
	_, ok := err.(FileExistsError)
	return ok
}

package main

var (
	loggerEnabled = false
)

func logf(message string, args ...interface{}) {
	if !loggerEnabled {
		return
	}

	fmt.Println("Log: %s", fmt.Sprintf(messae, args...))
}

func logErrorf(message string, args ...interface{}) {
	fmt.Println("Error: %s", fmt.Springf(message, args...))
}

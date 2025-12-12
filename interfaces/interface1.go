package interfaces

import "fmt"

type Logger interface {
	Log(message string)
}

// Concrete type : Console Logger
type consoleLogger struct{}

func (cl consoleLogger) Log(message string) {
	fmt.Println("[CONSOLE] " + message)
}

// Concrete type : File Logger
// FileLogger writes messages to a specific file.
type fileLogger struct {
    Filename string
}

// Log implements the Logger interface for FileLogger.
func (fl fileLogger) Log(message string) {
    // In a real application, you'd handle file opening/writing properly.
    // For simplicity, we'll just simulate the action.
    fmt.Printf("[FILE:%s] Writing: %s\n", fl.Filename, message)
}

func Interface1() {
	// nil interface
	var l Logger
	// l.Log("Hello") => We cannot call methods using a nil interface

	cl := consoleLogger{}
	l = cl // Without even explicit type casting we can assign concrete types to a interface
	l.Log("Hello")

	fl := fileLogger{Filename: "sample.txt"}
	l = fl
	l.Log("World")
}
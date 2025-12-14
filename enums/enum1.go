package enums

import "fmt"

// step 1 : The Custom Type Definition for type safety
type ServerState int

// Step 2 : Constant Definition using 'iota'
const (
	// iota starts at 0. StateIdle = 0.
	// This is the default zero-value for the ServerState type.
	StateIdle ServerState = iota 

	// iota increments automatically. StateConnected = 1.
	StateConnected 

	// StateError = 2.
	StateError 

	// StateRetrying = 3.
	StateRetrying
)

// Step 3 : The String Mapping (Readability Layer)
// This is necessary for outputting meaningful messages to users or logs.
var serverStateNames = map[ServerState]string{
	StateIdle:      "Idle",
	StateConnected: "Connected",
	StateError:     "Error",
	StateRetrying:  "Retrying",
}

// Step 4 : Implement the Stringer Interface
func (s ServerState) String() string {
	return serverStateNames[s]
}

// Uage :
func Enum1(){
	var state ServerState = StateConnected

	switch state {
	case StateConnected:
		fmt.Println("Server is ready to receive requests.")
	case StateError:
		fmt.Println("Server has failed!")
	case StateRetrying:
		fmt.Println("Server is retrying...")
	case StateIdle:
		fmt.Println("Server is idle")
	}

	 fmt.Printf("Current State: %s (Raw Value: %d)\n", state, state) 
}
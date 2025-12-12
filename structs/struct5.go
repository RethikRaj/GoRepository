// Empty structs => Smallest possible type in go ( 0 Bytes )

package structs

import (
	"fmt"
	"time"
)

type typedEmpty struct{}

// Channel Signalling Using empty structs
func ChannelSignallingUsingEmptyStructs() {
	done := make(chan struct{})
	go func(){
		fmt.Println("Task runninng...")
		time.Sleep(time.Second*10)
		done <- struct{}{} // send signal (no memory)
	}()
	<-done              // wait
}

func Struct5() {
	// Empty struct
	empty := struct{}{} // anonymous empty struct
	e2 := typedEmpty{}
	fmt.Println(empty, e2)

	// Set Implementation 
	visited := make(map[string]struct{})
	visited["India"] = struct{}{}
	visited["America"] = struct{}{}

	fmt.Println(visited)

	ChannelSignallingUsingEmptyStructs()
}
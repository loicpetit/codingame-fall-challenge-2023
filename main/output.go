package main

type Writer struct{}

func (Writer) Write(action *Action) {
	if action == nil {
		// panic("No action to write")
	}
	WriteOutput("WAIT 1")
}

func NewWriter() OutputWriter[Action] {
	return Writer{}
}

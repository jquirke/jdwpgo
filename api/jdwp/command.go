package jdwp

// Command represents a command specification
type Command struct {
	Commandset     byte
	Command        byte
	HasCommandData bool
	HasReplyData   bool
}

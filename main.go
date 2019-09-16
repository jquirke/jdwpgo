package main

import (
	"fmt"
	"net"
	"time"

	"github.com/jquirke/jdwpgo/debuggercore"

	"github.com/jquirke/jdwpgo/jdwpsession"
)

func main() {

	/* hacky test playground for now */

	conn, err := net.Dial("tcp", "127.0.0.1:5006")
	if err != nil {
		fmt.Printf("error dial: %v\n", err)
		return
	}
	s := jdwpsession.New(conn)
	err = s.Start()
	if err != nil {
		fmt.Printf("error start: %v\n", err)
		return
	}

	debuggercore := debuggercore.NewFromJWDPSession(s)

	version, err := debuggercore.VMCommands().Version()

	if err != nil {
		fmt.Printf("err = %v\n", err)
	}
	fmt.Printf("version = %v\n", version)

	allClasses, err := debuggercore.VMCommands().AllClasses()

	if err != nil {
		fmt.Printf("err = %v\n", err)
	}
	fmt.Printf("allclasses = %v\n", allClasses)

	idSizes, err := debuggercore.VMCommands().IDSizes()

	if err != nil {
		fmt.Printf("err = %v\n", err)
	}
	fmt.Printf("idSizes = %v\n", idSizes)

	caps, err := debuggercore.VMCommands().Capabilities()

	if err != nil {
		fmt.Printf("err = %v\n", err)
	}
	fmt.Printf("caps = %v\n", caps)

	capsNew, err := debuggercore.VMCommands().CapabilitiesNew()

	if err != nil {
		fmt.Printf("err = %v\n", err)
	}
	fmt.Printf("caps = %v\n", capsNew)

	tlg, err := debuggercore.VMCommands().TopLevelThreadGroups()

	if err != nil {
		fmt.Printf("err = %v\n", err)
	}
	fmt.Printf("tlgs = %v\n", tlg)

	err = debuggercore.VMCommands().Resume()

	if err != nil {
		fmt.Printf("err = %v\n", err)
	}

	time.Sleep(time.Second)

	allThreads, err := debuggercore.VMCommands().AllThreads()

	if err != nil {
		fmt.Printf("err = %v\n", err)
	}
	fmt.Printf("allThreads = %v\n", allThreads)

	for idx, threadID := range allThreads.Threads {
		name, err := debuggercore.ThreadCommands().Name(threadID)

		if err != nil {
			fmt.Printf("err = %v\n", err)
		}
		fmt.Printf("thread idx = %v, tid= %s name=%s\n", idx, threadID.String(), (string)(name.ByteString))

	}

}

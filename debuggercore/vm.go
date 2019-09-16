package debuggercore

import (
	"github.com/jquirke/jdwpgo/protocol/vm"
)

// VMCommands expose the VM commands
type VMCommands interface {
	// Class
	AllClasses() (*vm.AllClassReply, error)
	// Thread ops
	AllThreads() (*vm.AllThreadsReply, error)
	TopLevelThreadGroups() (*vm.TopLevelThreadGroupsReply, error)
	// Bootstrap
	Version() (*vm.VersionReply, error)
	IDSizes() (*vm.IDSizesReply, error)
	Capabilities() (*vm.CapabilitiesReply, error)
	CapabilitiesNew() (*vm.CapabilitiesNewReply, error)
	//Control
	Suspend() error
	Resume() error
	HoldEvents() error
	ReleaseEvents() error
	Exit(int32) error
}

func (d *debuggercore) Version() (*vm.VersionReply, error) {
	var versionReply vm.VersionReply
	err := d.processCommand(vm.VersionCommand, nil, &versionReply)
	if err != nil {
		return nil, err
	}
	return &versionReply, nil
}

func (d *debuggercore) AllClasses() (*vm.AllClassReply, error) {
	var allclassesReply vm.AllClassReply
	err := d.processCommand(vm.AllClassesCommand, nil, &allclassesReply)
	if err != nil {
		return nil, err
	}
	return &allclassesReply, nil
}

func (d *debuggercore) AllThreads() (*vm.AllThreadsReply, error) {
	var allthreadsReply vm.AllThreadsReply
	err := d.processCommand(vm.AllThreadsCommand, nil, &allthreadsReply)
	if err != nil {
		return nil, err
	}
	return &allthreadsReply, nil
}

func (d *debuggercore) TopLevelThreadGroups() (*vm.TopLevelThreadGroupsReply, error) {
	var topLevelThreadGroupsReply vm.TopLevelThreadGroupsReply
	err := d.processCommand(vm.TopLevelThreadGroupsCommand, nil, &topLevelThreadGroupsReply)
	if err != nil {
		return nil, err
	}
	return &topLevelThreadGroupsReply, nil
}

func (d *debuggercore) IDSizes() (*vm.IDSizesReply, error) {
	var idsizesReply vm.IDSizesReply
	err := d.processCommand(vm.IDSizesCommand, nil, &idsizesReply)
	if err != nil {
		return nil, err
	}
	return &idsizesReply, nil
}

func (d *debuggercore) Capabilities() (*vm.CapabilitiesReply, error) {
	var capsReply vm.CapabilitiesReply
	err := d.processCommand(vm.CapabilitiesCommand, nil, &capsReply)
	if err != nil {
		return nil, err
	}
	return &capsReply, nil
}

func (d *debuggercore) CapabilitiesNew() (*vm.CapabilitiesNewReply, error) {
	var capsNewReply vm.CapabilitiesNewReply
	err := d.processCommand(vm.CapabilitiesNewCommand, nil, &capsNewReply)
	if err != nil {
		return nil, err
	}
	return &capsNewReply, nil
}

func (d *debuggercore) Suspend() error {
	err := d.processCommand(vm.SuspendCommand, nil, nil)
	if err != nil {
		return err
	}
	return nil
}

func (d *debuggercore) Resume() error {
	err := d.processCommand(vm.ResumeCommand, nil, nil)
	if err != nil {
		return err
	}
	return nil
}

func (d *debuggercore) HoldEvents() error {
	err := d.processCommand(vm.HoldEventsCommand, nil, nil)
	if err != nil {
		return err
	}
	return nil
}

func (d *debuggercore) ReleaseEvents() error {
	err := d.processCommand(vm.ReleaseEventsCommand, nil, nil)
	if err != nil {
		return err
	}
	return nil
}

func (d *debuggercore) Exit(code int32) error {
	exitCommandData := &vm.ExitCommandData{
		ExitCode: code,
	}
	err := d.processCommand(vm.ExitCommand, exitCommandData, nil)
	if err != nil {
		return err
	}
	return nil
}

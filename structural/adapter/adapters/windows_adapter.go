package adapters

import (
	"design-patterns/structural/adapter/devices"
	"fmt"
)

type WindowsAdapter struct {
	windowsMachine *devices.Windows
}

func NewWindowsAdapter(windowsMachine *devices.Windows) *WindowsAdapter {
	return &WindowsAdapter{
		windowsMachine: windowsMachine,
	}
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.windowsMachine.InsertIntoUSBPort()
}

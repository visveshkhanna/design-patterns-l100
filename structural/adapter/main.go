package main

import (
	"design-patterns/structural/adapter/adapters"
	"design-patterns/structural/adapter/client"
	"design-patterns/structural/adapter/devices"
)

func main() {

	c := &client.Client{}

	mac := &devices.Mac{}
	c.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &devices.Windows{}

	windowsAdapter := adapters.NewWindowsAdapter(windowsMachine)
	c.InsertLightningConnectorIntoComputer(windowsAdapter)
}

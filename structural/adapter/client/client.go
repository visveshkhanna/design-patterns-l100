package client

import (
	"design-patterns/structural/adapter/interfaces"
	"fmt"
)

type Client struct{}

func (c *Client) InsertLightningConnectorIntoComputer(computer interfaces.IComputer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	computer.InsertIntoLightningPort()
}

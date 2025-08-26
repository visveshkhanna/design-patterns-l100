package main

func main() {
	light := &Light{}

	on := &OnCommand{device: light}
	off := &OffCommand{device: light}

	onBtn := &Button{command: on}
	offBtn := &Button{command: off}

	onBtn.press()
	offBtn.press()
}

package main

func main() {
	p := &Player{state: PausedState{}}
	p.pressPlay() // Playing
	p.pressPlay() // Pausing
}

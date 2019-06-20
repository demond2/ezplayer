package main

import (
	"github.com/demond2/ezplayer/player"
)

func main() {
	p := player.NewPlayer("/Users/p1contractor9/02 Magic.mp3")
	p.Play()
	p.Wait()
	p.Close()
}

package main

import (
	"time"

	"github.com/demond2/ezplayer/player"
)

func main() {
	p := player.NewPlayer("/Users/p1contractor9/02 Magic.mp3")
	p.Play()
	<-time.After(5 * time.Second)
	p.FadeOut()
	<-time.After(10 * time.Second)
	//p.Wait()
	p.Close()
}

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/demond2/ezplayer/player"
)

func main() {
	p := player.NewPlayer("/Users/p1contractor9/02 Magic.mp3")
	if err := p.Seek(670000); err != nil {
		fmt.Printf("Got error during seek: %s\n", err.Error())
		os.Exit(1)
	}
	p.Play()
	<-time.After(5 * time.Second)
	p.FadeOut()
	<-time.After(10 * time.Second)
	//p.Wait()
	p.Close()
}

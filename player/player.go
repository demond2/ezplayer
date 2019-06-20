package player

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

type Player struct {
	fileName string
	fh       io.ReadWriteCloser
	stream   beep.StreamSeekCloser
	format   beep.Format
	done     chan struct{}
}

func NewPlayer(file string) *Player {
	p := &Player{}
	p.fileName = file
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	p.fh = f
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	p.stream = streamer
	p.format = format
	p.done = make(chan struct{})
	return p
}

func (p *Player) Play() {
	speaker.Init(p.format.SampleRate, p.format.SampleRate.N(time.Second/10))
	speaker.Play(beep.Seq(p.stream, beep.Callback(func() {
		p.Done()
	})))
}

func (p *Player) Close() {
	p.stream.Close()
	p.fh.Close()
}

func (p *Player) Wait() {
	<-p.done
}

func (p *Player) Done() {
	p.done <- struct{}{}
}

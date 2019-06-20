package player

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

type Player struct {
	fileName string
	fh       io.ReadWriteCloser
	stream   beep.StreamSeekCloser
	format   beep.Format
	volume   *effects.Volume
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
	ctrl := &beep.Ctrl{Streamer: beep.Seq(p.stream, beep.Callback(func() {
		p.Done()
	})), Paused: false}

	p.volume = &effects.Volume{
		Streamer: ctrl,
		Base:     2,
		Volume:   0,
		Silent:   false,
	}

	speaker.Play(p.volume)
}

func (p *Player) FadeOut() {
	go func() {
		for i := float64(0); i > -100; i -= 0.2 {
			speaker.Lock()
			p.volume.Volume = i
			speaker.Unlock()
			<-time.After(time.Millisecond * 100)
		}
	}()
}

func (p *Player) Seek(pos int) error {
	speaker.Lock()
	err := p.stream.Seek(pos)
	speaker.Unlock()
	return err
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

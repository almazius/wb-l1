package main

import (
	"errors"
	"fmt"
	"log"
)

type MediaPlayer interface {
	Play(extension, path string) error
}

type ProMediaPlayer interface {
	PlayVideo(path string) error
	PlayAudio(path string) error
}

type FirstPlayer struct{}

func (p *FirstPlayer) PlayVideo(path string) error {
	return nil
}

func (p *FirstPlayer) PlayAudio(path string) error {
	return nil
}

type Adapter struct {
	player ProMediaPlayer
}

func (p *Adapter) Play(extension, path string) error {
	if extension == "mp4" {
		err := p.player.PlayVideo(path)
		if err != nil {
			log.Fatal(err)
		}

		err = p.player.PlayAudio(path)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		return errors.New(fmt.Sprintf("bad extension: %s", extension))
	}
	return nil
}

func main() {
	fmt.Println("The magic is somewhere higher")
}

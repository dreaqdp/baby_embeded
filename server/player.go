package main

import (
	"fmt"
	"github.com/vburenin/nsync"
	"log"
	"os/exec"
)

type Player struct {
	proc *exec.Cmd
	lock *nsync.TryMutex
}

func NewPlayer() *Player {
	return &Player{
		nil,
		nsync.NewTryMutex(),
	}
}

func (pl Player) TryPlay(path string) bool {
	if pl.lock.TryLock() {
		pl.lock.Unlock()
		pl.Play(path)
		return true
	}
	return false
}

func (pl Player) StopPlay() {
	if pl.proc != nil {
		if err := pl.proc.Process.Kill(); err != nil {
			log.Print("Failed to kill process")
		}
	}
}
func (pl Player) Play(path string) {
	pl.StopPlay()
	pl.proc = exec.Command("mpg123", path)
	if err := pl.proc.Start(); err != nil {
		log.Print("Failed to play music")
		log.Print("Error was: ", err.Error())
	}
}

func (pl Player) PlayLock(path string) {
	pl.lock.Lock()
	pl.StopPlay()
	pl.proc = exec.Command("mpg123", path)
	if err := pl.proc.Run(); err != nil {
		log.Print("Failed to play music")
		log.Print("Error was: ", err.Error())
	}
	pl.lock.Unlock()
}

func (pl Player) PlayYoutube(url string) {
	if pl.lock.TryLock() {
		pl.lock.Unlock()
		pl.StopPlay()
		c := fmt.Sprintf("rm -rf /tmp/s.mp3 && youtube-dl -x --audio-format mp3 %s -o /tmp/s.%%(ext)s && mpg123 /tmp/s.mp3", url)

		log.Print("Executing ", c)
		pl.proc = exec.Command("/bin/sh", "-c", c)
		if err := pl.proc.Start(); err != nil {
			log.Print("Failed to play from youtube")
			log.Print("Error was: ", err.Error())
		}
	}
}

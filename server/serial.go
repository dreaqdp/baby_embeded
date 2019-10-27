package main

import (
	"github.com/huin/goserial"
	"io"
	"log"
	"time"
)

func GetArduinoSerial() io.ReadWriteCloser {
	c := &goserial.Config{Name: "/dev/ttyACM0", Baud: 9600}
	s, err := goserial.OpenPort(c)
	if err != nil {
		log.Println("Error opening serial port")
		log.Printf("Error was: %s", err.Error())
		//		log.Fatal("adiosito")
	}

	time.Sleep(1 * time.Second)
	return s
}

type ArduinoManager struct {
	serial    io.ReadWriteCloser
	toSend    chan byte
	music     chan string
	toRequest chan byte
	player    *Player
}

func NewArduinoManager() *ArduinoManager {
	return &ArduinoManager{
		serial:    GetArduinoSerial(),
		toSend:    make(chan (byte)),
		music:     make(chan (string)),
		toRequest: make(chan (byte)),
		player:    NewPlayer(),
	}
}

func (m *ArduinoManager) attendReq() {
	for {
		select {
		case req := <-m.toRequest:
			if req == 'c' {
				m.player.PlayLock("assets/cry.mp3")
			} else {
				log.Printf("Error command %s was not understand", req)
			}
		case mus := <-m.music:
			m.player.PlayYoutube(mus)

		}
	}
}
func (m *ArduinoManager) run() {
	var readBuf []byte
	for {
		//write commands to arduino
		var bar []byte
		select {
		case command := <-m.toSend:
			log.Print(command)
			if _, err := m.serial.Write(append(bar, command)); err != nil {
				log.Printf("Error sending command %s", err.Error())
			}
		}

		//read commands from arduino
		if n, err := m.serial.Read(readBuf); err != nil {
			log.Print("Error reading from arduino")
			log.Printf("Error was: %s", err.Error())
		} else {
			for i := 0; i < n; i++ {
				m.toRequest <- readBuf[i]
			}
		}
	}
}

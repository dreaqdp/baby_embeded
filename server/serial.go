package main

import (
	"github.com/huin/goserial"
	"io"
	"log"
	"time"
)

//func sendArduinoCommand ( command byte, argument float32, serialPort io.ReadWriteCloser) error {
//	if serialPort == nil {
//		return nil
//	}
//
//	bufOut := new(bytes.Buffer)
//	err := binary.Write(bufOut, binary.LittleEndian, argument)
//	if err != nil {
//		return err
//	}
//	// Transmit command and argument down the pipe.
//	for _, v := range [][]byte{[]byte{command}, bufOut.Bytes()} {
//		log.Print(v)
//		_, err = serialPort.Write(v)
//		if err != nil {
//			return err
//		}
//	}
//
//	return nil
//}

func GetArduinoSerial() io.ReadWriteCloser {
	c := &goserial.Config{Name: "/dev/ttyACM0", Baud: 9600}
	s, err := goserial.OpenPort(c)
	if err != nil {
		log.Println("Error opening serial port")
		log.Printf("Error was: %s", err.Error())
		log.Fatal("adiosito")
	}

	time.Sleep(1 * time.Second)
	return s
}

type ArduinoManager struct {
	serial    io.ReadWriteCloser
	toSend    chan byte
	toRequest chan byte
}

func NewArduinoManager() *ArduinoManager {
	return &ArduinoManager{
		serial:    GetArduinoSerial(),
		toSend:    make(chan (byte)),
		toRequest: make(chan (byte)),
	}
}

func (m *ArduinoManager) attendReq() {
	for {
		select {
		case req := <-m.toRequest:
			if req == 'c' {
				log.Print("Colision was detected by ardu :(")
			} else {
				log.Printf("Error command %s was not understand", req)
			}
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

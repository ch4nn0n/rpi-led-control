package rpi_led_control

import (
	"os"
)

const (
	LED0Path string = "/sys/class/leds/led0/trigger"
	LED1Path string = "/sys/class/leds/led1/trigger"

	LED0Default string = "mmc0"
	LED1Default string = "default-on"
	LEDOff      string = "none"
)

func write(path string, value string) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if err != nil {
		return err
	}
	_, err = file.Write([]byte(value))
	if err != nil {
		return err
	}

	return nil
}

func Enable() error {
	err := write(LED0Path, LED0Default)
	if err != nil {
		return err
	}
	err = write(LED1Path, LED1Default)
	if err != nil {
		return err
	}
	return nil
}

func Disable() error {
	err := write(LED0Path, LEDOff)
	if err != nil {
		return err
	}
	err = write(LED1Path, LEDOff)
	if err != nil {
		return err
	}
	return nil
}

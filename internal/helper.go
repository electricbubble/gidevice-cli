package internal

import (
	"errors"
	giDevice "github.com/electricbubble/gidevice"
	"log"
	"os"
)

func GetDeviceFromCommand(udid string) (d giDevice.Device, err error) {
	if len(udid) == 0 {
		d, err = getFirstDevice()
	} else {
		d, err = getDevice(udid)
	}
	return
}

func ErrorExit(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(0)
	}
}

func IsDeveloper(d giDevice.Device) bool {
	imageSignatures, err := d.Images()
	ErrorExit(err)
	if len(imageSignatures) != 0 {
		return true
	}
	return false
}

func getFirstDevice() (d giDevice.Device, err error) {
	var devices []giDevice.Device
	if devices, err = _devices(); err != nil {
		return nil, err
	}

	if len(devices) == 0 {
		return nil, errors.New("no device found")
	}

	d = devices[0]
	return
}

func getDevice(udid string) (d giDevice.Device, err error) {
	var devices []giDevice.Device
	if devices, err = _devices(); err != nil {
		return nil, err
	}

	if len(devices) == 0 {
		return nil, errors.New("no device found")
	}

	for _, dev := range devices {
		if dev.Properties().SerialNumber == udid {
			return dev, nil
		}
	}
	return
}

func _devices() (devices []giDevice.Device, err error) {
	var usbmux giDevice.Usbmux
	if usbmux, err = giDevice.NewUsbmux(); err != nil {
		return nil, err
	}
	if devices, err = usbmux.Devices(); err != nil {
		return nil, err
	}
	return
}

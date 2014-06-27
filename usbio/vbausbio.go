// This is wrapper for vbausbio.dll(unknown version)

package usbio

import (
	"syscall"
	"unsafe"
	"errors"
)

var (
	dll *syscall.DLL
	uio_out *syscall.Proc
	uio_inp *syscall.Proc
	uio_find *syscall.Proc
	uio_free *syscall.Proc
	uio_getdevs *syscall.Proc
	uio_seldev *syscall.Proc
)

func Start() (err error) {
	dll, err = syscall.LoadDLL("vbausbio.dll")
	if err != nil {
		return err
	}

	uio_out, err = dll.FindProc("uio_out")
	if err != nil {
		return err
	}
	uio_inp, err = dll.FindProc("uio_inp")
	if err != nil {
		return err
	}
	uio_find, err = dll.FindProc("uio_find")
	if err != nil {
		return err
	}
	uio_free, err = dll.FindProc("uio_free")
	if err != nil {
		return err
	}
	uio_getdevs, err = dll.FindProc("uio_getdevs")
	if err != nil {
		return err
	}
	uio_seldev, err = dll.FindProc("uio_seldev")
	if err != nil {
		return err
	}

	return nil
}

func Set(port uint8, data uint16, pulse uint8) error {
	if uio_out == nil {
		return errors.New("call Start() first")
	}
	if port > 1 {
		return errors.New("invalid port number")
	}
	r1, _, _ := uio_out.Call(uintptr(port), uintptr(data), uintptr(pulse))
	if r1 != 0 {
		return errors.New("could not set register of USB-IO device")
	}
	return nil
}

func Get(port uint8, pulse uint8) (uint16, error) {
	var(
		data uint16
		p = unsafe.Pointer(&data)
	)

	if uio_inp == nil {
		return 0, errors.New("call Start() first")
	}
	r1, _, _ := uio_inp.Call(uintptr(p), uintptr(port), uintptr(pulse))
	if r1 == 1 {
		return 0, errors.New("could not send instraction")
	}
	if r1 == 2{
		return 0, errors.New("could not read signals")
	}
	return data, nil
}

func Find() error {
	if uio_find == nil {
		return errors.New("call Start() first")
	}

	r1, _, _ := uio_find.Call()
	if r1 == 1 {
		return errors.New("unavilable valid driver for USB-IO devices")
	}
	if r1 == 2 {
		return errors.New("could not find valid USB-IO devices")
	}
	return nil
}

func Free() error {
	if uio_free == nil {
		return errors.New("call Start() first")
	}
	uio_free.Call()
	return nil
}

func NDevices() (int32, error) {
	if uio_getdevs == nil {
		return 0, errors.New("call Start() first")
	}
	r1, _, _ := uio_getdevs.Call()
	return int32(r1), nil
}

func Select(n uint32) error {
	if uio_seldev == nil {
		return errors.New("call Start() first")
	}
	uio_seldev.Call(uintptr(n))
	return nil
}

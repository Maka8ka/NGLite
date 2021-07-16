package command

import (
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"syscall"

	"golang.org/x/text/encoding/simplifiedchinese"
)

func NewCommand() Commander {
	var cmd Commander

	switch runtime.GOOS {
	case "linux":
		cmd = NewLinuxCommand()
	case "windows":
		cmd = NewWindowsCommand()
	default:
		cmd = NewLinuxCommand()
	}

	return cmd
}

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

type Commander interface {
	Exec(args ...string) (int, string, error)
	ExecAsync(stdout chan string, args ...string) int
	ExecIgnoreResult(args ...string) error
}

func ConvertByte2String(byte []byte, charset Charset) string {
	var str string
	switch charset {
	case GB18030:
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}
	return str
}

type WindowsCommand struct {
}

func NewWindowsCommand() *WindowsCommand {
	return &WindowsCommand{}
}

func (lc *WindowsCommand) Exec(args ...string) (int, string, error) {
	args = append([]string{"/c"}, args...)
	cmd := exec.Command("cmd", args...)

	cmd.SysProcAttr = &syscall.SysProcAttr{}

	outpip, err := cmd.StdoutPipe()
	defer outpip.Close()

	if err != nil {
		return 0, "", err
	}

	err = cmd.Start()
	if err != nil {
		return 0, "", err
	}

	out, err := ioutil.ReadAll(outpip)
	if err != nil {
		return 0, "", err
	}
	cmdout := ConvertByte2String(out, "GB18030")
	return cmd.Process.Pid, string(cmdout), nil
}

func (lc *WindowsCommand) ExecAsync(stdout chan string, args ...string) int {
	var pidChan = make(chan int, 1)

	go func() {
		args = append([]string{"/c"}, args...)
		cmd := exec.Command("cmd", args...)

		cmd.SysProcAttr = &syscall.SysProcAttr{}

		outpip, err := cmd.StdoutPipe()
		defer outpip.Close()

		if err != nil {
			panic(err)
		}

		err = cmd.Start()
		if err != nil {
			panic(err)
		}

		pidChan <- cmd.Process.Pid

		out, err := ioutil.ReadAll(outpip)
		if err != nil {
			panic(err)
		}

		stdout <- string(out)
	}()

	return <-pidChan
}

func (lc *WindowsCommand) ExecIgnoreResult(args ...string) error {
	args = append([]string{"/c"}, args...)
	cmd := exec.Command("cmd", args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{}

	err := cmd.Run()

	return err
}

type LinuxCommand struct {
}

func NewLinuxCommand() *LinuxCommand {
	return &LinuxCommand{}
}

func (lc *LinuxCommand) Exec(args ...string) (int, string, error) {
	args = append([]string{"-c"}, args...)
	cmd := exec.Command(os.Getenv("SHELL"), args...)

	cmd.SysProcAttr = &syscall.SysProcAttr{}

	outpip, err := cmd.StdoutPipe()
	defer outpip.Close()

	if err != nil {
		return 0, "", err
	}

	err = cmd.Start()
	if err != nil {
		return 0, "", err
	}

	out, err := ioutil.ReadAll(outpip)
	if err != nil {
		return 0, "", err
	}

	return cmd.Process.Pid, string(out), nil
}

func (lc *LinuxCommand) ExecAsync(stdout chan string, args ...string) int {
	var pidChan = make(chan int, 1)

	go func() {
		args = append([]string{"-c"}, args...)
		cmd := exec.Command(os.Getenv("SHELL"), args...)

		cmd.SysProcAttr = &syscall.SysProcAttr{}

		outpip, err := cmd.StdoutPipe()
		defer outpip.Close()

		if err != nil {
			panic(err)
		}

		err = cmd.Start()
		if err != nil {
			panic(err)
		}

		pidChan <- cmd.Process.Pid

		out, err := ioutil.ReadAll(outpip)
		if err != nil {
			panic(err)
		}

		stdout <- string(out)
	}()

	return <-pidChan
}

func (lc *LinuxCommand) ExecIgnoreResult(args ...string) error {

	args = append([]string{"-c"}, args...)
	cmd := exec.Command(os.Getenv("SHELL"), args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{}

	err := cmd.Run()

	return err
}

package main

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/shirou/gopsutil/v3/process"
)

type IInstruction interface {
	DoSomething() error
}

type ChangeDirectoryCommand struct {
	Directory string
}

func (c *ChangeDirectoryCommand) DoSomething() error {
	return os.Chdir(c.Directory)
}

type PwdCommand struct{}

func (c *PwdCommand) DoSomething() error {
	path, err := c.ShowCurrentDirectory()
	io.WriteString(os.Stdout, path)
	return err
}
func (c *PwdCommand) ShowCurrentDirectory() (string, error) {
	return os.Getwd()
}

type EchoCommand struct {
	Args []string
}

func (c *EchoCommand) DoSomething() error {
	for _, str := range c.Args {
		_, err := io.WriteString(os.Stdout, str+" ")
		if err != nil {
			return err
		}
	}
	return nil
}

type KillCommand struct {
	ProcessID int
}

func (c *KillCommand) DoSomething() error {
	return c.KillProcess(c.ProcessID)
}
func (c *KillCommand) KillProcess(pid int) error {
	handle, err := os.FindProcess(c.ProcessID)
	if err != nil {
		return errors.New("invalid processID!")
	}

	err = handle.Kill()

	if err != nil {
		return fmt.Errorf("could not kill a process with ID = %d", c.ProcessID)
	}
	return nil
}

type PsCommand struct{}

func (c *PsCommand) DoSomething() error {
	return c.ShowInfo()
}
func (c *PsCommand) ShowInfo() error {
	processes, err := process.Processes()

	if err != nil {
		return errors.New("unable to get processes info")
	}

	for _, process := range processes {
		name, _ := process.Name()
		id := strconv.Itoa((int)(process.Pid))
		cpu, _ := process.MemoryPercent()
		memory := strconv.FormatFloat(float64(cpu), 'f', 2, 64) + "% CPU"

		io.WriteString(os.Stdout, strings.Join([]string{id, name, memory}, "\t\t\t\t"))
		io.WriteString(os.Stdout, "\n")
	}

	return nil
}

type ExecCommand struct {
	Filepath string
	Args     []string
}

func (c *ExecCommand) DoSomething() error {
	return c.execProcess(c.Filepath, c.Args...)
}
func (c *ExecCommand) execProcess(filepath string, args ...string) error {
	handle := exec.Command(filepath, args...)
	if err := handle.Run(); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

type ForkCommand struct {
	ProcessID int
}

func (c *ForkCommand) DoSomething() error {
	return c.forkProcess(c.ProcessID)
}
func (c *ForkCommand) forkProcess(pid int) error {
	cmd := exec.Command("cmd", "/c", "start", strconv.Itoa(pid))
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func parseStringToInstruction(args []string) IInstruction {
	if len(args) == 0 {
		return nil
	}

	switch args[0] {
	case "cd":
		return &ChangeDirectoryCommand{Directory: args[1]}
	case "pwd":
		return &PwdCommand{}
	case "echo":
		return &EchoCommand{Args: args[1:]}
	case "kill":
		if pid, err := strconv.Atoi(args[1]); err == nil {
			return &KillCommand{ProcessID: pid}
		}
	case "ps":
		return &PsCommand{}
	case "exec":
		return &ExecCommand{Filepath: args[1], Args: args[2:]}
	case "fork":
		if pid, err := strconv.Atoi(args[1]); err == nil {
			return &ForkCommand{ProcessID: pid}
		}
	}

	return nil
}

func main() {
	io.WriteString(os.Stdout, "Welcome to my UNIX-shell! Please type some commands!\n")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		line := scanner.Text()
		commands := strings.Split(line, "|")

		for _, commandLine := range commands {
			if len(commandLine) == 0 {
				continue
			}

			args := strings.Split(strings.TrimSpace(commandLine), " ")
			instruction := parseStringToInstruction(args)
			instruction.DoSomething()
		}
	}
}

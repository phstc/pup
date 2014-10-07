package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

type Config struct {
	Apps []App `json:"apps"`
}

type App struct {
	Name    string `json:"name"`
	Cmd     string `json:"cmd"`
	Chdir   string `json:"chdir"`
	PidFile string `json:"pidfile"`
	Out     string `json:"out"`
	Err     string `json:"err"`
}

func (app App) CmdName() string {
	cmdargs := strings.Fields(app.Cmd)
	return cmdargs[0]
}

func (app *App) CmdArgs() []string {
	cmdargs := strings.Fields(app.Cmd)
	return cmdargs[1:len(cmdargs)]
}

func writePid(app App, pid int) {
	w, _ := os.Create(app.PidFile)
	defer w.Close()
	if _, err := w.WriteString(strconv.Itoa(pid)); err != nil {
		fmt.Printf("Failed to write: %s, message: %s\n", app.PidFile, err)
	}
	w.Sync()
}

func config() Config {
	file, _ := os.Open("/Users/pablo/pup.json")
	decoder := json.NewDecoder(file)
	config := Config{}
	err := decoder.Decode(&config)
	if err != nil {
		panic(err)
	}

	return config
}

func start(app App) {
	cmd := exec.Command(app.CmdName(), app.CmdArgs()...)
	cmd.Dir = app.Chdir

	// detach the process
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	cmd.Start()

	fmt.Printf("starting %s with pid %d\n", app.Name, cmd.Process.Pid)
	writePid(app, cmd.Process.Pid)

	cmd.Process.Release()
}

func stop(app App) {
	pid := pid(app)
	pgid, err := syscall.Getpgid(pid)
	if err == nil {
		// err := syscall.Kill(-pgid, syscall.SIGTERM)
		err := syscall.Kill(-pgid, syscall.SIGINT)
		if err == nil {
			fmt.Printf("stopping %s with pid %d\n", app.Name, pid)
		} else {
			fmt.Printf("%s %s\n", app.Name, err)
		}
	} else {
		fmt.Printf("%s %s\n", app.Name, err)
	}
}

func status(app App) {
	pid := pid(app)
	if err := syscall.Kill(pid, 0); err == nil {
		fmt.Printf("%s is running with pid %d\n", app.Name, pid)
	} else {
		fmt.Printf("%s %s\n", app.Name, err)
	}
}

func pid(app App) int {
	p, _ := ioutil.ReadFile(app.PidFile)
	pid, _ := strconv.Atoi(string(p))
	return pid
}

func printUsage() {
	fmt.Println("usage: pup (start|stop|restart|status) app_name")
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	var appName string

	cmd := os.Args[1]

	if len(os.Args) > 2 {
		appName = os.Args[2]
	}

	count := 0

Parser:
	for _, app := range config().Apps {
		if appName != "" && appName != app.Name {
			continue
		}

		count++

		switch cmd {
		case "start":
			start(app)
		case "stop":
			stop(app)
		case "status":
			status(app)
		case "restart":
			start(app)
			stop(app)
		default:
			printUsage()
			break Parser
		}
	}

	if appName != "" && count == 0 {
		fmt.Printf("%s not found\n", appName)
	}
}
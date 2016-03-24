package main

import (
	"github.com/chzyer/readline"
	"runtime"
	"flag"
	"fmt"
	"os"
)

func main() {
	runtime.GOMAXPROCS(1)
	var configFile string
	flag.StringVar(&configFile, "c", "riline.conf", "config file help")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s -c riline.conf\n", os.Args[0])
	}
	if configFile == "" {
		flag.Usage()
		os.Exit(1)
	}
	flag.Parse()

	rl, err := readline.NewEx(&readline.Config{
		Prompt:          "\033[31m»\033[0m ",
		HistoryFile:     "/tmp/readline.tmp",
		InterruptPrompt: "\nInterrupt, Press Ctrl+D to exit",
		EOFPrompt:       "exit",
	})
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil {
			// io.EOF, readline.ErrInterrupt
			break
		}
		println(line)
	}
}

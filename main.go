package main

import (
    "flag"
    "fmt"
	"os"
	"strings"
)

func main() {
	initCmd := flag.NewFlagSet("init", flag.ExitOnError)
	initDryRun := initCmd.Bool("dry-run", false, "Don't write config file")
	initDir := initCmd.String("dir", "~/.aws", "AWS Directory")
    initBaseProfile := initCmd.String("base", "default", "Base profile")

	envCmd := flag.NewFlagSet("env", flag.ExitOnError)

	if len(os.Args) < 2 {
        fmt.Println("Expected 'init' or 'set' subcommands")
        os.Exit(1)
    }

	switch os.Args[1] {
	case "init":
		initCmd.Parse(os.Args[2:])
		args := initCmd.Args()
		doInit(*initDryRun, *initDir, *initBaseProfile, args)
	case "env":
		envCmd.Parse(os.Args[2:])
		for _, e := range os.Environ() {
			pair := strings.SplitN(e, "=", 2)
			fmt.Println(pair[0], pair[1])
		}
	default:
		fmt.Println("Expected 'init' or 'set' subcommands")
		os.Exit(1)
	}

}

func check(e error) {
    if e != nil {
        panic(e)
    }
}


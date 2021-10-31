package main

import (
    "flag"
    "fmt"
	"os"
)

func main() {
	initCmd := flag.NewFlagSet("init", flag.ExitOnError)
	initDryRun := initCmd.Bool("dry-run", false, "Don't write config file")
	initDir := initCmd.String("dir", "~/.aws", "AWS Directory")
    initBaseProfile := initCmd.String("base", "default", "Base profile")

	if len(os.Args) < 2 {
        fmt.Println("Expected 'init' or 'set' subcommands")
        os.Exit(1)
    }

	switch os.Args[1] {
		case "init":
			initCmd.Parse(os.Args[2:])
			args := initCmd.Args()
			doInit(*initDryRun, *initDir, *initBaseProfile, args)
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


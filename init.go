package main

import (
    "fmt"
	"strings"
	"gopkg.in/ini.v1"
)

func doInit(initDryRun bool, initDir string, initBaseProfile string, args []string) {
	fmt.Println("subcommand 'init'")
	fmt.Println("  dry-run:", initDryRun)
	fmt.Println("  dir:", initDir)
	fmt.Println("  base:", initBaseProfile)
	fmt.Println("  tail:", args)

	aws_access_key_id, aws_secret_access_key := readCredentials(initDir, initBaseProfile)

	fmt.Println(aws_access_key_id, aws_secret_access_key)

	childProfiles := findChildProfiles(initDir, initBaseProfile)
	for j := 0; j < len(childProfiles); j++ {
		fmt.Println(childProfiles[j])
	}
}

func readCredentials(initDir string, initBaseProfile string) (string, string) {
	credentials, err := ini.Load(initDir + "/credentials")
	check(err)

	profile, err := credentials.GetSection(initBaseProfile)
	check(err)
	aws_access_key_id, err := profile.GetKey("aws_access_key_id")
	check(err)
	aws_secret_access_key, err := profile.GetKey("aws_secret_access_key")
	check(err)

	return aws_access_key_id.String(), aws_secret_access_key.String()
}

func findChildProfiles(initDir string, initBaseProfile string) ([]string) {
	profileFile, err := ini.Load(initDir + "/config")
	check(err)

	names := profileFile.SectionStrings()
	childProfiles := make([]string, 0)
    for j := 0; j < len(names); j++ {
		profile, err := profileFile.GetSection(names[j])
		check(err)
		source_profile, err := profile.GetKey("source_profile")
		if err == nil && source_profile.String() == initBaseProfile {
			childProfiles = append(childProfiles, strings.Split(names[j], " ")[1])
		}
    }
	
	return childProfiles
}
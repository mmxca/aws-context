package main

import (
	"encoding/json"
	"strings"
	"gopkg.in/ini.v1"
	"fmt"
)

func doInit(initDryRun bool, initDir string, initBaseProfile string, args []string) {

	aws_access_key_id, aws_secret_access_key := initReadCredentials(initDir, initBaseProfile)
	mfa_serial := initGetMfaSerial(initDir, initBaseProfile)
	childProfiles := initFindChildProfiles(initDir, initBaseProfile)

	profiles := make([]Profile, 0)
	for j := 0; j < len(childProfiles); j++ {
		role_arn := initGetRoleArn(initDir, childProfiles[j])
		profile := &Profile{
			Name:   		childProfiles[j],
			RoleArn: 		role_arn}
		profiles = append(profiles, *profile)
	}

    config := &Config{
        AccessKeyId:   		aws_access_key_id,
        SecretAccessKey: 	aws_secret_access_key,
		MfaSerial:			mfa_serial,
		Profiles:			profiles}
    config_json, _ := json.Marshal(config)

	json_WriteConfig(initDir, initBaseProfile, config_json)

	test := json_ReadConfig(initDir, initBaseProfile)

	fmt.Println(test.AccessKeyId)
}


func initReadCredentials(initDir string, initBaseProfile string) (string, string) {
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

func initFindChildProfiles(initDir string, initBaseProfile string) ([]string) {
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

func initGetRoleArn(initDir string, profileName string) (string) {
	profileFile, err := ini.Load(initDir + "/config")
	check(err)

	profile, err := profileFile.GetSection("profile " + profileName)
	check(err)

	role_arn, err := profile.GetKey("role_arn")
	check(err)

	return role_arn.String()
}

func initGetMfaSerial(initDir string, initBaseProfile string) (string) {
	profileFile, err := ini.Load(initDir + "/config")
	check(err)

	profile, err := profileFile.GetSection("profile " + initBaseProfile)
	check(err)

	mfa_serial, err := profile.GetKey("mfa_serial")
	check(err)

	return mfa_serial.String()
}
package main

import (
	// "fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

func doList(configDir string, baseProfile string, args []string) {
	configs := json_ReadConfig(configDir, baseProfile)

	profiles := configs.Profiles
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Profile", "Active", "Expires"})

	counter := 0
	for _, profile := range profiles {
		counter = counter + 1
		assumedRoleCredentials =
		t.AppendRow(table.Row{counter, profile.Name, profile.AssumedRole.Credentials.AccessKeyId, profile.AssumedRole.Credentials.Expiration})
	}

	// t.AppendRows([]table.Row{
	//     {1, "Arya", "Stark", 3000},
	//     {20, "Jon", "Snow", 2000, "You know nothing, Jon Snow!"},
	// })
	// t.AppendSeparator()
	// t.AppendRow([]interface{}{300, "Tyrion", "Lannister", 5000})
	// t.AppendFooter(table.Row{"", "", "Total", 10000})
	t.Render()
}

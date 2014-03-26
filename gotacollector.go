package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/mattbaird/gota-api/api"
	"os"
	"strings"
)

func main() {
	address := ""
	commands := []cli.Command{
		{
			Name:        "gota",
			ShortName:   "g",
			Usage:       "./gotacollector collect",
			Description: "Collect latest DOTA stats from API",
			Flags: []cli.Flag{
				cli.StringFlag{"key", "", "DOTA API Key"},
				cli.StringFlag{"command", "", "one of [GetMatchHistory, GetMatchDetails, GetPlayerSummaries, GetLeagueListing, GetLiveLeagueGames, GetTeamInfoByTeamID, GetHeroes]"},
				// the following are optional
				cli.StringFlag{"startAt", "", ""},
			},
			Action: func(c *cli.Context) {
				startAt := strings.TrimSpace(c.String("startAt"))
				gotaApi, err := api.NewGotaAPI("", "en_us")
				if err != nil {
					fmt.Printf("Problem: %v", err)
					os.Exit(1)
				}
				//gotaApi.GetMatchHistory("", -1, "", "", "", "", -1, -1, "", -1, 1, false)
				//gotaApi.GetMatchDetails(580670098)
				gotaApi.GetHeroes()
				fmt.Printf("startAt:%v", startAt)
				fmt.Printf("url:%v", gotaApi.URL())
				fmt.Println("FINISHED:SUCCESS")
			},
		},
	}
	app := cli.NewApp()
	app.Commands = commands
	app.Name = "gota"
	app.Version = "0.0.3"

	app.Action = func(ctx *cli.Context) {
		if len(ctx.Args()) == 0 {
			cli.ShowAppHelp(ctx)
			os.Exit(1)
		}
		address = ctx.Args()[0]
		console := cli.NewApp()
		console.Commands = commands
		console.Action = func(c *cli.Context) {
			fmt.Println("Command not found. Type 'help' for a list of commands.")
		}
	}
	app.Run(os.Args)
	os.Exit(0)
}

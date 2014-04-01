package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/mattbaird/gota-api/api"
	"os"
	//	"strings"
)

func main() {
	address := ""
	commands := []cli.Command{
		{
			Name:        "GetMatchHistory",
			ShortName:   "mh",
			Usage:       "./gota-api mh [playerName] [heroId] [gameMode] [skill] [dateMin] [dateMax] [minPlayers] [accountId] [leagueId] [matchesRequested] [tournamentGamesOnly]",
			Description: "Request Match history",
			Flags: []cli.Flag{
				cli.IntFlag{"startAt", -1, "The match ID to start retrieval. Optional. Example 27110133"},
				// the following are optional
				cli.StringFlag{"playerName", "", ""},
				cli.IntFlag{"heroId", -1, ""},
				cli.StringFlag{"gameMode", "", ""},
				cli.StringFlag{"skill", "", ""},
				cli.StringFlag{"dateMin", "", ""},
				cli.StringFlag{"dateMax", "", ""},
				cli.IntFlag{"minPlayers", -1, ""},
				cli.IntFlag{"accountId", -1, ""},
				cli.StringFlag{"leagueId", "", ""},
				cli.IntFlag{"matchesRequested", -1, ""},
				cli.BoolFlag{"tournamentGamesOnly", "false"},
			},
			Action: func(c *cli.Context) {
				startAtMatchId := c.Int("startAt")
				playerName := c.String("playerName")
				heroId := c.Int("heroId")
				gameMode := c.String("gameMode")
				skill := c.String("skill")
				dateMin := c.String("dateMin")
				dateMax := c.String("dateMax")
				minPlayers := c.Int("minPlayers")
				accountId := c.Int("accountId")
				leagueId := c.String("leagueId")
				matchesRequested := c.Int("matchesRequested")
				tournamentGamesOnly := c.Bool("tournamentGamesOnly")
				gotaApi, err := api.NewGotaAPI("", "en_us")
				if err != nil {
					fmt.Printf("Problem: %v", err)
					os.Exit(1)
				}
				matchHistory, err := gotaApi.GetMatchHistory(playerName, heroId, gameMode, skill, dateMin, dateMax,
					minPlayers, accountId, leagueId, startAtMatchId, matchesRequested, tournamentGamesOnly)
				if err != nil {
					fmt.Printf("Error:%v", err)
					os.Exit(1)
				}
				fmt.Printf("Number of results: %v", matchHistory.NumberOfResults)
			},
		},
		{
			Name:        "GetMatchDetails",
			ShortName:   "md",
			Usage:       "./gota-api md [matchId]",
			Description: "Request Match Details",
			Flags: []cli.Flag{
				cli.IntFlag{"matchId", -1, "The match ID to get details for. Example 27110133"},
			},
			Action: func(c *cli.Context) {
				matchId := c.Int("matchId")
				gotaApi, err := api.NewGotaAPI("", "en_us")
				if err != nil {
					fmt.Printf("Problem: %v", err)
					os.Exit(1)
				}
				matchDetails, err := gotaApi.GetMatchDetails(matchId)
				if err != nil {
					fmt.Printf("Error:%v", err)
					os.Exit(1)
				}
				fmt.Printf("Duration: %v", matchDetails.Duration)
			},
		},
		{
			Name:        "GetMatchHistoryBySequenceNum",
			ShortName:   "mhs",
			Usage:       "./gota-api mhs [startAtMatchSeqNum] [matchesRequested]",
			Description: "Get heroes",
			Flags: []cli.Flag{
				cli.IntFlag{"startAtMatchSeqNum", 1, "The match ID to start at. Example 27110133"},
				cli.IntFlag{"matchesRequested", -1, "The number of matches to get"},
			},
			Action: func(c *cli.Context) {
				startAtMatchSeqNum := c.Int("startAtMatchSeqNum")
				matchesRequested := c.Int("matchesRequested")
				gotaApi, err := api.NewGotaAPI("", "en_us")
				if err != nil {
					fmt.Printf("Problem: %v", err)
					os.Exit(1)
				}
				matchHistory, err := gotaApi.GetMatchHistoryBySequenceNum(startAtMatchSeqNum, matchesRequested)
				if err != nil {
					fmt.Printf("Error:%v", err)
					os.Exit(1)
				}
				fmt.Printf("matchHistory: %v", matchHistory)
			},
		},
		{
			Name:        "GetHeroes",
			ShortName:   "h",
			Usage:       "./gota-api h",
			Description: "Get heroes",
			Flags:       []cli.Flag{},
			Action: func(c *cli.Context) {
				gotaApi, err := api.NewGotaAPI("", "en_us")
				if err != nil {
					fmt.Printf("Problem: %v", err)
					os.Exit(1)
				}
				heroes, err := gotaApi.GetHeroes()
				if err != nil {
					fmt.Printf("Error:%v", err)
					os.Exit(1)
				}
				fmt.Printf("# Heroes: %v", heroes.Count)
			},
		},
		{
			Name:        "GetLeagueListing",
			ShortName:   "ll",
			Usage:       "./gota-api ll",
			Description: "Get League Listings",
			Flags:       []cli.Flag{},
			Action: func(c *cli.Context) {
				gotaApi, err := api.NewGotaAPI("", "en_us")
				if err != nil {
					fmt.Printf("Problem: %v", err)
					os.Exit(1)
				}
				leagueListings, err := gotaApi.GetLeagueListing()
				if err != nil {
					fmt.Printf("Error:%v", err)
					os.Exit(1)
				}
				fmt.Printf("# leagueListings: %v", len(leagueListings.Leagues))
			},
		},
		{
			Name:        "GetLiveLeagueGames",
			ShortName:   "llg",
			Usage:       "./gota-api llg",
			Description: "Get live league games",
			Flags:       []cli.Flag{},
			Action: func(c *cli.Context) {
				gotaApi, err := api.NewGotaAPI("", "en_us")
				if err != nil {
					fmt.Printf("Problem: %v", err)
					os.Exit(1)
				}
				liveLeagueGames, err := gotaApi.GetLiveLeagueGames()
				if err != nil {
					fmt.Printf("Error:%v", err)
					os.Exit(1)
				}
				fmt.Printf("liveLeagueGames: %v", liveLeagueGames)
			},
		},
		{
			Name:        "GetTeamInfoByTeamID",
			ShortName:   "ti",
			Usage:       "./gota-api ti [startAtTeamId] [teamsRequested]",
			Description: "Get Team Information by ID",
			Flags: []cli.Flag{
				cli.IntFlag{"startAtTeamId", -1, "Which Team ID to start at."},
				cli.IntFlag{"teamsRequested", -1, "Number of teams to retrieve"},
			},
			Action: func(c *cli.Context) {
				startAtTeamId := c.Int("startAtTeamId")
				teamsRequested := c.Int("teamsRequested")
				gotaApi, err := api.NewGotaAPI("", "en_us")
				if err != nil {
					fmt.Printf("Problem: %v", err)
					os.Exit(1)
				}
				team, err := gotaApi.GetTeamInfoByTeamID(startAtTeamId, teamsRequested)
				if err != nil {
					fmt.Printf("Error:%v", err)
					os.Exit(1)
				}
				fmt.Printf("team: %v", team)
			},
		},
		{
			Name:        "GetPlayerSummaries",
			ShortName:   "ps",
			Usage:       "./gota-api ps [ids]",
			Description: "Get Play Summaries",
			Flags: []cli.Flag{
				cli.IntSliceFlag{"ids", nil, "Comma separated array of user ids"},
			},
			Action: func(c *cli.Context) {
				ids := c.IntSlice("ids")
				gotaApi, err := api.NewGotaAPI("", "en_us")
				if err != nil {
					fmt.Printf("Problem: %v", err)
					os.Exit(1)
				}
				playerSummaries, err := gotaApi.GetPlayerSummaries(ids...)
				if err != nil {
					fmt.Printf("Error:%v", err)
					os.Exit(1)
				}
				fmt.Printf("playerSummaries: %v", playerSummaries)
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

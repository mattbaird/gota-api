package main

import (
	"bufio"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/mattbaird/gota-api/api"
	"os"
	"text/tabwriter"
)

func main() {
	address := ""
	commands := []cli.Command{
		{
			Name:        "WriteMatchDetails",
			ShortName:   "wmd",
			Usage:       "./gotawriter wmd --file [file]",
			Description: "Write a hive compatible tab delimited file",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "file", Value: "", Usage: ""},
			},
			Action: func(c *cli.Context) {
				f := c.String("file")
				gotaApi, err := api.NewGotaAPI("", "en_us")
				if err != nil {
					fmt.Printf("Error:%v\n", err)
				}
				matchDetails, err := gotaApi.GetMatchDetails(834245356)
				writeFile(f, []api.MatchDetail{matchDetails})
				fmt.Printf("File Written to %s\n", f)
			},
		},
	}
	app := cli.NewApp()
	app.Commands = commands
	app.Name = "gota-writer"
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
func writeFile(fName string, matches []api.MatchDetail) {

	f, err := os.Create(fName)

	if err != nil {
		panic(err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	defer w.Flush()
	// New Tsv writer
	writer := new(tabwriter.Writer)
	defer writer.Flush()
	writer.Init(w, 1, 1, 1, ' ', 0)
	// Headers
	writeHeader := func(RadiantWin, Duration, StartTime, MatchId, SequenceNumber, RadiantTowerStatus,
		DireTowerStatus, RadiantBarracksStatus, DireBarracksStatus, Cluster, FirstBloodTime, HumanPlayers,
		LobbyType, LeagueId, PositiveVotes, NegativeVotes, GameMode, RadiantCaptain, DireCaptain string) {
		fmt.Fprintf(writer, "%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\n",
			RadiantWin, Duration, StartTime, MatchId, SequenceNumber, RadiantTowerStatus,
			DireTowerStatus, RadiantBarracksStatus, DireBarracksStatus, Cluster, FirstBloodTime, HumanPlayers,
			LobbyType, LeagueId, PositiveVotes, NegativeVotes, GameMode, RadiantCaptain, DireCaptain)
	}
	write := func(RadiantWin bool, Duration, StartTime, MatchId, SequenceNumber, RadiantTowerStatus,
		DireTowerStatus, RadiantBarracksStatus, DireBarracksStatus, Cluster, FirstBloodTime, HumanPlayers,
		LobbyType, LeagueId, PositiveVotes, NegativeVotes, GameMode, RadiantCaptain, DireCaptain int) {
		fmt.Fprintf(writer, "%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\n",
			RadiantWin, Duration, StartTime, MatchId, SequenceNumber, RadiantTowerStatus,
			DireTowerStatus, RadiantBarracksStatus, DireBarracksStatus, Cluster, FirstBloodTime, HumanPlayers,
			LobbyType, LeagueId, PositiveVotes, NegativeVotes, GameMode, RadiantCaptain, DireCaptain)
	}
	writeHeader("radiant_win", "duration", "start_time", "match_id", "match_seq_num", "tower_status_radiant",
		"tower_status_dire", "barracks_status_radiant", "barracks_status_dire", "cluster", "first_blood_time", "lobby_type", "human_players",
		"leagueid", "positive_votes", "negative_votes", "game_mode", "radiant_captain", "dire_captain")
	for _, match := range matches {
		write(match.RadiantWin, match.Duration, match.StartTime, match.MatchId, match.SequenceNumber,
			match.RadiantTowerStatus, match.DireTowerStatus, match.RadiantBarracksStatus, match.DireBarracksStatus,
			match.Cluster, match.FirstBloodTime, match.LobbyType, match.HumanPlayers, match.LeagueId, match.PositiveVotes,
			match.NegativeVotes, match.GameMode, match.RadiantCaptain, match.DireCaptain)
	}
}

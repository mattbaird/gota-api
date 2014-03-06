package main

import (
	"encoding/json"
	"io/ioutil"
	//	"log"
	"net/http"
)

type MatchResult struct {
	Result MatchHistory `json:"result"`
}

type MatchHistory struct {
	Status           int     `json:"status"`
	NumberOfResults  int     `json:"num_results"`
	TotalResults     int     `json:"total_results"`
	ResultsRemaining int     `json:"results_remaining"`
	Matches          []Match `json:"matches"`
}

type Match struct {
	Id             int      `json:"match_id"`
	SequenceNumber int      `json:"match_seq_num"`
	StartTime      int      `json:"start_time"`
	LobbyType      int      `json:"lobby_type"`
	Players        []Player `json:"players"`
}

type Player struct {
	Id         int `json:"account_id"`
	PlayerSlot int `json:"player_slot"`
	HeroId     int `json:"hero_id"`
}

type MatchDetailResult struct {
	Result MatchDetail `json:"result"`
}

type MatchDetail struct {
	RadiantWin            bool           `json:"radiant_win"`
	Duration              int            `json:"duration"`
	StartTime             int            `json:"start_time"`
	MatchId               int            `json:"match_id"`
	SequenceNumber        int            `json:"match_seq_num"`
	RadiantTowerStatus    int            `json:"tower_status_radiant"`
	DireTowerStatus       int            `json:"tower_status_dire"`
	RadiantBarracksStatus int            `json:"barracks_status_radiant"`
	DireBarracksStatus    int            `json:"barracks_status_dire"`
	Cluster               int            `json:"cluster"`
	FirstBloodTime        int            `json:"first_blood_time"`
	HumanPlayers          int            `json:"human_players"`
	LobbyTYpe             int            `json:"lobby_type"`
	LeagueId              int            `json:"leagueid"`
	PositiveVotes         int            `json:"positive_votes"`
	NegativeVotes         int            `json:"negative_votes"`
	GameMode              int            `json:"game_mode"`
	Players               []PlayerDetail `json:"players"`
}

type PlayerDetail struct {
	Id           int `json:"account_id"`
	PlayerSlot   int `json:"player_slot"`
	HeroId       int `json:"hero_id"`
	Item0        int `json:"item_0"`
	Item1        int `json:"item_1"`
	Item2        int `json:"item_2"`
	Item3        int `json:"item_3"`
	Item4        int `json:"item_4"`
	Item5        int `json:"item_5"`
	Kills        int `json:"kills"`
	Deaths       int `json:"deaths"`
	Assists      int `json:"assists"`
	LeaverStatus int `json:"leaver_status"`
	Gold         int `json:"gold"`
	LastHits     int `json:"last_hits"`
	Denies       int `json:"denies"`
	GoldPerMin   int `json:"gold_per_min"`
	XpPerMin     int `json:"xp_per_min"`
	GoldSpent    int `json:"gold_spent"`
	HeroDamage   int `json:"hero_damage"`
	TowerDamage  int `json:"tower_damage"`
	HeroHealing  int `json:"hero_healing"`
	Level        int `json:"level"`
}

func ReadMatchHistory(requestUrl string) (MatchHistory, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", requestUrl, nil)
	resp, err := client.Do(req)
	if err != nil {
		return MatchHistory{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return MatchHistory{}, err
	}
	var retval MatchResult
	err = json.Unmarshal(body, &retval)
	return retval.Result, err
}

func ReadMatchDetails(requestUrl string) (MatchDetail, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", requestUrl, nil)
	resp, err := client.Do(req)
	if err != nil {
		return MatchDetail{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return MatchDetail{}, err
	}
	var retval MatchDetailResult
	err = json.Unmarshal(body, &retval)
	return retval.Result, err
}

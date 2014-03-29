package api

import (
	"encoding/json"
	"fmt"
	"strings"
)

type APIError struct {
	Err  string `json:"error"`
	Code int
}

func (e APIError) Error() string {
	return fmt.Sprintf("%v: %v", e.Code, e.Err)
}

type Result struct {
	Result json.RawMessage `json:"result"`
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

type Heroes struct {
	Heroes []Hero `json:"heroes"`
	Count  int    `json:"count"`
}

type Hero struct {
	Name          string `json:"name"`
	Id            int    `json:"id"`
	LocalizedName string `json:"localized_name,omitempty"`
}

func (h *Hero) GetFriendlyName() string {
	//npc_dota_hero_earth_spirit
	temp := h.Name[14:]
	temp = strings.Replace(temp, "_", " ", -1)
	return temp
}

type Leagues struct {
	Leagues []League `json:"leagues"`
}

type League struct {
	Name          string `json:"name"`
	Id            int    `json:"leagueid"`
	Description   string `json:"description"`
	TournamentUrl string `json:"tournament_url"`
	ItemDef       int    `json:"itemdef,omitempty"`
}

type LeagueGames struct {
	Games []Game `json:"games"`
}

type Game struct {
	Players    []LeaguePlayer `json:"players"`
	Radiant    Team           `json:"radiant_team"`
	Dire       Team           `json:"dire_team"`
	LobbyId    int64          `json:"lobby_id"`
	Spectators int            `json:"spectators"`
	TowerState uint           `json:"tower_state"`
	LeagueId   int            `json:"league_id"`
}

type LeaguePlayer struct {
	AccountId int    `json:"account_id"`
	Name      string `json:"name"`
	HeroId    int    `json:"hero_id"`
	Team      int    `json:"team"`
}
type Teams struct {
	Status int    `json:"status"`
	Teams  []Team `json:"teams"`
}
type Team struct {
	Name                         string `json:"team_name"`
	Id                           int    `json:"team_id"`
	Logo                         int64  `json:"team_logo"`
	Complete                     bool   `json:"complete,omitempty"`
	Tag                          string `json:"tag,omitempty"`
	CreatedAt                    int64  `json:"time_created,omitempty"`
	Rating                       string `json:"rating,omitempty"`
	Sponsor                      int64  `json:"logo_sponsor,omitempty"`
	CountryCode                  string `json:"country_code,omitempty"`
	Url                          string `json:"url,omitempty"`
	GamesPlayedWithCurrentRoster int    `json:"games_played_with_current_roster,omitempty"`
	Player1Id                    int    `json:"player_1_account_id,omitempty"`
	Player2Id                    int    `json:"player_2_account_id,omitempty"`
	Player3Id                    int    `json:"player_3_account_id,omitempty"`
	Player4Id                    int    `json:"player_4_account_id,omitempty"`
	Player5Id                    int    `json:"player_5_account_id,omitempty"`
	AdminAccountId               int    `json:"admin_account_id,omitempty"`
	LeagueId0                    int    `json:"league_id_0,omitempty"`
	LeagueId1                    int    `json:"league_id_1,omitempty"`
	LeagueId2                    int    `json:"league_id_2,omitempty"`
	LeagueId3                    int    `json:"league_id_3,omitempty"`
	LeagueId4                    int    `json:"league_id_4,omitempty"`
	LeagueId5                    int    `json:"league_id_5,omitempty"`
	LeagueId6                    int    `json:"league_id_6,omitempty"`
	LeagueId7                    int    `json:"league_id_7,omitempty"`
	LeagueId8                    int    `json:"league_id_8,omitempty"`
	LeagueId9                    int    `json:"league_id_9,omitempty"`
	LeagueId10                   int    `json:"league_id_10,omitempty"`
	LeagueId11                   int    `json:"league_id_11,omitempty"`
	LeagueId12                   int    `json:"league_id_12,omitempty"`
	LeagueId13                   int    `json:"league_id_13,omitempty"`
	LeagueId14                   int    `json:"league_id_14,omitempty"`
	LeagueId15                   int    `json:"league_id_15,omitempty"`
	LeagueId16                   int    `json:"league_id_16,omitempty"`
	LeagueId17                   int    `json:"league_id_17,omitempty"`
	LeagueId18                   int    `json:"league_id_18,omitempty"`
	LeagueId19                   int    `json:"league_id_19,omitempty"`
	LeagueId20                   int    `json:"league_id_20,omitempty"`
	LeagueId21                   int    `json:"league_id_21,omitempty"`
	LeagueId22                   int    `json:"league_id_22,omitempty"`
	LeagueId23                   int    `json:"league_id_23,omitempty"`
	LeagueId24                   int    `json:"league_id_24,omitempty"`
	LeagueId25                   int    `json:"league_id_25,omitempty"`
}

type SteamUsers struct {
	Users []SteamUser `json:"players"`
}

type SteamUser struct {
	Id                int64  `json:"steamid"`
	VisibilityState   int    `json:"communityvisibilitystate,omitempty"`
	ProfileState      int    `json:"profilestate,omitempty"`
	PersonaName       int64  `json:"personaname,omitempty"`
	LastLogoff        int64  `json:"lastlogoff,omitempty"`
	ProfileUrl        string `json:"profileurl,omitempty"`
	Avatar            string `json:"avatar,omitempty"`
	AvatarMedium      string `json:"avatarmedium,omitempty"`
	AvatarFull        string `json:"avatarfull,omitempty"`
	PersonaState      int    `json:"personastate,omitempty"`
	RealName          string `json:"realname,omitempty"`
	PrimaryClanId     string `json:"primaryclanid,omitempty"`
	TimeCreated       int64  `json:"timecreated,omitempty"`
	PersonaStateFlags int    `json:"personastateflags,omitempty"`
	CountryCode       string `json:"loccountrycode,omitempty"`
	State             string `json:"locstatecode,omitempty"`
	CityId            int    `json:"loccityid,omitempty"`
}

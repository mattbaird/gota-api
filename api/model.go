package api

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
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

type Response struct {
	Response json.RawMessage `json:"response"`
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
	SequenceNumber int64    `json:"match_seq_num"`
	StartTime      int64    `json:"start_time"`
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
	RadiantWin            bool  `json:"radiant_win"`
	Duration              int   `json:"duration"`
	StartTime             int64 `json:"start_time"`
	MatchId               int   `json:"match_id"`
	SequenceNumber        int64 `json:"match_seq_num"`
	RadiantTowerStatus    int   `json:"tower_status_radiant"`
	DireTowerStatus       int   `json:"tower_status_dire"`
	RadiantBarracksStatus int   `json:"barracks_status_radiant"`
	DireBarracksStatus    int   `json:"barracks_status_dire"`
	Cluster               int   `json:"cluster"`
	FirstBloodTime        int   `json:"first_blood_time"`
	HumanPlayers          int   `json:"human_players"`
	LobbyType             int   `json:"lobby_type"`
	LeagueId              int   `json:"leagueid"`
	PositiveVotes         int   `json:"positive_votes"`
	NegativeVotes         int   `json:"negative_votes"`
	GameMode              int   `json:"game_mode"`
	//The following fields are only included if there were teams applied to radiant and dire (i.e. this is a league match in a private lobby)
	RadiantName         string `json:"radiant_name,omit_empty"`
	RadiantLogo         int64  `json:"radiant_logo,omit_empty"`
	RadiantLogoUrl      string `json:"-"`
	RadiantTeamComplete bool   `json:"radiant_team_complete,omit_empty"`
	DireName            string `json:"dire_name,omit_empty"`
	DireLogo            int64  `json:"dire_logo,omit_empty"`
	DireLogoUrl         string `json:"-"`
	DireTeamComplete    bool   `json:"dire_team_complete,omit_empty"`
	//serialize players into embedded structure
	Players []PlayerDetail `json:"players"`
}

func (md *MatchDetail) Region() string {
	switch md.Cluster {
	case 111:
		return "US West"
	case 112:
		return "US West"
	case 114:
		return "US West"
	case 121:
		return "US East"
	case 122:
		return "US East"
	case 123:
		return "US East"
	case 124:
		return "US East"
	case 131:
		return "Europe West"
	case 132:
		return "Europe West"
	case 133:
		return "Europe West"
	case 134:
		return "Europe West"
	case 135:
		return "Europe West"
	case 136:
		return "Europe West"
	case 142:
		return "South Korea"
	case 143:
		return "South Korea"
	case 151:
		return "Southeast Asia"
	case 152:
		return "Southeast Asia"
	case 153:
		return "Southeast Asia"
	case 161:
		return "China"
	case 163:
		return "China"
	case 171:
		return "Australia"
	case 181:
		return "Russia"
	case 182:
		return "Russia"
	case 183:
		return "Russia"
	case 184:
		return "Russia"
	case 191:
		return "Europe East"
	case 200:
		return "South America"
	case 204:
		return "South America"
	case 211:
		return "South America"
	case 212:
		return "South America"
	case 213:
		return "South America"
	case 221:
		return "China"
	case 222:
		return "China" //probably
	case 223:
		return "China" //probably
	case 231:
		return "China" //probably
	default:
		return "unknown"
	}
}

func (md *MatchDetail) GameModeString() string {
	switch md.Cluster {
	case 0:
		return "unknown"
	case 1:
		return "All Pick"
	case 2:
		return "Captains Mode"
	case 3:
		return "Random Draft"
	case 4:
		return "Single Draft"
	case 5:
		return "All Random"
	case 6:
		return "?? INTRO/DEATH ??"
	case 7:
		return "The Diretide"
	case 8:
		return "Reverse Captains Mode"
	case 9:
		return "Greeviling"
	case 10:
		return "Tutorial"
	case 11:
		return "Mid Only"
	case 12:
		return "Least Played"
	case 13:
		return "New Player Pool"
	case 14:
		return "Compendium Matchmaking"
	case 15:
		return "Custom"
	case 16:
		return "Captains Draft"
	case 17:
		return "Balanced Draft"
	case 18:
		return "Ability Draft"
	case 19:
		return "?? Event ??"
	case 20:
		return "All Random Death Match"
	case 21:
		return "1vs1 Solo Mid"
	default:
		return "unknown"
	}
}

func (md *MatchDetail) SV(separator string) string {
	var matchDetails []string
	for _, playa := range md.Players {
		var inputs []string
		inputs = append(inputs, writeNumeric(md.RadiantWin))
		inputs = append(inputs, writeNumeric(md.Duration))
		inputs = append(inputs, writeNumeric(md.StartTime))
		inputs = append(inputs, writeNumeric(convertToYYYYMMDDHH(md.StartTime)))
		inputs = append(inputs, writeNumeric(convertToYYYYMMDDHHMM(md.StartTime)))
		//	inputs = append(inputs, writeNumeric(convertToYYYYMMDDHHMM(md.StartTime+int64(md.Duration))))
		inputs = append(inputs, writeNumeric(convertToYYYYMMDDHHMM(time.Now().Unix())))
		inputs = append(inputs, writeNumeric(md.MatchId))
		inputs = append(inputs, writeNumeric(md.SequenceNumber))
		inputs = append(inputs, writeNumeric(md.RadiantTowerStatus))
		inputs = append(inputs, writeNumeric(md.DireTowerStatus))
		inputs = append(inputs, writeNumeric(md.RadiantBarracksStatus))
		inputs = append(inputs, writeNumeric(md.DireBarracksStatus))
		inputs = append(inputs, writeNumeric(md.Cluster))
		inputs = append(inputs, writeNumeric(md.FirstBloodTime))
		inputs = append(inputs, writeNumeric(md.HumanPlayers))
		inputs = append(inputs, writeNumeric(md.LobbyType))
		inputs = append(inputs, writeNumeric(md.LeagueId))
		inputs = append(inputs, writeNumeric(md.PositiveVotes))
		inputs = append(inputs, writeNumeric(md.NegativeVotes))
		inputs = append(inputs, writeNumeric(md.GameMode))
		inputs = append(inputs, writeNumeric(md.RadiantName))
		inputs = append(inputs, writeNumeric(md.RadiantLogoUrl))
		inputs = append(inputs, writeNumeric(md.RadiantTeamComplete))
		inputs = append(inputs, writeNumeric(md.DireName))
		inputs = append(inputs, writeNumeric(md.DireLogoUrl))
		inputs = append(inputs, writeNumeric(md.DireTeamComplete))
		inputs = append(inputs, writeNumeric(playa.Id))
		inputs = append(inputs, writeNumeric(playa.PlayerSlot))
		inputs = append(inputs, writeString(getTeam(playa.PlayerSlot))) //team
		inputs = append(inputs, writeNumeric(playa.HeroId))
		inputs = append(inputs, writeNumeric(playa.Item0))
		inputs = append(inputs, writeNumeric(playa.Item1))
		inputs = append(inputs, writeNumeric(playa.Item2))
		inputs = append(inputs, writeNumeric(playa.Item3))
		inputs = append(inputs, writeNumeric(playa.Item4))
		inputs = append(inputs, writeNumeric(playa.Item5))
		inputs = append(inputs, writeNumeric(playa.Kills))
		inputs = append(inputs, writeNumeric(playa.Deaths))
		inputs = append(inputs, writeNumeric(playa.Assists))
		inputs = append(inputs, writeNumeric(playa.LeaverStatus))
		inputs = append(inputs, writeNumeric(playa.Gold))
		inputs = append(inputs, writeNumeric(playa.LastHits))
		inputs = append(inputs, writeNumeric(playa.Denies))
		inputs = append(inputs, writeNumeric(playa.GoldPerMin))
		inputs = append(inputs, writeNumeric(playa.XpPerMin))
		inputs = append(inputs, writeNumeric(playa.GoldSpent))
		inputs = append(inputs, writeNumeric(playa.HeroDamage))
		inputs = append(inputs, writeNumeric(playa.TowerDamage))
		inputs = append(inputs, writeNumeric(playa.HeroHealing))
		inputs = append(inputs, writeNumeric(playa.Level))
		inputs = append(inputs, playa.SteamSV(","))
		record := strings.Join(inputs, separator)
		matchDetails = append(matchDetails, record)
	}
	return strings.Join(matchDetails, "\n")
}

// hero wins if match field radiant_win is true and its player_slot < 5 OR radiant_win is false and its player_slot > 5.
func getTeam(slot int) string {
	if slot < 5 {
		return "Radiant"
	} else {
		return "Dire"
	}
}

func convertToYYYYMMDDHH(seconds int64) int {
	t := time.Unix(seconds, 0)
	// use data formatting hack
	stringRepresentation := t.Format("2006010215")
	// now convert string to int
	yyyymmddhh, err := strconv.Atoi(stringRepresentation)
	if err != nil {
		fmt.Printf("Error parsing string number:%v\n", err)
	}
	return yyyymmddhh
}

func convertToYYYYMMDDHHMM(seconds int64) int {
	t := time.Unix(seconds, 0)
	// use data formatting hack
	stringRepresentation := t.Format("200601021504")
	// now convert string to int
	yyyymmddhh, err := strconv.Atoi(stringRepresentation)
	if err != nil {
		fmt.Printf("Error parsing string number:%v\n", err)
	}
	return yyyymmddhh
}

func (md *MatchDetail) PlayersArray() []int {
	var retval []int
	for _, pd := range md.Players {
		retval = append(retval, pd.Id)
	}
	return retval
}

type PlayerDetail struct {
	Id              int              `json:"account_id"`
	PlayerSlot      int              `json:"player_slot"`
	HeroId          int              `json:"hero_id"`
	Item0           int              `json:"item_0"`
	Item1           int              `json:"item_1"`
	Item2           int              `json:"item_2"`
	Item3           int              `json:"item_3"`
	Item4           int              `json:"item_4"`
	Item5           int              `json:"item_5"`
	Kills           int              `json:"kills"`
	Deaths          int              `json:"deaths"`
	Assists         int              `json:"assists"`
	LeaverStatus    int              `json:"leaver_status"`
	Gold            int              `json:"gold"`
	LastHits        int              `json:"last_hits"`
	Denies          int              `json:"denies"`
	GoldPerMin      int              `json:"gold_per_min"`
	XpPerMin        int              `json:"xp_per_min"`
	GoldSpent       int              `json:"gold_spent"`
	HeroDamage      int              `json:"hero_damage"`
	TowerDamage     int              `json:"tower_damage"`
	HeroHealing     int              `json:"hero_healing"`
	Level           int              `json:"level"`
	AbilityUpgrades []AbilityUpgrade `json:"ability_upgrades,omitempty"`
	SteamUser       *SteamUser
}

type AbilityUpgrade struct {
	Ability int `json:"ability"`
	Time    int `json:"time"`
	Level   int `json:"level"`
}

func (pd *PlayerDetail) SV(separator string) string {
	var inputs []string
	inputs = append(inputs, writeKeyValue("account_id", pd.Id))
	inputs = append(inputs, writeKeyValue("player_slot", pd.PlayerSlot))
	inputs = append(inputs, writeKeyValue("hero_id", pd.HeroId))
	inputs = append(inputs, writeKeyValue("item_0", pd.Item0))
	inputs = append(inputs, writeKeyValue("item_1", pd.Item1))
	inputs = append(inputs, writeKeyValue("item_2", pd.Item2))
	inputs = append(inputs, writeKeyValue("item_3", pd.Item3))
	inputs = append(inputs, writeKeyValue("item_4", pd.Item4))
	inputs = append(inputs, writeKeyValue("item_5", pd.Item5))
	inputs = append(inputs, writeKeyValue("kills", pd.Kills))
	inputs = append(inputs, writeKeyValue("deaths", pd.Deaths))
	inputs = append(inputs, writeKeyValue("assists", pd.Assists))
	inputs = append(inputs, writeKeyValue("leaver_status", pd.LeaverStatus))
	inputs = append(inputs, writeKeyValue("gold", pd.Gold))
	inputs = append(inputs, writeKeyValue("last_hits", pd.LastHits))
	inputs = append(inputs, writeKeyValue("denies", pd.Denies))
	inputs = append(inputs, writeKeyValue("gold_per_minute", pd.GoldPerMin))
	inputs = append(inputs, writeKeyValue("experience_per_minute", pd.XpPerMin))
	inputs = append(inputs, writeKeyValue("gold_spent", pd.GoldSpent))
	inputs = append(inputs, writeKeyValue("hero_damage", pd.HeroDamage))
	inputs = append(inputs, writeKeyValue("tower_damage", pd.TowerDamage))
	inputs = append(inputs, writeKeyValue("hero_healing", pd.HeroHealing))
	inputs = append(inputs, writeKeyValue("level", pd.Level))
	if pd.SteamUser != nil {
		inputs = append(inputs, writeKeyValue("steam_id", pd.SteamUser.Id))
		inputs = append(inputs, writeKeyValue("visibility_state", pd.SteamUser.VisibilityState))
		inputs = append(inputs, writeKeyValue("profile_state", pd.SteamUser.ProfileState))
		inputs = append(inputs, writeKeyValue("persona_name", pd.SteamUser.PersonaName))
		inputs = append(inputs, writeKeyValue("last_logoff", pd.SteamUser.LastLogoff))
		inputs = append(inputs, writeKeyValue("profile_url", pd.SteamUser.ProfileUrl))
		inputs = append(inputs, writeKeyValue("avatar", pd.SteamUser.Avatar))
		inputs = append(inputs, writeKeyValue("avatar_medium", pd.SteamUser.AvatarMedium))
		inputs = append(inputs, writeKeyValue("avatar_full", pd.SteamUser.AvatarFull))
		inputs = append(inputs, writeKeyValue("persona_state", pd.SteamUser.PersonaState))
		inputs = append(inputs, writeKeyValue("real_name", pd.SteamUser.RealName))
		inputs = append(inputs, writeKeyValue("primary_clan_id", pd.SteamUser.PrimaryClanId))
		inputs = append(inputs, writeKeyValue("time_created", pd.SteamUser.TimeCreated))
		inputs = append(inputs, writeKeyValue("persona_state_flags", pd.SteamUser.PersonaStateFlags))
		inputs = append(inputs, writeKeyValue("country_code", pd.SteamUser.CountryCode))
		inputs = append(inputs, writeKeyValue("state", pd.SteamUser.State))
		inputs = append(inputs, writeKeyValue("city_id", pd.SteamUser.CityId))
	}
	return strings.Join(inputs, separator)
}

func (pd *PlayerDetail) SteamSV(separator string) string {
	var inputs []string
	if pd.SteamUser != nil {
		inputs = append(inputs, writeKeyValue("steam_id", pd.SteamUser.Id))
		inputs = append(inputs, writeKeyValue("visibility_state", pd.SteamUser.VisibilityState))
		inputs = append(inputs, writeKeyValue("profile_state", pd.SteamUser.ProfileState))
		inputs = append(inputs, writeKeyValue("persona_name", pd.SteamUser.PersonaName))
		inputs = append(inputs, writeKeyValue("last_logoff", pd.SteamUser.LastLogoff))
		inputs = append(inputs, writeKeyValue("profile_url", pd.SteamUser.ProfileUrl))
		inputs = append(inputs, writeKeyValue("avatar", pd.SteamUser.Avatar))
		inputs = append(inputs, writeKeyValue("avatar_medium", pd.SteamUser.AvatarMedium))
		inputs = append(inputs, writeKeyValue("avatar_full", pd.SteamUser.AvatarFull))
		inputs = append(inputs, writeKeyValue("persona_state", pd.SteamUser.PersonaState))
		inputs = append(inputs, writeKeyValue("real_name", pd.SteamUser.RealName))
		inputs = append(inputs, writeKeyValue("primary_clan_id", pd.SteamUser.PrimaryClanId))
		inputs = append(inputs, writeKeyValue("time_created", pd.SteamUser.TimeCreated))
		inputs = append(inputs, writeKeyValue("persona_state_flags", pd.SteamUser.PersonaStateFlags))
		inputs = append(inputs, writeKeyValue("country_code", pd.SteamUser.CountryCode))
		inputs = append(inputs, writeKeyValue("state", pd.SteamUser.State))
		inputs = append(inputs, writeKeyValue("city_id", pd.SteamUser.CityId))
	}
	tmp := strings.Join(inputs, separator)
	return tmp
}

type Heroes struct {
	Heroes []Hero `json:"heroes"`
	Status int    `json:"status"`
	Count  int    `json:"count"`
}

func (hs *Heroes) SV(separator string) []string {
	var retval []string
	for _, hero := range hs.Heroes {
		retval = append(retval, hero.SV(separator))
	}
	return retval
}

type Hero struct {
	Name          string `json:"name"`
	Id            int    `json:"id"`
	LocalizedName string `json:"localized_name,omitempty"`
}

func (h *Hero) SV(separator string) string {
	var inputs []string
	inputs = append(inputs, writeNumeric(h.Id))
	inputs = append(inputs, writeString(h.Name))
	inputs = append(inputs, writeString(h.LocalizedName))
	return strings.Join(inputs, separator)
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
	Users []*SteamUser `json:"players,omitempty"`
}

type SteamUser struct {
	Id                string `json:"steamid"`
	VisibilityState   int    `json:"communityvisibilitystate,omitempty"`
	ProfileState      int    `json:"profilestate,omitempty"`
	PersonaName       string `json:"personaname,omitempty"`
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

type UGCFileDetails struct {
	Filename string `json:"filename"`
	Url      string `json:"url"`
	Size     int    `json:"size"`
}

// from: https://raw.githubusercontent.com/Holek/steam-friends-countries/master/data/steam_countries.json

type GeoData struct {
	Countries map[string]*Country
}

type Country struct {
	Name                     string
	States                   map[string]*State
	Longitude                float64
	Latitude                 float64
	CoordinatesAccuracyLevel string
}

type State struct {
	Name                     string
	Cities                   map[string]*City
	Longitude                float64
	Latitude                 float64
	CoordinatesAccuracyLevel string
}

type City struct {
	Name                     string
	Longitude                float64
	Latitude                 float64
	CoordinatesAccuracyLevel string
}

type Items struct {
	Items []Item `json:"items"`
}

type Item struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (i *Items) GetById(id int) (Item, bool) {
	for _, item := range i.Items {
		if item.Id == id {
			return item, true
		}
	}
	return Item{}, false
}

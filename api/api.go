package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const dota_api_host string = "api.steampowered.com"
const dota_api_path string = "/DOTA_API_PREFIX/DOTA_API_COMMAND/DOTA_API_VERSION/"
const gota_api_key_param = "key={{key}}"
const version string = "v0001"
const api_remote_storage_prefix string = "ISteamRemoteStorage"
const api_match_prefix string = "IDOTA2Match_570"
const api_match_prefix_debug string = "IDOTA2Match_205790"
const api_econ_prefix string = "IEconDOTA2_570"
const debug bool = false
const use_debug_service = true

func getApiMatchPrefix() string {
	if debug || use_debug_service {
		return api_match_prefix_debug
	} else {
		return api_match_prefix
	}
}
func NewGotaAPI(apiKey string, lang string) (*GotaAPI, error) {
	path := strings.Replace(dota_api_path, "DOTA_API_VERSION", version, -1)
	u := url.URL{}
	u.Scheme = "https"
	u.Host = dota_api_host
	u.Path = path
	v := url.Values{}
	var realKey string = apiKey
	if apiKey == "" {
		realKey = os.Getenv("DOTA_API_KEY")
	}
	if realKey == "" {
		return nil, fmt.Errorf("Please either pass dota key, or set it in OS environment variable 'DOTA_API_KEY'")
	}
	if debug {
		log.Printf("key:%s", realKey)
	}

	v.Set("key", realKey)
	u.RawQuery = v.Encode()
	var validLang string = "en_us"
	if len(lang) == 0 {
		validLang = lang
	}
	return &GotaAPI{Key: realKey, endpoint: u.String(), language: validLang}, nil
}

type GotaAPI struct {
	Key       string
	Transport http.RoundTripper
	endpoint  string
	language  string
}

func runGota(api *GotaAPI, command string, apiPrefix string, parameters map[string]interface{}) (Result, error) {
	// set parameters
	vals := url.Values{}
	for k, v := range parameters {
		s := fmt.Sprintf("%v", v)
		if debug {
			log.Printf("%s:%s", k, s)
		}
		vals.Add(k, s)
	}
	if debug {
		log.Printf("vals:%s", vals.Encode())
	}
	// do command/DOTA_API_COMMAND substitution
	requestUrl := strings.Replace(api.endpoint, "DOTA_API_COMMAND", command, -1)
	requestUrl = strings.Replace(requestUrl, "DOTA_API_PREFIX", apiPrefix, -1)
	requestUrl = requestUrl + "&" + vals.Encode()
	if debug {
		log.Printf("Request URL:%s", requestUrl)
	}
	var result Result
	err := runGotaRaw(api, requestUrl, &result)
	return result, err
}

func runGotaRaw(api *GotaAPI, requestUrl string, r interface{}) error {
	client := &http.Client{Transport: api.Transport}
	resp, err := client.Get(requestUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if debug {
		log.Printf("Response Body:%s", string(body))
	}
	if err = errorCheck(body); err != nil {
		return err
	}
	err = json.Unmarshal(body, &r)
	return err
}

func errorCheck(body []byte) error {
	var e APIError
	json.Unmarshal(body, &e)
	if e.Err != "" || e.Code != 0 {
		return e
	}
	return nil
}

func (api *GotaAPI) URL() string {
	return api.endpoint
}

func (api *GotaAPI) GetMatchHistory(playerName string, heroId int, gameMode string, skill string, dateMin string, dateMax string,
	minPlayers int, accountId int, leagueId string, startAtMatchId int, matchesRequested int, tournamentGamesOnly bool) (MatchHistory, error) {
	parameters := make(map[string]interface{})
	if playerName != "" {
		parameters["player_name"] = playerName
	}
	if heroId > 0 {
		parameters["hero_id"] = heroId
	}
	if gameMode != "" {
		parameters["game_mode"] = gameMode
	}
	if skill != "" {
		parameters["skill"] = skill
	}
	if dateMin != "" {
		parameters["date_min"] = dateMin
	}
	if dateMax != "" {
		parameters["date_max"] = dateMax
	}
	if minPlayers > 0 {
		parameters["min_players"] = minPlayers
	}
	if accountId > 0 {
		parameters["account_id"] = accountId
	}
	if leagueId != "" {
		parameters["league_id"] = leagueId
	}
	if startAtMatchId > 0 {
		parameters["start_at_matchId"] = startAtMatchId
	}
	if matchesRequested > 0 {
		parameters["matches_requested"] = matchesRequested
	}
	if tournamentGamesOnly {
		parameters["tournament_games_only"] = tournamentGamesOnly
	}
	var retval MatchHistory = MatchHistory{}
	var err error = nil
	var result Result
	result, err = runGota(api, "GetMatchHistory", getApiMatchPrefix(), parameters)
	if err == nil {
		err = json.Unmarshal(result.Result, &retval)
	}
	return retval, err
}

func (api *GotaAPI) GetMatchDetails(matchId int) (MatchDetail, error) {
	var retval MatchDetail = MatchDetail{}
	parameters := make(map[string]interface{})
	if matchId <= 0 {
		return retval, fmt.Errorf("invalid matchId :%v", matchId)
	}
	parameters["match_id"] = matchId

	var err error = nil
	var result Result
	result, err = runGota(api, "GetMatchDetails", getApiMatchPrefix(), parameters)
	if err == nil {
		err = json.Unmarshal(result.Result, &retval)
	}
	return retval, err
}

func (api *GotaAPI) GetMatchHistoryBySequenceNum(startAtMatchSeqNum int64, matchesRequested int) (MatchHistory, error) {
	var retval MatchHistory = MatchHistory{}
	parameters := make(map[string]interface{})
	if startAtMatchSeqNum <= 0 {
		return retval, fmt.Errorf("invalid startAtMatchSeqNum :%v", startAtMatchSeqNum)
	}
	parameters["start_at_match_seq_num"] = startAtMatchSeqNum
	if matchesRequested > 0 {
		parameters["matches_requested"] = matchesRequested
	}
	var err error = nil
	var result Result
	result, err = runGota(api, "GetMatchHistoryBySequenceNum", getApiMatchPrefix(), parameters)
	if err == nil {
		err = json.Unmarshal(result.Result, &retval)
	}
	return retval, err
}

func (api *GotaAPI) GetHeroes() (Heroes, error) {
	var retval Heroes = Heroes{}
	parameters := make(map[string]interface{})
	parameters["language"] = "en_us"
	var err error = nil
	var result Result
	result, err = runGota(api, "GetHeroes", api_econ_prefix, parameters)
	if err == nil {
		err = json.Unmarshal(result.Result, &retval)
	}
	return retval, err
}

func (api *GotaAPI) GetLeagueListing() (Leagues, error) {
	var retval Leagues = Leagues{}
	parameters := make(map[string]interface{})
	var err error = nil
	var result Result
	result, err = runGota(api, "GetLeagueListing", getApiMatchPrefix(), parameters)
	if err == nil {
		err = json.Unmarshal(result.Result, &retval)
	}
	return retval, err
}

func (api *GotaAPI) GetLiveLeagueGames() (LeagueGames, error) {
	var retval LeagueGames = LeagueGames{}
	parameters := make(map[string]interface{})
	var err error = nil
	var result Result
	result, err = runGota(api, "GetLiveLeagueGames", getApiMatchPrefix(), parameters)
	if err == nil {
		err = json.Unmarshal(result.Result, &retval)
	}
	return retval, err
}

func (api *GotaAPI) GetTeamInfoByTeamID(startAtTeamId int, teamsRequested int) (Teams, error) {
	var retval Teams = Teams{}
	parameters := make(map[string]interface{})
	if startAtTeamId <= 0 {
		return retval, fmt.Errorf("invalid startAtTeamId :%v", startAtTeamId)
	}
	parameters["start_at_team_id"] = startAtTeamId
	if teamsRequested > 0 {
		parameters["teams_requested"] = teamsRequested
	}
	var err error = nil
	var result Result
	result, err = runGota(api, "GetTeamInfoByTeamID", getApiMatchPrefix(), parameters)
	if err == nil {
		err = json.Unmarshal(result.Result, &retval)
	}
	return retval, err
}

func (api *GotaAPI) GetPlayerSummaries(ids ...int) (SteamUsers, error) {
	var retval SteamUsers = SteamUsers{}
	var idArray string
	for i, id := range ids {
		if i == 0 {
			idArray = fmt.Sprintf("%v", TranslateSteamId(id))
		} else {
			idArray += "," + fmt.Sprintf("%v", TranslateSteamId(id))
		}
	}
	requestUrl := fmt.Sprintf("https://%s/ISteamUser/GetPlayerSummaries/v0002/?key=%s&steamids=%s", dota_api_host, api.Key, idArray)
	if debug {
		fmt.Printf("GetPlayerSummaries:%s\n", requestUrl)
	}
	var err error = nil
	var response Response
	err = runGotaRaw(api, requestUrl, &response)
	if err == nil {
		err = json.Unmarshal(response.Response, &retval)
	}
	return retval, err
}

func (api *GotaAPI) GetRemoveFileDetails(ugcId int64) (UGCFileDetails, error) {
	var retval UGCFileDetails = UGCFileDetails{}
	parameters := make(map[string]interface{})
	parameters["ugcid"] = ugcId
	var err error = nil
	var result Result
	result, err = runGota(api, "GetUGCFileDetails", api_remote_storage_prefix, parameters)
	if err == nil {
		err = json.Unmarshal(result.Result, &retval)
	}
	return retval, err
}

func TranslateSteamId(id int) string {
	return fmt.Sprintf("%v", id+76561197960265728)
}

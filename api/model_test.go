package api

import (
	"testing"
)

func Test_Matches(t *testing.T) {
	url := "https://api.steampowered.com/IDOTA2Match_570/GetMatchHistory/V001/?key=KEY"
	matches, err := ReadMatchHistory(url)
	if err != nil {
		t.Errorf("fail:%v", err)
	}
	if len(matches.Matches) == 0 {
		t.Error("did not load")
	}
}

func Test_MatchDetails(t *testing.T) {
	url := "https://api.steampowered.com/IDOTA2Match_570/GetMatchDetails/V001/?match_id=27110133&key=KEY"
	match, err := ReadMatchDetails(url)
	if err != nil {
		t.Errorf("fail:%v", err)
	}
	if len(match.Players) == 0 {
		t.Error("did not load")
	}
}

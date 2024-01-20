package nba

import (
	"encoding/json"
	"io"
	"net/http"
)

// FetchNBASchedule fetches the NBA schedule data from the provided JSON endpoint
func FetchNBASchedule() ([]NBAGameData, error) {
	resp, err := http.Get("https://cdn.nba.com/static/json/staticData/scheduleLeagueV2.json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var nbaData struct {
		LeagueSchedule struct {
			GameDates []struct {
				Games []NBAGameData `json:"games"`
			} `json:"gameDates"`
		} `json:"leagueSchedule"`
	}

	err = json.Unmarshal(body, &nbaData)
	if err != nil {
		return nil, err
	}

	var schedule []NBAGameData
	for _, date := range nbaData.LeagueSchedule.GameDates {
		schedule = append(schedule, date.Games...)
	}

	return schedule, nil
}

package xmltv

import (
	"fmt"
	"nba-xmltv-schedule/nba"
	"time"

	"github.com/gosimple/slug"
)

func ConvertToXMLTV(schedule []nba.NBAGameData, iconBaseURL string) TV {
	var tv TV

	// Convert games to programs
	for _, game := range schedule {
		fullTeamName := fmt.Sprintf("%s %s", game.HomeTeam.TeamCity, game.HomeTeam.TeamName)
		fullAwayTeamName := fmt.Sprintf("%s %s", game.AwayTeam.TeamCity, game.AwayTeam.TeamName)

		// Add channel if not already added
		if !channelExists(tv.Channels, getChannelID(fullTeamName)) && getChannelID(fullTeamName) != "" {
			channel := Channel{
				ID:          getChannelID(fullTeamName),
				DisplayName: fullTeamName,
				Icon: Icon{
					Src: fmt.Sprintf("%s/%s.png", iconBaseURL, getChannelID(fullTeamName)),
				},
			}
			tv.Channels = append(tv.Channels, channel)
		}

		program := Program{
			Start:   game.GameTime.Format("20060102150405 -0700"),
			Stop:    game.GameTime.Add(2 * time.Hour).Format("20060102150405 -0700"), // Assuming each game lasts 2 hours
			Channel: getChannelID(fullTeamName),
			Categories: []Category{
				{
					Lang: "en",
					Text: "Sports",
				},
			},
			Desc: fmt.Sprintf("%s - %s, %s, %s", game.WeekName, game.ArenaName, game.ArenaCity, game.ArenaState),
			Title: Title{
				Text: fmt.Sprintf("%s vs %s", fullTeamName, fullAwayTeamName),
				Lang: "en",
			},
			Date: game.GameTime.Format("20060102"),
			Audio: Audio{
				Stereo: "stereo",
			},
			Subtitles: Subtitles{
				Type: "teletext",
			},
		}

		awayProgram := Program{
			Start:   game.GameTime.Format("20060102150405 -0700"),
			Stop:    game.GameTime.Add(2 * time.Hour).Format("20060102150405 -0700"), // Assuming each game lasts 2 hours
			Channel: getChannelID(fullAwayTeamName),
			Categories: []Category{
				{
					Lang: "en",
					Text: "Sports",
				},
			},
			Desc: fmt.Sprintf("%s - %s, %s, %s", game.WeekName, game.ArenaName, game.ArenaCity, game.ArenaState),
			Title: Title{
				Text: fmt.Sprintf("%s vs %s", fullTeamName, fullAwayTeamName),
				Lang: "en",
			},
			Date: game.GameTime.Format("20060102"),
			Audio: Audio{
				Stereo: "stereo",
			},
			Subtitles: Subtitles{
				Type: "teletext",
			},
		}

		if getChannelID(fullTeamName) == "" {
			continue
		}

		tv.Programs = append(tv.Programs, program)
		tv.Programs = append(tv.Programs, awayProgram)
	}

	tv.SourceInfoName = "NBA XMLTV"
	tv.SourceInfoUrl = "https://nba.almerolnet.ovh/"
	tv.GeneratorInfoName = "XMLTV/$Id: tv_grab_na_dd.in,v 1.70 2008/03/03 15:21:41 rmeden Exp $"
	tv.GeneratorInfoUrl = "http://www.xmltv.org/"
	return tv
}

func getChannelID(teamName string) string {
	// Add your logic to map team names to channel IDs
	// For simplicity, we'll return a default channel ID
	return slug.Make(teamName)
}

func channelExists(channels []Channel, channelID string) bool {
	for _, channel := range channels {
		if channel.ID == channelID {
			return true
		}
	}
	return false
}

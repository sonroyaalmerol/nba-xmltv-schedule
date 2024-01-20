package nba

import (
	"time"
)

// Event represents a single game event
type Event struct {
	Title     string    `xml:"title"`
	StartTime time.Time `xml:"start"`
	EndTime   time.Time `xml:"end"`
}

// TeamSchedule represents the schedule for a specific NBA team
type TeamSchedule struct {
	TeamName string  `xml:"teamName"`
	Events   []Event `xml:"event"`
}

// NBASchedule represents the schedule for all NBA teams
type NBASchedule struct {
	Teams []TeamSchedule `xml:"teamSchedule"`
}

// NBAGameData represents the game data from the JSON endpoint
type NBAGameData struct {
	GameId     string    `json:"gameId"`
	HomeTeam   Team      `json:"homeTeam"`
	AwayTeam   Team      `json:"awayTeam"`
	GameStatus int       `json:"gameStatus"`
	GameTime   time.Time `json:"gameDateTimeEst"`
	WeekName   string    `json:"weekName"`
	ArenaName  string    `json:"arenaName"`
	ArenaState string    `json:"arenaState"`
	ArenaCity  string    `json:"arenaCity"`
}

// Team represents an NBA team
type Team struct {
	TeamName string `json:"teamName"`
	TeamCity string `json:"teamCity"`
}

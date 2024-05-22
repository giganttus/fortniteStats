package models

import "time"

type Response struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
	Data   Player `json:"data"`
}

type Player struct {
	Account Account `json:"account"`
	Stats   Stats   `json:"stats"`
}

type Account struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Stats struct {
	All ModeStats `json:"all"`
}

type ModeStats struct {
	Overall ModeInfo `json:"overall"`
	Solo    ModeInfo `json:"solo"`
	Duo     ModeInfo `json:"duo"`
	Trio    ModeInfo `json:"trio"`
	Squad   ModeInfo `json:"squad"`
}

type ModeInfo struct {
	Wins         int       `json:"wins"`
	Top3         int       `json:"top3"`
	Top5         int       `json:"top5"`
	Top6         int       `json:"top6"`
	Top10        int       `json:"top10"`
	Top12        int       `json:"top12"`
	Top25        int       `json:"top25"`
	Kills        int       `json:"kills"`
	Matches      int       `json:"matches"`
	LastModified time.Time `json:"lastModified"`
}

type FinalStats struct {
	ID                  string    `gorm:"primaryKey" json:"id"`
	Name                string    `json:"name"`
	OverallWins         int       `json:"overallWins"`
	OverallTop3         int       `json:"overallTop3"`
	OverallTop5         int       `json:"overallTop5"`
	OverallTop6         int       `json:"overallTop6"`
	OverallTop10        int       `json:"overallTop10"`
	OverallTop12        int       `json:"overallTop12"`
	OverallTop25        int       `json:"overallTop25"`
	OverallKills        int       `json:"overallKills"`
	OverallMatches      int       `json:"overallMatches"`
	OverallLastModified time.Time `gorm:"default:'2000-01-01 00:00:00'" json:"overallLastModified"`

	// Solo stats
	SoloWins         int       `json:"soloWins"`
	SoloTop3         int       `json:"soloTop3"`
	SoloTop5         int       `json:"soloTop5"`
	SoloTop6         int       `json:"soloTop6"`
	SoloTop10        int       `json:"soloTop10"`
	SoloTop12        int       `json:"soloTop12"`
	SoloTop25        int       `json:"soloTop25"`
	SoloKills        int       `json:"soloKills"`
	SoloMatches      int       `json:"soloMatches"`
	SoloLastModified time.Time `gorm:"default:'2000-01-01 00:00:00'" json:"soloLastModified"`

	// Duo stats
	DuoWins         int       `json:"duoWins"`
	DuoTop3         int       `json:"duoTop3"`
	DuoTop5         int       `json:"duoTop5"`
	DuoTop6         int       `json:"duoTop6"`
	DuoTop10        int       `json:"duoTop10"`
	DuoTop12        int       `json:"duoTop12"`
	DuoTop25        int       `json:"duoTop25"`
	DuoKills        int       `json:"duoKills"`
	DuoMatches      int       `json:"duoMatches"`
	DuoLastModified time.Time `gorm:"default:'2000-01-01 00:00:00'" json:"duoLastModified"`

	// Trio stats
	TrioWins         int       `json:"trioWins"`
	TrioTop3         int       `json:"trioTop3"`
	TrioTop5         int       `json:"trioTop5"`
	TrioTop6         int       `json:"trioTop6"`
	TrioTop10        int       `json:"trioTop10"`
	TrioTop12        int       `json:"trioTop12"`
	TrioTop25        int       `json:"trioTop25"`
	TrioKills        int       `json:"trioKills"`
	TrioMatches      int       `json:"trioMatches"`
	TrioLastModified time.Time `gorm:"default:'2000-01-01 00:00:00'" json:"trioLastModified"`

	// Squad stats
	SquadWins         int       `json:"squadWins"`
	SquadTop3         int       `json:"squadTop3"`
	SquadTop5         int       `json:"squadTop5"`
	SquadTop6         int       `json:"squadTop6"`
	SquadTop10        int       `json:"squadTop10"`
	SquadTop12        int       `json:"squadTop12"`
	SquadTop25        int       `json:"squadTop25"`
	SquadKills        int       `json:"squadKills"`
	SquadMatches      int       `json:"squadMatches"`
	SquadLastModified time.Time `gorm:"default:'2000-01-01 00:00:00'" json:"squadLastModified"`
}

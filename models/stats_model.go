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
	KeyboardMouse ModeStats `json:"keyboardMouse"`
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
	ID                  string    `json:"id" bson:"id"`
	Name                string    `json:"name" bson:"name"`
	OverallWins         int       `json:"overallWins" bson:"overallWins"`
	OverallTop3         int       `json:"overallTop3" bson:"overallTop3"`
	OverallTop5         int       `json:"overallTop5" bson:"overallTop5"`
	OverallTop6         int       `json:"overallTop6" bson:"overallTop6"`
	OverallTop10        int       `json:"overallTop10" bson:"overallTop10"`
	OverallTop12        int       `json:"overallTop12" bson:"overallTop12"`
	OverallTop25        int       `json:"overallTop25" bson:"overallTop25"`
	OverallKills        int       `json:"overallKills" bson:"overallKills"`
	OverallMatches      int       `json:"overallMatches" bson:"overallMatches"`
	OverallLastModified time.Time `json:"overallLastModified" bson:"overallLastModified"`

	// Solo stats
	SoloWins         int       `json:"soloWins" bson:"soloWins"`
	SoloTop3         int       `json:"soloTop3" bson:"soloTop3"`
	SoloTop5         int       `json:"soloTop5" bson:"soloTop5"`
	SoloTop6         int       `json:"soloTop6" bson:"soloTop6"`
	SoloTop10        int       `json:"soloTop10" bson:"soloTop10"`
	SoloTop12        int       `json:"soloTop12" bson:"soloTop12"`
	SoloTop25        int       `json:"soloTop25" bson:"soloTop25"`
	SoloKills        int       `json:"soloKills" bson:"soloKills"`
	SoloMatches      int       `json:"soloMatches" bson:"soloMatches"`
	SoloLastModified time.Time `json:"soloLastModified" bson:"soloLastModified"`

	// Duo stats
	DuoWins         int       `json:"duoWins" bson:"duoWins"`
	DuoTop3         int       `json:"duoTop3" bson:"duoTop3"`
	DuoTop5         int       `json:"duoTop5" bson:"duoTop5"`
	DuoTop6         int       `json:"duoTop6" bson:"duoTop6"`
	DuoTop10        int       `json:"duoTop10" bson:"duoTop10"`
	DuoTop12        int       `json:"duoTop12" bson:"duoTop12"`
	DuoTop25        int       `json:"duoTop25" bson:"duoTop25"`
	DuoKills        int       `json:"duoKills" bson:"duoKills"`
	DuoMatches      int       `json:"duoMatches" bson:"duoMatches"`
	DuoLastModified time.Time `json:"duoLastModified" bson:"duoLastModified"`

	// Trio stats
	TrioWins         int       `json:"trioWins" bson:"trioWins"`
	TrioTop3         int       `json:"trioTop3" bson:"trioTop3"`
	TrioTop5         int       `json:"trioTop5" bson:"trioTop5"`
	TrioTop6         int       `json:"trioTop6" bson:"trioTop6"`
	TrioTop10        int       `json:"trioTop10" bson:"trioTop10"`
	TrioTop12        int       `json:"trioTop12" bson:"trioTop12"`
	TrioTop25        int       `json:"trioTop25" bson:"trioTop25"`
	TrioKills        int       `json:"trioKills" bson:"trioKills"`
	TrioMatches      int       `json:"trioMatches" bson:"trioMatches"`
	TrioLastModified time.Time `json:"trioLastModified" bson:"trioLastModified"`

	// Squad stats
	SquadWins         int       `json:"squadWins" bson:"squadWins"`
	SquadTop3         int       `json:"squadTop3" bson:"squadTop3"`
	SquadTop5         int       `json:"squadTop5" bson:"squadTop5"`
	SquadTop6         int       `json:"squadTop6" bson:"squadTop6"`
	SquadTop10        int       `json:"squadTop10" bson:"squadTop10"`
	SquadTop12        int       `json:"squadTop12" bson:"squadTop12"`
	SquadTop25        int       `json:"squadTop25" bson:"squadTop25"`
	SquadKills        int       `json:"squadKills" bson:"squadKills"`
	SquadMatches      int       `json:"squadMatches" bson:"squadMatches"`
	SquadLastModified time.Time `json:"squadLastModified" bson:"squadLastModified"`
}

package main

import (
	"io"
	"os"
	"slices"
)

type Name string
type Team struct {
	Name    Name
	Players []string
}

type League struct {
	Teams []Team
	Wins  map[Name]int
}

func (l *League) MatchResult(teamOne, teamTwo Name, scoreOne, scoreTwo int) {
	if scoreOne > scoreTwo {
		l.Wins[teamOne]++
	} else if scoreTwo > scoreOne {
		l.Wins[teamTwo]++
	}
	// do nothing for draws
}

func (l League) Ranking() []Team {
	comparator := func(a, b Team) int {
		return l.Wins[b.Name] - l.Wins[a.Name]
	}
	slices.SortFunc(l.Teams, comparator)
	return l.Teams
}

type Ranker interface {
	Ranking() []Team
}

func RankPrinter(r Ranker, w io.Writer) {
	results := r.Ranking()
	for _, v := range results {
		io.WriteString(w, string(v.Name))
		w.Write([]byte("\n"))
	}
}

func main() {
	teams := []Team{{Name: "Liverpool", Players: []string{}},
		{Name: "Man U", Players: []string{}},
		{Name: "Chelsea", Players: []string{}},
	}
	league := League{Teams: teams, Wins: map[Name]int{}}
	league.MatchResult("Liverpool", "Chelsea", 1, 0)
	league.MatchResult("Man U", "Liverpool", 2, 7)
	league.MatchResult("Chelsea", "Man U", 2, 1)

	RankPrinter(league, os.Stdout)
}

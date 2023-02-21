package main

import (
	"strings"
	"time"
)

type API struct{}

type Song struct {
	Name      string   `json:"name"`
	Album     string   `json:"album"`
	Artists   []string `json:"artists"`
	Duration  int      `json:"duration"`
	Slug      string
	AddedOn   time.Time
	UpdatedOn time.Time
}

type Error struct {
	Status    int       `json:"status_code"`
	Error     string    `json:"error"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

func (s *Song) Create() {
	lower := strings.ToLower(s.Name)
	s.Slug = strings.ReplaceAll(lower, " ", "-")
	s.AddedOn = time.Now()
	s.UpdatedOn = s.AddedOn
}

func (s *Song) Update(u Song) {
	s.Name = u.Name
	s.Artists = u.Artists
	s.Album = u.Album
	s.Duration = u.Duration
	s.UpdatedOn = time.Now()
}

var DB = []Song{
	{
		Name:     "I Wanna Be Yours",
		Slug:     "i-wanna-be-yours",
		Album:    "AM",
		Artists:  []string{"Arctic Monkeys"},
		Duration: 183,
		AddedOn:  time.Now(),
	},
	{
		Name:     "Family Ties",
		Slug:     "family-ties",
		Album:    "The Melodic Blue",
		Artists:  []string{"Baby Keem", "Kendrick Lamar"},
		Duration: 252,
		AddedOn:  time.Now(),
	},
}

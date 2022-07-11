package models

import "time"

type About struct {
	About          string                       `yaml:"about" json:"about"`
	Education      map[string]map[string]string `yaml:"education" json:"education"`
	Experience     map[string]map[string]string `yaml:"experience" json:"experience"`
	Certifications map[string][]string          `yaml:"certifications" json:"certifications"`
	Timestamp      time.Time                    `json:"timestamp"`
}

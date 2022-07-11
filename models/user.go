package models

import "time"

type User struct {
	Name         string                         `yaml:"name" json:"name"`
	Year         int                            `yaml:"year" json:"-"`
	Month        int                            `yaml:"month" json:"-"`
	Date         int                            `yaml:"date" json:"-"`
	Age          int                            `json:"age"`
	Profession   string                         `yaml:"profession" json:"profession"`
	Organization string                         `yaml:"organization" json:"organization"`
	Country      string                         `yaml:"country" json:"country"`
	Email        string                         `yaml:"email" json:"email"`
	Languages    []string                       `yaml:"languages" json:"languages"`
	TechStack    map[string]map[string][]string `yaml:"tech_stack" json:"tech_stack"`
	Socials      map[string]string              `yaml:"socials" json:"socials"`
	Timestamp    time.Time                      `json:"timestamp"`
}

package main

type DraftMetadata struct {
	ID            string              `json:"id" bson:"_id"`
	CurrentPicker string              `json:"current_picker" bson:"current_picker"`
	Pickers       []string            `json:"pickers" bson:"pickers"`
	Started       bool                `json:"started" bson:"started"`
	Picks         map[string][]string `json:"picks" bson:"-"`
	Pool          map[string][]string `json:"pool" bson:"-"`
}

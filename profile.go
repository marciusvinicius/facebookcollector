package main

type Profile struct {
	Name        string
	FacebokID   string
	PhoneNumber string
}

type Profiles []Profile

type Event struct {
	Lat float32
	Lng float32
}

type Events []Event

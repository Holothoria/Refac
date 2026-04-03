package main


import (
          "time"
          



)


type User struct {
	ID              int
	Username        string
	CreatedAt       time.Time
}

type Room struct {
	ID              int
	Name            string
	CreatedAt       time.Time
}

type Message struct {
	ID               int
	RoomID           int
	UserID           int
	Content          string
	CreatedAt        time.Time
}

package main

import (
    "time"
)

type Note struct {
    Id        int           `json:"id"`
    Name      string        `json:"name"`
    CreatedAt time.Time     `json:"created-at"`
}

type Notes []Note

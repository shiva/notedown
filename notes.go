package main

import (
    "time"
)

type Note struct {
    Name      string        `json:"name"`
    CreatedAt time.Time     `json:"created-at"`
}

type Notes []Note

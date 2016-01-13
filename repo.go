package main

import (
    "fmt"
    "time"
)

var currentId int

var notes Notes

func init() {
    RepoCreateNote(Note{Name: "Write presentation"})
    RepoCreateNote(Note{Name: "Host meetup"})
}

func RepoFindNote(id int) Note {
    for _, t := range notes {
        if t.Id == id {
            return t
        }
    }
    // return empty Note if not found
    return Note{}
}

func RepoCreateNote(t Note) Note {
    currentId += 1
    t.Id = currentId
    if t.CreatedAt.IsZero() {
        t.CreatedAt = time.Now()
    }

    notes = append(notes, t)
    return t
}

func RepoDestroyNote(id int) error {
    for i, t := range notes {
        if t.Id == id {
            notes = append(notes[:i], notes[i+1:]...)
            return nil
        }
    }
    return fmt.Errorf("Could not find Note with id of %d to delete", id)
}

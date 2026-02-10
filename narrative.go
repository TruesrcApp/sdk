package core

import (
	"time"
)

type Narrative struct {
	ID        string    `json:"id"`
	Content   string    `json:"content"`
	Hash      string    `json:"hash"`
	CreatedAt time.Time `json:"created_at"`
}

func PrepareNarrative(n Narrative) Narrative {
	n.CreatedAt = time.Now()
	n.Hash = Hash(n.Content)
	n.ID = n.Hash[:12]
	return n
}

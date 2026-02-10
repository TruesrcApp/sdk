package storage

import "truesrc/core"

type Store interface {
	Save(core.Narrative)
	Get(string) (core.Narrative, bool)
}

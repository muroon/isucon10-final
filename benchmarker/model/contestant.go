package model

import (
	"sync/atomic"

	"github.com/isucon/isucandar/agent"
	"github.com/isucon/isucon10-final/benchmarker/random"
)

const (
	CONTESTANT_ID_LENGTH       = 12
	CONTESTANT_PASSWORD_LENGTH = 64
)

type Contestant struct {
	ID                   string
	Password             string
	Name                 string
	IsStudent            bool
	latestNotificationID int64

	Agent *agent.Agent
}

func NewContestant() (*Contestant, error) {
	a, err := agent.NewAgent()
	if err != nil {
		return nil, err
	}

	id := random.Alphabet(CONTESTANT_ID_LENGTH)

	return &Contestant{
		ID:                   id,
		Password:             id, // random.Alphabet(CONTESTANT_PASSWORD_LENGTH),
		Name:                 random.Alphabet(20),
		IsStudent:            random.Pecentage(1, 10),
		latestNotificationID: 0,
		Agent:                a,
	}, nil
}

func NewAdmin() (*Contestant, error) {
	admin, err := NewContestant()
	if err != nil {
		return nil, err
	}

	admin.ID = "admin"
	admin.Password = "admin"
	return admin, nil
}

func (c *Contestant) LatestNotificationID() int64 {
	return atomic.LoadInt64(&c.latestNotificationID)
}

func (c *Contestant) UpdateLatestNotificationID(id int64) {
	if c.LatestNotificationID() < id {
		atomic.StoreInt64(&c.latestNotificationID, id)
	}
}

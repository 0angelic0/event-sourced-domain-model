package main

import (
	"time"

	"gorm.io/datatypes"
)

// PetORM a orm struct mapping to a table in database, implementing active record pattern
type PetORM struct {
	ID        string
	AggrID    string
	EventID   uint32
	EventType string
	CreatedAt time.Time
	EventBody datatypes.JSON
}

// TableName overrides the table name used by PetORM
func (PetORM) TableName() string {
	return "pets"
}

// NewPetORM a PetORM's constructor
func NewPetORM(e PetEvent) *PetORM {
	return &PetORM{
		ID:        e.GetEventUUID().String(),
		AggrID:    e.GetAggrID().String(),
		EventID:   e.GetEventID(),
		EventType: e.GetEventType(),
		CreatedAt: time.Now(),
		EventBody: datatypes.JSON(e.GetEventBody()),
	}
}

func toPetEvent(petORM *PetORM) PetEvent {
	switch petORM.EventType {
	case EventTypeInitialized:
		return NewPetEventInitializedByEventBody(petORM.EventBody)
	case EventTypeChangedName:
		return NewPetEventChangedNameByEventBody(petORM.EventBody)
	case EventTypeSold:
		return NewPetEventSoldByEventBody(petORM.EventBody)
	case EventTypeReturned:
		return NewPetEventReturnedByEventBody(petORM.EventBody)
	default:
		panic("invalid event type")
	}
}

func toPetEvents(petORMs []PetORM) []PetEvent {
	events := make([]PetEvent, 0, 10)
	for i := 0; i < len(petORMs); i++ {
		events = append(events, toPetEvent(&petORMs[i]))
	}
	return events
}

func toMapPetEvents(petORMs []PetORM) map[string][]PetEvent {
	m := make(map[string][]PetEvent)
	for i := 0; i < len(petORMs); i++ {
		_, ok := m[petORMs[i].AggrID]
		if ok == false {
			m[petORMs[i].AggrID] = make([]PetEvent, 0, 10)
		}
		m[petORMs[i].AggrID] = append(m[petORMs[i].AggrID], toPetEvent(&petORMs[i]))
	}
	return m
}

func toPetORM(e PetEvent) *PetORM {
	return NewPetORM(e)
}

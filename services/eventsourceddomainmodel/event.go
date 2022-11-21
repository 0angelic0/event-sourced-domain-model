package main

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

//////////////////////////////////////////////////
// Interface
//////////////////////////////////////////////////

// PetEvent interface
type PetEvent interface {
	GetEventUUID() uuid.UUID
	GetAggrID() uuid.UUID
	GetEventType() string
	GetEventID() uint32
	SetEventID(eventID uint32)
	GetTimeStamp() time.Time
	GetEventBody() []byte
}

//////////////////////////////////////////////////
// Base Event
//////////////////////////////////////////////////

// BaseEvent is a base event implementation
type BaseEvent struct {
	eventUUID uuid.UUID
	aggrID    uuid.UUID
	eventType string
	eventID   uint32
	timestamp time.Time
}

// NewBaseEvent is a constructor
func NewBaseEvent(aggrID uuid.UUID, eventType string) *BaseEvent {
	return &BaseEvent{
		eventUUID: uuid.New(),
		aggrID:    aggrID,
		eventType: eventType,
		timestamp: time.Now(),
	}
}

// GetEventUUID return event's uuid
func (e *BaseEvent) GetEventUUID() uuid.UUID {
	return e.eventUUID
}

// GetAggrID return event's aggregate's uuid
func (e *BaseEvent) GetAggrID() uuid.UUID {
	return e.aggrID
}

// GetEventType return event's type
func (e *BaseEvent) GetEventType() string {
	return e.eventType
}

// GetEventID return event's id
func (e *BaseEvent) GetEventID() uint32 {
	return e.eventID
}

// SetEventID set the event's id
func (e *BaseEvent) SetEventID(eventID uint32) {
	e.eventID = eventID
}

// GetTimeStamp return event's timestamp
func (e *BaseEvent) GetTimeStamp() time.Time {
	return e.timestamp
}

//////////////////////////////////////////////////
// Domain Events
//////////////////////////////////////////////////

const (
	// EventTypeInitialized is when a new pet initialized
	EventTypeInitialized = "pet-initialized"
	// EventTypeChangedName is when a pet was changed name
	EventTypeChangedName = "changed-name"
	// EventTypeSold is when a pet was sold
	EventTypeSold = "sold"
	// EventTypeReturned is when a pet was returned
	EventTypeReturned = "returned"
)

//------------------------------------------------
// PetEventInitialized
//------------------------------------------------

// PetEventInitialized is an event
type PetEventInitialized struct {
	*BaseEvent
	name    string
	age     int32
	addedAt time.Time
}

type petEventInitializedEventBody struct {
	ID        string `json:"id"`
	PetID     string `json:"pet-id"`
	EventID   uint32 `json:"event-id"`
	EventType string `json:"event-type"`
	TimeStamp string `json:"timestamp"`
	Name      string `json:"name"`
	Age       int32  `json:"age"`
	AddedAt   string `json:"added-at"`
}

// NewPetEventInitialized is a constructor for new event
func NewPetEventInitialized(aggrID uuid.UUID, name string, age int32, addedAt time.Time) *PetEventInitialized {
	return &PetEventInitialized{
		BaseEvent: NewBaseEvent(aggrID, EventTypeInitialized),
		name:      name,
		age:       age,
		addedAt:   addedAt,
	}
}

// NewPetEventInitializedByEventBody is a constructor for old event
func NewPetEventInitializedByEventBody(eventBody []byte) *PetEventInitialized {
	var eb petEventInitializedEventBody
	err := json.Unmarshal(eventBody, &eb)
	if err != nil {
		panic(err)
	}

	timestamp, err := time.Parse(time.RFC3339Nano, eb.TimeStamp)
	if err != nil {
		panic(err)
	}

	addedAt, err := time.Parse(time.RFC3339Nano, eb.AddedAt)
	if err != nil {
		panic(err)
	}

	return &PetEventInitialized{
		BaseEvent: &BaseEvent{
			eventUUID: uuid.MustParse(eb.ID),
			aggrID:    uuid.MustParse(eb.PetID),
			eventType: eb.EventType,
			eventID:   eb.EventID,
			timestamp: timestamp.UTC(),
		},
		name:    eb.Name,
		age:     eb.Age,
		addedAt: addedAt.UTC(),
	}
}

// GetEventBody return event body as a JSON bytes
func (p *PetEventInitialized) GetEventBody() []byte {
	body, err := json.Marshal(&petEventInitializedEventBody{
		ID:        p.eventUUID.String(),
		PetID:     p.aggrID.String(),
		EventID:   p.eventID,
		EventType: p.eventType,
		TimeStamp: p.timestamp.UTC().Format(time.RFC3339Nano),
		Name:      p.name,
		Age:       p.age,
		AddedAt:   p.addedAt.UTC().Format(time.RFC3339Nano),
	})
	if err != nil {
		panic(err)
	}

	return body
}

//------------------------------------------------
// PetEventChangedName
//------------------------------------------------

// PetEventChangedName is an event
type PetEventChangedName struct {
	*BaseEvent
	name string
}

type petEventChangedNameEventBody struct {
	ID        string `json:"id"`
	PetID     string `json:"pet-id"`
	EventID   uint32 `json:"event-id"`
	EventType string `json:"event-type"`
	TimeStamp string `json:"timestamp"`
	Name      string `json:"name"`
}

// NewPetEventChangedName is a constructor for new event
func NewPetEventChangedName(aggrID uuid.UUID, name string) *PetEventChangedName {
	return &PetEventChangedName{
		BaseEvent: NewBaseEvent(aggrID, EventTypeChangedName),
		name:      name,
	}
}

// NewPetEventChangedNameByEventBody is a constructor for old event
func NewPetEventChangedNameByEventBody(eventBody []byte) *PetEventChangedName {
	var eb petEventChangedNameEventBody
	err := json.Unmarshal(eventBody, &eb)
	if err != nil {
		panic(err)
	}

	timestamp, err := time.Parse(time.RFC3339Nano, eb.TimeStamp)
	if err != nil {
		panic(err)
	}

	return &PetEventChangedName{
		BaseEvent: &BaseEvent{
			eventUUID: uuid.MustParse(eb.ID),
			aggrID:    uuid.MustParse(eb.PetID),
			eventType: eb.EventType,
			eventID:   eb.EventID,
			timestamp: timestamp.UTC(),
		},
		name: eb.Name,
	}
}

// GetEventBody return event body as a JSON bytes
func (p *PetEventChangedName) GetEventBody() []byte {
	body, err := json.Marshal(&petEventChangedNameEventBody{
		ID:        p.eventUUID.String(),
		PetID:     p.aggrID.String(),
		EventID:   p.eventID,
		EventType: p.eventType,
		TimeStamp: p.timestamp.UTC().Format(time.RFC3339Nano),
		Name:      p.name,
	})
	if err != nil {
		panic(err)
	}

	return body
}

//------------------------------------------------
// PetEventSold
//------------------------------------------------

// PetEventSold is an event
type PetEventSold struct {
	*BaseEvent
	soldAt time.Time
}

type petEventSoldEventBody struct {
	ID        string `json:"id"`
	PetID     string `json:"pet-id"`
	EventID   uint32 `json:"event-id"`
	EventType string `json:"event-type"`
	TimeStamp string `json:"timestamp"`
	SoldAt    string `json:"sold-at"`
}

// NewPetEventSold is a constructor for new event
func NewPetEventSold(aggrID uuid.UUID, soldAt time.Time) *PetEventSold {
	return &PetEventSold{
		BaseEvent: NewBaseEvent(aggrID, EventTypeSold),
		soldAt:    soldAt,
	}
}

// NewPetEventSoldByEventBody is a constructor for old event
func NewPetEventSoldByEventBody(eventBody []byte) *PetEventSold {
	var eb petEventSoldEventBody
	err := json.Unmarshal(eventBody, &eb)
	if err != nil {
		panic(err)
	}

	timestamp, err := time.Parse(time.RFC3339Nano, eb.TimeStamp)
	if err != nil {
		panic(err)
	}

	soldAt, err := time.Parse(time.RFC3339Nano, eb.SoldAt)
	if err != nil {
		panic(err)
	}

	return &PetEventSold{
		BaseEvent: &BaseEvent{
			eventUUID: uuid.MustParse(eb.ID),
			aggrID:    uuid.MustParse(eb.PetID),
			eventType: eb.EventType,
			eventID:   eb.EventID,
			timestamp: timestamp.UTC(),
		},
		soldAt: soldAt.UTC(),
	}
}

// GetEventBody return event body as a JSON bytes
func (p *PetEventSold) GetEventBody() []byte {
	body, err := json.Marshal(&petEventSoldEventBody{
		ID:        p.eventUUID.String(),
		PetID:     p.aggrID.String(),
		EventID:   p.eventID,
		EventType: p.eventType,
		TimeStamp: p.timestamp.UTC().Format(time.RFC3339Nano),
		SoldAt:    p.soldAt.UTC().Format(time.RFC3339Nano),
	})
	if err != nil {
		panic(err)
	}

	return body
}

//------------------------------------------------
// PetEventReturned
//------------------------------------------------

// PetEventReturned is an event
type PetEventReturned struct {
	*BaseEvent
}

type petEventReturnedEventBody struct {
	ID        string `json:"id"`
	PetID     string `json:"pet-id"`
	EventID   uint32 `json:"event-id"`
	EventType string `json:"event-type"`
	TimeStamp string `json:"timestamp"`
}

// NewPetEventReturned is a constructor for new event
func NewPetEventReturned(aggrID uuid.UUID) *PetEventReturned {
	return &PetEventReturned{
		BaseEvent: NewBaseEvent(aggrID, EventTypeReturned),
	}
}

// NewPetEventReturnedByEventBody is a constructor for old event
func NewPetEventReturnedByEventBody(eventBody []byte) *PetEventReturned {
	var eb petEventReturnedEventBody
	err := json.Unmarshal(eventBody, &eb)
	if err != nil {
		panic(err)
	}

	timestamp, err := time.Parse(time.RFC3339Nano, eb.TimeStamp)
	if err != nil {
		panic(err)
	}

	return &PetEventReturned{
		BaseEvent: &BaseEvent{
			eventUUID: uuid.MustParse(eb.ID),
			aggrID:    uuid.MustParse(eb.PetID),
			eventType: eb.EventType,
			eventID:   eb.EventID,
			timestamp: timestamp.UTC(),
		},
	}
}

// GetEventBody return event body as a JSON bytes
func (p *PetEventReturned) GetEventBody() []byte {
	body, err := json.Marshal(&petEventReturnedEventBody{
		ID:        p.eventUUID.String(),
		PetID:     p.aggrID.String(),
		EventID:   p.eventID,
		EventType: p.eventType,
		TimeStamp: p.timestamp.UTC().Format(time.RFC3339Nano),
	})
	if err != nil {
		panic(err)
	}

	return body
}

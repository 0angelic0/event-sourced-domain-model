package main

import (
	"log"
	"time"

	"github.com/google/uuid"
)

//////////////////////////////////////////////////
// State Projector
//////////////////////////////////////////////////

// PetState is a default state projector for Pet aggregate
type PetState struct {
	id      uuid.UUID
	name    string
	age     int32
	addedAt time.Time
	isSold  bool
	soldAt  *time.Time
	version uint32
}

// NewPetState is a constructor
func NewPetState() *PetState {
	return &PetState{}
}

// ApplyEvent is event applier, this is the only way to do state transition
func (s *PetState) ApplyEvent(event PetEvent) {
	switch e := event.(type) {
	case *PetEventInitialized:
		s.applyEventInitialized(e)
	case *PetEventChangedName:
		s.applyEventChangedName(e)
	case *PetEventSold:
		s.applyEventSold(e)
	case *PetEventReturned:
		s.applyEventReturned(e)
	default:
		panic("Apply invalid event")
	}
}

func (s *PetState) applyEventInitialized(event *PetEventInitialized) {
	log.Printf("applyEventInitialized %v\n", *event)
	s.id = event.aggrID
	s.name = event.name
	s.age = event.age
	s.addedAt = event.addedAt
	s.isSold = false
	s.soldAt = nil
	s.version = 0
}

func (s *PetState) applyEventChangedName(event *PetEventChangedName) {
	log.Printf("applyEventChangedName %v\n", *event)
	s.name = event.name
	s.version++
}

func (s *PetState) applyEventSold(event *PetEventSold) {
	log.Printf("applyEventSold %v\n", *event)
	s.isSold = true
	s.soldAt = &event.soldAt
	s.version++
}

func (s *PetState) applyEventReturned(event *PetEventReturned) {
	log.Printf("applyEventReturned %v\n", *event)
	s.isSold = false
	s.soldAt = nil
	s.version++
}

//////////////////////////////////////////////////
// Model
//////////////////////////////////////////////////

// Pet represents a pet
type Pet struct {
	state             *PetState
	events            []PetEvent
	uncommittedEvents []PetEvent
}

// NewPet is a Pet's constructor for new object
func NewPet(name string, age int32) *Pet {
	p := &Pet{
		state:             NewPetState(),
		events:            make([]PetEvent, 0, 10),
		uncommittedEvents: make([]PetEvent, 0, 10),
	}
	initializedEvent := NewPetEventInitialized(uuid.New(), name, age, time.Now())
	p.appendEvent(initializedEvent)
	p.uncommittedEvents = append(p.uncommittedEvents, initializedEvent)
	return p
}

// NewPetByEvents is a Pet's constructor for old object
func NewPetByEvents(events []PetEvent) *Pet {
	p := &Pet{
		state:             NewPetState(),
		events:            make([]PetEvent, 0, 10),
		uncommittedEvents: make([]PetEvent, 0, 10),
	}
	for i := range events {
		p.appendEvent(events[i])
	}
	return p
}

func (p *Pet) appendEvent(event PetEvent) {
	p.events = append(p.events, event)
	p.state.ApplyEvent(event)
}

// ChangeName change a pet's name
func (p *Pet) ChangeName(name string) {
	changedNameEvent := NewPetEventChangedName(p.state.id, name)
	p.appendEvent(changedNameEvent)
	changedNameEvent.SetEventID(p.state.version)
	p.uncommittedEvents = append(p.uncommittedEvents, changedNameEvent)
}

// Sell pet from the pet shop
func (p *Pet) Sell() {
	soldEvent := NewPetEventSold(p.state.id, time.Now())
	p.appendEvent(soldEvent)
	soldEvent.SetEventID(p.state.version)
	p.uncommittedEvents = append(p.uncommittedEvents, soldEvent)
}

// Return pet to the pet shop
func (p *Pet) Return() {
	returnedEvent := NewPetEventReturned(p.state.id)
	p.appendEvent(returnedEvent)
	returnedEvent.SetEventID(p.state.version)
	p.uncommittedEvents = append(p.uncommittedEvents, returnedEvent)
}

// ClearUnCommittedEvents reset the uncommitted event list
func (p *Pet) ClearUnCommittedEvents() {
	p.uncommittedEvents = make([]PetEvent, 0, 10)
}

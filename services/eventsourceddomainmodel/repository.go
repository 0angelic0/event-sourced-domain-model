package main

// PetRepository interface to retrieve pets from data store
type PetRepository interface {
	FindByID(id string) (*Pet, error)
	FindAll() ([]*Pet, error)
	Save(p *Pet, originalVersion uint32) error
}

// PetRepositoryEventStore implement mysql data store
type PetRepositoryEventStore struct {
	es PetEventStore
}

// NewPetRepositoryEventStore create new instance of PetRepositoryMySQL
func NewPetRepositoryEventStore() *PetRepositoryEventStore {
	return &PetRepositoryEventStore{
		es: NewPetEventStoreMySQL(),
	}
}

// FindByID find a pet by id
func (r *PetRepositoryEventStore) FindByID(id string) (*Pet, error) {
	events, err := r.es.FetchByID(id)
	if err != nil {
		return nil, err
	}

	pet := NewPetByEvents(events)

	return pet, nil
}

// FindAll find all pets
func (r *PetRepositoryEventStore) FindAll() ([]*Pet, error) {
	aggrEvents, err := r.es.FetchAll()
	if err != nil {
		return nil, err
	}

	pets := make([]*Pet, 0, len(aggrEvents))
	for _, events := range aggrEvents {
		pet := NewPetByEvents(events)
		pets = append(pets, pet)
	}

	return pets, nil
}

// Save do save a pet
func (r *PetRepositoryEventStore) Save(p *Pet, originalVersion uint32) error {
	err := r.es.Append(p.state.id.String(), p.uncommittedEvents, originalVersion)
	if err != nil {
		return err
	}

	p.ClearUnCommittedEvents()
	return nil
}

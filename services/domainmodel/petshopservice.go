package main

// PetShopService is a stateless service object implementing transaction scripts
type PetShopService struct {
	repo PetRepository
}

// NewPetShopService create new instance of a PetShopService
func NewPetShopService(repo PetRepository) *PetShopService {
	return &PetShopService{
		repo: repo,
	}
}

// GetPetByID return a pet by id
func (s *PetShopService) GetPetByID(id string) (*Pet, error) {
	pet, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return pet, nil
}

// GetPetList return a list of pets
func (s *PetShopService) GetPetList() ([]*Pet, error) {
	pets, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return pets, nil
}

// AddPet add a new pet to the shop
func (s *PetShopService) AddPet(name string, age int32) (*Pet, error) {
	pet := NewPet(name, age)

	err := s.repo.Save(pet)
	if err != nil {
		return nil, err
	}

	return pet, nil
}

// ChangePetName change a pet's name
func (s *PetShopService) ChangePetName(id string, name string) error {
	pet, err := s.GetPetByID(id)
	if err != nil {
		return err
	}

	pet.ChangeName(name)

	err = s.repo.Save(pet)
	if err != nil {
		return err
	}
	return nil
}

// SellPet sell a pet
func (s *PetShopService) SellPet(id string) error {
	pet, err := s.GetPetByID(id)
	if err != nil {
		return err
	}

	pet.Sell()

	err = s.repo.Save(pet)
	if err != nil {
		return err
	}
	return nil
}

// ReturnPet return a pet to the shop
func (s *PetShopService) ReturnPet(id string) error {
	pet, err := s.GetPetByID(id)
	if err != nil {
		return err
	}

	pet.Return()

	err = s.repo.Save(pet)
	if err != nil {
		return err
	}
	return nil
}

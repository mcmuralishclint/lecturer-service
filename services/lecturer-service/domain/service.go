package domain

type Service interface {
	Find(name_map string) Skill
	Create() (bool, error)
	Delete() (bool, error)
	FindAll() []Skill
}

type Repository interface {
	Find(name_map string) Skill
	Create() (bool, error)
	Delete() (bool, error)
	FindAll() []Skill
}

type service struct {
	skillrepo Repository
}

func NewSkillService(skillrepo Repository) Service {
	return &service{skillrepo: skillrepo}
}

func (s *service) Find(name_map string) Skill {
	return s.skillrepo.Find(name_map)
}
func (s *service) Create() (bool, error) {
	return s.skillrepo.Create()
}
func (s *service) Delete() (bool, error) {
	return s.skillrepo.Delete()
}
func (s *service) FindAll() []Skill {
	return s.skillrepo.FindAll()
}

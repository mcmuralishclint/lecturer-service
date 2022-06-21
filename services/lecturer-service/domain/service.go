package domain

type Service interface {
	Find(name_map string) Skill
	Create(skill Skill) (Skill, error)
	Delete(name_map string) (bool, error)
	FindAll() []Skill
}

type Repository interface {
	Find(name_map string) Skill
	Create(skill Skill) (Skill, error)
	Delete(name_map string) (bool, error)
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
func (s *service) Create(skill Skill) (Skill, error) {
	return s.skillrepo.Create(skill)
}
func (s *service) Delete(name_map string) (bool, error) {
	return s.skillrepo.Delete(name_map)
}
func (s *service) FindAll() []Skill {
	return s.skillrepo.FindAll()
}

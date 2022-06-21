package domain

type Service interface {
	FindSkill(name_map string) Skill
	CreateSkill(skill Skill) (Skill, error)
	DeleteSkill(name_map string) (bool, error)
	FindAllSkills() []Skill
}

type Repository interface {
	FindSkill(name_map string) Skill
	CreateSkill(skill Skill) (Skill, error)
	DeleteSkill(name_map string) (bool, error)
	FindAllSkills() []Skill
}

type service struct {
	skillrepo Repository
}

func NewSkillService(skillrepo Repository) Service {
	return &service{skillrepo: skillrepo}
}

func (s *service) FindSkill(name_map string) Skill {
	return s.skillrepo.FindSkill(name_map)
}
func (s *service) CreateSkill(skill Skill) (Skill, error) {
	return s.skillrepo.CreateSkill(skill)
}
func (s *service) DeleteSkill(name_map string) (bool, error) {
	return s.skillrepo.DeleteSkill(name_map)
}
func (s *service) FindAllSkills() []Skill {
	return s.skillrepo.FindAllSkills()
}

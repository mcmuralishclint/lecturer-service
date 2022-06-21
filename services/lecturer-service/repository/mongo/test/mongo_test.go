package mongo_test

import (
	"github.com/mcmuralishclint/personal_tutor/mocks"
	"github.com/mcmuralishclint/personal_tutor/services/lecturer-service/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindSkills(t *testing.T) {
	service := &mocks.Service{}
	service.On("FindAll").Return([]domain.Skill{})
	res := service.FindAll()
	assert.Equal(t, res, []domain.Skill{})
}

func TestFindSkill(t *testing.T) {
	service := &mocks.Service{}
	skill := domain.Skill{}
	service.On("Find", "good").Return(skill)

	res := service.Find("good")
	assert.Equal(t, res.NameMap, skill.NameMap)
}

package forms

import (
	"bitbucket.org/SummerCampDev/summercamp/models"
)

// TODO: consider MaxSize and Length values
type Vacancy struct {
	FormModel
	ID int `json:"id" valid:"Required;"`
	Name string `json:"name" valid:"Required; MaxSize(100)"`
	Description string `json:"description" valid:"Required; MaxSize(1000)"`
	Skills []int `json:"skills" valid:"Required; Length(10)"`
	Spheres []int `json:"spheres" valid:"Required; Length(3)"`
	TeamID int `json:"team_id" valid:"Required;"`
}

func (v *Vacancy) Save() (*Vacancy, bool) {
	vacancy := models.Vacancy{
		ID:v.ID,
		Name:v.Name,
		Description:v.Description,
		TeamID:v.TeamID,
	}
	// TODO: vacancy skill
	// TODO: vacancy spheres
	ok := vacancy.Save()
	return vacancy, ok
}

type VacancyUpdate struct {
	FormModel
	Field string
	Value string
}

func (vu *VacancyUpdate) Update() bool {
	switch vu.Field {
	case "skill":
	case "sphere":
		default:

	}
	return true
}




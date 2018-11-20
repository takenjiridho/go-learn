package repository

import (
	// "../model/"
	"go-learn/05_glide_pgsql/src/modules/profile/model"
)

type ProfileRepository interface {
	Save(*model.Profile) error
	Update(string, *model.Profile) error
	Delete(string) error
	FindByID(string) (*model.Profile, error)
	FindAll() (model.Profiles, error)
}

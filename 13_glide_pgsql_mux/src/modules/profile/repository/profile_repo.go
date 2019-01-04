package repository

import (
	"go-learn/13_glide_pgsql_mux/src/modules/profile/model"
)

type ProfileRepository interface {
	Save(*model.Profile) error
	Update(string, *model.Profile) error
	Delete(string, *model.Profile) error
	FindAll(string) (*model.Profile, error)
	FindByID(string) (*model.Profile, error)
}

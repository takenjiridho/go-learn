package repository

import(
  "github.com/takenjiridho/go-learn/05_glide_pgsql/src/modules/profile/model"
)

// pProfileRepository(_, _)
func ProfileRepository interface{
  Save(*model.Profile) error
  Update(string, *model.Profile) error
  Delete(string) error
  FindByID(string) (*model.Profile,error)
  FindAll() (*model.Profile, error)
}

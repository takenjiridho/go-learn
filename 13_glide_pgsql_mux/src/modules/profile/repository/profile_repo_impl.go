package repository

import (
	"database/sql"
	"go-learn/05_glide_pgsql/src/modules/profile/model"
)

type profileImpl struct {
	db *sql.DB
}

// NNewProfileImpl(db)
func NewProfileImpl(db *sql.DB) *profileImpl {
	return &profileImpl{db}
}

func (r *profileImpl) Save(profile *model.Profile) error {
	query := `insert into "profile"("id","first_name", "last_name","email","password","created_at","updated_at")
            values($1,$2,$3,$4,$5,$6,$7)`

	stat, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer stat.Close()

	_, err = stat.Exec(profile.ID, profile.FirstName, profile.LastName, profile.Email, profile.Password, profile.CreatedAt, profile.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (r *profileImpl) Update(id string, profile *model.Profile) error {
	query := `UPDATE "profile" set "first_name"=$1, "last_name"=$2, "email"=$3, "password"=$4,"updated_at"=$5 where "id"=$6 `

	stat, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer stat.Close()

	_, err = stat.Exec(profile.FirstName, profile.LastName, profile.Email, profile.Password, profile.UpdatedAt, id)

	if err != nil {
		return err
	}

	return nil

}

func (r *profileImpl) Delete(id string) error {
	query := `DELETE FROM "profile" where "id"=$1`

	stat, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer stat.Close()

	_, err = stat.Exec(id)

	if err != nil {
		return err
	}

	return nil
}

func (r *profileImpl) FindByID(id string) (*model.Profile, error) {
	query := `SELECT id, first_name, last_name, email, password, created_at, updated_at  FROM "profile" where "id"=$1`

	var profile model.Profile

	stat, err := r.db.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer stat.Close()

	err = stat.QueryRow(id).Scan(&profile.ID, &profile.FirstName, &profile.LastName, &profile.Email, &profile.Password, &profile.CreatedAt, &profile.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &profile, nil

}

func (r *profileImpl) FindAll() (model.Profiles, error) {
	query := `SELECT id, first_name, last_name, email, password, created_at, updated_at FROM "profile"`

	var profiles model.Profiles

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		var profile model.Profile

		err = rows.Scan(&profile.ID, &profile.FirstName, &profile.LastName, &profile.Email, &profile.Password, &profile.CreatedAt, &profile.UpdatedAt)

		// fmt.Println(" nama ", profile.FirstName)

		profiles = append(profiles, profile)
	}

	return profiles, nil
}

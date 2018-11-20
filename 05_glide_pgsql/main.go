package main

import (
	"fmt"

	"go-learn/05_glide_pgsql/config"
	"go-learn/05_glide_pgsql/src/modules/profile/model"
	"go-learn/05_glide_pgsql/src/modules/profile/repository"
)

func main() {

	fmt.Println("glide , docker , curd with pgsql and pattern ")

	db, err := config.GetPostgresDB()

	if err != nil {
		fmt.Println(err)
	}

	// err = db.Ping()

	profileImpl := repository.NewProfileImpl(db)

	f := model.NewProfile()
	f.ID = "P3"
	f.FirstName = "farid"
	f.LastName = "wjd"
	f.Email = "farid@gmail.com"
	f.Password = "123456"

	b, err := saveProfile(f, profileImpl)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("message : ", b)

	// profile, err := getProfile("P2", profileImpl)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("ID : ", profile.ID, "Nama :  ", profile.FirstName, profile.LastName, " Email : ", profile.Email)

	// u := model.NewProfile()
	// u.ID = "1"
	// u.FirstName = "farid aja deh"
	// u.LastName = "wijdan"
	// u.Email = "farid-ganteng-lagi@gmail.com"
	// u.Password = "123456-ganbate"
	//
	// bd, err := updateProfile(u.ID, u, profileImpl)
	//
	// if err != nil {
	// 	fmt.Println(err)
	// }
	//
	// fmt.Println("message : ", bd)

	// b, err := deleteProfile("1", profileImpl)
	//
	// if err != nil {
	// 	fmt.Println(err)
	// }
	//
	// fmt.Println("message : ", b)

	profiles, err := getProfiles(profileImpl)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("=========================")

	for _, v := range profiles {
		// fmt.Println(v)
		fmt.Println("ID : ", v.ID, "Nama :  ", v.FirstName, v.LastName, " Email : ", v.Email)
	}

}

func saveProfile(p *model.Profile, r repository.ProfileRepository) (string, error) {
	err := r.Save(p)

	if err != nil {
		return "", err
	}

	b := "data berhasil diinsert kan bang!!!"

	return b, err
}

func getProfile(id string, repo repository.ProfileRepository) (*model.Profile, error) {
	profile, err := repo.FindByID(id)

	if err != nil {
		return nil, err
	}

	return profile, nil
}

func updateProfile(id string, p *model.Profile, r repository.ProfileRepository) (string, error) {
	err := r.Update(id, p)

	if err != nil {
		return "", err
	}

	b := "data berhasil disimpan bang!!"

	return b, err
}

func deleteProfile(id string, r repository.ProfileRepository) (string, error) {
	err := r.Delete(id)

	if err != nil {
		return "", err
	}

	b := "data berhasil dihapus !!"

	return b, err

}

func getProfiles(r repository.ProfileRepository) (model.Profiles, error) {
	profiles, err := r.FindAll()

	if err != nil {
		return nil, err
	}

	return profiles, err

}

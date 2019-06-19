package service

import (
	"fmt"
	"Makanan/repository"
	"Makanan/model"
	"Makanan/driver"
)

func ReadAllUsers() []model.UserModel{
	db, err := driver.Connect()

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	defer db.Close()

    var result []model.UserModel

	items, err := db.Query(repository.ReadAll())
	if err  != nil {
		fmt.Println(err.Error())
		return nil
	}

	for items.Next() {
		var each = model.UserModel{}
		var err = items.Scan(&each.Id, &each.NamaUser,&each.Username,&each.Password,&each.IdRole,&each.Phone,&each.Email)

		if err != nil {
			fmt.Println(err.Error())
		}

		result = append(result, each)

	}

	if err = items.Err(); err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return result

}
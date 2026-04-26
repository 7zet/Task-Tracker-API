package internal

import (
	"encoding/json"
	"os"
)

var Users []User
var Tasks []Task



// 1. faylga saqlash funksiyasi
func SaveToFile() error {
	// sliceni json formatga ogiramiz
	data, err := json.MarshalIndent(Users, "", "  ")
	if err != nil {
		return err
	}
	
	// users.json degan faylga yozamiz(0644 oqish huquqi)
	return os.WriteFile("users.json", data, 0644)

}


// 2. fayldan oqish funksiyasi(dastur yonganda bir marta ishga tushiriladi)
func LoadUserFromFile() error {
	//faylni ochamiz
	data, err := os.ReadFile("users.json")
	if err != nil {
		// agar fayl yoq bolsa xato bolmasligi uchun
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	// jsonni yana slicega qaytaramiz
	return json.Unmarshal(data, &Users)
}


func TaskSaver() error {
	data, err := json.MarshalIndent(Tasks, "", "  ")

	if err != nil {
		return err
	}

	return os.WriteFile("tasks.json", data, 0644)

}

func LoadTaskFromFile() error {
	data, err := os.ReadFile("tasks.json")
	if err != nil{
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	return json.Unmarshal(data, &Users)
}

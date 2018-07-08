package data

import (
	"fmt"
)

//Profile структура для хранения записи адресной книги
type Profile struct {
	Name        string `json:"name"`
	Nickname    string `json:"nickname"`
	Email       string `json:"email"`
	Description string `json:"description"`
	Birthday    string `json:"birthday"`
	Photos      string `json:"photos"`
}

// profile хранимый список профилей
var profiles []Profile

// GetProfiles возращает список профилей
func GetProfiles() []Profile {
	return profiles
}

// AddProfile добавляет профиль profile в конец списка и возращает id
func AddProfile(profile Profile) int {
	id := len(profiles)
	profiles = append(profiles, profile)
	return id
}

// EditProfile изменяет профиль c id на profile
func EditProfile(profile Profile, id int) error {
	if id < 0 || id >= len(profiles) {
		return fmt.Errorf("incorrect ID")
	}
	profiles[id] = profile
	return nil
}

// RemoveProfile удаляет профиль по id
func RemoveProfile(id int) error {
	if id < 0 || id >= len(profiles) {
		return fmt.Errorf("incorrect ID")
	}
	profiles = append(profiles[:id], profiles[id+1:]...)
	return nil
}

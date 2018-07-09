package data

import (
	"fmt"
	"image/color"
)

// Profile структура для хранения записи адресной книги
type Profile struct {
	Name        string `json:"name"`
	Nickname    string `json:"nickname"`
	Email       string `json:"email"`
	Description string `json:"description"`
	Birthday    string `json:"birthday"`
	Photos      string `json:"photos"`
}

// Photo структура для хранения фотографий
type Photo struct {
	ColorModel    color.Model
	Width, Height int
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

// photo хранимый список фотографий
var photos []Photo

// GetPhotos возращает список фотографий
func GetPhotos() []Photo {
	return photos
}

// AddPhoto добавляет фотогрфии photo в конец списка и возращает id
func AddPhoto(photo Photo) int {
	id := len(photos)
	profiles = append(photos, photo)
	return id
}

// EditPhoto изменяет профиль c id на profile
func EditPhoto(photo Photo, id int) error {
	if id < 0 || id >= len(photos) {
		return fmt.Errorf("incorrect ID")
	}
	photos[id] = photo
	return nil
}

// RemovePhoto удаляет фотографии по id
func RemovePhoto(id int) error {
	if id < 0 || id >= len(photos) {
		return fmt.Errorf("incorrect ID")
	}
	photos = append(photos[:id], photos[id+1:]...)
	return nil
}

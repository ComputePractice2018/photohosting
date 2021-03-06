package data

import (
	"fmt"
	"image/color"
)

// Profile структура для хранения записи профиля
type Profile struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Nickname    string `json:"nickname"`
	Email       string `json:"email"`
	Description string `json:"description"`
	Birthday    string `json:"birthday"`
	Photos      string `json:"photos"`
}

// ProfileList структрура для списка профилей
type ProfileList struct {
	profiles []Profile
}

// Editable интерфейс для работы со списком записей
type Editable interface {
	GetProfiles() []Profile
	AddProfile(profile Profile) int
	EditProfile(profile Profile, id int) error
	RemoveProfile(id int) error
	GetPhotos() []Photo
	AddPhoto(photo Photo) int
	EditPhoto(photo Photo, id int) error
	RemovePhoto(id int) error
}

// NewProfileList конструктор списка профилей
func NewProfileList() *ProfileList {
	return &ProfileList{}
}

// GetProfiles возращает список профилей
func (cl *ProfileList) GetProfiles() []Profile {
	return cl.profiles
}

// AddProfile добавляет профиль profile в конец списка и возращает id
func (cl *ProfileList) AddProfile(profile Profile) int {
	id := len(cl.profiles) + 1
	profile.ID = id
	cl.profiles = append(cl.profiles, profile)
	return id
}

// EditProfile изменяет профиль c id на profile
func (cl *ProfileList) EditProfile(profile Profile, id int) error {
	if id < 1 || id > len(cl.profiles) {
		return fmt.Errorf("incorrect ID")
	}
	cl.profiles[id-1] = profile
	return nil
}

// RemoveProfile удаляет профиль по id
func (cl *ProfileList) RemoveProfile(id int) error {
	if id < 1 || id > len(cl.profiles) {
		return fmt.Errorf("incorrect ID")
	}
	id--
	cl.profiles = append(cl.profiles[:id], cl.profiles[id+1:]...)
	return nil
}

// Photo структура для хранения фотографий
type Photo struct {
	ColorModel    color.Model
	Width, Height int
}

// PhotoList структрура для списка фотографий
type PhotoList struct {
	photos []Photo
}

// NewPhotoList конструктор списка фотографий
func NewPhotoList() *PhotoList {
	return &PhotoList{}
}

// GetPhotos возращает список фотографий
func (cl *PhotoList) GetPhotos() []Photo {
	return cl.photos
}

// AddPhoto добавляет фотогрфии photo в конец списка и возращает id
func (cl *PhotoList) AddPhoto(photo Photo) int {
	id := len(cl.photos) + 1
	photo.ID = id
	cl.photos = append(cl.photos, photo)
	return id
}

// EditPhoto изменяет фотографию c id на photo
func (cl *PhotoList) EditPhoto(photo Photo, id int) error {
	if id < 1 || id > len(cl.photos) {
		return fmt.Errorf("incorrect ID")
	}
	cl.photos[id-1] = photo
	return nil
}

// RemovePhoto удаляет фотографии по id
func (cl *PhotoList) RemovePhoto(id int) error {
	if id < 1 || id > len(cl.photos) {
		return fmt.Errorf("incorrect ID")
	}
	id--
	cl.photos = append(cl.photos[:id], cl.photos[id+1:]...)
	return nil
}

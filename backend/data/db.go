package data

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type DBProfileList struct {
	db *gorm.DB
}

func NewDBProfileList(connection string) (*DBProfileList, error) {
	db, err := gorm.Open("mysql", connection)
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&Profile{})
	return &DBProfileList{db: db}, db.Error
}

func (cl *DBProfileList) Close() {
	cl.db.Close()
}

func (cl *DBProfileList) GetProfiles() []Profile {
	var profiles []Profile
	cl.db.Find(&profiles)
	return profiles
}

func (cl *DBProfileList) AddProfile(profile Profile) int {
	cl.db.Create(&profile)
	return profile.ID
}

func (cl *DBProfileList) EditProfile(profile Profile, id int) error {
	var profiles []Profile
	cl.db.Where("id = ?", id).Find(&profiles)
	if len(profiles) == 0 {
		return fmt.Errorf("can't find record #%d", id)
	}

	profile.ID = profiles[0].ID
	cl.db.Save(&profile)
	return cl.db.Error
}

func (cl *DBProfileList) RemoveProfile(id int) error {
	var profiles []Profile
	cl.db.Where("id = ?", id).Find(&profiles)
	if len(profiles) == 0 {
		return fmt.Errorf("can't find record #%d", id)
	}
	cl.db.Delete(&profiles[0])
	return cl.db.Error
}

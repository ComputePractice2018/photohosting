package data

import "testing"

var testProfiles = []Profile{
	{
		Name:        "Ivan",
		Nickname:    "Ivivan44",
		Email:       "iivanov@mail.com",
		Description: "О себе",
		Birthday:    "9.2.1979",
		Photos:      "Фотографии",
	},
	{
		Name:        "Sergey",
		Nickname:    "Serg0vich",
		Email:       "ssergeev@mail.com",
		Description: "О себе",
		Birthday:    "28.11.1981",
		Photos:      "Фотографии",
	},
}

func TestAddProfile(t *testing.T) {
	cl := NewProfileList()

	cl.AddProfile(testProfiles[0])

	if cl.GetProfiles()[0] != testProfiles[0] {
		t.Errorf("AddProfile is not working")
	}
}

func TestEditProfile(t *testing.T) {
	cl := NewProfileList()
	cl.AddProfile(testProfiles[0])

	err := cl.EditProfile(testProfiles[1], 0)

	if cl.GetProfiles()[0] != testProfiles[1] {
		t.Errorf("EditProfile is not working")
	}
	if err != nil {
		t.Errorf("Unexpected EditProfile error")
	}

	err = cl.EditProfile(testProfiles[1], -1)
	if err == nil {
		t.Errorf("Nothandled out of range error")
	}

}

func TestRemoveProfile(t *testing.T) {
	cl := NewProfileList()
	cl.AddProfile(testProfiles[0])
	cl.AddProfile(testProfiles[1])

	err := cl.RemoveProfile(0)

	if cl.GetProfiles()[0] != testProfiles[1] {
		t.Errorf("RemoveProfile is not working")
	}
	if err != nil {
		t.Errorf("Unexpected RemoveProfile error")
	}

	err = cl.RemoveProfile(-1)
	if err == nil {
		t.Errorf("Nothandled out of range error")
	}

}

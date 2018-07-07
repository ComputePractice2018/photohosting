package data

//Profile структура для хранения записи адресной книги
type Profile struct {
	Name        string `json:"name"`
	Nickname    string `json:"nickname"`
	Email       string `json:"email"`
	Description string `json:"description"`
	Birthday    string `json:"birthday"`
	Photos      string `json:"photos"`
}

//ProfileList список профилей
var ProfileList = []Profile{Profile{

	Name:        "Имя",
	Nickname:    "Псевдоним",
	Email:       "user@domain.com",
	Description: "О себе",
	Birthday:    "День рождения",
	Photos:      "Фотографии:"}}

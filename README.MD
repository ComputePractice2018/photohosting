Usecases


1. Как пользователь, я хочу иметь возможность видеть все свои фотографии
1. Как пользователь, я хочу иметь возможность добавлять свои фотографии
1. Как пользователь, я хочу иметь возможность удалять свои фотографии, чтобы иметь желаемый вид профиля
1. Как пользователь, я хочу иметь возможность редактировать информацию о себе в своём профиле (имя, псевдоним, почта, дата рождения, описание)
1. Как пользователь, я хочу иметь возможность редактировать описание своих фотографий


REST API


GET /api/photohosting/profile
Ответ: 200 OK

    [{
        "name": "Имя",
        "nickname": "Псевдоним",
        "email": "user@domain.com",
        "description": "О себе",
        "birthday": "День рождения",
        "photos": "Фотографии:",
    }]
POST /api/photohosting/profile
Тело запроса:

    {
        "name": "Имя",
        "nickname": "Псевдоним",
        "email": "user@domain.com",
        "description": "О себе",
        "birthday": "День рождения",
        "github" "user"
    }
Ответ: 201 created

Location: /api/photohosting/profile/1

PUT /api/photohosting/profile/1
Тело запроса:

    {
        "name": "Имя",
        "nickname": "Псевдоним",
        "email": "user@domain.com",
        "description": "О себе",
        "birthday": "День рождения",
        "github" "user"
    }
Ответ: 202 accepted

Location: /api/photohosting/profile/1

DELETE /api/photohosting/profile/1
Ответ: 204 No content

POST /api/photohosting/profile/1/photos
Тело запроса:

    {
        "photos": "Фотографии:",
    }
Ответ: 200 OK Location: /api/photohosting/profile/1/photos/1

DELETE /api/photohosting/profile/1/photos/1
Ответ: 200 OK

GET /api/photohosting/profile/1/photos
Тело запроса:

    {
        "photos": "Фотографии:",
    }
Ответ: 200 OK

PUT /api/photohosting/profile/1/photos/1/description
Тело запроса:

    {
        "description": "Описание:",
    }
Ответ: 200 OK Location: /api/photohosting/profile/1/photos/1/description
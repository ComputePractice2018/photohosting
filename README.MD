Usecases

Как пользователь, я хочу иметь возможность видеть все свои фотографии
Как пользователь, я хочу иметь возможность добавлять свои фотографии
Как пользователь, я хочу иметь возможность удалять свои фотографии, чтобы иметь желаемый вид профиля
Как пользователь, я хочу иметь возможность редактировать информацию о себе в своём профиле (имя, псевдоним, почта, дата рождения, описание)
Как пользователь, я хочу иметь возможность редактировать описание своих фотографий
REST API

GET /api/photohosting/profile Ответ: 200 OK

[{
    "name": "Имя",
    "nickname": "Псевдоним",
    "email": "user@domain.com",
    "description": "О себе",
    "birthday": "День рождения",
    "photos": "Фотографии:",
}]
POST /api/photohosting/profile Тело запроса:

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

PUT /api/photohosting/profile/1 Тело запроса:

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

DELETE /api/photohosting/profile/1 Ответ: 204 No content

POST /api/photohosting/profile/1/photos Тело запроса:

{
    "photos": "Фотографии:",
}
Ответ: 200 OK Location: /api/photohosting/profile/1/photos/1

DELETE /api/photohosting/profile/1/photos/1 Ответ: 200 OK

GET /api/photohosting/profile/1/photos Тело запроса:

{
    "photos": "Фотографии:",
}
Ответ: 200 OK

PUT /api/photohosting/profile/1/photos/1/description Тело запроса:

{
    "description": "Описание:",
}
Ответ: 200 OK Location: /api/photohosting/profile/1/photos/1/description

## Как собрать и запустить

Backend:

```bat
cd backend
docker build -f Dockerfile -t photohostingbackend:<имя ветки> .
docker run --rm --name photohosting -e NAME=<параметр приложения> photohostingbackend:<имя ветки>
```

Frontend:

```bat
cd backend
docker build -f Dockerfile -t photohostingfrontend:<имя ветки> .
docker run -d --rm --name photohostingfrontend -p 80:80 photohostingfrontend:<имя ветки>
```
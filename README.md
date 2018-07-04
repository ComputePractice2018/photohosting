# photohosting

## Usecases

1. Как пользователь, я хочу иметь возможность видеть все свои фотографии
1. Как пользователь, я хочу иметь возможность добавлять свои фотографии
1. Как пользователь, я хочу иметь возможность удалять свои фотографии, чтобы иметь желаемый вид профиля
1. Как пользователь, я хочу иметь возможность видеть фотографии всех своих друзей на которых я подписан
1. Как пользователь, я хочу иметь возможность управлять своими подписками (подписаться, отписаться)
1. Как пользователь, я хочу иметь возможность редактировать информацию о себе в своём профиле (имя, псевдоним, почта, дата рождения, описание)
1. Как пользователь, я хочу иметь возможность видеть список ллюдей на которых я подписан


## REST API

### GET /api/photohosting/profile

Ответ: 200 OK

```json
    [{
        "name": "Имя",
        "nickname": "Псевдоним",
        "e-mail": "user@domain.com",
        "description": "О себе",
        "birthday": "День рождения",
        "photos": "Фотографии:",
        
    }]

```

### POST /api/photohosting/profile

Тело запроса:

```json
    {
        "name": "Имя",
        "nickname": "Псевдоним",
        "e-mail": "user@domain.com",
        "description": "О себе",
        "birthday": "День рождения",
        "github" "user"
    }
```

Ответ: 201 created

Location: /api/photohosting/profile/1


 ###  PUT /api/photohosting/profile/1

Тело запроса:

```json
    {
        "name": "Имя",
        "nickname": "Псевдоним",
        "e-mail": "user@domain.com",
        "description": "О себе",
        "birthday": "День рождения",
        "github" "user"
    }
```

Ответ: 202 accepted

Location: /api/photohosting/profile/1



###  DELETE /api/photohosting/profile/1
Ответ: 204 No content


### POST /api/photohosting/profile/1/photos

Тело запроса:

```json
    {
        "photos": "Фотографии:",
    }
```

Ответ: 200 OK
Location: /api/photohosting/profile/1/photos/1


### DELETE /api/photohosting/profile/1/photos/1
Ответ: 200 OK



### GET /api/photohosting/profile/1/photos/1


Тело запроса:

```json
    {
        "photos": "Фотографии:",
    }
```
Ответ: 200 OK


### GET /api/photohosting/profile/1/friends/1


Тело запроса:

```json
    [{
        "name": "Имя",
        "nickname": "Псевдоним",
        "e-mail": "user@domain.com",
        "description": "О себе",
        "birthday": "День рождения",
        "photos": "Фотографии:",
        
    }]

Ответ: 200 OK
Location: /api/photohosting/profile/1/friends/1



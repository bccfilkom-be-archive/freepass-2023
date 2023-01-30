# freepass-2023 bcc_university

In this digital age, transparency and management of information is greatly needed, especially in the field of education. At some point, there is an investor who establishes a university, while the university requires a system for managing information from its students.

This application is intended for university students to be able to access information provided by the university. Developed in Golang gin backend framework and uses Mongodb as the database management service, mapped using mongo go models (mgm)

## Setup API Service and Container Image
```
docker compose --env-file app.env up
```

## API

### `POST /register`
User register

Example request body
```
{
    "username": "naruto"
    "email": "narutouzumaki@gmail.com",
    "password": "hokage123"
}
```
Response
```
{
  "username": "naruto"
}
```
### `POST /login`
User login using JWT auth and returns the jwt token

Example request body
```
{
    "email": "narutouzumaki@gmail.com",
    "password": "hokage123"
}
```
Response
```
{
  "status": "success"
  "token": jwt_token_goes_here
}
```

### `GET /user/profile`
Get current user profile

Response
```
{
    "username": "naruto",
    "email": "narutouzumaki@gmail.com",
    "first_name": "Naruto",
    "last_name": "Uzumaki",
    "classes_enrolled": {
        "63d460a6f1081c5158698c74": 3,
        "63d460d0f1081c5158698c76": 2
    }
}
```

### `GET /class`
Get all class

Response
```
{
    "classes": [
        {
            "id": "63d46056f1081c5158698c73",
            "created_at": "0001-01-01T00:00:00Z",
            "updated_at": "0001-01-01T00:00:00Z",
            "title": "Learn Python: The Complete Python Programming Course",
            "sks": 3,
            "participants": {
                "63d3240132c417a6c2594175": true
            }
        },
        {
            "id": "63d460a6f1081c5158698c74",
            "created_at": "0001-01-01T00:00:00Z",
            "updated_at": "0001-01-01T00:00:00Z",
            "title": "Public Relations: Media Crisis Communications",
            "sks": 3,
            "participants": {}
        },
        {
            "id": "63d460d0f1081c5158698c76",
            "created_at": "0001-01-01T00:00:00Z",
            "updated_at": "2023-01-29T09:28:00.714Z",
            "title": "Speak Like a Pro: Public Speaking for Professionals",
            "sks": 2,
            "participants": {
                "63d43adbbad27e360a6e67b0": true
            }
        }
    ]
}
```

### `PUT /user/profile/edit`
Edit and update current user profile

Example request body
```
{
    "edit_map": {
        "first_name": "Sasuke",
        "last_name": "Uchiha"
        "email": "uchihasasuke@gmail.com",
    }
}
```
Response
```
{
    "message": "successfully edited"
}
```

### `POST /myclass/add-class/{classId}`
Add new class by extracting the classId parameter

Response
```
{
    "message": "successfully added new class"
}
```

### `DELETE /myclass/delete-class/{classId}`
Drop class by extracting the classId parameter

Response
```
{
    "message": "successfully dropped a class"
}
```

### `GET /myclass/{classId}/participants`
View all class participants by extracting the classId parameter

Response
```
{
    "participants": [
        {
            "id": "63d43adbbad27e360a6e67b0",
            "created_at": "2023-01-27T20:58:03.71Z",
            "updated_at": "2023-01-29T09:28:00.648Z",
            "username": "clark_kent",
            "email": "clarkkent@gmail.com",
            "password": "$2a$10$fh4vJraOT6BiqPVfhwl2vuCIHuCPeGHxm.sN2tKI38UJAuAoXVXVe",
            "groups": "user",
            "first_name": "Clark",
            "last_name": "Kent",
            "classes_enrolled": {
                "63d460a6f1081c5158698c74": 3,
                "63d460d0f1081c5158698c76": 2
            },
            "rem_sks": 19
        }
    ]
}
```

### `DELETE /class/{classId}/delete-user/{userId}`
Admin endpoint to delete user from class, removing the class from users classes enrolled, and also adding the class sks into users remaining sks

Response
```
{
    "message": "successfully deleted user from class"
}
```

### `POST /class/{classId}/delete-user/{userId}`
Admin endpoint to add user from class, add the class to users classes enrolled, and also substracting the users sks by class remaining sks

Response
```
{
    "message": "successfully added user to class"
}
```

### `POST /class/create`
Admin endpoint to create new class

Example Request Body
```
{
    "title": "Advanced Machine Learning"
    "sks": 3
}
```

Response
```
{
    "message": "successfully created class"
}
```

### `POST /class/{classId}/edit`
Admin endpoint to edit class

Example Request Body
```
{
    "edit_map": {
        "title": "Python Machine Learning"
        "sks": 2
    }
}
```

Response
```
{
    "message": "successfully edited"
}
```

### `POST /class/delete/{classId}`
Admin endpoint to delete class

Response
```
{
    "message": "successfully deleted class"
}
```












# MyGram Final Bootcamp Golang Project
* [Photo](#photo)
* [User](#user)
* [Comment](#comment)
* [Social Media](#social-media)

## [Postman Collection](MyGram.postman_collection.json)
* Run postman documentation

[![Postman Documentation](https://www.postman.com/prabowoteguh/workspace/teguhworkspace/collection/11831418-b3f838f0-d693-422a-980f-6f1577d7be45?action=share&creator=11831418)


# Authentication

- HTTP Authentication, scheme: bearer

# Photo

## POST Create Photo

POST /photos

> Body Parameters

```yaml
title: Beach
photo_url: https://images.unsplash.com/photo-1707343843598-39755549ac9a?q=80&w=2832&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDF8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D
caption: Foto Pantai

```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|
|» title|body|string| yes |none|
|» photo_url|body|string| yes |none|
|» caption|body|string| yes |none|

> Response Examples

> Create First Photo

```json
{
  "ID": 1,
  "CreatedAt": "2024-02-16T09:20:56.460191+07:00",
  "UpdatedAt": "2024-02-16T09:20:56.460191+07:00",
  "DeletedAt": null,
  "title": "Foto Pertama",
  "photo_url": "https://images.unsplash.com/photo-1679065949530-7bb1fba3ccb3?q=80&w=2787&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
  "caption": "ini adalah foto pertama aink cuy",
  "UserID": 1,
  "User": null
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|201|[Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)|Create First Photo|Inline|

### Responses Data Schema

HTTP Status Code **201**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» ID|integer|true|none||none|
|» CreatedAt|string|true|none||none|
|» UpdatedAt|string|true|none||none|
|» DeletedAt|null|true|none||none|
|» title|string|true|none||none|
|» photo_url|string|true|none||none|
|» caption|string|true|none||none|
|» UserID|integer|true|none||none|
|» User|null|true|none||none|

## GET Get All Photo

GET /photos

> Response Examples

> Success Get All Photo

```json
[
  {
    "ID": 1,
    "CreatedAt": "2024-02-16T09:20:56.460191+07:00",
    "UpdatedAt": "2024-02-16T09:20:56.460191+07:00",
    "DeletedAt": null,
    "title": "Foto Pertama",
    "photo_url": "https://images.unsplash.com/photo-1679065949530-7bb1fba3ccb3?q=80&w=2787&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
    "caption": "ini adalah foto pertama aink cuy",
    "UserID": 1,
    "User": null
  },
  {
    "ID": 2,
    "CreatedAt": "2024-02-16T09:23:28.78337+07:00",
    "UpdatedAt": "2024-02-16T09:23:28.78337+07:00",
    "DeletedAt": null,
    "title": "Beach",
    "photo_url": "https://images.unsplash.com/photo-1707343843598-39755549ac9a?q=80&w=2832&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDF8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
    "caption": "Foto Pantai",
    "UserID": 1,
    "User": null
  }
]
```

> Failed not login

```json
{
  "error_message": "sign in to proceed",
  "error_status": "Unauthorized"
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|201|[Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)|Success Get All Photo|Inline|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|Failed not login|Inline|

### Responses Data Schema

HTTP Status Code **201**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» ID|integer|true|none||none|
|» CreatedAt|string|true|none||none|
|» UpdatedAt|string|true|none||none|
|» DeletedAt|null|true|none||none|
|» title|string|true|none||none|
|» photo_url|string|true|none||none|
|» caption|string|true|none||none|
|» UserID|integer|true|none||none|
|» User|null|true|none||none|

HTTP Status Code **401**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» error_message|string|true|none||none|
|» error_status|string|true|none||none|

## GET Get One Photo

GET /photos/1

> Response Examples

> Get One Photo

```json
{
  "ID": 1,
  "CreatedAt": "2024-02-16T09:20:56.460191+07:00",
  "UpdatedAt": "2024-02-16T09:20:56.460191+07:00",
  "DeletedAt": null,
  "title": "Foto Pertama",
  "photo_url": "https://images.unsplash.com/photo-1679065949530-7bb1fba3ccb3?q=80&w=2787&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
  "caption": "ini adalah foto pertama aink cuy",
  "UserID": 1,
  "User": null
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|201|[Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)|Get One Photo|Inline|

### Responses Data Schema

HTTP Status Code **201**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» ID|integer|true|none||none|
|» CreatedAt|string|true|none||none|
|» UpdatedAt|string|true|none||none|
|» DeletedAt|null|true|none||none|
|» title|string|true|none||none|
|» photo_url|string|true|none||none|
|» caption|string|true|none||none|
|» UserID|integer|true|none||none|
|» User|null|true|none||none|

## PUT Update Photo

PUT /photos/1

> Body Parameters

```yaml
title: Judulnya ganti
caption: Captionnya ganti
photo_url: https://images.unsplash.com/photo-1707696199186-d26b0f779f30?q=80&w=1828&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D

```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|
|» title|body|string| yes |none|
|» caption|body|string| yes |none|
|» photo_url|body|string| yes |none|

> Response Examples

> Update Photo

```json
{
  "ID": 1,
  "CreatedAt": "0001-01-01T00:00:00Z",
  "UpdatedAt": "2024-02-16T09:34:12.39539+07:00",
  "DeletedAt": null,
  "title": "Judulnya ganti",
  "photo_url": "https://images.unsplash.com/photo-1707696199186-d26b0f779f30?q=80&w=1828&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
  "caption": "Captionnya ganti",
  "UserID": 1,
  "User": null
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|201|[Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)|Update Photo|Inline|

### Responses Data Schema

HTTP Status Code **201**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» ID|integer|true|none||none|
|» CreatedAt|string|true|none||none|
|» UpdatedAt|string|true|none||none|
|» DeletedAt|null|true|none||none|
|» title|string|true|none||none|
|» photo_url|string|true|none||none|
|» caption|string|true|none||none|
|» UserID|integer|true|none||none|
|» User|null|true|none||none|

## DELETE Delete Photo

DELETE /photos/3

> Response Examples

> Delete Photo

```json
{
  "message": "photo with id 3 has been successfully deleted"
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|201|[Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)|Delete Photo|Inline|

### Responses Data Schema

HTTP Status Code **201**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» message|string|true|none||none|

# User

## POST Register MyGram

POST /users/register

> Body Parameters

```yaml
username: farras
email: farras@mail.com
password: "123456"
age: "26"

```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|
|» username|body|string| yes |none|
|» email|body|string| yes |none|
|» password|body|string| yes |none|
|» age|body|string| yes |none|

> Response Examples

> Register MyGram

```json
{
  "age": 26,
  "email": "farras@mail.com",
  "id": 2,
  "username": "farras"
}
```

> Failed Register MyGram

```json
{
  "error": "Bad Request",
  "message": "Age must be greater than 8"
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|201|[Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)|Register MyGram|Inline|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Failed Register MyGram|Inline|

### Responses Data Schema

HTTP Status Code **201**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» age|integer|true|none||none|
|» email|string|true|none||none|
|» id|integer|true|none||none|
|» username|string|true|none||none|

HTTP Status Code **400**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» error|string|true|none||none|
|» message|string|true|none||none|

## POST Login MyGram

POST /users/login

> Body Parameters

```yaml
email: alfa@mail.com
password: "123456"

```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|
|» email|body|string| yes |none|
|» password|body|string| yes |none|

> Response Examples

> Login MyGram

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFsZmFAbWFpbC5jb20iLCJpZCI6MX0.CruDbZWyA_ygWV_1TwVEwwBR8n26V7ePvNcI5vWtSVg"
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Login MyGram|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» token|string|true|none||none|

# Comment

## POST Post Comment

POST /comments/1

> Body Parameters

```yaml
message: coba comment foto pertama

```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|
|» message|body|string| yes |none|

> Response Examples

> Post Comment

```json
{
  "ID": 18,
  "CreatedAt": "2024-02-16T10:44:31.243673+07:00",
  "UpdatedAt": "2024-02-16T10:44:31.243673+07:00",
  "DeletedAt": null,
  "message": "coba comment foto pertama",
  "UserID": 1,
  "User": null,
  "PhotoID": 1,
  "Photo": null
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|201|[Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)|Post Comment|Inline|

### Responses Data Schema

HTTP Status Code **201**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» ID|integer|true|none||none|
|» CreatedAt|string|true|none||none|
|» UpdatedAt|string|true|none||none|
|» DeletedAt|null|true|none||none|
|» message|string|true|none||none|
|» UserID|integer|true|none||none|
|» User|null|true|none||none|
|» PhotoID|integer|true|none||none|
|» Photo|null|true|none||none|

## GET Get All Comment

GET /comments

> Body Parameters

```yaml
{}

```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|

> Response Examples

> Get All Comment

```json
{
  "comments": [
    {
      "ID": 14,
      "CreatedAt": "2024-02-16T10:36:28.712273+07:00",
      "UpdatedAt": "2024-02-16T10:36:28.712273+07:00",
      "DeletedAt": null,
      "message": "coba comment foto pertama",
      "UserID": 1,
      "User": {
        "ID": 1,
        "CreatedAt": "2024-02-13T16:18:36.717295+07:00",
        "UpdatedAt": "2024-02-13T16:18:36.717295+07:00",
        "DeletedAt": null,
        "username": "alfa",
        "email": "alfa@mail.com",
        "password": "$2a$08$xunMy1yQGACQSwKTzwWusuCM1v4c13IlQKXkxEpfoez93JqYC4ndm",
        "age": 9,
        "photos": null,
        "comments": null,
        "social_media": null
      },
      "PhotoID": 1,
      "Photo": {
        "ID": 1,
        "CreatedAt": "2024-02-16T09:20:56.460191+07:00",
        "UpdatedAt": "2024-02-16T09:34:12.39539+07:00",
        "DeletedAt": null,
        "title": "Judulnya ganti",
        "photo_url": "https://images.unsplash.com/photo-1707696199186-d26b0f779f30?q=80&w=1828&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
        "caption": "Captionnya ganti",
        "UserID": 1,
        "User": null
      }
    },
    {
      "ID": 16,
      "CreatedAt": "2024-02-16T10:42:44.593117+07:00",
      "UpdatedAt": "2024-02-16T11:04:04.363957+07:00",
      "DeletedAt": null,
      "message": "Update comment dulu",
      "UserID": 1,
      "User": {
        "ID": 1,
        "CreatedAt": "2024-02-13T16:18:36.717295+07:00",
        "UpdatedAt": "2024-02-13T16:18:36.717295+07:00",
        "DeletedAt": null,
        "username": "alfa",
        "email": "alfa@mail.com",
        "password": "$2a$08$xunMy1yQGACQSwKTzwWusuCM1v4c13IlQKXkxEpfoez93JqYC4ndm",
        "age": 9,
        "photos": null,
        "comments": null,
        "social_media": null
      },
      "PhotoID": 1,
      "Photo": {
        "ID": 1,
        "CreatedAt": "2024-02-16T09:20:56.460191+07:00",
        "UpdatedAt": "2024-02-16T09:34:12.39539+07:00",
        "DeletedAt": null,
        "title": "Judulnya ganti",
        "photo_url": "https://images.unsplash.com/photo-1707696199186-d26b0f779f30?q=80&w=1828&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
        "caption": "Captionnya ganti",
        "UserID": 1,
        "User": null
      }
    },
    {
      "ID": 17,
      "CreatedAt": "2024-02-16T10:42:47.696695+07:00",
      "UpdatedAt": "2024-02-16T10:42:47.696695+07:00",
      "DeletedAt": null,
      "message": "coba comment foto pertama",
      "UserID": 1,
      "User": {
        "ID": 1,
        "CreatedAt": "2024-02-13T16:18:36.717295+07:00",
        "UpdatedAt": "2024-02-13T16:18:36.717295+07:00",
        "DeletedAt": null,
        "username": "alfa",
        "email": "alfa@mail.com",
        "password": "$2a$08$xunMy1yQGACQSwKTzwWusuCM1v4c13IlQKXkxEpfoez93JqYC4ndm",
        "age": 9,
        "photos": null,
        "comments": null,
        "social_media": null
      },
      "PhotoID": 1,
      "Photo": {
        "ID": 1,
        "CreatedAt": "2024-02-16T09:20:56.460191+07:00",
        "UpdatedAt": "2024-02-16T09:34:12.39539+07:00",
        "DeletedAt": null,
        "title": "Judulnya ganti",
        "photo_url": "https://images.unsplash.com/photo-1707696199186-d26b0f779f30?q=80&w=1828&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
        "caption": "Captionnya ganti",
        "UserID": 1,
        "User": null
      }
    },
    {
      "ID": 18,
      "CreatedAt": "2024-02-16T10:44:31.243673+07:00",
      "UpdatedAt": "2024-02-16T10:44:31.243673+07:00",
      "DeletedAt": null,
      "message": "coba comment foto pertama",
      "UserID": 1,
      "User": {
        "ID": 1,
        "CreatedAt": "2024-02-13T16:18:36.717295+07:00",
        "UpdatedAt": "2024-02-13T16:18:36.717295+07:00",
        "DeletedAt": null,
        "username": "alfa",
        "email": "alfa@mail.com",
        "password": "$2a$08$xunMy1yQGACQSwKTzwWusuCM1v4c13IlQKXkxEpfoez93JqYC4ndm",
        "age": 9,
        "photos": null,
        "comments": null,
        "social_media": null
      },
      "PhotoID": 1,
      "Photo": {
        "ID": 1,
        "CreatedAt": "2024-02-16T09:20:56.460191+07:00",
        "UpdatedAt": "2024-02-16T09:34:12.39539+07:00",
        "DeletedAt": null,
        "title": "Judulnya ganti",
        "photo_url": "https://images.unsplash.com/photo-1707696199186-d26b0f779f30?q=80&w=1828&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
        "caption": "Captionnya ganti",
        "UserID": 1,
        "User": null
      }
    },
    {
      "ID": 19,
      "CreatedAt": "2024-02-16T10:51:25.601451+07:00",
      "UpdatedAt": "2024-02-16T10:51:25.601451+07:00",
      "DeletedAt": null,
      "message": "Komentar kedua nih",
      "UserID": 1,
      "User": {
        "ID": 1,
        "CreatedAt": "2024-02-13T16:18:36.717295+07:00",
        "UpdatedAt": "2024-02-13T16:18:36.717295+07:00",
        "DeletedAt": null,
        "username": "alfa",
        "email": "alfa@mail.com",
        "password": "$2a$08$xunMy1yQGACQSwKTzwWusuCM1v4c13IlQKXkxEpfoez93JqYC4ndm",
        "age": 9,
        "photos": null,
        "comments": null,
        "social_media": null
      },
      "PhotoID": 1,
      "Photo": {
        "ID": 1,
        "CreatedAt": "2024-02-16T09:20:56.460191+07:00",
        "UpdatedAt": "2024-02-16T09:34:12.39539+07:00",
        "DeletedAt": null,
        "title": "Judulnya ganti",
        "photo_url": "https://images.unsplash.com/photo-1707696199186-d26b0f779f30?q=80&w=1828&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
        "caption": "Captionnya ganti",
        "UserID": 1,
        "User": null
      }
    }
  ]
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Get All Comment|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» comments|[object]|true|none||none|
|»» ID|integer|true|none||none|
|»» CreatedAt|string|true|none||none|
|»» UpdatedAt|string|true|none||none|
|»» DeletedAt|null|true|none||none|
|»» message|string|true|none||none|
|»» UserID|integer|true|none||none|
|»» User|object|true|none||none|
|»»» ID|integer|true|none||none|
|»»» CreatedAt|string|true|none||none|
|»»» UpdatedAt|string|true|none||none|
|»»» DeletedAt|null|true|none||none|
|»»» username|string|true|none||none|
|»»» email|string|true|none||none|
|»»» password|string|true|none||none|
|»»» age|integer|true|none||none|
|»»» photos|null|true|none||none|
|»»» comments|null|true|none||none|
|»»» social_media|null|true|none||none|
|»» PhotoID|integer|true|none||none|
|»» Photo|object|true|none||none|
|»»» ID|integer|true|none||none|
|»»» CreatedAt|string|true|none||none|
|»»» UpdatedAt|string|true|none||none|
|»»» DeletedAt|null|true|none||none|
|»»» title|string|true|none||none|
|»»» photo_url|string|true|none||none|
|»»» caption|string|true|none||none|
|»»» UserID|integer|true|none||none|
|»»» User|null|true|none||none|

## DELETE Delete Comment

DELETE /comments/15

> Response Examples

> Delete Comment

```json
{
  "message": "comment with id 15 has been successfully deleted"
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|201|[Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)|Delete Comment|Inline|

### Responses Data Schema

HTTP Status Code **201**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» message|string|true|none||none|

## PUT Update Comment

PUT /comments/16

> Body Parameters

```yaml
message: Update comment dulu

```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|
|» message|body|string| yes |none|

> Response Examples

> Update Comment

```json
{
  "ID": 16,
  "CreatedAt": "0001-01-01T00:00:00Z",
  "UpdatedAt": "2024-02-16T11:04:04.363957+07:00",
  "DeletedAt": null,
  "message": "Update comment dulu",
  "UserID": 1,
  "User": null,
  "PhotoID": 0,
  "Photo": null
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|201|[Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)|Update Comment|Inline|

### Responses Data Schema

HTTP Status Code **201**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» ID|integer|true|none||none|
|» CreatedAt|string|true|none||none|
|» UpdatedAt|string|true|none||none|
|» DeletedAt|null|true|none||none|
|» message|string|true|none||none|
|» UserID|integer|true|none||none|
|» User|null|true|none||none|
|» PhotoID|integer|true|none||none|
|» Photo|null|true|none||none|

## GET Get One Comment

GET /comments/14

> Body Parameters

```yaml
{}

```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|

> Response Examples

> Get One Comment

```json
{
  "comment": {
    "ID": 14,
    "CreatedAt": "2024-02-16T10:36:28.712273+07:00",
    "UpdatedAt": "2024-02-16T10:36:28.712273+07:00",
    "DeletedAt": null,
    "message": "coba comment foto pertama",
    "UserID": 1,
    "User": {
      "ID": 1,
      "CreatedAt": "2024-02-13T16:18:36.717295+07:00",
      "UpdatedAt": "2024-02-13T16:18:36.717295+07:00",
      "DeletedAt": null,
      "username": "alfa",
      "email": "alfa@mail.com",
      "password": "$2a$08$xunMy1yQGACQSwKTzwWusuCM1v4c13IlQKXkxEpfoez93JqYC4ndm",
      "age": 9,
      "photos": null,
      "comments": null,
      "social_media": null
    },
    "PhotoID": 1,
    "Photo": {
      "ID": 1,
      "CreatedAt": "2024-02-16T09:20:56.460191+07:00",
      "UpdatedAt": "2024-02-16T09:34:12.39539+07:00",
      "DeletedAt": null,
      "title": "Judulnya ganti",
      "photo_url": "https://images.unsplash.com/photo-1707696199186-d26b0f779f30?q=80&w=1828&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
      "caption": "Captionnya ganti",
      "UserID": 1,
      "User": null
    }
  }
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Get One Comment|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» comment|object|true|none||none|
|»» ID|integer|true|none||none|
|»» CreatedAt|string|true|none||none|
|»» UpdatedAt|string|true|none||none|
|»» DeletedAt|null|true|none||none|
|»» message|string|true|none||none|
|»» UserID|integer|true|none||none|
|»» User|object|true|none||none|
|»»» ID|integer|true|none||none|
|»»» CreatedAt|string|true|none||none|
|»»» UpdatedAt|string|true|none||none|
|»»» DeletedAt|null|true|none||none|
|»»» username|string|true|none||none|
|»»» email|string|true|none||none|
|»»» password|string|true|none||none|
|»»» age|integer|true|none||none|
|»»» photos|null|true|none||none|
|»»» comments|null|true|none||none|
|»»» social_media|null|true|none||none|
|»» PhotoID|integer|true|none||none|
|»» Photo|object|true|none||none|
|»»» ID|integer|true|none||none|
|»»» CreatedAt|string|true|none||none|
|»»» UpdatedAt|string|true|none||none|
|»»» DeletedAt|null|true|none||none|
|»»» title|string|true|none||none|
|»»» photo_url|string|true|none||none|
|»»» caption|string|true|none||none|
|»»» UserID|integer|true|none||none|
|»»» User|null|true|none||none|

# Social Media

## POST Create Social Media

POST /socialmedias

> Body Parameters

```yaml
name: Alfa
social_media_url: alfa@instagram.com

```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|
|» name|body|string| yes |none|
|» social_media_url|body|string| yes |none|

> Response Examples

> Create Social Media

```json
{
  "ID": 1,
  "CreatedAt": "2024-02-16T11:32:53.000673+07:00",
  "UpdatedAt": "2024-02-16T11:32:53.000673+07:00",
  "DeletedAt": null,
  "name": "Alfa",
  "social_media_url": "alfa@instagram.com",
  "UserID": 1,
  "User": null
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|201|[Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)|Create Social Media|Inline|

### Responses Data Schema

HTTP Status Code **201**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» ID|integer|true|none||none|
|» CreatedAt|string|true|none||none|
|» UpdatedAt|string|true|none||none|
|» DeletedAt|null|true|none||none|
|» name|string|true|none||none|
|» social_media_url|string|true|none||none|
|» UserID|integer|true|none||none|
|» User|null|true|none||none|

## GET Get All Social Media

GET /socialmedias

> Response Examples

> Get All Social Media

```json
[
  {
    "ID": 1,
    "CreatedAt": "2024-02-16T11:32:53.000673+07:00",
    "UpdatedAt": "2024-02-16T11:32:53.000673+07:00",
    "DeletedAt": null,
    "name": "Alfa",
    "social_media_url": "alfa@instagram.com",
    "UserID": 1,
    "User": null
  },
  {
    "ID": 2,
    "CreatedAt": "2024-02-16T12:03:01.119968+07:00",
    "UpdatedAt": "2024-02-16T12:03:01.119968+07:00",
    "DeletedAt": null,
    "name": "anjoy",
    "social_media_url": "anjoy@google.com",
    "UserID": 1,
    "User": null
  },
  {
    "ID": 3,
    "CreatedAt": "2024-02-16T12:03:13.324375+07:00",
    "UpdatedAt": "2024-02-16T12:05:06.03972+07:00",
    "DeletedAt": null,
    "name": "bimo",
    "social_media_url": "bimo@gmail.com",
    "UserID": 1,
    "User": null
  }
]
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|201|[Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)|Get All Social Media|Inline|

### Responses Data Schema

HTTP Status Code **201**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» ID|integer|true|none||none|
|» CreatedAt|string|true|none||none|
|» UpdatedAt|string|true|none||none|
|» DeletedAt|null|true|none||none|
|» name|string|true|none||none|
|» social_media_url|string|true|none||none|
|» UserID|integer|true|none||none|
|» User|null|true|none||none|

## GET Get One Social Media

GET /socialmedias/1

> Response Examples

> Get One Social Media

```json
{
  "ID": 1,
  "CreatedAt": "2024-02-16T11:32:53.000673+07:00",
  "UpdatedAt": "2024-02-16T11:32:53.000673+07:00",
  "DeletedAt": null,
  "name": "Alfa",
  "social_media_url": "alfa@instagram.com",
  "UserID": 1,
  "User": null
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|201|[Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)|Get One Social Media|Inline|

### Responses Data Schema

HTTP Status Code **201**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» ID|integer|true|none||none|
|» CreatedAt|string|true|none||none|
|» UpdatedAt|string|true|none||none|
|» DeletedAt|null|true|none||none|
|» name|string|true|none||none|
|» social_media_url|string|true|none||none|
|» UserID|integer|true|none||none|
|» User|null|true|none||none|

## PUT Update Social Media

PUT /socialmedias/3

> Body Parameters

```yaml
name: bimo
social_media_url: bimo@gmail.com

```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|
|» name|body|string| yes |none|
|» social_media_url|body|string| yes |none|

> Response Examples

> Update Social Media

```json
{
  "ID": 3,
  "CreatedAt": "0001-01-01T00:00:00Z",
  "UpdatedAt": "2024-02-16T12:05:06.03972+07:00",
  "DeletedAt": null,
  "name": "bimo",
  "social_media_url": "bimo@gmail.com",
  "UserID": 1,
  "User": null
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|201|[Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)|Update Social Media|Inline|

### Responses Data Schema

HTTP Status Code **201**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» ID|integer|true|none||none|
|» CreatedAt|string|true|none||none|
|» UpdatedAt|string|true|none||none|
|» DeletedAt|null|true|none||none|
|» name|string|true|none||none|
|» social_media_url|string|true|none||none|
|» UserID|integer|true|none||none|
|» User|null|true|none||none|

## DELETE Delete Social Media

DELETE /socialmedias/2

> Response Examples

> Delete Social Media

```json
{
  "message": "socialmedia with id 2 has been successfully deleted"
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|201|[Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)|Delete Social Media|Inline|

### Responses Data Schema

HTTP Status Code **201**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» message|string|true|none||none|
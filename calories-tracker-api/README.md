# Table of contents

- [Overview](#overview)
- [Technologies](#technologies)
- [API Reference](#api-reference)
- [Run Locally](#run-locally)

# Overview

I create api where you can delete, update, create and also get data food calories and used [gin](https://github.com/gin-gonic/gin) for building API endpoints and microservices.

# Technologies

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![MongoDB](https://img.shields.io/badge/MongoDB-%234ea94b.svg?style=for-the-badge&logo=mongodb&logoColor=white)
![Render](https://img.shields.io/badge/Render-%46E3B7.svg?style=for-the-badge&logo=render&logoColor=white)

# API Reference

`BASE_URL = https://tracking-calories.onrender.com`

### Get one by id

```http
  GET /entry/:id
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `number` | **Required**. Item id             |


#### Get search page videos

```http
  POST /entry/create
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
|  `none`  | `none`    | **Required**. None         |

#### Get all data

```http
  GET /entries
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `none` | `none` | **Required**. None |

#### Get ingredient

```http
  GET /ingredient/:ingredient
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `ingredient`      | `string` | **Required**. Item ingredient            |

#### update item

```http
  PUT /entry/update/:id
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `number` | **Required**. Item id             |


#### update ingredient

```http
  PUT /ingredient/update/:id
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `number` | **Required**. Item id             |

#### Delete item

```http
  DELETE /entry/delete/:id
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `number` | **Required**. Item id             |

# Run Locally

Clone the project

```bash
  git clone git@github.com:Rokiis1/calories-tracker-api.git
```

Start the server

```bash
  go run main.go
```




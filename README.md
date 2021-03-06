# jakeri-backend

> Backend has been built using Golang and the Gin framework. It is using Mongodb as the database

## Dependencies

- [gin](https://github.com/gin-gonic/gin)
- [MongoDB](https://github.com/mongodb/mongo)

## Structure

```
.
└── authorizations
    ├── groups.go
    ├── reviews.go
    ├── cards.go
    ├── profiles.go
    ├── users.go
    └── utils.go

└── controllers
    ├── groups.go
    ├── confirmation.go
    ├── reviews.go
    ├── cards.go
    ├── profiles.go
    ├── recovery.go
    ├── session.go
    ├── status.go
    └── users.go

└── middleware
    ├── authorization.go
    ├── cognito.go
    └── cors.go

└── models
    ├── audit.go
    ├── groups.go
    ├── reviews.go
    ├── cards.go
    ├── profiles.go
    └── users.go

└── routers
    ├── groups.go
    ├── confirmation.go
    ├── reviews.go
    ├── cards.go
    ├── profiles.go
    ├── recovery.go
    ├── routers.go
    ├── session.go
    ├── status.go
    └── users.go

└── utils
    ├── api.go
    ├── db.go
    └── env.go

└── validations
    ├── groups.go
    ├── confirmation.go
    ├── reviews.go
    ├── cards.go
    ├── profiles.go
    ├── recovery.go
    ├── session.go
    └── users.go

├── .gitignore
├── main.go

```

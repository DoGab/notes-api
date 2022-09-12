# notes-api
Notes API written in Go using the Fiber web framework

- [notes-api](#notes-api)
  - [Project setup](#project-setup)
  - [Todos](#todos)

## Project setup

1. Create the project
```bash
mkdir notes-api
cd notes-api
go mod init github.com/dogab/notes-api
# Open Editor in in the directory where the go.mod file is
go get github.com/gofiber/fiber/v2
touch main.go
go run main.go
```

2. Used for the makefile to hot reload the server
```
go install github.com/cespare/reflex@latest
touch Makefile
make watch
```

3. Setup env file for storing the config variables
```
touch .env
go get github.com/joho/godotenv
mkdir config
touch config/config.go
```

4. Setup database
```
touch docker-compose.yaml
mkdir database
touch database/database.go

go get gorm.io/gorm
go get gorm.io/driver/postgres
```

5. Create Models
```
mkdir model
touch model/model.go

go get github.com/google/uuid
```

6. Create router
```
mkdir router
touch router/router.go
```

7. Add routes
```
mkdir routes
mkdir routes/note
touch routes/note/note.go
```

8. Add handlers
```
mkdir handlers
mkdir handlers/note
touch handlers/note/note.go
```

9. Swagger
Add comments to the handler functions. https://github.com/arsmn/fiber-swagger
```
# Add comments to the files
go get -u github.com/swaggo/swag/cmd/swag
swag init
go get -u github.com/arsmn/fiber-swagger/v2
import "github.com/arsmn/fiber-swagger/v2"
import "github.com/dogab/docs"
```

10. Encrypt passwords
Add encrypted passwords.
```
go get golang.org/x/crypto/bcrypt
```

11. Use JWT Auth
Use this in the package for the jwt function.
```
go get github.com/golang-jwt/jwt/v4
```

Use this in the package where the Middleware should be used for the routes.
```
go get  github.com/gofiber/jwt/v3
```

## Todos

- Update readme with usage and routes
- check how to use secret key as global var, https://stackoverflow.com/questions/35038864/how-to-access-global-variables
- check how to embed json with token into other message format: https://github.com/gofiber/fiber/issues/164
- validator https://www.youtube.com/watch?v=5q_wsashJZA, https://dev.to/koddr/build-a-restful-api-on-go-fiber-postgresql-jwt-and-swagger-docs-in-isolated-docker-containers-475j#folder-with-business-logic-only
- add swagger docu for endpoints
- add utils functions for jwt https://dev.to/koddr/build-a-restful-api-on-go-fiber-postgresql-jwt-and-swagger-docs-in-isolated-docker-containers-475j#useful-utilities)
- fix csrf
- add more endpoints
- create migration with initial user
- validate POST input for user update
- test update user endpoints
- create user delete endpoint
  
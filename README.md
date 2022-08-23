## Features

- Gorm
- Echo



# Directory structure
```
.
├── cmd
│   └── go-rest-api-cleancode
│       └── main.go
├── configs
│   └── db-config.go    // Setup Database Connection
├── controller
│   ├── auth.go
│   ├── product.go
│   └── user.go
├── dto                 // data transfer objects
│   ├── login.go   
│   ├── product.go 
│   ├── register.go 
│   └── user.go     
├── helper
│   └── response.go     // BuildResponse, BuildErrorResponse and EmptyObj
├── middleware
│   └── jwt-auth.go
├── models
|   ├── Product.go       
|   └── User.go   
├── repository
|   ├── product.go
|   └── user.go
├── service
|   ├── auth.go
|   ├── jwt.go
|   ├── product.go
|   └── user.go
|
...
```
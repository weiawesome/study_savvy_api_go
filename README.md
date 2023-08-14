# Server for study-savvy in go-lang version
## How to start
```
docker-compose up -d
```
## Project Structure
Layered Architecture(分層式架構)
```markdown
├── api
│   ├── handler
│   ├── middleware
│   ├── request
│   ├── response
│   ├── model
│   ├── routes
│   ├── utils
│
├── internal
│   ├── repository
│   ├── service
│
├── docs
│
├── go.mod
├── go.sum
│
├── test
│
├── main.go
```
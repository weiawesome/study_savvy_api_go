# Server for study-savvy in go-lang version 

## How to start
```
docker-compose up -d
```

## Project Structure - Layered Architecture
```markdown
├── api
│   ├── handler
│   ├── middleware
│   ├── request
│   ├── response
│   ├── routes
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

> ### [Presentation Layer](./api)
> [The api directory](./api) encompasses the presentation layer, responsible for handling user interactions and requests.
> > routes
> > > This directory contains the configuration of API routes. All API endpoints are defined here. 
> 
> > handler
> > > The handler directory connects with the service layer and generates responses for incoming requests.
> 
> > middleware
> > > The middleware directory contains pre-processing logic before the request reaches the handler. It includes tasks such as Jwt validation, content-type filtering, and request content validation.
> 
> > request
> > > The request directory defines the structure and types of incoming requests. It helps ensure proper handling and validation of user inputs.
> 
> > response
> > > The response directory defines the structure and types of outgoing responses. It ensures consistency in the data sent back to clients.

> ### [Business Logic Layer](./internal/service)
> [The internal/service directory](./internal/service) serves as the Business Logic layer of the application.</br>
> Each subdirectory and service file within this directory corresponds to a specific handler in the presentation layer.</br> 
> These components are responsible for carrying out intricate internal logical operations and orchestrating data manipulation.</br>
> Once the necessary business logic has been executed, </br>
> the service returns a well-formed response to the respective handler in the presentation layer.</br>
> This seamless interaction ensures that the core business rules are effectively implemented and maintained.

> ### [Data Access Layer](./internal/repository)
> [The internal/repository directory](./internal/repository)` corresponds to the Data Access Layer, responsible for interacting with various data sources.
> > model
> > > The model subdirectory contains the definitions of all data models used in the SQL Database. These models define the structure and relationships of the data.
> 
> > sql
> > > The sql subdirectory houses all SQL Database manipulation operations. This includes CRUD operations for all database tables. The services within the application utilize the repository to retrieve and manipulate information stored in the SQL Database.
> 
> > redis
> > > The redis subdirectory handles all interactions with the Redis Database. Here, tasks like setting blacklists, validating mail verifications, and adding Celery tasks are managed. Redis plays a vital role in caching and handling specific tasks efficiently. 
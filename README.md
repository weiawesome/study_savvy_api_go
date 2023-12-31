# Server for study-savvy Go version

## Description the service

> This is the back-end service for [Study-Savvy](#study-savvy).
> 
> The back-end server can get client's request and reply it.
> 
> Furthermore, it can cooperate with Celery in Python which gain tasks-queue from Redis set by the server.
> </br> The corresponding celery repository can see in [Study-Savvy-AsyncWork-Celery](https://github.com/weiawesome/study_savvy_asyncwork_celery/tree/master)
> 
> In the project will set My-SQL and Redis and Go-server, and it should combine with other project to have whole service.
> 
> Furthermore, the logs will send into InfluxDB and display on the Grafana.
> > Logs stored in InfluxDB and display by Grafana
> 
> ![img.png](resource/logs.png)
 

## Study-Savvy
> Supply main two service
> * Audio content summarize and assistant especially in education zone.
> * Article (from graph or text) content improver by giving some advice especially in high-school student writing.
> 
> Two application
> * App made by Flutter
> * Website made by Next.js
> 
> > ### [Website](https://study-savvy.com)
> > ![img.png](resource/website.png)
> >
> > #### URL : https://study-savvy.com
>
> > ### APP
> > ![img.png](resource/app.png)
> 
> > ### [API document](https://study-savvy.com/api/docs/)
> > ![img.png](resource/apidocs.png)
> >
> > (Image just part of document) 
> > 
> > Details in document can see the [docs/openapi.yaml](./docs/openapi.yaml)
> >
> > or visit the API-Document 
> > 
> > URL : https://study-savvy.com/api/docs/
### System Architecture in Study-Savvy
![](resource/system_architecture.png)
* This project is the "Back-end Subsystem" in Study-Savvy

## How to install
```bash
git clone https://github.com/weiawesome/study_savvy_api_go.git
```
## How to start
```bash
docker-compose up -d
```
## How to set the environment
```bash
# Move to test directory
cd ./test

# Set all environment
docker-compose up -d
```


## Project Structure - Layered Architecture
### [There Layered Architecture](https://en.wikipedia.org/wiki/Multitier_architecture#Three-tier_architecture)
- #### Presentation Layer
- #### Business Logic Layer
- #### Data Access Layer

```markdown
Project Architecture

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
├── main.go
```

### Presentation Layer
> [The api directory](./api) encompasses the presentation layer, responsible for handling user interactions and requests.
> > #### routes
> * This directory contains the configuration of API routes. All API endpoints are defined here.
>
> > #### handler
> * The handler directory connects with the service layer and generates responses for incoming requests.
>
> > #### middleware
> * The middleware directory contains pre-processing logic before the request reaches the handler. It includes tasks such as Jwt validation, content-type filtering, and request content validation.
>
> > #### request
> * The request directory defines the structure and types of incoming requests. It helps ensure proper handling and validation of user inputs.
>
> > #### response
> * The response directory defines the structure and types of outgoing responses. It ensures consistency in the data sent back to clients.
> ### Business Logic Layer
> [The internal/service directory](./internal/service) serves as the Business Logic layer of the application.</br>
> Each subdirectory and service file within this directory corresponds to a specific handler in the presentation layer.</br> 
> These components are responsible for carrying out intricate internal logical operations and orchestrating data manipulation.</br>
> Once the necessary business logic has been executed, </br>
> the service returns a well-formed response to the respective handler in the presentation layer.</br>
> This seamless interaction ensures that the core business rules are effectively implemented and maintained.

### Data Access Layer
> [The internal/repository directory](./internal/repository) corresponds to the Data Access Layer, responsible for interacting with various data sources.
> > #### model
> * The model subdirectory contains the definitions of all data models used in the SQL Database. These models define the structure and relationships of the data.
> 
> > #### sql
> * The sql subdirectory houses all SQL Database manipulation operations. This includes CRUD operations for all database tables. The services within the application utilize the repository to retrieve and manipulate information stored in the SQL Database.
> 
> > #### redis
> * The redis subdirectory handles all interactions with the Redis Database. Here, tasks like setting blacklists, validating mail verifications, and adding Celery tasks are managed. Redis plays a vital role in caching and handling specific tasks efficiently. 
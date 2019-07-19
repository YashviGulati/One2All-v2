# One2All-v2

AIM: The main objective is to create a REST API application with GoLang which will do CRUD operations around any cloud resource.

DETAILS: One2All is an application made using Amazon SNS resource, that coordinates and manages the delivery of messages to subscribing endpoints.
- Cloud Resource: AWS
- Language used: Golang
- Added Basic Authentication
- Added Unit Testing

FILE STRUCTURE:
- main.go : Index file
- middlewares: It contains basic authentication go file
- handlecontrol: It contains all the functions of the handlers.
- main_test.go: It is the unit testing file.
- One2All_swagger.json: swagger file in json format 
- One2All_swagger.yaml: swagger file in yaml format

DEPENDENCIES:
- gorilla/mux
- aws-sdk-go
- stretchr/testify/assert
- gorilla/handlers

RUN:

- To run the code, first add AKID and SECRET_KEY provided by the amazon web services and then run the following command.

`go run main.go`

- To run the unit testing file, following is the command

`go test`

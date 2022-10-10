# Participant Registry Microservice

[![CircleCI](https://circleci.com/gh/yemiwebby/participants-registry.svg?style=svg)](https://circleci.com/gh/yemiwebby/participants-registry)

## Prerequisite

This project requires [Go](https://go.dev/) runtime installed on your computer. Follow the instructions [here to download](https://go.dev/doc/install) and install it for your computer's operatiing system.

## Clone the project

```bash
git clone https://github.com/yemiwebby/participants-registry.git
```

Navigate into the project

```bash
cd participants-registry
```

## Run the application

```bash
make server
```

Alternatively, you can also run the project using:

```bash
go run main.go
```

The preceding command will automatically start the project on [http://localhost:8080](http://localhost:8080)

![Default Homepage](./screenshots/homepage.png)

**Note**: If you encouter an error about `gotest`, run the command below to install fix it:

```bash
go install github.com/rakyll/gotest
```

## Run test

Issue the following command from the root of the application to run unit test:

```bash
make test
```

Alternatively, you can also use the following command:

```bash
gotest -v -cover ./api
```

Or

```bash
go test -v -cover ./api
```

You will see an output similar to the following:

```bash
GIN_MODE=release gotest -v -cover ./api
=== RUN   TestHomepageHandler
[GIN] 2022/10/10 - 09:00:02 | 200 |      87.333µs |                 | GET      "/"
--- PASS: TestHomepageHandler (0.00s)
=== RUN   TestCreateParticipant
[GIN] 2022/10/10 - 09:00:02 | 201 |     262.625µs |                 | POST     "/participant"
    participants_test.go:73: Expected to get status 201 is the same as 201
--- PASS: TestCreateParticipant (0.00s)
=== RUN   TestGetParticipants
[GIN] 2022/10/10 - 09:00:02 | 200 |      29.667µs |                 | GET      "/participants"
--- PASS: TestGetParticipants (0.00s)
=== RUN   TestGetParticipant
[GIN] 2022/10/10 - 09:00:02 | 200 |      13.541µs |                 | GET      "/participant/ED34"
[GIN] 2022/10/10 - 09:00:02 | 404 |         125ns |                 | PUT      "/participant/KH32"
--- PASS: TestGetParticipant (0.00s)
=== RUN   TestUpdateParticipant
[GIN] 2022/10/10 - 09:00:02 | 200 |      39.625µs |                 | PUT      "/participant/ED34"
[GIN] 2022/10/10 - 09:00:02 | 404 |      10.292µs |                 | PUT      "/participant/KH32"
--- PASS: TestUpdateParticipant (0.00s)
=== RUN   TestDeleteParticipant
[GIN] 2022/10/10 - 09:00:02 | 204 |       1.875µs |                 | DELETE   "/participant/ED34"
[GIN] 2022/10/10 - 09:00:02 | 404 |         125ns |                 | PUT      "/participant/KH32"
--- PASS: TestDeleteParticipant (0.00s)
PASS
coverage: 68.4% of statements
ok      participant-project/api (cached)        coverage: 68.4% of statements
```

## Test the application

Use API testing tools such as [Postman](https://www.postman.com/) or [Insomnia](https://insomnia.rest/) to test the application locally.

### Get the list of participants

![Get Participants](./screenshots/get-participants.png)

### Create a participant

![Create Participant](./screenshots/create-participant.png)

### Get participant

![Get Participant](./screenshots/get-participant.png)

### Update participant

![Get Participants](./screenshots/update-participant.png)

### Delete participant

![Get Participants](./screenshots/delete-participant.png)

# Cushon API

## Description

Cushon offers pensions to customers that are employees of companies, but now they also would like to offer pensions directly to other customers (retail). I will refer to these customers as `employed customers` and `retail customers` respectively.

When customers want to invest, they can choose between different `funds`. Once the fund is selected the user can choose the amount they want to invest, I will be referring to this (depositing money into a fund) as an `investment`.

The customer must be able to retrieve all the `investments` they have made.

## Proposed Solution

I've assumed the customer interacts with Cushon through a frontend (web, mobile app...), that will be calling a backend. I will create a service that this backend can query to manage customers, investments and funds.

I decided to implement the functionality to handle all these entities in just one service with different endpoints due to time limitations. The goal was to be able to have a working vertical slice of a complete use case.

Another option would've been to create one microservice to handle each entity, ie: one for customers, another one for funds...

The statement of the task was very open ended, so I had to make a choice between focusing more on the code (what I've done, having a solution that runs) or focusing on system design (just providing a document with the design of my solution).

I still had to cut corners, like not using a DB or using unsafe practices (certificates, api keys...), but I've tried to explain what I would've done in a real scenario across the document.


## Implementation Details

I've structured the code the following way:

```
handlers --> services --> repositories
```

The `handlers` are HTTP handlers that receive a request and get the data from it, to pass it into the `services` that will execute business logic with that data. Services make use of `repositories` to store the data.

Services and repositories provide an interface so that different versions can be implemented.

For example, different services that implement the same interface can use the same data for use cases that are slightly different, and we could swap the service without having to make big changes to the handlers.

In the case of the repositories, I've implemented in-memory repositories that store the data in memory, but they could be replaced with a Postgres implementation, for example.
(I would choose a relation database like Postgres, due to the nature of the problem and the data. I can image we would want complex queries making different relation between entities. Also we have to guarantee strong data consistency because we are working with sensitive information. 
So I think an ACID database makes more sense in this context.)

I've also implemented mocks for all the interfaces I had to use in unit tests.

For example the **employer** repository defines the interface:
```go
type EmployerRepository interface {
	CreateEmployer(name string) (*model.Employer, error)
}
```
That is implemented by a **real implementation**
```go
type InMemoryEmployerRepository struct {
	employers map[uint]*model.Employer
	nextID    uint
}
```
And also a **mock**

```go
type EmployerRepository struct {
	MockEmployer *model.Employer
	MockErr      error
}
```

## Project Structure

```
├── cmd/
│   └── api/                    
│       └── main.go         
├── internal/
│   ├── handler/             # HTTP handlers
│   │   ├── customer_handler.go
│   │   ├── employer_handler.go
│   │   ├── fund_handler.go
│   │   └── investments_handler.go
│   ├── middleware/         
│   │   └── auth.go            
│   ├── model/              # Data models
│   │   ├── customer.go
│   │   ├── employer.go
│   │   ├── fund.go
│   │   └── investment.go
│   ├── repository/         # Data storage
│   │   ├── customer.go
│   │   ├── employer.go
│   │   ├── fund.go
│   │   └── investment.go
│   └── service/           # Business logic
│       ├── customer.go
│       ├── employer.go
│       ├── fund.go
│       └── investment.go
└── mocks/                
    ├── customer_repository.go
    ├── employer_repository.go
    ├── fund_repository.go
    ├── investment_repository.go
    ├── customer_service.go
    ├── employer_service.go
    ├── fund_service.go
    └── investment_service.go
```

There is a `certs` folder for the certificates used for TLS since the server uses HTTPS. I'm including them in the repo just for simplicity, but I'm aware 
this is not safe.

The `scripts` folder contains scripts to run end to end tests that call the API endpoints using a Python client. Instructions to run these can be found in the `How to Run` section.


## Authentication

I've created a `middleware` for authentication that checks the API key in every request. This middleware uses a simplified `apiKeyRepository` to save and retrieve API keys.

The key is hardcoded just as a demonstration, but in a real scenario we would have a config file or an environment variable to point to the API key, and it would NOT be pushed to the repo.

## Testing

### Unit tests

I added unit tests for the services, repositories and handlers. I've created mocks for the dependencies.

### End to end tests
I've provided a Python script `test_api.py` that can be considered as end-to-end/acceptance tests. 

If I had more time I would've structured it in a better way, with different separated cases. 

In a real scenario I would've used a BDD framework and described the scenarios using Gherkin.

In order to run these end-to-end cases I've written a Python client that calls the API to perform the following operations:
- Create an employer to associate it with a customer so that we can have both `employed customers` and `retail customers`
- Create `two customers`, one retail and one employed
- Create 3 funds: `Fund1`, `Fund2`, `Fund3`
- Retrieve the funds created: this is used to display all the possible funds to the user so that they can pick one
- Create `investments` for both customers:
  - Retail customer invests `2000 in Fund1`, `1500 in Fund3`
  - Employed customer invests `3000 in Fund2`
- Retrieve the investments we've created one by one
- Retrieve the investments associated with each customer

## Improvements

- Proper logging for monitoring
- Add a rate limit
- Better inputs validation / edge cases
- Use microservices instead communicated with events
- More complex authentication, authorization
- Proper API docs with something like Swagger

## How to Run

### Prerequisites
1. Go must be installed (version 1.16 or later recommended)
2. Python must be installed to run the end-to-end test

### Running the Server

From the root of the project:

```bash
go run cmd/api/main.go
```

The server will start on port 8443.

### Example API Calls

```bash
# Create an employer
curl -k -X POST https://localhost:8443/api/employers \
  -H "X-API-Key: test-api-key" \
  -H "Content-Type: application/json" \
  -d '{"name": "Acme Corp"}'

# Create a customer (employed)
curl -k -X POST https://localhost:8443/api/customers \
  -H "X-API-Key: test-api-key" \
  -H "Content-Type: application/json" \
  -d '{"name": "Jane Smith", "employer_id": 1}'

# Create a fund
curl -k -X POST https://localhost:8443/api/funds \
  -H "X-API-Key: test-api-key" \
  -H "Content-Type: application/json" \
  -d '{"name": "Fund1"}'

# Get all funds
curl -k https://localhost:8443/api/funds \
  -H "X-API-Key: test-api-key"
```

> **Note**: Use `-k` flag to skip SSL certificate verification since we're using a self-signed certificate.

### Running the End-to-End Tests

Run the `scripts/run_tests.sh` script from the root of the project. This will:
1. Create a virtualenv
2. Run the Go server
3. Make the Python calls to the API

> **Note**:You might have to give execution permissions to the script to run it: `chmod +x scripts/run_tests.sh`


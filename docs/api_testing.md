# Testing

## Unit Test Suite

This project includes unit tests primarily for the **usecase** layer of Task and User management, ensuring the business logic functions correctly with mocked repository dependencies.

### What’s Covered

#### TaskUsecase:
- Create, get (all & by ID), update, and delete tasks.
- Validation of ObjectID strings.

#### UserUsecase:
- User creation, authentication success/failure.
- Role promotion validation.
- User deletion.

- Error handling for invalid input and repository failures.

---

## How to Run Tests

### Run All Tests

```
go test ./... -v
```

### Run Tests with Coverage

```
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
```

## Test Organization
- Test files are located in the /tests directory.
- Mocks for repository interfaces live in /tests/mocks.
- Tests use the testify package for assertions and mocking.



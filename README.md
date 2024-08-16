# A CRUD app template

This is a template for a CRUD application. It's a simple Todo app.

## Project Structure

- `cmd/main.go`: entry point of the application
- `pkg/domain`: domain layer. It contains the domain models that shape the system.
- `internal/router`: the http router of the application
- `internal/controller`: the controller layer. It checks the incoming requests and delegates the business logic to the service layer.
- `internal/service`: the service layer. It contains the business logic of the application.
- `internal/repository`: the repository layer. It contains the data access logic of the application. The repository layer is responsible for the persistence of the domain models. The service layer depends on the repository layer to retrieve the data from the database.

## How to run

### Dev Container

**Please make sure you have [Docker](https://docs.docker.com/get-docker/) installed.**

1. Open the root path of the project in VSCode.
2. Open command palette (Ctrl+Shift+P) and run `Dev Containers: Reopen in Container`.
3. Happy coding!

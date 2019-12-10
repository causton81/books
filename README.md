# Books App

# Tests
All tests should pass:
```bash
go test ./...
```
They are mostly mocked unit tests, but one hits Google Books API.

# Run the app
```bash
go run ./cmd/books
```

# Design
The app is structured around 3 use cases:
- Query Books
- Add to Reading List
- View Reading List

You can find those under [interactor](interactor)

## Boundaries
The goal was to put boundaries between the UI, interactors, external service and persistence.
The interactors do not know anything about consoles or Google. The "persistence" layer is currently
just an in-memory fake, but it could be replaced with some real persistence client fairly easily.

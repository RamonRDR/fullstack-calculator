# Technical Design

## 1. Overview

This project is a small full-stack calculator. A React + TypeScript frontend collects two operands and one of four required operations. A Go REST API validates the request, executes the calculation through a separate domain package, and returns JSON.

## 2. Architecture

```text
Browser
  -> React + TypeScript frontend
  -> POST /api/calculate
  -> Go HTTP handler
  -> calculator domain logic
```

The frontend uses local component state because the application has one form and no shared client state. API access is separated into `frontend/src/api/calculator.ts`.

The backend has two meaningful layers:

- `internal/calculator`: calculation rules and domain errors
- `internal/httpapi`: HTTP routing, JSON validation, error mapping, and responses

`cmd/server` only starts the HTTP server.

## 3. API Design

The API uses one endpoint:

`POST /api/calculate`

Request:

```json
{
  "operation": "add",
  "a": 10,
  "b": 5
}
```

Success:

```json
{
  "result": 15
}
```

Error:

```json
{
  "error": "division by zero"
}
```

Supported operations are `add`, `subtract`, `multiply`, and `divide`.

A single endpoint avoids repetitive handlers while keeping the request contract explicit. Separate endpoints per operation would add routing code without providing a meaningful benefit for this assessment.

## 4. Error Handling

The frontend validates that both entered values are present and finite before making a request. API or network errors are displayed in a user-readable message.

The backend rejects malformed JSON, unknown JSON fields, multiple JSON values, missing fields, unsupported operations, division by zero, and calculations that cannot produce a finite JSON number. Client errors return `400`; unsupported HTTP methods return `405`; unexpected server failures are mapped to `500` without exposing internals.

## 5. Testing Strategy

Backend domain tests cover all four operations, division by zero, unsupported operations, invalid operands, and non-finite results. HTTP tests cover a successful calculation plus request validation and method handling.

Frontend tests exercise user behavior: entering operands, selecting an operation, submitting to the API, rendering a result, preventing invalid submissions, displaying backend errors, and displaying network errors.

Coverage commands are included for both applications, but the implementation does not chase a target percentage. Behavioral confidence is more important than maximizing a metric.

## 6. Design Decisions

- Go standard library HTTP avoids a framework dependency for one endpoint.
- Calculator rules are isolated from HTTP so they are straightforward to test and explain.
- React local state is enough for the UI and avoids unnecessary state management.
- The frontend calls a relative `/api` path. Vite proxies it to the Go server in development.
- Optional calculator operations are not implemented to keep the assessment focused on the required behavior.
- Docker is not included because Go + Node local setup is already small and direct.

## 7. Assumptions

- Every supported operation uses two operands.
- Inputs and results must be finite numbers representable by JSON.
- Development uses the Vite proxy and backend port `8080` unless `PORT` is set.
- A production deployment would provide same-origin routing or an equivalent reverse proxy.

## 8. Trade-offs

This is intentionally not an enterprise architecture. There is no database, authentication, dependency-injection container, generalized service layer, request framework, generated API schema, or deployment platform configuration.

Those additions could be reasonable in a larger product, but in a 2–4 hour technical assessment they would create more code to explain and test without improving the required calculator behavior.

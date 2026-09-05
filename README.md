# Full-Stack Calculator

A small full-stack calculator built for a Junior Software Engineer technical assessment. The React frontend collects two operands and an operation, then sends the calculation to a Go REST API.

The project intentionally favors correctness, readability, testability, and a small amount of well-justified structure over extra features.

## Architecture

```text
Browser
  |
  v
React + TypeScript
  |
  | POST /api/calculate (JSON)
  v
Go HTTP API
  |
  v
Calculator domain logic
```

The Go calculator package contains the business rules and has no HTTP dependency. The HTTP package owns JSON validation and status codes. On the frontend, API communication is isolated in `src/api/calculator.ts` while `App.tsx` handles form state and user feedback.

See [`docs/technical-design.md`](docs/technical-design.md) for the concise design rationale.

## Technology stack

- React 18
- TypeScript
- Vite
- Vitest + Testing Library
- Go 1.26+
- Go standard library (`net/http`, `encoding/json`, `testing`)
- GitHub Actions for automated quality checks

No database, authentication, state-management library, UI framework, or backend web framework is used.

## Prerequisites

- Go 1.26 or newer
- Node.js 20.19+ or 22.12+
- npm

## Run locally

### 1. Start the backend

```bash
cd backend
go run ./cmd/server
```

The API listens on `http://localhost:8080` by default. Set the `PORT` environment variable to use another port.

### 2. Start the frontend

In another terminal:

```bash
cd frontend
npm install
npm run dev
```

Open the local URL printed by Vite. During development, Vite proxies `/api` requests to `http://localhost:8080`, so no separate CORS configuration is required.

## API

### `POST /api/calculate`

Request:

```json
{
  "operation": "add",
  "a": 10,
  "b": 5
}
```

Supported operations:

- `add`
- `subtract`
- `multiply`
- `divide`

Success response (`200 OK`):

```json
{
  "result": 15
}
```

Validation error (`400 Bad Request`):

```json
{
  "error": "division by zero"
}
```

Example with `curl`:

```bash
curl -X POST http://localhost:8080/api/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation":"multiply","a":6,"b":7}'
```

The endpoint rejects malformed JSON, unknown fields, missing operands, missing operations, unsupported operations, division by zero, and non-finite calculations. Other HTTP methods return `405 Method Not Allowed`.

## Tests and quality checks

### Backend tests

```bash
cd backend
go test ./...
```

Backend coverage report:

```bash
go test ./... -coverprofile=coverage.out
go tool cover -func=coverage.out
```

Backend build:

```bash
go build ./cmd/server
```

### Frontend tests

After `npm install`:

```bash
cd frontend
npm test
```

Frontend coverage report:

```bash
npm run coverage
```

The HTML report is generated under `frontend/coverage/` and is intentionally ignored by Git.

Type checking:

```bash
npm run typecheck
```

Production build:

```bash
npm run build
```

The project does not configure a separate frontend linter. TypeScript strict checking, behavioral tests, production builds, and Go formatting checks provide the intentionally small quality gate for this assessment.

GitHub Actions runs Go formatting, backend tests and coverage, the backend build, frontend type checking, frontend tests and coverage, and the frontend production build on pushes and pull requests targeting `main`.

## Important design decisions

- **One calculation endpoint:** `POST /api/calculate` keeps the API small and consistent for all required operations.
- **Business logic separated from HTTP:** calculator rules can be tested without JSON or HTTP concerns.
- **Same-origin frontend API path:** the frontend always calls `/api/calculate`; Vite supplies a development proxy. A production deployment would normally place both applications behind the same origin or reverse proxy.
- **No extra state library:** React local state is sufficient for a single form.
- **Required operations only:** exponentiation, square root, and percentage were intentionally not implemented because they are optional and do not improve the core assessment goals.
- **No Docker:** local setup requires only Go and Node, so Docker would add files and concepts without solving a requirement for this small exercise.

## Assumptions

- Both operands are required for every supported operation.
- The service accepts finite JSON numbers and returns finite JSON results.
- Development runs the frontend and backend as separate processes, with Vite proxying API requests.
- Production hosting and cross-origin deployment configuration are outside the assessment scope.

## AI-assisted development

AI tools were explicitly allowed for this assessment and were used as an engineering assistant for planning, architecture, implementation, testing, documentation, and review.

Meaningful AI prompts are preserved under [`ai-prompts/`](ai-prompts/). The initial project prompt is recorded in full rather than creating artificial prompt history for small implementation decisions.

All AI-generated suggestions were reviewed, adapted, tested, and validated before being included in the final solution.

<div align="center">

# Full-Stack Calculator 🧮

**A lightweight full-stack calculator powered by React, TypeScript, and a Go REST API.**

<p>
  <img alt="Go 1.26" src="https://img.shields.io/badge/Go-1.26-00ADD8?logo=go&amp;logoColor=white" />
  <img alt="TypeScript 5.7.2" src="https://img.shields.io/badge/TypeScript-5.7.2-3178C6?logo=typescript&amp;logoColor=white" />
  <img alt="Node.js 24 CI" src="https://img.shields.io/badge/Node.js-24%20CI-339933?logo=nodedotjs&amp;logoColor=white" />
  <img alt="React 18.3.1" src="https://img.shields.io/badge/React-18.3.1-149ECA?logo=react&amp;logoColor=white" />
  <a href="https://github.com/RamonRDR/fullstack-calculator/actions/workflows/quality.yml"><img alt="GitHub Actions Quality" src="https://github.com/RamonRDR/fullstack-calculator/actions/workflows/quality.yml/badge.svg" /></a>
</p>

<p>
  <img alt="Vite 7.3.6" src="https://img.shields.io/badge/Vite-7.3.6-646CFF?logo=vite&amp;logoColor=white" />
  <img alt="HTML5 UI" src="https://img.shields.io/badge/HTML5-UI-E34F26?logo=html5&amp;logoColor=white" />
  <img alt="CSS3 responsive" src="https://img.shields.io/badge/CSS3-responsive-1572B6?logo=css3&amp;logoColor=white" />
  <img alt="npm package manager" src="https://img.shields.io/badge/npm-package%20manager-CB3837?logo=npm&amp;logoColor=white" />
  <img alt="Go net/http REST API" src="https://img.shields.io/badge/Go%20net%2Fhttp-REST%20API-00ADD8?logo=go&amp;logoColor=white" />
</p>

<p>
  <a href="https://github.com/RamonRDR/fullstack-calculator/actions/runs/33995777489"><img alt="Go tests passing" src="https://img.shields.io/badge/Go%20tests-passing-00ADD8?logo=go&amp;logoColor=white" /></a>
  <a href="https://github.com/RamonRDR/fullstack-calculator/actions/runs/33995777489"><img alt="Vitest 4 passing" src="https://img.shields.io/badge/Vitest-4%20passing-6E9F18?logo=vitest&amp;logoColor=white" /></a>
  <img alt="React Testing Library 16.1.0" src="https://img.shields.io/badge/React%20Testing%20Library-16.1.0-E33332?logo=testinglibrary&amp;logoColor=white" />
  <img alt="User Event 14.5.2" src="https://img.shields.io/badge/User%20Event-14.5.2-E33332" />
  <img alt="jsdom 30.0.1" src="https://img.shields.io/badge/jsdom-30.0.1-6C757D" />
  <img alt="V8 Coverage enabled" src="https://img.shields.io/badge/V8%20Coverage-enabled-7C4DFF" />
</p>

<p>
  <a href="https://github.com/RamonRDR/fullstack-calculator/actions/runs/33995777489"><img alt="Go coverage 86.0%" src="https://img.shields.io/badge/Go%20coverage-86.0%25-2E7D32?logo=go&amp;logoColor=white" /></a>
  <a href="https://github.com/RamonRDR/fullstack-calculator/actions/runs/33995777489"><img alt="Frontend coverage 92.85%" src="https://img.shields.io/badge/Frontend%20coverage-92.85%25-2E7D32" /></a>
  <a href="https://github.com/RamonRDR/fullstack-calculator/actions/runs/33995777489"><img alt="gofmt checked" src="https://img.shields.io/badge/gofmt-checked-00ADD8?logo=go&amp;logoColor=white" /></a>
  <a href="https://github.com/RamonRDR/fullstack-calculator/actions/runs/33995777489"><img alt="TypeScript strict" src="https://img.shields.io/badge/TypeScript-strict-3178C6?logo=typescript&amp;logoColor=white" /></a>
  <a href="https://github.com/RamonRDR/fullstack-calculator/actions/runs/33995777489"><img alt="Go build passing" src="https://img.shields.io/badge/Go%20build-passing-2E7D32?logo=go&amp;logoColor=white" /></a>
  <a href="https://github.com/RamonRDR/fullstack-calculator/actions/runs/33995777489"><img alt="Vite build passing" src="https://img.shields.io/badge/Vite%20build-passing-2E7D32?logo=vite&amp;logoColor=white" /></a>
  <a href="https://github.com/RamonRDR/fullstack-calculator/actions/runs/33995777489"><img alt="Integration smoke PASS" src="https://img.shields.io/badge/Integration%20smoke-PASS-2E7D32" /></a>
</p>

<p>
  <a href="https://github.com/RamonRDR/fullstack-calculator/tree/main/ai-prompts"><img alt="ChatGPT / Codex AI assisted" src="https://img.shields.io/badge/ChatGPT%20%2F%20Codex-AI%20assisted-412991?logo=openai&amp;logoColor=white" /></a>
</p>

<sub>Test counts and coverage are verified snapshots of <a href="https://github.com/RamonRDR/fullstack-calculator/actions/runs/33995777489">Quality #7 (2026-09-05)</a>: Vitest 4 passing, Go coverage 86.0%, frontend statement coverage 92.85%. The CI badge is dynamic; other quality badges reflect that run. Node.js 24 is the CI runtime (minimum declared Node.js version: 20.19.0).</sub>

</div>

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

GitHub Actions runs Go formatting, backend tests and coverage, the backend build, frontend type checking, frontend tests and coverage, and the frontend production build on pushes and pull requests targeting `main`. A final smoke job starts both development servers, verifies the Go endpoint directly, and repeats the calculation through Vite's `/api` proxy to validate the full development path.

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

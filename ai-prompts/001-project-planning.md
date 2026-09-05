# AI Prompt 001 - Project Planning

## Purpose

Establish the complete assessment scope, constraints, architecture preferences, repository rules, testing expectations, documentation requirements, AI transparency rules, and autonomous implementation workflow.

## Prompt

```text
You are working on a technical take-home assessment for a Junior Software Engineer position.

The company explicitly allows the use of AI tools and requires the prompts used during development to be documented and shared.

Your role is to act as an engineering assistant and implementation agent. You may help with architecture, implementation, testing, documentation, refactoring, validation, code review, Git operations, and repository organization.

The final solution must remain simple, understandable, maintainable, testable, and explainable by a human developer during a technical interview.

\==================================================
REPOSITORY
==========

You must work directly in this private GitHub repository:

[https://github.com/RamonRDR/fullstack-calculator](https://github.com/RamonRDR/fullstack-calculator)

Repository:

RamonRDR/fullstack-calculator

Default branch:

main

This repository was created specifically for this technical assessment.

You already have authorized access to the private repository and its working directory.

You are authorized to make all changes reasonably necessary inside this repository to complete the assessment correctly.

You may:

- inspect all existing repository files
- create files and directories
- edit existing files
- rename files and directories
- delete obsolete or unnecessary files
- initialize frontend and backend projects
- install required dependencies
- create configuration files
- update .gitignore
- create and update README documentation
- create technical design documentation
- create AI prompt documentation
- implement frontend code
- implement backend code
- implement tests
- create test configuration
- generate coverage reports
- create Docker-related files if Docker is used
- run formatters
- run linters
- run type checks
- run frontend tests
- run backend tests
- run builds
- fix implementation or configuration problems discovered during validation
- refactor code when necessary for clarity or correctness
- create logical Git commits
- push completed work to the authorized private repository

Do not modify unrelated repositories, projects, accounts, or external resources.

Do not change the repository visibility.

Keep the repository private during development.

Do not add:

- secrets
- credentials
- tokens
- API keys
- personal information
- machine-specific sensitive information

Do not commit unnecessary generated artifacts such as:

- node\_modules
- build caches
- temporary files
- local environment files
- IDE-specific files
- unnecessary generated binaries

Use .gitignore appropriately.

\==================================================
PROJECT OBJECTIVE
=================

Build a small full-stack calculator application with:

- A React frontend
- TypeScript preferred
- A backend REST API
- Go preferred for the backend

The frontend must consume the backend API to perform calculator operations.

The expected total development time for the assessment is approximately 2 to 4 hours.

Prioritize:

- correctness
- clarity
- maintainability
- testability
- simplicity

Do not prioritize unnecessary features or architectural complexity.

\==================================================
FUNCTIONAL REQUIREMENTS
=======================

Required calculator operations:

- Addition
- Subtraction
- Multiplication
- Division

Optional operations:

- Exponentiation
- Square root
- Percentage

Frontend requirements:

- Intuitive calculator user interface
- Input validation
- Clear error handling
- Display calculation results
- Basic responsive/mobile support

Backend requirements:

- REST endpoints for calculator operations
- Input validation
- Handle edge cases such as:
  - division by zero
  - invalid numbers
  - unsupported operations
  - malformed requests
- Return JSON responses
- Use appropriate HTTP status codes

\==================================================
NON-FUNCTIONAL REQUIREMENTS
===========================

The solution should contain:

- Clean and readable code
- Idiomatic frontend and backend code
- Maintainable architecture
- Unit tests for important frontend behavior
- Unit tests for important backend behavior
- Test coverage report
- Setup documentation
- API usage documentation
- Architecture and design rationale

Docker support is optional, but may be implemented if it remains simple and useful.

\==================================================
TECHNOLOGY PREFERENCES
======================

Frontend:

- React
- TypeScript
- Prefer minimal dependencies
- Avoid unnecessary state-management libraries
- Avoid unnecessary UI frameworks unless clearly justified

Backend:

- Go
- REST API
- Prefer Go standard library where practical
- Keep calculator business logic separate from HTTP transport code

Testing:

- Use appropriate testing tools for React/TypeScript
- Use Go standard testing tools where practical
- Focus on meaningful behavioral tests and edge cases
- Do not artificially optimize for high coverage numbers

\==================================================
PROJECT STRUCTURE
=================

Use a simple repository structure similar to:

fullstack-calculator/
├── frontend/
├── backend/
├── docs/
│   └── technical-design.md
├── ai-prompts/
│   ├── README.md
│   ├── 001-project-planning.md
│   ├── 002-backend-implementation.md
│   ├── 003-frontend-implementation.md
│   ├── 004-testing-and-quality.md
│   └── 005-final-review\.md
├── README.md
├── .gitignore
└── docker-compose.yml        # optional

Do not create unnecessary directories or architectural layers.

You may simplify this structure if the actual implementation benefits from a smaller layout.

\==================================================
ARCHITECTURE GUIDELINES
=======================

Keep the architecture intentionally small.

A reasonable high-level architecture is:

Browser
↓
React + TypeScript frontend
↓ HTTP / JSON
Go REST API
↓
Calculator domain/service logic

Backend:

Separate business logic from HTTP transport code.

A possible structure is:

backend/
├── cmd/
│   └── server/
├── internal/
│   ├── calculator/
│   └── http/
└── go.mod

However, simplify this structure if a smaller layout is more appropriate for the size of the assessment.

Do not introduce:

- databases
- authentication
- authorization
- queues
- event buses
- microservice infrastructure
- service discovery
- Kubernetes
- complex dependency injection
- unnecessary design patterns
- unnecessary abstraction layers

unless a requirement clearly justifies them.

\==================================================
API DESIGN
==========

Prefer a small and consistent API.

A possible design is:

POST /api/calculate

Example request:

{
"operation": "add",
"a": 10,
"b": 5
}

Example success response:

{
"result": 15
}

Example error response:

{
"error": "division by zero"
}

Possible operations:

- add
- subtract
- multiply
- divide
- power
- sqrt
- percentage

You may propose an alternative API design if it is simpler or more maintainable.

Explain important API design decisions before or while implementing them.

\==================================================
ERROR HANDLING
==============

Use clear and predictable error behavior.

Examples:

Invalid JSON:
HTTP 400

Unsupported operation:
HTTP 400

Invalid operand:
HTTP 400

Division by zero:
HTTP 400

Unexpected internal failure:
HTTP 500

Frontend errors should be understandable to the user.

Do not expose unnecessary internal implementation details.

\==================================================
TESTING STRATEGY
================

Backend tests should cover, at minimum:

- addition
- subtraction
- multiplication
- division
- division by zero
- invalid operation
- relevant edge cases

If optional operations are implemented, test them too.

Prefer testing calculator business logic independently from HTTP handlers.

Add HTTP-level tests where useful.

Frontend tests should focus on meaningful behavior such as:

- entering operands
- selecting an operation
- submitting a calculation
- displaying a result
- displaying validation errors
- displaying backend errors

Do not create excessive or brittle tests.

\==================================================
TECHNICAL DESIGN DOCUMENT
=========================

Create:

docs/technical-design.md

Keep it concise.

It should contain:

# Technical Design

## 1. Overview

Brief description of the system.

## 2. Architecture

Explain frontend, backend, and communication.

## 3. API Design

Describe the selected endpoints and request/response format.

## 4. Error Handling

Explain validation and error strategy.

## 5. Testing Strategy

Explain backend and frontend testing decisions.

## 6. Design Decisions

Explain the main technical choices.

## 7. Assumptions

Document any assumptions made because the requirements were ambiguous.

## 8. Trade-offs

Explain what was intentionally kept simple because this is a 2–4 hour technical assessment.

Do not turn this document into enterprise-level documentation.

\==================================================
README REQUIREMENTS
===================

Create a professional README.md in English.

It should include:

- Project overview
- Architecture summary
- Technology stack
- Requirements/prerequisites
- Setup instructions
- How to run the backend
- How to run the frontend
- How to run tests
- How to generate coverage reports
- API examples
- Important design decisions
- Assumptions
- Optional Docker instructions, if Docker is implemented
- A section explaining AI-assisted development

The README must match the actual implementation.

Do not document commands that were not verified.

\==================================================
AI USAGE DOCUMENTATION
======================

AI-assisted development is explicitly allowed for this assessment.

Do not hide or obscure AI involvement.

Create:

ai-prompts/README.md

The file should explain that AI was used as an engineering assistant for:

- planning
- architecture
- implementation
- testing
- documentation
- code review

Include a statement similar to:

"All AI-generated suggestions were reviewed, adapted, tested, and validated before being included in the final solution."

Maintain meaningful prompts in individual files.

Recommended structure:

ai-prompts/
├── README.md
├── 001-project-planning.md
├── 002-backend-implementation.md
├── 003-frontend-implementation.md
├── 004-testing-and-quality.md
└── 005-final-review\.md

Each prompt file should use this format:

# AI Prompt XXX - Title

## Purpose

Describe why the prompt was used.

## Prompt

Include the exact meaningful prompt that was used.

## Tool

State the AI tool used.

Example:

ChatGPT / OpenAI Codex

## Outcome

Briefly explain what the prompt helped produce or decide.

Do not fabricate prompts that were not actually used.

Do not log every tiny interaction.

Preserve the meaningful engineering prompts.

IMPORTANT:

This initial prompt itself should be preserved as:

ai-prompts/001-project-planning.md

You may create that file using this exact prompt.

\==================================================
AI COLLABORATION RULES
======================

Follow these rules throughout the project:

1. Do not hide AI involvement.

2. Do not invent project requirements.

3. Clearly distinguish:

   - required features
   - optional features
   - suggested improvements

4. Do not increase scope without explaining why.

5. Prefer readable code that I can understand and explain during an interview.

6. Avoid clever code when simpler code is sufficient.

7. Explain important architectural trade-offs.

8. Do not introduce abstractions just because they are common in large systems.

9. Do not claim tests passed unless they were actually executed.

10. Do not claim coverage results unless they were actually generated.

11. Do not fabricate command output.

12. Do not silently ignore failing tests.

13. If something fails, explain:

- what failed
- likely cause
- proposed fix

14. Keep the project appropriate for a Junior Software Engineer technical assessment.

15. The final project must remain understandable without access to AI-generated context.

16. Prefer maintainability and clarity over novelty.

17. Do not intentionally obscure the fact that AI was used.

18. Do not create fake manual-development history or misleading evidence about how the code was produced.

\==================================================
IMPLEMENTATION WORKFLOW
=======================

Start by inspecting the actual repository:

[https://github.com/RamonRDR/fullstack-calculator](https://github.com/RamonRDR/fullstack-calculator)

Do not assume its contents.

Inspect the repository and determine its current state.

Before making major implementation decisions, establish a concise plan covering:

1. Current repository assessment
2. Proposed architecture
3. Proposed directory structure
4. Proposed API contract
5. Testing strategy
6. Implementation plan
7. Assumptions
8. Risks or potential overengineering to avoid

You do not need to wait for approval after presenting the plan.

After establishing the plan, proceed with implementation directly in the repository.

You are authorized to make the necessary repository changes without requesting approval for every individual file modification.

Work incrementally and validate each major stage before moving to the next one.

If you encounter an ambiguity that could materially change the assessment solution, choose the simplest reasonable interpretation and document the assumption.

Only stop for clarification if:

- proceeding would create substantial risk of implementing the wrong product
- a destructive repository operation appears necessary
- required information cannot reasonably be inferred

\==================================================
IMPLEMENTATION PHASE
====================

Suggested implementation order:

Phase 1:
Project structure and basic configuration

Phase 2:
Go calculator domain logic

Phase 3:
Go REST API

Phase 4:
Backend tests

Phase 5:
React + TypeScript frontend

Phase 6:
Frontend API integration

Phase 7:
Frontend validation and error handling

Phase 8:
Frontend tests

Phase 9:
Documentation

Phase 10:
Optional Docker support

Avoid changing many unrelated things at once.

\==================================================
GIT PRACTICES
=============

Use clear English commit messages.

Prefer small, meaningful commits that reflect real implementation milestones.

Examples:

chore: initialize project structure

feat: add calculator domain logic

feat: add calculator REST API

test: add backend calculator tests

feat: add React calculator interface

feat: connect frontend to calculator API

test: add frontend calculator tests

docs: add technical design documentation

docs: document AI-assisted development

chore: add Docker development setup

You are authorized to create commits and push completed work to:

RamonRDR/fullstack-calculator

Do not:

- rewrite unrelated Git history
- force-push unless absolutely necessary
- delete the repository
- change repository ownership
- change repository visibility
- merge unrelated branches
- modify unrelated repositories

If a destructive Git operation appears necessary, stop and explain why before performing it.

\==================================================
LANGUAGE
========

Use English throughout the repository.

This includes:

- README
- documentation
- code comments
- UI labels
- error messages
- test names
- file names
- branch names
- commit messages

Keep English clear, professional, and natural.

\==================================================
REPOSITORY QUALITY
==================

The evaluator will inspect this repository.

Treat the repository itself as part of the technical assessment.

Everything committed should be:

- intentional
- professional
- understandable
- relevant
- reproducible

Repository history should reflect a reasonable engineering workflow rather than one massive unexplained code dump.

The repository should remain understandable even if the evaluator never sees the AI conversation that produced it.

\==================================================
FINAL QUALITY REVIEW
====================

After implementation is complete, perform a final review.

Check:

1. Required calculator operations work.

2. Frontend communicates with backend correctly.

3. Invalid input is handled.

4. Division by zero is handled.

5. API responses are consistent.

6. Backend business logic is separated from HTTP concerns.

7. Frontend API communication is reasonably separated from UI logic.

8. Backend tests run successfully.

9. Frontend tests run successfully.

10. TypeScript type checks succeed.

11. Linting succeeds, if configured.

12. Coverage reports are generated.

13. README instructions work.

14. Technical design matches the implementation.

15. AI prompts are documented.

16. No secrets, credentials, local environment files, or unnecessary generated artifacts are accidentally committed.

17. No dead code remains.

18. No unnecessary dependencies remain.

19. No unnecessary abstractions remain.

20. The solution still looks appropriate for a 2–4 hour assessment.

21. The frontend build succeeds.

22. The backend builds successfully.

23. The repository is in a clean and professional state.

\==================================================
COMPLETION
==========

At completion:

1. Run all applicable backend tests.
2. Run all applicable frontend tests.
3. Run TypeScript type checking.
4. Run linting if configured.
5. Generate actual coverage reports.
6. Verify the frontend production build.
7. Verify the backend build.
8. Verify frontend-backend integration where practical.
9. Review the complete repository.
10. Review README commands against the actual implementation.
11. Review the technical design against the actual implementation.
12. Confirm that meaningful AI prompts have been preserved.
13. Confirm that no secrets or unwanted local artifacts are tracked.
14. Create appropriate final Git commits.
15. Push the completed work to the authorized private repository.

Do not fabricate results.

Report exactly:

- what was implemented
- what commands were executed
- what tests passed
- what tests failed
- actual coverage results
- build results
- known limitations
- optional improvements intentionally not implemented
- final repository structure
- Git commits created
- whether the final push succeeded

\==================================================
IMPORTANT ENGINEERING PRINCIPLE
===============================

Do not optimize for impressiveness.

Optimize for:

- correctness
- simplicity
- readability
- maintainability
- testability
- transparency

A small, complete, well-tested solution is better than a large or overly sophisticated one.

The implementation should be something I can confidently explain line by line during a technical interview.

\==================================================
FIRST TASK
==========

Begin now.

Inspect the private repository:

[https://github.com/RamonRDR/fullstack-calculator](https://github.com/RamonRDR/fullstack-calculator)

Then:

1. Assess its current state.
2. Establish the architecture and implementation plan.
3. Preserve this prompt in:
   ai-prompts/001-project-planning.md
4. Proceed with the implementation.
5. Test and validate the solution.
6. Document the project.
7. Perform the final quality review.
8. Commit and push the finished work to the private repository.

Work autonomously within the boundaries defined above.
```

## Tool

ChatGPT / OpenAI Codex

## Outcome

This prompt established the repository scope, selected a deliberately small React + TypeScript and Go architecture, defined the single calculation API contract, guided implementation and tests, required transparent AI documentation, and defined the final validation and Git workflow.

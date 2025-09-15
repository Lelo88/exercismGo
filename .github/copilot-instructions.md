# Copilot Instructions for Exercism Go Track

## Overview
This repository contains multiple Go exercises, each in its own directory. Each exercise is self-contained with its own `README.md`, `go.mod`, and test files. The goal of this repository is to help developers practice Go by solving small, focused problems.

## Repository Structure
- Each exercise is in its own directory (e.g., `isogram/`, `dnd-character/`).
- Common files in each exercise:
  - `README.md`: Describes the problem and requirements.
  - `*.go`: Contains the implementation.
  - `*_test.go`: Contains the test cases.
  - `go.mod`: Defines the Go module for the exercise.

## Development Workflow
1. **Select an Exercise**: Navigate to the directory of the exercise you want to work on.
2. **Read the Problem Statement**: Refer to the `README.md` file in the exercise directory.
3. **Implement the Solution**: Edit the `*.go` file to implement the solution.
4. **Run Tests**: Use the following command to run tests for the exercise:
   ```bash
   go test ./<exercise-directory>
   ```
   Example:
   ```bash
   go test ./isogram
   ```
5. **Debugging**: Use print statements or a debugger to troubleshoot issues.

## Project-Specific Conventions
- Each exercise is independent and does not depend on other exercises.
- Follow idiomatic Go practices (e.g., naming conventions, error handling).
- Tests are the primary way to validate your solution.

## Key Files and Patterns
- **Test Files**: Test files are named `*_test.go` and use Go's `testing` package.
- **Modules**: Each exercise has its own `go.mod` file, making it a standalone Go module.
- **Examples**: Refer to the `README.md` files for examples and expected behavior.

## External Dependencies
- No external dependencies are used; all exercises rely on Go's standard library.

## Tips for AI Agents
- Focus on the specific exercise directory when implementing or testing.
- Use the examples in `README.md` to understand the expected behavior.
- Ensure all tests pass before considering the solution complete.

## Example Commands
- Run tests for all exercises:
  ```bash
  go test ./...
  ```
- Run tests for a specific exercise:
  ```bash
  go test ./<exercise-directory>
  ```

## Additional Resources
- Refer to the `HELP.md` files in each exercise directory for more guidance.
- Visit the [Exercism Go Track](https://exercism.org/tracks/go) for more information.
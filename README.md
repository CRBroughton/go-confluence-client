# Go Confluence Client

A basic client for the Confluence V2 API. The client is still in rapid development
and should be considered unstable.

The below is a table of working and in-progress
features:

| Feature                   |  Supported  |
| ------------------------- |  :-------:  |
| Pages - Get page by ID    |     ✅      |
| Pages - Update page by ID |     ✅      |
| Spaces - Get spaces       |     ✅      |
| Spaces - Get space by ID  |     ✅      |

## Testing

To run the tests, you can run `go test -v ./...` which will run both
the unit and integration tests. Before running integration tests

## Contributing

Before contributing, please raise an issue that outlines the change(s) you'd like
to make. Once the change has been approved, then proceed with the following
steps.

To contribute to this repository, you'll need the following installed:

- https://go.dev - The Go Programming Language
- https://marketplace.visualstudio.com/items?itemName=golang.go - The Go VSCode Extension
- https://pnpm.io - PNPM for changesets and versioning, then you can run `pnpm i`

Each Pull Request should feature a changeset file, ideally on the same commit
for the majority of the feature you are introducing into the codebase. This
project follows the SemVer versioning system, so please follow accordingly.
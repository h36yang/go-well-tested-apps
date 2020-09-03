# go-well-tested-apps
Practice materials for "Creating Well-Tested Applications in Go" course

# Introduction to Testing

## Testing in Go
* Full-featured testing and profiling tools
* Complete API for controlling test execution
* Little support for assertions (intentionally)

## Test Types in Go
* Test - unit, integration, e2e tests
* Benchmark - performance profiling
* Example - documentation

## Testing Related Packages in Go Standard Library
* `testing` - https://golang.org/pkg/testing
  * Main testing package in Go
* `testing/quick` - https://golang.org/pkg/testing/quick
  * Designed to simplify black box testing
* `testing/iotest` - https://golang.org/pkg/testing/iotest
  * Provides easy ways to create readers and writers for testing input and output
* `net/http/httptest` - https://golang.org/pkg/net/http/httptest
  * Allows to simulate HTTP requests, record HTTP responses, and set up test servers

## Useful Community Projects
* Testify - https://github.com/stretchr/testify
  * Provides the complete assertion experience
* Ginkgo - https://github.com/onsi/ginkgo
  * Behaviour Driven Development (BDD) Testing Framework
* GoConvey - https://github.com/smartystreets/goconvey
  * Generates results in browser based format
* httpexpect - https://github.com/gavv/httpexpect
  * Designed to simplify e2e web service testing
* GoMock - https://github.com/golang/mock
  * Mocking Framework
* go-sqlmock - https://github.com/DATA-DOG/go-sqlmock
  * Mocking Library implementing `sql/driver` as in-memory mockable database

# Deep Dives
1. [Creating and Running Tests](01_Creating_and_Running_Tests.md)
2. [Benchmarking and Profiling](02_Benchmarking_And_Profiling.md)

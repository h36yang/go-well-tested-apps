# Creating and Running Tests

## Naming Conventions
* `_test` suffix in file names
* `Test` prefix in test functions followed by an upper case letter
* Accept one parameter `t *testing.T`
* Same package name for whitebox tests
* Add `_test` suffix to package name for blackbox tests

## Writing Tests
* Receive `*testing.T` object
* Write tests using normal Go code - no special assertion API
* Report errors using methods on `*testing.T` object

## Reporting Test Failures
|Immediate Failures|Non-Immediate Failures|
|---:|:---|
|`t.FailNow()`|`t.Fail()`|
|`t.Fatal(args ...interface{})`|`t.Error(args ...interface{})`|
|`t.Fatalf(format string, args ...interface{})`|`t.Errorf(format string, args ...interface{})`|

## Running Tests
* `go test` - Run all tests in current directory
* `go test {pkg1} {pkg2} ... {pkgn}` - Test specified packages
* `go test ./...` - Run tests in current package and descendants
* `go test -v` - Generate verbose output
* `go test -run {regexp}` - Run only tests matching `{regexp}`

## Generating Test Coverage
* `go test -cover` - Generate overall test coverage in console
* `go test -coverprofile {filename}.out` - Generate more detailed test coverage to file
  * `go tool cover -func {filename}.out` - Analyze and display the test coverage file in console
  * `go tool cover -html {filename}.out` - Analyze and display the test coverage file in webpage
* `go test -coverprofile {filename}.out -covermode count` - Generate more detailed test coverage to file and also indicate whether a function has relatively high or low coverage in HTML report

## Other Useful Functions
* `Log` and `Logf` - writing output without failing
* `Helper` - indicating a Test function is actually a Helper function so test runner skips it
* `Skip`, `Skipf` and `SkipNow` - skipping Test functions
* `Run` - allowing top level Test function to run sub Test functions
* `Parallel` - indicating the Test function can be executed in parallel

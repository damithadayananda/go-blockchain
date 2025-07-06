#unit test
gotestsum --junitfile results.xml -- -v -coverprofile=coverage.out ../...
#get total unit test coverage
go tool cover -func=coverage.out | grep total
#unit test coverage as html representation
go tool cover -html=coverage.out -o coverage.html



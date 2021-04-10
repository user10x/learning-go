# Adds Quiz questions from the CSV File
# Uses anonymous go func
# adds timer package
# uses command cli flags with go flags package
## example usages below
### go run quiz.go -csv test
### go run  -csv=ab c.csv
### go run  quiz.go -csv="test"
### go run  quiz.go -csv=test.csv
### go run  run quiz.go -csv=problems.csv -limit=20
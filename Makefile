.PHONY: range
## range: run range.go file
range:
	go run week-one/range/range.go

.PHONY: function
## function: run function.go file
function:
	go run week-one/function/function.go

.PHONY: method
## method: run method.go file
method:
	go run week-one/method/main.go

.PHONY: interface
## interface: run interface.go file
interface:
	go run week-one/interface/main.go

.PHONY: maps
## maps: run maps.go file
maps:
	go run week-one/maps/main.go

.PHONY: array-interview-questions
## array-interview-questions: run array-interview-questions.go file
array-interview-questions:
	go run week-one/array-interview-questions/main.go

.PHONY: recursive
## recursive: run recursive.go file
recursive:
	go run week-one/recursive/main.go


.PHONY: exercises
## exercises: run main.go file
exercises:
	go run week-one/exercises/main.go


.PHONY: help
## help: prints this help message
help:

	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
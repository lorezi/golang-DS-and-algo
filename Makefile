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


.PHONY: help
## help: prints this help message
help:

	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
installDependencies:
	- go get -u github.com/gofiber/fiber/v2
	- go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	- go install github.com/cosmtrek/air@latest

runDev:
	- air


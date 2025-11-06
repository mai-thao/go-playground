# book-server
This is a simple book server that exposes a RESTful API to interact with its book resource. Its data are stored in memory using a slice. 

It uses the [Gin](https://gin-gonic.com/en/docs/) web framework and follows the opinionated [golang-standards](https://github.com/golang-standards/project-layout) project layout.

### How to Run the App
1) Start the PostgreSQL Docker container with: `docker compose up -d`
2) Execute the command: `go run ./cmd/memory`
3) Interact with the endpoints exposed by this API using [Insomnia](https://insomnia.rest/features/api-testing) (or [Postman](https://www.postman.com/product/what-is-postman/), [curl](https://curl.se/))
4) Terminate the app with: Ctrl + C 
5) Stop the Docker container with: `docker compose down`

### How to Get Dependencies
1) Execute the command: `go get .` (resolve and download dependencies for code in the current directory)
   1) `go mod tidy` to add missing modules and remove unused modules
   2) `go clean -modcache` to remove entire module download cache

### Wishlist
* Write tests
* API specification with OpenAPI or [swaggo/swag](https://github.com/swaggo/swag)

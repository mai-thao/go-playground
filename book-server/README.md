# book-server
This is a simple book server that exposes a RESTful API to interact with its book resource. Its data are stored in memory
with a slice. It uses the [Gin](https://gin-gonic.com/en/docs/) web framework.

### How to Run the App
1) Execute the command: `go run ./cmd/memory`
2) Terminate the app with: Ctrl + C 

### How to Get Dependencies
1) Execute the command: `go get .` (resolve and download dependencies for code in the current directory)

### Wishlist
* Write tests
* Persist to a database instead of a slice?
* API specification with OpenAPI or [swaggo/swag](https://github.com/swaggo/swag)

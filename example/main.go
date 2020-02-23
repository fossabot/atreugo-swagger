package main

import (
	atreugoswagger "github.com/Nerzal/atreugo-swagger"
	_ "github.com/Nerzal/atreugo-swagger/example/docs" // docs is generated by Swag CLI, you have to import it.
	"github.com/savsgio/atreugo/v10"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	config := &atreugo.Config{
		Addr: "0.0.0.0:1337",
	}

	a := atreugo.New(config)

	// This will server all swagger files under the /docs/* path.
	a.GET("/docs/*doc", atreugoswagger.AtreugoWrapHandler())

	// Start the server
	a.ListenAndServe()
}

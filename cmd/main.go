package main

import (
	"github.com/backend-magang/eniqilo-store/config"
	"github.com/backend-magang/eniqilo-store/driver"
	"github.com/backend-magang/eniqilo-store/internal/handler/api"
	"github.com/backend-magang/eniqilo-store/internal/repository/postgres"
	"github.com/backend-magang/eniqilo-store/internal/usecase"
	"github.com/backend-magang/eniqilo-store/middleware"
	"github.com/backend-magang/eniqilo-store/router"
	"github.com/backend-magang/eniqilo-store/utils/pkg"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// @title           Swagger Backend Magang - Project 2
// @version         1.0
// @description     This is a documentation of Backend Magang - Project 2
/* MAIN FUNCTION:
This is the main function, which serves as the entry point of the program.
It's where the execution starts when the program is run. */
func main() {
	/* ECHO SERVER INITIALIZATION: This line initializes a new instance of an Echo server.
	Echo is a web framework for Go, used for building web applications and APIs.
	It simplifies the process of handling HTTP requests and responses.*/
	server := echo.New()

	// Load Config
	/* CONFIGURATION LOADING:
	This line loads the configuration settings for the application using the Load() function from the config package.
	Configuration settings typically include things like database connection details, server port,
	and other environment-specific settings.*/
	cfg := config.Load()
	/* LOGGER INITIALIZATION:
	This line initializes a new instance of a logger using Logrus, a popular logging library for Go.
	Logging is essential for recording events and debugging information during the execution of the application.*/
	logger := logrus.New()

	/* DATABASE INITIALIZATION:
	This line initializes the PostgreSQL database client by calling the InitPostgres() function
	from the driver package, passing the application configuration (cfg) as an argument.
	This function establishes a connection to the PostgreSQL database using the settings from the configuration.
	*/
	dbClient := driver.InitPostgres(cfg)

	// Set Transaction
	/* TRANSACTION SERVICE INITIALIZATION:
	These lines initialize a new SQL transaction service using the initialized database client.
	The transaction service is used to manage database transactions.
	The cfg.SqlTrx field is updated in the application configuration with this transaction service. */
	sqlTrx := pkg.NewSqlWithTransactionService(dbClient)
	cfg.SqlTrx = sqlTrx

	/* REPOSITORY INITIALIZATION:
	This line initializes a new instance of a repository using PostgreSQL.
	It calls the NewRepository() function from the postgres package,
	passing the application configuration, database client, and logger as arguments.
	The repository is responsible for interacting with the database. */
	postgresRepository := postgres.NewRepository(cfg, dbClient, logger)
	/* USECASE INITIALIZATION:
	This line initializes a new instance of a use case.
	A use case contains the application's business logic.
	It calls the NewUsecase() function from the usecase package,
	passing the application configuration, logger, and repository as arguments.*/
	usecase := usecase.NewUsecase(cfg, logger, postgresRepository)
	/* HANDLER INITIALIZATION:
	This line initializes a new instance of a handler for processing HTTP requests.
	The handler is responsible for receiving HTTP requests,
	invoking the appropriate use case methods, and returning HTTP responses.
	It calls the NewHandler() function from the api package, passing the logger and use case as arguments.*/
	handler := api.NewHandler(logger, usecase)

	/* ROUTER INITIALIZATION:
	This line initializes the router for the Echo server.
	The router maps HTTP requests to the appropriate handlers.
	It calls the InitRouter() function from the router package,
	passing the Echo server instance and handler as arguments.*/
	router.InitRouter(server, handler)
	/* MIDDLEWARE INITIALIZATION:
	This line initializes middleware for the Echo server.
	Middleware functions intercept HTTP requests and responses,
	allowing you to perform operations like logging, authentication,
	or request modification before passing them to the handler.
	It calls the InitMiddlewares() function from the middleware package,
	passing the Echo server instance as an argument. */
	middleware.InitMiddlewares(server)

	// host := fmt.Sprintf("%s:8080", cfg.AppHost)

	// SERVER START: This line starts the Echo server, causing it to listen for incoming HTTP requests on port 8080.
	server.Start(":8080")
}

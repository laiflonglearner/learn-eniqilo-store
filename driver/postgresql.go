package driver

import (
	// cf is the alias referring to the config package, which contains the application configuration.
	cf "github.com/backend-magang/eniqilo-store/config"
	"github.com/jmoiron/sqlx" //  provides a set of extensions on top of Go's built-in database/sql package. It's used for working with SQL databases.

	_ "github.com/lib/pq"
	// blank import (_) is used to import the pq package anonymously.
	// The pq package is a PostgreSQL driver for Go's database/sql package.

	"log" // for logging messages.
)

func InitPostgres(config cf.Config) *sqlx.DB {
	/* (-) This line calls the GetDSN() method on the config struct passed as an argument.
	   (-) The GetDSN() method generates a Database Source Name (DSN) string based on the configuration settings,
	       which contains information such as database host, port, name, username, and password.*/
	psqlInfo := config.GetDSN()

	/* (-) `sqlx.Connect()` is called to establish a connection to the PostgreSQL database.
	       The first argument specifies the driver name, which in this case is "postgres" indicating the PostgreSQL driver.
	       The second argument is the DSN string generated earlier.
	   (-) If there's an error during the connection attempt, it's captured in the err variable.
	       If an error occurs, the function logs the error using log.Fatal() and returns nil. */
	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	log.Println("[Database] initialized...")

	/* (-) After connecting to the database, the function calls db.Ping() to verify that the connection is active
		   and the database is accessible.
	   (-) If Ping() returns an error, it indicates that the connection failed or the database is not accessible.
	   In this case, the error is logged, and the function returns nil. */
	err = db.Ping()
	if err != nil {
		log.Println("[Database] failed to connect to database: ", err)
		return nil
	}

	log.Println("[Database] successfully connected")

	/* Finally, if all steps are successful, the function returns the connected sqlx.DB object,
	   which represents the database connection.
	   This allows other parts of the application to use this connection for executing database queries. */
	return db
}

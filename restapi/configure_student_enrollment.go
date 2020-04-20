// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"

	"github.com/Cinojose/StudentEnrollmentApi/restapi/operations"
	"github.com/Cinojose/StudentEnrollmentApi/restapi/operations/student"
)

//go:generate swagger generate server --target ../../StudentEnrollmentApi --name StudentEnrollment --spec ../swagger.yaml

func configureFlags(api *operations.StudentEnrollmentAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.StudentEnrollmentAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()
	api.XMLConsumer = runtime.XMLConsumer()

	api.JSONProducer = runtime.JSONProducer()
	api.XMLProducer = runtime.XMLProducer()

	mydb := dbInit(readEnv())

	studenthandlerobj := student.StudentOrm{mydb}

	api.StudentAddStudentHandler = student.AddStudentHandlerFunc(studenthandlerobj.AddStudentHandler)
	api.StudentUpdateStudentHandler = student.UpdateStudentHandlerFunc(studenthandlerobj.UpdateStudentHandler)
	api.StudentDeleteStudentHandler = student.DeleteStudentHandlerFunc(studenthandlerobj.DeleteStudentHandler)
	api.StudentFindStudentsByStatusHandler = student.FindStudentsByStatusHandlerFunc(studenthandlerobj.FindStudentsHandler)

	if api.StudentAddStudentHandler == nil {
		api.StudentAddStudentHandler = student.AddStudentHandlerFunc(func(params student.AddStudentParams) middleware.Responder {
			return middleware.NotImplemented("operation student.AddStudent has not yet been implemented")
		})
	}
	if api.StudentDeleteStudentHandler == nil {
		api.StudentDeleteStudentHandler = student.DeleteStudentHandlerFunc(func(params student.DeleteStudentParams) middleware.Responder {
			return middleware.NotImplemented("operation student.DeleteStudent has not yet been implemented")
		})
	}
	if api.StudentFindStudentsByStatusHandler == nil {
		api.StudentFindStudentsByStatusHandler = student.FindStudentsByStatusHandlerFunc(func(params student.FindStudentsByStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation student.FindStudentsByStatus has not yet been implemented")
		})
	}
	if api.StudentUpdateStudentHandler == nil {
		api.StudentUpdateStudentHandler = student.UpdateStudentHandlerFunc(func(params student.UpdateStudentParams) middleware.Responder {
			return middleware.NotImplemented("operation student.UpdateStudent has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}

// Function to read the env variables from file and return a string of variables
func readEnv() (string, string, string, string) {

	// Reading the env variables
	viper.BindEnv("DB_USER")
	viper.BindEnv("DB_PASSWORD")
	viper.BindEnv("DB_HOST")
	viper.BindEnv("DB_NAME")

	//log.Printf("env variables user %s, pass %s, host %s, db %s", user, password, host, db)
	return viper.GetString("DB_USER"), viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_HOST"), viper.GetString("DB_NAME")
}

// Function to intialise the DB connection
func dbInit(user string, password string, host string, dbname string) *gorm.DB {

	var db *gorm.DB
	var err error

	connectionString := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local&tls=skip-verify", user, password, host, dbname)

	db, err = gorm.Open("mysql", connectionString)

	if err != nil {
		log.Fatal("Error connecting to database", err)
	}
	log.Print("DB connection established")
	return db

}

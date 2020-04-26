![Build Status](https://codebuild.ap-southeast-1.amazonaws.com/badges?uuid=eyJlbmNyeXB0ZWREYXRhIjoiUWlHeDJVZlBkSUpGZzYzMDNsTVN0U01UWVc4a1hZZFBsVTBvdlhRRDJHVzJLa1AyYjVRVXJWL3oyN2hjdmJPY3BaSGV4WlptaVJjYjd4YUlYdUhrQVdnPSIsIml2UGFyYW1ldGVyU3BlYyI6IjBBUkNFVlFCOW51cll6djUiLCJtYXRlcmlhbFNldFNlcmlhbCI6MX0%3D&branch=master)

## Student Enrollment Api Server

A student enrollment api server using [Go Swagger](https://goswagger.io/)

# Prerequisites

* Running mysql database with a table created based on the sql/student.sql
* Docker engine


# How to run ?

```bash
 # Download source code
  git clone https://github.com/Cinojose/StudentEnrollmentApi.git
 #setup the env variables
  export DB_USER=<db-user>
  export DB_PASSWORD=<db-password>
  export DB_NAME=<database-name>
  export DB_HOST=<database-host-address>
 # build the docker image
  docker build -t studentenrollment .
 # Run the docker image
  docker run -it -p 8001:8001 --env DB_USER --env DB_PASSWORD --env DB_NAME --env DB_HOST studentenrollment studentenrollment 
 # The app can  be accessed via following URL upon successful run
 # http://localhost:8001/internal/apidocs
```

# How to run tests ?

Start the APP server using the above instructions

```bash
   # set up the env variables
   export API_PORT=<aplication port>
   export API_HOST=<application-host>

   # Build and run docker
   docker build -t unitest -f DockerfileTest . && docker rm unitest && docker run -it --env API_HOST --env API_PORT--name unitest unitest
   
```

# Generating server stub using swagger
```bash
   # Generate a server stub using the swagger yml
   cd github.com/StudentEntollmentApi
   alias swagger="docker run --rm -it -e GOPATH=$HOME/go:/go -v $HOME:$HOME -w $(pwd) quay.io/goswagger/swagger"
   # Update the swagger.yml with your new changes
   swagger generate server -f swagger.yaml -A StudentEnrollmentApi
   # Run the above docker commands to build the code
```

# App Deployment

![CICD diagram](/studentenrollment.png)

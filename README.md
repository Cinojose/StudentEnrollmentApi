## Student Enrollment Api Server

# How to run ?

```bash
  docker build -t studentenrollment .
  docker run -it -p 8001:8001 studentenrollment
```

# Swagger Setup
```bash
   # Generate a server stub using the swagger yml
   cd github.com/StudentEntollmentApi
   alias swagger="docker run --rm -it -e GOPATH=$HOME/go:/go -v $HOME:$HOME -w $(pwd) quay.io/goswagger/swagger"
   # Update the swagger.yml with your new changes
   swagger generate server -f swagger.yaml -A StudentEnrollmentApi
   # Run the above docker commands to build the code
```
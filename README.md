## Student Enrollment Api Server

A student enrollment api server using (Go Swagger)[https://goswagger.io/]

# How to run ?

```bash
  docker build -t studentenrollment .
  docker run -it -p 8001:8001 studentenrollment
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

# How app is getting deployed
![CICD diagram](/studentenrollment.png)

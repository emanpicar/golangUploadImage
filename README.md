# golangUploadImage
Is a simple file upload project using Go Language (go1.10). 

# Requirements
- Installed Golang
    - "go get github.com/lib/pq"
- Installed Docker
    - "docker image pull postgres:10.5"

# Environment Variables
- TESTAPP_DBHOST - host ip of postgres is required
- TESTAPP_TOKEN - this token is used to authenticate the post request when uploading a file

# How to run the app
- We need to run the db first, run it using this command:
    - docker run --name postgres-go -e POSTGRES_PASSWORD=password -d -p 5432:5432 postgres:10.5
    - note that we are not attaching/using volumes here but if you need to store the data, you can just add something like this "-v ./postgres-data:/var/lib/postgresql/data"
- And for the server, it's simple as:
    - go run main.go
    - Open localhost:8080 in browser
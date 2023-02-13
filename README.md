# ChaosCampGo Project

This is a project for the ChaosCamp Golang Course.\
There is a client side and a server side to it. <br/>

In order to run the server, you need to enter the server folder and to create an "app.env" file with the following structure:\
`HTTP_ADDRESS = "..."`\
`DB_ADDRESS = "..."`\
`DB_NAME = "..."`\
`TOKEN_SECRET = "..."`\
`SENDGRID_API_KEY = "..."` <br/>

HTTP_ADDRESS - the address on which the server will be running, eg. "0.0.0.0:8888"\
DB_ADDRESS - the address the database is running on, eg. "mongodb://localhost:27017"\
DB_NAME - the name of the database\
TOKEN_SECRET - secret for generating JWTs, any string will work\
SENDGRID_API_KEY - an API key generated from the email sending third-party app Sendgrid <br/>

After configuring this, you can run the server with `go run main.go` but you can specify if you want a logger with the `-log` flag as follows: `go run main.go -log="."`. This specifies that the path where the logger will be generated will be the current directory and the name of the file will be "logs.txt". <br/>

If you would want to run the unit tests, you need to open the corresponding folder and run the `go test -cover` command. If you want to extract the result to a file run: `go test -coverprofile=coverage.out` and then `go tool cover -html=coverage.out`. This will open an .html file listing the test coverage of the functionality. <br/>

In order to run the database tests, in the server folder create folder "test" and a file "app.env" in it. Then enter the database folder and run `go test -cover`. This "app.env" file can only have the following fields:\
`HTTP_ADDRESS = "0.0.0.0:8888"`\
`DB_ADDRESS = "mongodb://localhost:27017"`\
`DB_NAME = "mockgocc"`\
`SENDGRID_API_KEY = "..."` <br/>

The tests run will only affect the database specified by this configuration. <br/>

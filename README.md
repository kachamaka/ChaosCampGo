# ChaosCampGo Project

This is a project for the ChaosCamp Golang Course.\
The project represents an online scheduler for all kinds of people that have the same schedule for a long period of time.\
To use it you need to register and then enter what kind of events you have on which weekday. There is also a possibility to add a reminder for an event some time before it starts and receive an email saying how much time remains. 

##

There is a client side and a server side to it. <br/>

In order to run the server, you need to enter the server folder and to create an "app.env" file with the following structure:\
`HTTP_ADDRESS = "..."`\
`DB_ADDRESS = "..."`\
`DB_NAME = "..."`\
`TOKEN_SECRET = "..."`\
`SENDGRID_API_KEY = "..."` <br/>

HTTP_ADDRESS - the address on which the server will be running, eg. `0.0.0.0:8888`\
DB_ADDRESS - the address the database is running on, eg. `mongodb://localhost:27017`\
DB_NAME - the name of the database, eg. `testDB`\
TOKEN_SECRET - secret for generating JWTs, any string will work, eg. `2sdGzJ6rKkyZjPU04SWEqEK4Uwho8NDp`\
SENDGRID_API_KEY - an API key generated from the email sending third-party app Sendgrid <br/>

After configuring this, you can run the server with `go run main.go` but you can specify if you want a logger with the `-log` flag as follows: `go run main.go -log="."`. This specifies that the path where the logger will be generated will be the current directory and the name of the file will be "logs.txt". You can also the "chaosgo.exe" file. <br/>

If you would want to run the unit tests, you need to open the corresponding folder and run the `go test -cover` command.\
If you want to extract the result to a file run:\
`go test -coverprofile=coverage.out` and then `go tool cover -html=coverage.out`.\
This will open an .html file listing the test coverage of the functionality. <br/>

In order to run the database tests, create a folder "test" and a file "app.env" in the server folder. Then, enter the database folder and run `go test -cover`. This "app.env" file needs to have at least the following fields:\
`DB_ADDRESS = "mongodb://localhost:27017"`\
`DB_NAME = "mockgocc"`\
`SENDGRID_API_KEY = "..."` <br/>

The tests run will only affect the database specified by the test configuration. <br/>

##

In order to run the client, you need to first have AngularCLI installed. Then enter the app folder and run `npm install` to download all packages and dependencies.\
Finally you need to run `ng s --o` and it will open a new tab in your browser where the application will be served.

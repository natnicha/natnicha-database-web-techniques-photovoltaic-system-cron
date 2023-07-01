# Photovoltaic System Cron

## Prerequisite - Need to know
There are 4 services in this system including
1. Photovoltaic System Services
2. Photovoltaic System App
3. Photovoltaic System Cron
4. Photovoltaic System Batch

Work flows look like this
1. `App` call to `Services` 
2. `Services` call to `Batch`
3. `Cron` call to `Services`

You should setup/run the services in this sequence:
1. `Batch`
2. `Services`
3. `Cron` - no dependencies can setup either `App` or `Cron` first
3. `App` - no dependencies can setup either `App` or `Cron` first


## Installing the project on your machine

1 - Install Go by following this instruction
~~~
https://go.dev/doc/install
~~~

2 - Install dependencies using this command
~~~
go build
~~~

3 - Tidy your project by this command. It will remove dependencies and tidy your project
~~~
go mod tidy
~~~
or tidy using Makefile on Windows
~~~
MinGW32-make tidy 
~~~

## Setting up the database destination
1 - Set a database destination in a .env file 


## Running the project on your machine
1 - Run this command
~~~
go run main.go
~~~
or run using Makefile on OSX
~~~
make run 
~~~
or run using Makefile on Windows
~~~
MinGW32-make run 
~~~
2 - Enjoy the service!

| Please note that this project is a job scheduler. It will automatically calls a request when reached the time specified in Cron string |
| --- |

| Remark: This service works together with Photovoltaic System Services. Please run it first, then followed by this project! |
| --- |

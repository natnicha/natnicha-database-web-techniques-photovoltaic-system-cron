# Photovoltaic System Cron

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

> **Please note that this project is a job scheduler. It will automatically calls a request when reached the time specified in Cron string**

| Remark: This service works together with Photovoltaic System Services. Please run it first, then followed by this project! |
| --- |

# Simple Rest API with Golang and MySQL


( )



## Table of Contents
* [Built With](#built-with)
* [Getting Started](#getting-started)
* [Not Work! What to do?](#not-work-what-to-do)
* [Additional information](#additional-information)



## Built With ##
* [GO](https://go.dev/) - . | version 1.21.3
* [MySQL](https://www.mysql.com/) - An open-source relational database management system. | version 0.0.0

[[top]](#table-of-contents)



## Getting started

### Create database
Create SQL Database on your local.

### Setup database enviroment
Edit `setup.go` file in **`models/setup.go`**
```
	USER := "(yourhostnameusername)"
	PASS := "(yourhostnamepassword)"
	HOST := "(yourhostname)"
	PORT := "(yourhostport)"
	DBNAME := "(yourdatabasename)"
```

### Run server
Run command **`go run main.go`**.

[[top]](#table-of-contents)



## Not Work! What to do?

[[top]](#table-of-contents)



## Additional information

### Meng-ignore file untuk di local saja tanpa harus di push ke Git
Akses folder **`.git/info`**, ubah file **`exclude`**, tambahkan file atau folder yang ingin di ignore di local masing-masing, contoh **`notes/`**.

File ini berfungsi sama dengan **`.gitignore`** tapi hanya berlaku di local masing-masing dan tidak akan ter-push ke git.

[[top]](#table-of-contents)

---

# go-API-REST
### A rest api made it in golang that connects to a mongodb database

![alt text](https://miro.medium.com/fit/c/262/262/1*yh90bW8jL4f8pOTZTvbzqw.png)


#### Dependencies
  * https://github.com/gin-gonic/gin
  * go.mongodb.org/mongo-driver/mongo

#### Run the API
In order to run the API without problems, you must have a config file with the connection variables to the database in the path `config/config.json`. Then you can execute:

```
go run main.go
```

#### Compile the API to binary and run it
```
go build
./go-API-REST
```

```
Go Gin Crud
```

```
# init the mod file
 go mod init "packge name"
```
# get the go gin package 
# $ go get -u github.com/gin-gonic/gin

# mongo driver
# go get go.mongodb.org/mongo-driver/mongo

# run monogo on docker
# docker run -it --rm --name mongodb_container -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INIT_ROOT_PASSWORD=admin -v mongodata:/data/db  -p 27017:27017 mongo

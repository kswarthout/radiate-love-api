# Radiate Love API
An API for Radiate Love Foundation's Fund Management web app. The API was implemented in Go and uses MongoDB for data storage.

# API Documentation
You can view the YAML document describing the API design [here](https://app.swaggerhub.com/apis/kari_bullard/Radiate-Love-Foundation/1.0.0)

# Dependencies
This project uses the following dependencies that can be retrieved using ```go get```:
- github.com/gorilla/mux
- go.mongodb.org/mongo-driver/mongo

# MongoDB
The API uses the on-premise MongoDB Community Server to store application data. 

### Running MongoDB Locally
To run the MongoDB service on your local machine, find the directory where MongoDB was installed (usually ```C:\Program Files\MongoDB\Server\<version>```) and open the bin directory. 

To use the DB locally, you need to start the MongoDB service and client from the command-line

#### Starting the Service

cd to ```C:\Program Files\MongoDB\Server\<version>\bin```

Service:  type ```mongod``` and hit enter

###### First time starting the MongoDB Service
When you start the db service using mongod, it will look for the directory ```C:\data\db```. 
- If running just one instance, you need to create the directory ```C:\data\db``` for the command to execute successfully
- If you have multiple DB instances, you will need to specify which directory to use by executing the ```mongod``` command described above using flags as descried [here](https://stackoverflow.com/questions/15124610/multiple-instances-of-mongo-db-on-same-server).

###### Automatically Starting the Service
If you selected to run MongoDB as a service during installation, the MongoDB will automatically start when your computer starts up. This means you won't have to explicitly start it each time to work with the service, and instead you can just start the client as described below.

#### Starting the Client
Once the service is running, in a new terminal, cd to ```C:\Program Files\MongoDB\Server\<version>\bin```

Client: type  ```mongo``` and hit enter

Once running, you can view all local dbs by executing ```show dbs``` from the client console.

### Creating a new DB
In the client console, execute ```use <db_name>```, for example: ```use radiatelove```

If successful, you'll see the message ```switched to db <db_name>```

For the newly created databade to appear in the database list shown when ```show dbs``` is executed, you need to add some data.

### Adding Data
From the client console, with the appropriate DB selected, execute ```db.<collection_name>.insert(<data>)```
- <collection_name> is the name for the particular collection, like a table name in SQL
- <data> is the JSON data you'd like to insert for that collection

You can find more about designing MongoDB schemas [here](https://docs.mongodb.com/manual/core/data-model-design/). 

Once data has been added, you can view collections and documents by executing the following from the client console:
- All DB collections: ```show collections;```
- Specific collection documents: ```db.<collection_name>.find()```

### Adding mongod to environment
To execute ```mongod``` and ```mongo``` from any directory, add ```C:\Program Files\MongoDB\Server\<version>\bin``` to your PATH environment variable.

Once added, restart your terminal session, and you should be able to execute these commands without changing directories.





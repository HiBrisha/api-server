package monggoDb

import (
	"context"

	logger "module/pkg/utils/logs"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Server struct {
	ConnStr string
	Client  *mongo.Client
	Log     *logger.SysLog
}

func (sv *Server) Init() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(sv.ConnStr).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		sv.Log.Write("ERR", "Can't is connect to MongoDB: %v", err)
	}
	sv.Client = client

	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		sv.Log.Write("ERR", "Error ping MongoDB: %v", err)
	}

	sv.Log.Write("OK", "Successfully connected to MongoDB!")
}

func (sv *Server) Close() {
	if sv.Client == nil {
		sv.Log.Write("ERR", "MongoDB client is not initialized.")
	}

	if err := sv.Client.Disconnect(context.TODO()); err != nil {
		sv.Log.Write("ERR", "MongoDB client is error: %v", err)
	}
}

func (sv *Server) Insert(Db string, Col string, Doc interface{}) {
	if sv.Client == nil {
		sv.Log.Write("ERR", "MongoDB client is not initialized.")
	}
	coll := sv.Client.Database(Db).Collection(Col)

	// Insert the document
	_, err := coll.InsertOne(context.TODO(), Doc)
	if err != nil {
		sv.Log.Write("ERR", "Can't is insert the document: %v", err)
	}

	sv.Log.Write("OK", "Insert document successfull!")
}

func (sv *Server) Delete(Db string, Col string, Filter bson.M) {
	if sv.Client == nil {
		sv.Log.Write("ERR", "MongoDB client is not initialized.")
	}
	//Connect to collection
	coll := sv.Client.Database(Db).Collection(Col)

	// Delete document
	deleteResult, err := coll.DeleteMany(context.TODO(), Filter)
	if err != nil {
		sv.Log.Write("ERR", "Could not delete documents: %v", err)
	}

	sv.Log.Write("OK", "Deleted %v document(s) successfully!", deleteResult.DeletedCount)
}

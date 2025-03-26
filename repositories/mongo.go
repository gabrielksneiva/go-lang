package repositories

import (
	"context"
	"fmt"
	"log"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
var testCollection *mongo.Collection

func StartMongoContainer() string {
	// Configura o container do MongoDB
	req := testcontainers.ContainerRequest{
		Image:        "mongo:latest",
		ExposedPorts: []string{"27017/tcp"},
		WaitingFor:   wait.ForListeningPort("27017/tcp"),
	}
	mongoC, err := testcontainers.GenericContainer(context.Background(), testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		log.Fatalf("Erro ao iniciar o container do MongoDB: %v", err)
	}

	// Obtém a URL de conexão
	host, _ := mongoC.Host(context.Background())
	port, _ := mongoC.MappedPort(context.Background(), "27017")
	uri := fmt.Sprintf("mongodb://%s:%s", host, port)

	fmt.Println("MongoDB no container iniciado!")
	return uri
}

func ConnectToMongo(uri string) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("testdb").Collection("items")
	fmt.Println("Connected to MongoDB!")
}

func InsertItem(ctx *context.Context, user User) error {
	_, err := collection.InsertOne(*ctx, user)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("Item created:", user)
	return nil
}

func ReadItems(ctx *context.Context, filter bson.M) ([]User, error) {
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var items []User

	for cursor.Next(context.TODO()) {
		var item User
		err := cursor.Decode(&item)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func UpdateItem(id primitive.ObjectID, updatedData bson.M) {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": updatedData}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Item updated:", id)
}

func DeleteItem(id primitive.ObjectID) {
	filter := bson.M{"_id": id}

	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Item deleted:", id)
}

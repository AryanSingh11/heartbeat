package main

import(
	"fmt"
	"net"
	"log"
	//"context"

	//"go.mongodb.org/mongo-driver/mongo"
    //"go.mongodb.org/mongo-driver/mongo/options"
)

func handleErr(err error){
	if err != nil {
		log.Fatal(err)
	}
}

//var collection *mongo.Collection

func main(){
	//50051
	lis, err := net.Listen("tcp", "0.0.0.0:50051");
	handleErr(err)


	//client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	handleErr(err)
	
	//fmt.Println("Mongodb connected")
	fmt.Println(lis)	



}


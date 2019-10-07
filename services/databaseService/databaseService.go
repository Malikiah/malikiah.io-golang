package databaseservice

//clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")

//db =: client.Database("malikiah")

func Find(collection string, criteria string, criteriaValue string, findAll bool) {

	//client, err := mongo.Connect(context.TODO(), clientOptions)

	if criteria == "_id" {
		//db.Collection(collection).findOne({ "_id": new ObjectId(criteriaValue)}, func () {
		//	log.Println()
		//})
	} else if criteria == "" {

	} else if findAll == true {

	} else {

	}
}

func Insert(collection string,data) {
	//collection := dbName.Collection(collection)

	//insertResult, err =: collection.InsertOne(context.TODO(), data)

	//if err != nil {
	//	log.Fatal(err)
	//}

	//log.Println(insertResult.InsertedID)
}

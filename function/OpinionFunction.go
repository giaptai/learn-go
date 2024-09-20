package function

import (
	"context"
	// "encoding/json"
	"fmt"

	"firm.com/connectDB"
	"firm.com/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetOpinions() ([]models.Opinion, error) {
	var opinions []models.Opinion
	db := connectdb.Client.Database("db_comment").Collection("opinion")
	cur, err := db.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, fmt.Errorf("%v", err.Error())
	}
	for cur.Next(context.TODO()) {
		var opinion models.Opinion
		err := cur.Decode(&opinion) //giả mã bỏ vô opinion
		if err != nil {
			return nil, fmt.Errorf("%v", err.Error())
		}
		opinions = append(opinions, opinion)
	}
	if err := cur.Err(); err != nil {
		return nil, fmt.Errorf("%v", err.Error())
	}
	return opinions, nil
}

func GetFirmOpinions(firm_id string) ([]models.Opinion, error) {
	var intFirm int
	var err error
	var opinions []models.Opinion
	_, err = fmt.Sscan(firm_id, &intFirm)
	if err != nil {
		return nil, fmt.Errorf("%v", err.Error())
	}
	db := connectdb.Client.Database("db_comment").Collection("opinion")
	cur, err := db.Find(context.TODO(), bson.D{{"firm_id", &intFirm}})
	if err != nil {
		return nil, fmt.Errorf("%v", err.Error())
	}
	for cur.Next(context.TODO()) {
		var opinion models.Opinion
		err := cur.Decode(&opinion) //giả mã bỏ vô opinion
		if err != nil {
			return nil, fmt.Errorf("%v", err.Error())
		}
		opinions = append(opinions, opinion)
	}
	if err := cur.Err(); err != nil {
		return nil, fmt.Errorf("%v", err.Error())
	}
	return opinions, nil
}

func InsertOpinion(opinion *models.Opinion) (*models.Opinion, error) {
	*opinion = opinion.DefaultOpinion()
	db := connectdb.Client.Database("db_comment").Collection("opinion")
	result, err := db.InsertOne(context.TODO(), &opinion)
	if err != nil {
		return nil, err
	}
	// Prints the ID of the inserted document
	fmt.Printf("Document inserted with ID: %s\n", result.InsertedID)
	return opinion, nil
}

func GetSingleOpinion(id string) (models.Opinion, error) {
	var objID primitive.ObjectID
	var opinion models.Opinion

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return opinion, err
	}

	coll := connectdb.Client.Database("db_comment").Collection("opinion")

	result := coll.FindOne(context.TODO(), bson.M{"_id": objID})
	if err := result.Decode(&opinion); err != nil {
		return opinion, err
	}
	return opinion, nil
}

func UpdateOpinion(opinion *models.Opinion) (*models.Opinion, error) {
	coll := connectdb.Client.Database("db_comment").Collection("opinion")
	result, err := coll.UpdateOne(
		context.TODO(),
		bson.D{{Key: "_id", Value: opinion.ID}},
		bson.D{{Key: "$set", Value: opinion}},
	)
	if err != nil {
		fmt.Printf("Not found: %s\n", err.Error())
		return nil, err
	}
	fmt.Printf("The number of modified documents: %d\n", result.ModifiedCount)
	return opinion, nil
}

func DeleteOpinion(id string) error {
	coll := connectdb.Client.Database("db_comment").Collection("opinion")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	result, err := coll.DeleteOne(
		context.TODO(),
		bson.D{{"_id", objID}},
	)
	if err != nil {
		return err
	}
	fmt.Printf("The number of modified documents: %d\n", result.DeletedCount)
	return nil
}

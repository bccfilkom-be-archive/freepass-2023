package utils

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Contains(s []primitive.ObjectID, str primitive.ObjectID) bool {
	for _, v := range s {
		if v.Hex() == str.Hex() {
			return true
		}
	}

	return false
}

func ArrayMethod(m mgm.Model, id primitive.ObjectID, opr string, arrayName string, value any) error {
	filter := bson.D{{"_id", id}}
	update := bson.D{{opr, bson.D{{arrayName, value}}}}

	_, err := mgm.Coll(m).UpdateOne(mgm.Ctx(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

package models

import (
	"context"
	"jakeri-backend/utils"

	"github.com/chidiwilliams/flatbson"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

var usersCollection *mongo.Collection

func init() {
	usersCollection = utils.CollectionConnection("users")
}

type User struct {
	ID       *string `json:"_id"                     bson:"_id"                  binding:"required_without_all=Username Email"`
	Username *string `json:"username,omitempty"      bson:"username,omitempty"   binding:"requered_without=ID"`
	Email    *string `json:"email,omitempty"         bson:"email,omitempty"      binding:"required_without=ID"`
	Password *string `json:"password,omitempty"      bson:"-"                    binding:"-"`
	Audit    *Audit  `json:"audit,omitempty"         bson:"audit,omitempty"      binding:"-"`
	Decks    *Decks  `json:"decks,omitempty"         bson:"-"                    binding:"-"`
}

type Users []User

func (users *Users) Add() ([]interface{}, error) {
	var ctx context.Context
	data := make([]interface{}, 0)
	for _, user := range *users {
		user.Audit = &Audit{}
		user.Audit.Create()
		data = append(data, user)
	}
	val, err := usersCollection.InsertMany(ctx, data)
	if err != nil {
		return nil, err
	}
	return val.InsertedIDs, err
}

func (users *Users) Get(ids []string, embed map[string]interface{}) error {
	var ctx context.Context
	query := bson.M{}
	if len(ids) > 0 {
		query["_id"] = bson.M{"$in": ids}
	}
	cur, err := usersCollection.Find(ctx, query)
	cur.All(ctx, users)
	if err == nil {
		users.Load(embed)
	}
	return err
}

func (user *User) Get(id string, embed map[string]interface{}) error {
	var ctx context.Context
	query := bson.M{"_id": id}
	err := usersCollection.FindOne(ctx, query).Decode(&user)
	if err == nil {
		user.Load(embed)
	}
	return err
}

func (user *User) Update(id *string, modificatorId *string) (int, int, error) {
	var ctx context.Context
	query := bson.M{"_id": id}
	user.ID = id
	user.Audit = &Audit{}
	user.Audit.Modify()
	obj, _ := flatbson.Flatten(user)
	data := bson.M{"$set": obj}
	res, err := usersCollection.UpdateOne(ctx, query, data)
	return int(res.ModifiedCount), int(res.MatchedCount), err
}

func (users *Users) Delete(ids []string) (int, error) {
	var ctx context.Context
	query := bson.M{}
	if len(ids) > 0 {
		query["_id"] = bson.M{"$in": ids}
	}
	res, err := usersCollection.DeleteMany(ctx, query)
	return int(res.DeletedCount), err
}

func (user *User) Delete(id string) (int, error) {
	var ctx context.Context
	query := bson.M{"_id": id}
	res, err := usersCollection.DeleteOne(ctx, query)
	return int(res.DeletedCount), err
}

func (users *Users) Load(embed map[string]interface{}) {
	for i := 0; i < len(*users); i++ {
		(*users)[i].Load(embed)
	}
}

func (user *User) Load(embed map[string]interface{}) {
	if value, ok := embed["decks"]; ok {
		tmp := value.(map[string]interface{})
		user.Decks = &Decks{}
		user.Decks.Get(nil, user.ID, tmp)
	}
}

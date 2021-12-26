package models

import (
	"context"
	"jakeri-backend/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

var reviewsCollection *mongo.Collection

func init() {
	reviewsCollection = utils.CollectionConnection("reviews")
}

type Review struct {
	ID      *primitive.ObjectID `json:"_id"                     bson:"_id,omitempty"            binding:"required_without_all=Label Content"`
	Label   *string             `json:"label,omitempty"         bson:"label,omitempty"          binding:"required_without=ID"`
	Content *string             `json:"content,omitempty"       bson:"content,omitempty"        binding:"required_without=ID"`
	User    *User               `json:"user,omitempty"          bson:"-"                        binding:"-"`
	Audit   *Audit              `json:"audit,omitempty"         bson:"audit,omitempty"          binding:"-"`
}

type Reviews []Review

func (reviews *Reviews) Add(creatorId *string) ([]interface{}, error) {
	var ctx context.Context
	data := make([]interface{}, 0)
	for _, review := range *reviews {
		review.Audit = &Audit{}
		review.Audit.Create(creatorId)
		data = append(data, review)
	}
	val, err := reviewsCollection.InsertMany(ctx, data)
	if err != nil {
		return nil, err
	}
	return val.InsertedIDs, err
}

func (reviews *Reviews) Get(ids []primitive.ObjectID, embed map[string]interface{}) error {
	var ctx context.Context
	query := bson.M{}
	if len(ids) > 0 {
		query["_id"] = bson.M{"$in": ids}
	}
	cur, err := reviewsCollection.Find(ctx, query)
	cur.All(ctx, reviews)
	if err == nil {
		reviews.Load(embed)
	}
	return err
}

func (review *Review) Get(id *primitive.ObjectID, embed map[string]interface{}) error {
	var ctx context.Context
	query := bson.M{"_id": id}
	err := reviewsCollection.FindOne(ctx, query).Decode(&review)
	if err == nil {
		review.Load(embed)
	}
	return err
}

func (reviews *Reviews) Load(embed map[string]interface{}) {
	for i := 0; i < len(*reviews); i++ {
		(*reviews)[i].Load(embed)
	}
}

func (review *Review) Load(embed map[string]interface{}) {
	if _, ok := embed["user"]; ok {
		review.User = &User{}
		review.User.Get(*review.Audit.ModifiedBy, nil)
	}
}

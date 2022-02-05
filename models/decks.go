package models

import (
	"context"
	"jakeri-backend/utils"

	"github.com/chidiwilliams/flatbson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

var decksCollection *mongo.Collection

func init() {
	decksCollection = utils.CollectionConnection("decks")
}

type Deck struct {
	ID    *primitive.ObjectID `json:"_id"                     bson:"_id,omitempty"            binding:"required_without_all=User"`
	User  *User               `json:"user,omitempty"          bson:"user,omitempty"           binding:"required_without=ID"`
	Card  *Card               `json:"card,omitempty"          bson:"card,omitempty"           binding:"-"`
	Audit *Audit              `json:"audit,omitempty"         bson:"audit,omitempty"          binding:"-"`
}

type Decks []Deck

func (decks *Decks) Add(userId *string) ([]interface{}, error) {
	var ctx context.Context
	data := make([]interface{}, 0)
	for _, deck := range *decks {
		deck.Audit = &Audit{}
		deck.Audit.Create()
		data = append(data, deck)
	}
	val, err := decksCollection.InsertMany(ctx, data)
	if err != nil {
		return nil, err
	}
	return val.InsertedIDs, err
}

func (decks *Decks) Get(ids []primitive.ObjectID, userId *string, embed map[string]interface{}) error {
	var ctx context.Context
	query := bson.M{"user._id": userId}
	if len(ids) > 0 {
		query["_id"] = bson.M{"$in": ids}
	}
	cur, err := decksCollection.Find(ctx, query)
	cur.All(ctx, decks)
	if err == nil {
		decks.Load(embed)
	}
	return err
}

func (deck *Deck) Get(id *primitive.ObjectID, userId *string) error {
	var ctx context.Context
	query := bson.M{"_id": id, "user._id": userId}
	err := decksCollection.FindOne(ctx, query).Decode(&deck)
	return err
}

func (deck *Deck) Update(id *primitive.ObjectID, userId *string) (int, int, error) {
	var ctx context.Context
	query := bson.M{"_id": id, "user._id": userId}
	deck.ID = id
	deck.Audit = &Audit{}
	deck.Audit.Modify()
	obj, _ := flatbson.Flatten(deck)
	data := bson.M{"$set": obj}
	res, err := decksCollection.UpdateOne(ctx, query, data)
	return int(res.ModifiedCount), int(res.MatchedCount), err
}

func (decks *Decks) Delete(ids []primitive.ObjectID, userId *string) (int, error) {
	var ctx context.Context
	query := bson.M{"user._id": userId}
	if len(ids) > 0 {
		query["_id"] = bson.M{"$in": ids}
	}
	res, err := decksCollection.DeleteMany(ctx, query)
	return int(res.DeletedCount), err
}

func (deck *Deck) Delete(id *primitive.ObjectID, userId *string) (int, error) {
	var ctx context.Context
	query := bson.M{"_id": id, "user._id": userId}
	res, err := decksCollection.DeleteOne(ctx, query)
	return int(res.DeletedCount), err
}

func (decks *Decks) Load(embed map[string]interface{}) {
	for i := 0; i < len(*decks); i++ {
		(*decks)[i].Load(embed)
	}
}

func (deck *Deck) Load(embed map[string]interface{}) {
}

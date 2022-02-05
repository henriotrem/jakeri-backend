package models

import (
	"context"
	"jakeri-backend/utils"

	"github.com/chidiwilliams/flatbson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

var cardsCollection *mongo.Collection

func init() {
	cardsCollection = utils.CollectionConnection("cards")
}

type Card struct {
	ID          *primitive.ObjectID `json:"_id"                     bson:"_id,omitempty"            binding:"required_without_all=Label Description"`
	Label       *string             `json:"label,omitempty"         bson:"label,omitempty"          binding:"required_without=ID"`
	Description *string             `json:"description,omitempty"   bson:"description,omitempty"    binding:"required_without=ID"`
	Content     *string             `json:"content,omitempty"       bson:"content,omitempty"        binding:"required_without=ID"`
	Audit       *Audit              `json:"audit,omitempty"         bson:"audit,omitempty"          binding:"-"`
}

type Cards []Card

func (cards *Cards) Add(userId *string) ([]interface{}, error) {
	var ctx context.Context
	data := make([]interface{}, 0)
	for _, card := range *cards {
		card.Audit = &Audit{}
		card.Audit.Create(userId)
		data = append(data, card)
	}
	val, err := cardsCollection.InsertMany(ctx, data)
	if err != nil {
		return nil, err
	}
	return val.InsertedIDs, err
}

func (cards *Cards) Get(ids []primitive.ObjectID, embed map[string]interface{}) error {
	var ctx context.Context
	query := bson.M{}
	if len(ids) > 0 {
		query["_id"] = bson.M{"$in": ids}
	}
	cur, err := cardsCollection.Find(ctx, query)
	cur.All(ctx, cards)
	if err == nil {
		cards.Load(embed)
	}
	return err
}

func (card *Card) Get(id *primitive.ObjectID) error {
	var ctx context.Context
	query := bson.M{"_id": id}
	err := cardsCollection.FindOne(ctx, query).Decode(&card)
	return err
}

func (card *Card) Update(id *primitive.ObjectID, userId *string) (int, int, error) {
	var ctx context.Context
	query := bson.M{"_id": id, "audit.createdBy": userId}
	card.ID = id
	card.Audit = &Audit{}
	card.Audit.Modify(userId)
	obj, _ := flatbson.Flatten(card)
	data := bson.M{"$set": obj}
	res, err := cardsCollection.UpdateOne(ctx, query, data)
	return int(res.ModifiedCount), int(res.MatchedCount), err
}

func (cards *Cards) Delete(ids []primitive.ObjectID, userId *string) (int, error) {
	var ctx context.Context
	query := bson.M{"audit.createdBy": userId}
	if len(ids) > 0 {
		query["_id"] = bson.M{"$in": ids}
	}
	res, err := cardsCollection.DeleteMany(ctx, query)
	return int(res.DeletedCount), err
}

func (card *Card) Delete(id *primitive.ObjectID, userId *string) (int, error) {
	var ctx context.Context
	query := bson.M{"_id": id, "audit.createdBy": userId}
	res, err := cardsCollection.DeleteOne(ctx, query)
	return int(res.DeletedCount), err
}

func (cards *Cards) Load(embed map[string]interface{}) {
	for i := 0; i < len(*cards); i++ {
		(*cards)[i].Load(embed)
	}
}

func (card *Card) Load(embed map[string]interface{}) {
}

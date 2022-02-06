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
	ID      *primitive.ObjectID `json:"_id"                     bson:"_id,omitempty"            binding:"required_without_all=Label Description"`
	Content *string             `json:"content,omitempty"       bson:"content,omitempty"        binding:"required_without=ID"`
	Profile *Profile            `json:"profile,omitempty"       bson:"profile,omitempty"        binding:"-"`
	Audit   *Audit              `json:"audit,omitempty"         bson:"audit,omitempty"          binding:"-"`
}

type Cards []Card

func (cards *Cards) Add(profileId *string) ([]interface{}, error) {
	var ctx context.Context
	data := make([]interface{}, 0)
	for _, card := range *cards {
		card.Profile = &Profile{ID: profileId}
		card.Audit = &Audit{}
		card.Audit.Create()
		data = append(data, card)
	}
	val, err := cardsCollection.InsertMany(ctx, data)
	if err != nil {
		return nil, err
	}
	return val.InsertedIDs, err
}

func (cards *Cards) Get(ids []primitive.ObjectID, profileId *string, embed map[string]interface{}) error {
	var ctx context.Context
	query := bson.M{"profile._id": profileId}
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

func (card *Card) Get(id *primitive.ObjectID, embed map[string]interface{}) error {
	var ctx context.Context
	query := bson.M{"_id": id}
	err := cardsCollection.FindOne(ctx, query).Decode(&card)
	if err == nil {
		card.Load(embed)
	}
	return err
}

func (card *Card) Update(id *primitive.ObjectID, profileId *string) (int, int, error) {
	var ctx context.Context
	query := bson.M{"_id": id, "profile._id": profileId}
	card.ID = id
	card.Audit = &Audit{}
	card.Audit.Modify()
	obj, _ := flatbson.Flatten(card)
	data := bson.M{"$set": obj}
	res, err := cardsCollection.UpdateOne(ctx, query, data)
	return int(res.ModifiedCount), int(res.MatchedCount), err
}

func (cards *Cards) Delete(ids []primitive.ObjectID, profileId *string) (int, error) {
	var ctx context.Context
	query := bson.M{"profile._id": profileId}
	if len(ids) > 0 {
		query["_id"] = bson.M{"$in": ids}
	}
	res, err := cardsCollection.DeleteMany(ctx, query)
	return int(res.DeletedCount), err
}

func (card *Card) Delete(id *primitive.ObjectID, profileId *string) (int, error) {
	var ctx context.Context
	query := bson.M{"_id": id, "profile._id": profileId}
	res, err := cardsCollection.DeleteOne(ctx, query)
	return int(res.DeletedCount), err
}

func (cards *Cards) Load(embed map[string]interface{}) {
	for i := 0; i < len(*cards); i++ {
		(*cards)[i].Load(embed)
	}
}

func (card *Card) Load(embed map[string]interface{}) {
	if _, ok := embed["profile"]; ok {
		card.Profile.Get(*card.Profile.ID, nil)
	}
}

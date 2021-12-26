package models

import (
	"context"
	"jakeri-backend/utils"

	"github.com/chidiwilliams/flatbson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

var profilesCollection *mongo.Collection

func init() {
	profilesCollection = utils.CollectionConnection("profiles")
}

type Profile struct {
	ID    *primitive.ObjectID `json:"_id"                    bson:"_id,omitempty"        binding:"required_without_all=User JobTitle"`
	User  *User               `json:"user,omitempty"         bson:"user,omitempty"       binding:"required_without=ID"`
	Group *Group              `json:"group,omitempty"        bson:"group,omitempty"      binding:"-"`
	Audit *Audit              `json:"audit,omitempty"        bson:"audit,omitempty"      binding:"-"`
}

type Profiles []Profile

func (profiles *Profiles) Ids() (ids []primitive.ObjectID) {
	for _, profile := range *profiles {
		ids = append(ids, *profile.ID)
	}
	return
}

func (profiles *Profiles) Add(groupId *primitive.ObjectID, creatorId *string) ([]interface{}, error) {
	var ctx context.Context
	data := make([]interface{}, 0)
	for _, profile := range *profiles {
		profile.Group = &Group{ID: groupId}
		profile.Audit = &Audit{}
		profile.Audit.Create(creatorId)
		data = append(data, profile)
	}
	val, err := profilesCollection.InsertMany(ctx, data)
	if err != nil {
		return nil, err
	}
	return val.InsertedIDs, err
}

func (profiles *Profiles) Get(groupId *primitive.ObjectID, ids []primitive.ObjectID, userId *string, embed map[string]interface{}) error {
	var ctx context.Context
	query := bson.M{}
	if groupId != nil {
		query["group._id"] = groupId
	}
	if userId != nil {
		query["user._id"] = userId
	}
	if len(ids) > 0 {
		query["_id"] = bson.M{"$in": ids}
	}
	cur, err := profilesCollection.Find(ctx, query)
	cur.All(ctx, profiles)
	if err == nil {
		profiles.Load(embed)
	}
	return err
}

func (profile *Profile) Get(groupId *primitive.ObjectID, id *primitive.ObjectID, embed map[string]interface{}) error {
	var ctx context.Context
	query := bson.M{"group._id": groupId, "_id": id}
	err := profilesCollection.FindOne(ctx, query).Decode(&profile)
	if err == nil {
		profile.Load(embed)
	}
	return err
}

func (profile *Profile) Update(groupId *primitive.ObjectID, id *primitive.ObjectID, modificatorId *string) (int, int, error) {
	var ctx context.Context
	query := bson.M{"group._id": groupId, "_id": id}
	profile.ID = id
	profile.Audit = &Audit{}
	profile.Audit.Modify(modificatorId)
	obj, _ := flatbson.Flatten(profile)
	data := bson.M{"$set": obj}
	res, err := profilesCollection.UpdateOne(ctx, query, data)
	return int(res.ModifiedCount), int(res.MatchedCount), err
}

func (profiles *Profiles) Delete(groupId *primitive.ObjectID, ids []primitive.ObjectID) (int, error) {
	var ctx context.Context
	query := bson.M{"group._id": groupId}
	if len(ids) > 0 {
		query["_id"] = bson.M{"$in": ids}
	}
	res, err := profilesCollection.DeleteMany(ctx, query)
	return int(res.DeletedCount), err
}

func (profile *Profile) Delete(groupId *primitive.ObjectID, id *primitive.ObjectID) (int, error) {
	var ctx context.Context
	query := bson.M{"group._id": groupId, "_id": id}
	res, err := profilesCollection.DeleteOne(ctx, query)
	return int(res.DeletedCount), err
}

func (profiles *Profiles) Load(embed map[string]interface{}) {
	for i := 0; i < len(*profiles); i++ {
		(*profiles)[i].Load(embed)
	}
}

func (profile *Profile) Load(embed map[string]interface{}) {
	if value, ok := embed["group"]; ok {
		tmp := value.(map[string]interface{})
		profile.Group.Get(profile.Group.ID, tmp)
	}
	if _, ok := embed["user"]; ok {
		profile.User.Get(*profile.User.ID, nil)
	}
}

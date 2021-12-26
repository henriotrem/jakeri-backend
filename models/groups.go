package models

import (
	"context"
	"jakeri-backend/utils"

	"github.com/chidiwilliams/flatbson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

var groupsCollection *mongo.Collection

func init() {
	groupsCollection = utils.CollectionConnection("groups")
}

type Group struct {
	ID          *primitive.ObjectID `json:"_id"                     bson:"_id,omitempty"            binding:"required_without_all=Label Description Address"`
	Label       *string             `json:"label,omitempty"         bson:"label,omitempty"          binding:"required_without=ID"`
	Description *string             `json:"description,omitempty"   bson:"description,omitempty"    binding:"required_without=ID"`
	Content     *string             `json:"content,omitempty"       bson:"content,omitempty"        binding:"required_without=ID"`
	Audit       *Audit              `json:"audit,omitempty"         bson:"audit,omitempty"          binding:"-"`
	Profiles    *Profiles           `json:"profiles,omitempty"      bson:"-"                        binding:"-"`
}

type Groups []Group

func (groups *Groups) Ids() (ids []primitive.ObjectID) {
	for _, group := range *groups {
		ids = append(ids, *group.ID)
	}
	return
}

func (groups *Groups) Add(creatorId *string) ([]interface{}, error) {
	var ctx context.Context
	data := make([]interface{}, 0)
	for _, group := range *groups {
		group.Audit = &Audit{}
		group.Audit.Create(creatorId)
		data = append(data, group)
	}
	val, err := groupsCollection.InsertMany(ctx, data)
	if err != nil {
		return nil, err
	}
	return val.InsertedIDs, err
}

func (groups *Groups) Get(ids []primitive.ObjectID, embed map[string]interface{}) error {
	var ctx context.Context
	query := bson.M{}
	if len(ids) > 0 {
		query["_id"] = bson.M{"$in": ids}
	}
	cur, err := groupsCollection.Find(ctx, query)
	cur.All(ctx, groups)
	if err == nil {
		groups.Load(embed)
	}
	return err
}

func (group *Group) Get(id *primitive.ObjectID, embed map[string]interface{}) error {
	var ctx context.Context
	query := bson.M{"_id": id}
	err := groupsCollection.FindOne(ctx, query).Decode(&group)
	if err == nil {
		group.Load(embed)
	}
	return err
}

func (group *Group) Update(id *primitive.ObjectID, modificatorId *string) (int, int, error) {
	var ctx context.Context
	query := bson.M{"_id": id}
	group.ID = id
	group.Audit = &Audit{}
	group.Audit.Modify(modificatorId)
	obj, _ := flatbson.Flatten(group)
	data := bson.M{"$set": obj}
	res, err := groupsCollection.UpdateOne(ctx, query, data)
	return int(res.ModifiedCount), int(res.MatchedCount), err
}

func (group *Groups) Delete(ids []primitive.ObjectID) (int, error) {
	var ctx context.Context
	query := bson.M{}
	if len(ids) > 0 {
		query["_id"] = bson.M{"$in": ids}
	}
	res, err := groupsCollection.DeleteMany(ctx, query)
	return int(res.DeletedCount), err
}

func (group *Group) Delete(id *primitive.ObjectID) (int, error) {
	var ctx context.Context
	query := bson.M{"_id": id}
	res, err := groupsCollection.DeleteOne(ctx, query)
	return int(res.DeletedCount), err
}

func (groups *Groups) Load(embed map[string]interface{}) {
	for i := 0; i < len(*groups); i++ {
		(*groups)[i].Load(embed)
	}
}

func (group *Group) Load(embed map[string]interface{}) {
	if value, ok := embed["profiles"]; ok {
		tmp := value.(map[string]interface{})
		group.Profiles = &Profiles{}
		group.Profiles.Get(group.ID, nil, nil, tmp)
	}
}

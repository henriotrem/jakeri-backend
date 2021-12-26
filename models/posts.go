package models

import (
	"context"
	"jakeri-backend/utils"

	"github.com/chidiwilliams/flatbson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

var postsCollection *mongo.Collection

func init() {
	postsCollection = utils.CollectionConnection("posts")
}

type Post struct {
	ID          *primitive.ObjectID `json:"_id"                     bson:"_id,omitempty"            binding:"required_without_all=Label Description Address"`
	Label       *string             `json:"label,omitempty"         bson:"label,omitempty"          binding:"required_without=ID"`
	Description *string             `json:"description,omitempty"   bson:"description,omitempty"    binding:"required_without=ID"`
	Content     *string             `json:"content,omitempty"       bson:"content,omitempty"        binding:"required_without=ID"`
	Audit       *Audit              `json:"audit,omitempty"         bson:"audit,omitempty"          binding:"-"`
}

type Posts []Post

func (posts *Posts) Add(creatorId *string) ([]interface{}, error) {
	var ctx context.Context
	data := make([]interface{}, 0)
	for _, post := range *posts {
		post.Audit = &Audit{}
		post.Audit.Create(creatorId)
		data = append(data, post)
	}
	val, err := postsCollection.InsertMany(ctx, data)
	if err != nil {
		return nil, err
	}
	return val.InsertedIDs, err
}

func (posts *Posts) Get(ids []primitive.ObjectID, embed map[string]interface{}) error {
	var ctx context.Context
	query := bson.M{}
	if len(ids) > 0 {
		query["_id"] = bson.M{"$in": ids}
	}
	cur, err := postsCollection.Find(ctx, query)
	cur.All(ctx, posts)
	if err == nil {
		posts.Load(embed)
	}
	return err
}

func (post *Post) Get(id *primitive.ObjectID) error {
	var ctx context.Context
	query := bson.M{"_id": id}
	err := postsCollection.FindOne(ctx, query).Decode(&post)
	return err
}

func (post *Post) Update(id *primitive.ObjectID, modificatorId *string) (int, int, error) {
	var ctx context.Context
	query := bson.M{"_id": id}
	post.ID = id
	post.Audit = &Audit{}
	post.Audit.Modify(modificatorId)
	obj, _ := flatbson.Flatten(post)
	data := bson.M{"$set": obj}
	res, err := postsCollection.UpdateOne(ctx, query, data)
	return int(res.ModifiedCount), int(res.MatchedCount), err
}

func (posts *Posts) Delete(ids []primitive.ObjectID) (int, error) {
	var ctx context.Context
	query := bson.M{}
	if len(ids) > 0 {
		query["_id"] = bson.M{"$in": ids}
	}
	res, err := postsCollection.DeleteMany(ctx, query)
	return int(res.DeletedCount), err
}

func (post *Post) Delete(id *primitive.ObjectID) (int, error) {
	var ctx context.Context
	query := bson.M{"_id": id}
	res, err := postsCollection.DeleteOne(ctx, query)
	return int(res.DeletedCount), err
}

func (posts *Posts) Load(embed map[string]interface{}) {
	for i := 0; i < len(*posts); i++ {
		(*posts)[i].Load(embed)
	}
}

func (post *Post) Load(embed map[string]interface{}) {
}

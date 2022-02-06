package models

import (
	"context"

	"gopkg.in/mgo.v2/bson"
)

type Profile struct {
	ID       *string `json:"_id"                     bson:"_id"                  binding:"required_without_all=Profilename"`
	Username *string `json:"username,omitempty"      bson:"username,omitempty"   binding:"required_without=ID"`
	Audit    *Audit  `json:"audit,omitempty"         bson:"audit,omitempty"      binding:"-"`
	Cards    *Cards  `json:"cards,omitempty"         bson:"-"                    binding:"-"`
}

type Profiles []Profile

func (profiles *Profiles) Get(ids []string, embed map[string]interface{}) error {
	var ctx context.Context
	query := bson.M{}
	if len(ids) > 0 {
		query["_id"] = bson.M{"$in": ids}
	}
	cur, err := usersCollection.Find(ctx, query)
	cur.All(ctx, profiles)
	if err == nil {
		profiles.Load(embed)
	}
	return err
}

func (profile *Profile) Get(id string, embed map[string]interface{}) error {
	var ctx context.Context
	query := bson.M{"_id": id}
	err := usersCollection.FindOne(ctx, query).Decode(&profile)
	if err == nil {
		profile.Load(embed)
	}
	return err
}

func (profiles *Profiles) Load(embed map[string]interface{}) {
	for i := 0; i < len(*profiles); i++ {
		(*profiles)[i].Load(embed)
	}
}

func (profile *Profile) Load(embed map[string]interface{}) {
	if value, ok := embed["cards"]; ok {
		tmp := value.(map[string]interface{})
		profile.Cards = &Cards{}
		profile.Cards.Get(nil, profile.ID, tmp)
	}
}

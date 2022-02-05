package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Audit struct {
	CreatedAt  *primitive.DateTime `json:"createdAt"        bson:"createdAt,omitempty"`
	ModifiedAt *primitive.DateTime `json:"modifiedAt"       bson:"modifiedAt"`
}

func (audit *Audit) Create() {

	timeCurrent := primitive.NewDateTimeFromTime(time.Now())

	audit.CreatedAt = &timeCurrent
	audit.ModifiedAt = &timeCurrent
}

func (audit *Audit) Modify() {

	timeCurrent := primitive.NewDateTimeFromTime(time.Now())

	audit.ModifiedAt = &timeCurrent
}

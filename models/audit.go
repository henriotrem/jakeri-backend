package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Audit struct {
	CreatedBy  *string             `json:"createdBy"        bson:"createdBy,omitempty"`
	CreatedAt  *primitive.DateTime `json:"createdAt"        bson:"createdAt,omitempty"`
	ModifiedBy *string             `json:"modifiedBy"       bson:"modifiedBy"`
	ModifiedAt *primitive.DateTime `json:"modifiedAt"       bson:"modifiedAt"`
}

func (audit *Audit) Create(userId *string) {

	timeCurrent := primitive.NewDateTimeFromTime(time.Now())

	audit.CreatedBy = userId
	audit.CreatedAt = &timeCurrent
	audit.ModifiedBy = userId
	audit.ModifiedAt = &timeCurrent
}

func (audit *Audit) Modify(userId *string) {

	timeCurrent := primitive.NewDateTimeFromTime(time.Now())

	audit.ModifiedBy = userId
	audit.ModifiedAt = &timeCurrent
}

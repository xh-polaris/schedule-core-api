package group

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Group struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserId     string             `bson:"user_id" json:"userId"`
	Name       string             `bson:"name" json:"name"`
	Status     int64              `bson:"status" json:"status"`
	CreateTime time.Time          `bson:"create_time,omitempty" json:"createTime"`
	UpdateTime time.Time          `bson:"update_time,omitempty" json:"updateTime"`
	DeleteTime time.Time          `bson:"delete_time,omitempty" json:"deleteTime"`
}

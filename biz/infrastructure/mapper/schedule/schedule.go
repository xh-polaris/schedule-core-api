package schedule

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Schedule struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserId      string             `bson:"user_id" json:"userId"`
	Group       string             `bson:"group" json:"group"`
	Origin      string             `bson:"origin,omitempty" json:"origin,omitempty"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Done        int64              `bson:"done" json:"done"`
	Progress    int64              `bson:"progress" json:"progress"`
	Priority    int64              `bson:"priority" json:"priority"`
	Top         int64              `bson:"top" json:"top"`
	DDL         time.Time          `bson:"ddl,omitempty" json:"ddl,omitempty"`
	Status      int64              `bson:"status" json:"status"`
	CreateTime  time.Time          `bson:"create_time" json:"createTime"`
	UpdateTime  time.Time          `bson:"update_time" json:"updateTime"`
	DeleteTime  time.Time          `bson:"delete_time,omitempty" json:"deleteTime,omitempty"`
}

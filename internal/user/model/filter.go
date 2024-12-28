package model

import (
	"time"

	"github.com/core-go/search"
)

type UserFilter struct {
	*search.Filter
	Username       string            `yaml:"username" mapstructure:"username" json:"username" gorm:"column:username" bson:"username" dynamodbav:"username" firestore:"username" avro:"username" operator:">="`
	Version        []int32           `yaml:"version" mapstructure:"version" json:"version,omitempty" gorm:"column:version" bson:"version,omitempty" dynamodbav:"version" firestore:"version" avro:"version"`
	Email          *string           `yaml:"email" mapstructure:"email" json:"email" gorm:"column:email" bson:"email" dynamodbav:"email" firestore:"email" avro:"email" operator:">="`
	DateOfBirth    *search.TimeRange `yaml:"date_of_birth" mapstructure:"date_of_birth" json:"dateOfBirth" gorm:"column:date_of_birth" bson:"-" dynamodbav:"dateOfBirth" firestore:"dateOfBirth" avro:"dateOfBirth"`
	MinDateOfBirth *time.Time        `yaml:"date_of_birth" mapstructure:"date_of_birth" json:"minDateOfBirth" gorm:"column:date_of_birth" bson:"dateOfBirth" dynamodbav:"dateOfBirth" firestore:"dateOfBirth" avro:"dateOfBirth" operator:">="`
	MaxDateOfBirth *time.Time        `yaml:"date_of_birth" mapstructure:"date_of_birth" json:"maxDateOfBirth" gorm:"column:date_of_birth" bson:"dateOfBirth" dynamodbav:"dateOfBirth" firestore:"dateOfBirth" avro:"dateOfBirth" operator:"<="`
	Id             string            `yaml:"id" mapstructure:"id" json:"id" gorm:"column:id;primary_key" bson:"_id" dynamodbav:"id" firestore:"-" avro:"id" validate:"required,max=40" operator:"="`

	Phone string `yaml:"phone" mapstructure:"phone" json:"phone" gorm:"column:phone" bson:"phone" dynamodbav:"phone" firestore:"phone" avro:"phone" validate:"required,phone,max=18" operator:"like"`
}

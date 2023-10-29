package services

import (
	"REST-API/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	userCollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(userCollection *mongo.Collection, ctx context.Context) *UserServiceImpl {
	return &UserServiceImpl{
		userCollection: userCollection,
		ctx:            ctx,
	}
}

func (u *UserServiceImpl) CreatUser(user *models.User) error {

	_, err := u.userCollection.InsertOne(u.ctx, user)
	return err
}

func (u *UserServiceImpl) GetUser(name *string) (*models.User, error) {
	var user *models.User
	query := bson.D{bson.E{Key: "name", Value: name}}

	err := u.userCollection.FindOne(u.ctx, query).Decode(&user)

	return user, err
}

func (u *UserServiceImpl) GetAllUsers() ([]*models.User, error) {
	var users []*models.User

	cursor, err := u.userCollection.Find(u.ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(u.ctx) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if len(users) == 0 {
		return nil, mongo.ErrNoDocuments
	}
	
	return users, nil

}

func (u *UserServiceImpl) UpdateUser(user *models.User) error {
	filter := bson.D{bson.E{Key: "name", Value: user.Name}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "name", Value: user.Name}, bson.E{Key: "Email", Value: user.Email}, bson.E{Key: "Password", Value: user.Password}, bson.E{Key: "Address", Value: user.Address}}}}

	result,_:= u.userCollection.UpdateOne(u.ctx, filter, update)
	if  result.MatchedCount != 1 {
		return mongo.ErrNoDocuments
	}
	return nil

}

func (u *UserServiceImpl) DeleteUser(id *string) error {

	filter := bson.D{bson.E{Key: "id", Value: id}}

	result,_:= u.userCollection.DeleteOne(u.ctx, filter)
	if  result.DeletedCount != 1 {
		return mongo.ErrNoDocuments
	}
	return nil

}

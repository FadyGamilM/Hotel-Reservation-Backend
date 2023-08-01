package mongo

import (
	"context"

	"github.com/FadyGamilM/hotelreservationapi/internal/adapters/repository/mongo/dtos"
	"github.com/FadyGamilM/hotelreservationapi/internal/core/domain"
)

func (m *MongoDB) Insert(u *domain.User) {
	// define a context
	ctx := context.Background()

	// get the collection
	user_coll := m.Client.Database(m.DbName).Collection(m.CollectionsNames[Users_collection])

	user := dtos.UserDto{
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}

	res, err := user_coll.InsertOne(ctx, user)
	if err != nil {
		m.ErrLog.Fatalf("Error while inserting a new user : %v \n ", err)
	}

	m.InfoLog.Println(res)
}

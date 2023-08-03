package types

type User struct {
	// define the bson version as an omitempty so when we create an instance and not provide the id, it will not be passed as an empty so it won't be persisted as an empty id, instead it will be created via the db
	ID        string `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName string `bson:"first_name" json:"first_name"`
	LastName  string `bson:"last_name" json:"last_name"`
}

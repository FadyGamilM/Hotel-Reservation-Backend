package types

import (
	"golang.org/x/crypto/bcrypt"
)

var (
	COST                   = 12
	minFirstNameLength     = 2
	minLastNameLength      = 2
	minPasswordLength      = 8
	InvalidFirstNameErrMsg = "invalid first_name"
	InvalidLastNameErrMsg  = "invalid last_name"
	InvalidPasswordErrMsg  = "invalid password"
)

type User struct {
	// define the bson version as an omitempty so when we create an instance and not provide the id, it will not be passed as an empty so it won't be persisted as an empty id, instead it will be created via the db
	ID                string
	FirstName         string
	LastName          string
	Email             string
	EncryptedPassword string
}

func (user *User) Validate() error {
	if len(user.FirstName) < minFirstNameLength {
		return InvalidFirstNameErr{
			msg: InvalidFirstNameErrMsg,
		}
	}

	if len(user.LastName) < minLastNameLength {
		return InvalidLastNameErr{
			msg: InvalidLastNameErrMsg,
		}

	}

	if len(user.EncryptedPassword) < minPasswordLength {
		return InvalidPasswordErr{
			msg: InvalidPasswordErrMsg,
		}
	}

	return nil
}

type InvalidFirstNameErr struct {
	msg string
}

func (e InvalidFirstNameErr) Error() string {
	return e.msg
}

type InvalidLastNameErr struct {
	msg string
}

func (e InvalidLastNameErr) Error() string {
	return e.msg
}

type InvalidPasswordErr struct {
	msg string
}

func (e InvalidPasswordErr) Error() string {
	return e.msg
}

type UserMongoDb struct {
	// define the bson version as an omitempty so when we create an instance and not provide the id, it will not be passed as an empty so it won't be persisted as an empty id, instead it will be created via the db
	ID                string `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName         string `bson:"first_name" json:"first_name"`
	LastName          string `bson:"last_name" json:"last_name"`
	Email             string `bson:"email" json:"email"`
	EncryptedPassword string `bson:"encrypted_password" json:"-"` // _ in json because i won't return it in the json representation of this
}

// for handler usage, not for db usage
// thats how i will separate the logic and go towards more data-oriented approach
type CreateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type CreateUserResponse struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

// method to convert from web-req-data to domain-entity-data
func NewUserEntity(dto CreateUserRequest) (*User, error) {
	// encrypt the password
	hashed, err := bcrypt.GenerateFromPassword([]byte(dto.Password), COST)
	if err != nil {
		return nil, err
	}

	return &User{
		FirstName:         dto.FirstName,
		LastName:          dto.LastName,
		Email:             dto.Email,
		EncryptedPassword: string(hashed),
	}, nil
}

// method to convert form domain-entity-data to database-data type
func NewMongoDbUserEntity(domainEntity User) (*UserMongoDb, error) {
	return &UserMongoDb{
		FirstName:         domainEntity.FirstName,
		LastName:          domainEntity.LastName,
		Email:             domainEntity.Email,
		EncryptedPassword: domainEntity.EncryptedPassword,
	}, nil
}

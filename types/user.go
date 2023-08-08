package types

import (
	"fmt"
	"net"
	"reflect"
	"regexp"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var (
	COST                      = 12
	minFirstNameLength        = 2
	minLastNameLength         = 2
	minPasswordLength         = 8
	InvalidFirstNameErrMsg    = "invalid first_name"
	InvalidLastNameErrMsg     = "invalid last_name"
	InvalidPasswordErrMsg     = "invalid password"
	InvalidEmailErrMsg        = "invalid email"
	InvalidUpdateParameterMsg = "invalid update paramters"
)

type User struct {
	// define the bson version as an omitempty so when we create an instance and not provide the id, it will not be passed as an empty so it won't be persisted as an empty id, instead it will be created via the db
	ID                int64
	FirstName         string
	LastName          string
	Email             string
	EncryptedPassword string
}

type PostgresUser struct {
	ID                int64     `json:"id"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	Email             string    `json:"email"`
	EncryptedPassword string    `json:"encrypted_password"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func (pu *PostgresUser) IsSameDomainEntity(u *User) bool {
	if pu.FirstName == u.FirstName && pu.LastName == u.LastName && pu.Email == u.Email && pu.EncryptedPassword == u.EncryptedPassword {
		return true
	}
	return false
}

func (user *User) Validate() error {
	if len(user.FirstName) < minFirstNameLength {
		return InvalidFirstNameErr{
			Msg: InvalidFirstNameErrMsg,
		}
	}

	if len(user.LastName) < minLastNameLength {
		return InvalidLastNameErr{
			Msg: InvalidLastNameErrMsg,
		}

	}

	if len(user.EncryptedPassword) < minPasswordLength {
		return InvalidPasswordErr{
			Msg: InvalidPasswordErrMsg,
		}
	}

	if !isEmailValid(user.Email) {
		return InvalidEmailErr{
			Msg: InvalidEmailErrMsg,
		}
	}

	return nil
}

// domain entity specifies for our app which fields are allowed to be updated and which fields are not .. the powerful domain design is optained here !
func (u *User) Update(field string, value interface{}) error {
	fieldExists, ok := reflect.TypeOf(u).FieldByName(field)
	if ok {
		fmt.Println("field exists : ", fieldExists.Name)
		switch fieldExists.Name {
		case "FirstName":
			{
				correctType := reflect.TypeOf(value).Kind() == reflect.String
				if !correctType {
					return InvalidFirstNameErr{Msg: InvalidFirstNameErrMsg}
				}
				fmt.Printf("it was %v ", u.FirstName)
				u.FirstName = value.(string)
				fmt.Printf("it became %v \n", u.FirstName)
			}
		case "LastName":
			{
				correctType := reflect.TypeOf(value).Kind() == reflect.String
				if !correctType {
					return InvalidLastNameErr{Msg: InvalidLastNameErrMsg}
				}
				fmt.Printf("it was %v ", u.LastName)
				u.LastName = value.(string)
				fmt.Printf("it became %v \n", u.LastName)
			}
		}

	} else {
		return InvalidUpdateParameterErr{Msg: InvalidUpdateParameterMsg}
	}
	return nil
}

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func isEmailValid(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	if !emailRegex.MatchString(e) {
		return false
	}
	parts := strings.Split(e, "@")
	mx, err := net.LookupMX(parts[1])
	if err != nil || len(mx) == 0 {
		return false
	}
	return true
}

type InvalidFirstNameErr struct {
	Msg string
}

func (e InvalidFirstNameErr) Error() string {
	return e.Msg
}

type InvalidLastNameErr struct {
	Msg string
}

func (e InvalidLastNameErr) Error() string {
	return e.Msg
}

type InvalidPasswordErr struct {
	Msg string
}

func (e InvalidPasswordErr) Error() string {
	return e.Msg
}

type InvalidEmailErr struct {
	Msg string
}

func (e InvalidEmailErr) Error() string {
	return e.Msg
}

type InvalidUpdateParameterErr struct {
	Msg string
}

func (e InvalidUpdateParameterErr) Error() string {
	return e.Msg
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
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type UpdateUserResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type GetUserResponse struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type UserMongoDb struct {
	// define the bson version as an omitempty so when we create an instance and not provide the id, it will not be passed as an empty so it won't be persisted as an empty id, instead it will be created via the db
	ID                string `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName         string `bson:"first_name" json:"first_name"`
	LastName          string `bson:"last_name" json:"last_name"`
	Email             string `bson:"email" json:"email"`
	EncryptedPassword string `bson:"encrypted_password" json:"-"` // _ in json because i won't return it in the json representation of this
}

// to allow user to send one or more fields we set the fields as a map with string key and an interface as a value
type UpdateUserRequest struct {
	// btw the allowed fiedls to be updated are the firstName and lastName
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// method to convert from web-req-data to domain-entity-data
func NewUserEntityFromUserRequestDto(dto CreateUserRequest) (*User, error) {

	hashed, err := EncryptPassword(dto.Password)
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

func EncryptPassword(password string) (string, error) {
	// encrypt the password
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), COST)
	if err != nil {
		return password, err
	}
	return string(hashed), nil
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

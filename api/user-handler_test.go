package api

// import (
// 	"bytes"
// 	"database/sql"
// 	"encoding/json"
// 	"net/http/httptest"
// 	"testing"
// 	"time"

// 	"github.com/FadyGamilM/hotelreservationapi/db"
// 	"github.com/FadyGamilM/hotelreservationapi/db/postgres"
// 	"github.com/FadyGamilM/hotelreservationapi/types"
// 	"github.com/gofiber/fiber/v2"
// )

// const (
// 	TESTING_DB_DSN = "host=localhost port=2222 user=postgres password=postgres dbname=testhrdb sslmode=disable timezone=UTC connect_timeout=5"
// )

// // a testing database type
// type testDB struct {
// 	userRepo db.UserRepo
// }

// // setup a testing database
// func NewTestingDB(t *testing.T) *testDB {
// 	connPool, err := sql.Open("pgx", TESTING_DB_DSN)
// 	if err != nil {
// 		t.Errorf("Error while connecting to a testing db : %v \n", err)
// 	}

// 	// set pool conn attributes
// 	connPool.SetMaxOpenConns(10)
// 	connPool.SetMaxIdleConns(5)
// 	connPool.SetConnMaxLifetime(5 * time.Minute)

// 	// return the testingdb with the appropriate repos
// 	return &testDB{
// 		userRepo: &postgres.UserPostgresRepo{
// 			// DB: connPool,
// 		},
// 	}
// }

// type testRequests struct {
// 	method   string
// 	route    string
// 	response interface{}
// }

// func TestCreateUserHandler(t *testing.T) {
// 	// define test cases
// 	tRequests := []testRequests{
// 		testRequests{
// 			method: "POST",
// 			route:  "/users",
// 			response: types.CreateUserResponse{
// 				FirstName: "fady",
// 				LastName:  "gamil",
// 				Email:     "fadygamil@gmail.com",
// 			},
// 		},
// 	}

// 	// setup a testing database
// 	testDB := NewTestingDB(t)

// 	// drop the tables you created in your test database again after finishing the test case

// 	// setup an app instance
// 	testApp := fiber.New()

// 	// define a user handler
// 	createUserHandler := NewUserHandler(testDB.userRepo).HandleCreateUser

// 	// define the routes on your app instance
// 	for _, tRequest := range tRequests {
// 		testApp.Post(tRequest.route, createUserHandler)
// 	}

// 	// simulate the defined routes to perform the request
// 	userReqDto := types.CreateUserRequest{
// 		FirstName: "fady",
// 		LastName:  "gamil",
// 		Email:     "fadygamil@gmail.com",
// 		Password:  "fadygamil1234",
// 	}

// 	// prepare the data for the request
// 	reqParamsBytes, err := json.Marshal(userReqDto)
// 	if err != nil {
// 		t.Errorf("Error while marshleing the request params : %v \n", err)
// 	}

// 	// compare the results with the expected results
// 	// => define httptest request
// 	req := httptest.NewRequest(tRequests[0].method, tRequests[0].route, bytes.NewReader(reqParamsBytes))
// 	// => set req headrs and body
// 	req.Header.Add("Content-Type", "application/json")
// 	// => get the response
// 	res, err := testApp.Test(req)
// 	if err != nil {
// 		t.Errorf("Error while testing the request :%v \n", err)
// 	}

// 	responseDto := new(types.CreateUserResponse)
// 	json.NewDecoder(res.Body).Decode(&responseDto)
// 	if responseDto.FirstName != userReqDto.FirstName ||
// 		responseDto.LastName != userReqDto.LastName ||
// 		responseDto.Email != userReqDto.Email {
// 		t.Errorf("expected %v \n but got %v \n", userReqDto, responseDto)
// 	}

// }

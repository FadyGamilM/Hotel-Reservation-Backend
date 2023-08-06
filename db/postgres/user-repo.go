package postgres

import (
	"errors"
	"log"
	"strconv"

	"github.com/FadyGamilM/hotelreservationapi/types"
)

// GetUsers(context.Context) ([]*types.User, error)
// CreateUser(context.Context, *types.UserMongoDb) (*types.User, error)
// GetUserById(context.Context, string) (*types.User, error)
// DeleteUserById(context.Context, string) error
// UpdateUserById(context.Context, string, types.UpdateUserRequest) error

func (upr *UserPostgresRepo) GetUsers() ([]*types.User, error) {
	ctx, cancel := CreateContext()
	defer cancel()

	query := `
		SELECT * FROM users 
	`

	rows, err := upr.db.QueryContext(ctx, query)
	if err != nil {
		log.Printf("[REPO] | Error while fetching users from database : %v \n", err)
		return nil, err
	}

	var users []*types.User
	for rows.Next() {
		var user *types.User
		err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.EncryptedPassword,
		)
		if err != nil {
			log.Printf("[REPO] | Error while scanning user from db to domain entity type : %v \n", err)
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (upr *UserPostgresRepo) GetUserById(id string) (*types.User, error) {
	ctx, cancel := CreateContext()
	defer cancel()

	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Printf("[REPO] | error while converting user id from string to int64 to execute db query")
		return nil, err
	}

	query := `
		SELECT * FROM users 
		WHERE id = $1
	`

	row := upr.db.QueryRowContext(ctx, query, userID)

	var user *types.User
	err = row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.EncryptedPassword,
	)
	if err != nil {
		log.Printf("[REPO] | Error while scanning user from db to domain entity type : %v \n", err)
		return nil, err
	}

	return user, nil
}

func (upr *UserPostgresRepo) CreateUser(domainUser *types.User) (*types.User, error) {
	ctx, cancel := CreateContext()
	defer cancel()

	query := `
	INSERT INTO users (first_name, last_name, email, encrypted_password)
	VALUES ($1, $2, $3, $4)
	`

	res, err := upr.db.ExecContext(ctx, query, domainUser.FirstName, domainUser.LastName, domainUser.Email, domainUser.EncryptedPassword)
	if err != nil {
		log.Printf("[REPO] | Error while inserting user to database : %v \n", err)
		return nil, err
	}
	insertedUserID, err := res.LastInsertId()
	if err != nil {
		log.Printf("[REPO] | Error while trying to fetch the id of last inserted user to database : %v \n", err)
		return nil, err
	}
	var dbUser *types.PostgresUser
	fetchUserByIdQuery := `SELECT * FROM users WHERE id = $1`
	insertedUserRow := upr.db.QueryRowContext(ctx, fetchUserByIdQuery, insertedUserID)
	err = insertedUserRow.Scan(
		&dbUser.ID,
		&dbUser.FirstName,
		&dbUser.LastName,
		&dbUser.Email,
		&dbUser.EncryptedPassword,
	)
	if err != nil {
		log.Printf("[REPO] | Error while trying to fetch the last inserted user to database : %v \n", err)
		return nil, err
	}
	if dbUser.IsSameDomainEntity(domainUser) {
		return domainUser, nil
	} else {
		log.Printf("[REPO] | Error while trying to fetch the last inserted user to database : %v \n", err)
		return nil, errors.New("[REPO] | Not the same user is persisted in the database ! ")
	}
}

func (upr *UserPostgresRepo) UpdateUserById(id string, updatedValues *types.UpdateUserRequest) (*types.User, error) {
	ctx, cancel := CreateContext()
	defer cancel()

	fetchUserByIdQuery := `SELECT * FROM users WHERE id = $1`
	var updatedUser *types.User
	row := upr.db.QueryRowContext(ctx, fetchUserByIdQuery, id)
	err := row.Scan(
		&updatedUser.ID,
		&updatedUser.FirstName,
		&updatedUser.LastName,
		&updatedUser.Email,
		&updatedUser.EncryptedPassword,
	)
	if err != nil {
		log.Printf("[REPO] | Error while trying to fetch the user which should be updated from database : %v \n", err)
		return nil, err
	}

	updateFirstNameQuery := `
		UPDATE users SET first_name = $1
		WHERE id = $2
	`
	updateLastNameQuery := `
		UPDATE users SET last_name = $1
		WHERE id = $2
	`
	updateFirstAndLastNameQuery := `
		UPDATE users SET first_name = $1, last_name = $2
		WHERE id = $3
	`

	valuesToBeUpdated := make(map[string]string)
	if len(updatedValues.FirstName) > 0 {
		valuesToBeUpdated["first_name"] = updatedValues.FirstName
	}
	if len(updatedValues.LastName) > 0 {
		valuesToBeUpdated["last_name"] = updatedValues.LastName
	}

	updatedFirstName, firstNameProvided := valuesToBeUpdated["first_name"]
	updatedLastName, lastNameProvided := valuesToBeUpdated["last_name"]
	if firstNameProvided && !lastNameProvided {
		res, err := upr.db.ExecContext(ctx, updateFirstNameQuery, updatedFirstName, id)
		if err != nil {
			log.Printf("[REPO] | error while updating the user")
		}
		insertedUserID, err := res.LastInsertId()
		if err != nil {
			log.Printf("[REPO] | Error while trying to fetch the id of last inserted user to database : %v \n", err)
			return nil, err
		}
		updatedUserRow := upr.db.QueryRowContext(ctx, fetchUserByIdQuery, insertedUserID)
		err = updatedUserRow.Scan(
			&updatedUser.ID,
			&updatedUser.FirstName,
			&updatedUser.LastName,
			&updatedUser.Email,
			&updatedUser.EncryptedPassword,
		)
		if err != nil {
			log.Printf("[REPO] | Error while trying to fetch the last inserted user to database : %v \n", err)
			return nil, err
		}
		return updatedUser, nil
	}
	if !firstNameProvided && lastNameProvided {
		res, err := upr.db.ExecContext(ctx, updateLastNameQuery, updatedLastName, id)
		if err != nil {
			log.Printf("[REPO] | error while updating the user")
		}
		insertedUserID, err := res.LastInsertId()
		if err != nil {
			log.Printf("[REPO] | Error while trying to fetch the id of last inserted user to database : %v \n", err)
			return nil, err
		}
		updatedUserRow := upr.db.QueryRowContext(ctx, fetchUserByIdQuery, insertedUserID)
		err = updatedUserRow.Scan(
			&updatedUser.ID,
			&updatedUser.FirstName,
			&updatedUser.LastName,
			&updatedUser.Email,
			&updatedUser.EncryptedPassword,
		)
		if err != nil {
			log.Printf("[REPO] | Error while trying to fetch the last inserted user to database : %v \n", err)
			return nil, err
		}
		return updatedUser, nil
	}
	if firstNameProvided && lastNameProvided {
		res, err := upr.db.ExecContext(ctx, updateFirstAndLastNameQuery, updatedFirstName, updatedLastName, id)
		if err != nil {
			log.Printf("[REPO] | error while updating the user")
		}
		insertedUserID, err := res.LastInsertId()
		if err != nil {
			log.Printf("[REPO] | Error while trying to fetch the id of last inserted user to database : %v \n", err)
			return nil, err
		}
		updatedUserRow := upr.db.QueryRowContext(ctx, fetchUserByIdQuery, insertedUserID)
		err = updatedUserRow.Scan(
			&updatedUser.ID,
			&updatedUser.FirstName,
			&updatedUser.LastName,
			&updatedUser.Email,
			&updatedUser.EncryptedPassword,
		)
		if err != nil {
			log.Printf("[REPO] | Error while trying to fetch the last inserted user to database : %v \n", err)
			return nil, err
		}
		return updatedUser, nil
	}

	return updatedUser, nil
}

func (upr *UserPostgresRepo) DeleteUserById(id string) error {
	ctx, cancel := CreateContext()
	defer cancel()

	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Printf("[REPO] | error while converting user id from string to int64 to execute db query")
		return err
	}

	query := `
		DELETE FROM users 
		WHERE id = $1
	`

	_, err = upr.db.ExecContext(ctx, query, userID)
	if err != nil {
		log.Printf("[REPO] | Error while trying to delete user from database : %v \n", err)
		return err
	}

	return nil
}

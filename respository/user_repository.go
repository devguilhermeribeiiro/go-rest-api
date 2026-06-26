package respository

import (
	"database/sql"
	"fmt"

	"users-api/model"
)

type UserRepository struct {
	dbConnection *sql.DB
}

func NewUserRepository(connection *sql.DB) UserRepository {
	return UserRepository{
		dbConnection: connection,
	}
}

func (ur *UserRepository) GetUsers() ([]model.User, error) {
	query := "SELECT * FROM users"
	rows, err := ur.dbConnection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.User{}, err
	}

	defer rows.Close()

	var userList []model.User
	var userObj model.User

	for rows.Next() {
		err = rows.Scan(
			&userObj.ID,
			&userObj.Name,
			&userObj.Balance,
		)
		if err != nil {
			fmt.Println(err)
			return []model.User{}, err
		}

		userList = append(userList, userObj)
	}

	return userList, nil
}

func (ur *UserRepository) CreateUser(user model.User) (int, error) {
	var id int

	query, err := ur.dbConnection.Prepare("INSERT INTO users (name, balance) VALUES ($1, $2) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	defer query.Close()

	err = query.QueryRow(user.Name, user.Balance).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return id, nil
}

func (ur *UserRepository) GetUserByID(id int) (*model.User, error) {
	user := model.User{}

	query, err := ur.dbConnection.Prepare("SELECT id, name, balance FROM users WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer query.Close()

	err = query.QueryRow(id).Scan(&user.ID, &user.Name, &user.Balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		fmt.Println(err)
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) UpdateUser(user model.User) (*model.User, error) {
	updatedUser := model.User{}

	query, err := ur.dbConnection.Prepare("UPDATE users SET name = $1, balance = $2 WHERE id = $3 RETURNING *")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer query.Close()

	err = query.QueryRow(user.Name, user.Balance, user.ID).Scan(&updatedUser.ID, &updatedUser.Name, &updatedUser.Balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		fmt.Println(err)
		return nil, err
	}

	return &updatedUser, nil
}

func (ur *UserRepository) DeleteUserByID(id int) (*model.User, error) {
	deletedUser := model.User{}

	query, err := ur.dbConnection.Prepare("DELETE FROM users WHERE id = $1 RETURNING *")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer query.Close()

	err = query.QueryRow(id).Scan(&deletedUser.ID, &deletedUser.Name, &deletedUser.Balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		fmt.Println(err)
		return nil, err
	}

	return &deletedUser, nil
}

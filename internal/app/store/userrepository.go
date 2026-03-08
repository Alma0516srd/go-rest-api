package store

import "test/internal/app/model"

type UserRepository struct {
	store *Store
}

// scan мапит возвращаемое значение в переменную
func (r *UserRepository) CreateUser(u *model.User) (*model.User, error) {
	err := r.store.db.QueryRow(
		"insert into users(email, encrypted_password) "+
			"values ($1, $2) returning id)", u.Email, u.EncryptedPassword).Scan(&u.ID)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *UserRepository) findByEmail(email string) (*model.User, error) {
	return nil, nil
}

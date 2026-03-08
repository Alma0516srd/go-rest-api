package store

import "test/internal/app/model"

type UserRepository struct {
	store *Store
}

// scan мапит возвращаемое значение в переменную
func (r *UserRepository) CreateUser(u *model.User) (*model.User, error) {
	err := r.store.db.QueryRow(
		"insert into users(email, encrypted_password) "+
			"values ($1, $2) returning id", u.Email, u.EncryptedPassword).Scan(&u.ID)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	user := model.User{}
	err := r.store.db.QueryRow("select * from users where email = $1", email).
		Scan(&user.ID, &user.Email, &user.EncryptedPassword) // sca разложит значения в структуру
	if err != nil {
		return nil, err
	}
	return &user, nil
}

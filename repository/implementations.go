package repository

import (
	"context"
	"database/sql"

	"github.com/SawitProRecruitment/UserService/utils"
)

func (r *Repository) GetTestById(ctx context.Context, input GetTestByIdInput) (output GetTestByIdOutput, err error) {
	err = r.Db.QueryRowContext(ctx, "SELECT name FROM test WHERE id = $1", input.Id).Scan(&output.Name)
	if err != nil {
		return
	}
	return
}

func (r *Repository) RegisterUser(ctx context.Context, phoneNumber, fullName, passwordHash string) (out int, err error) {
	// db, err := sql.Open("postgres", "user=username dbname=mydb sslmode=disable")
	// err = r.Db.Conn()(ctx, "SELECT name FROM test WHERE id = $1", input.Id).Scan(&output.Name)
	// if err != nil {
	//     return 0, err
	// }
	defer r.Db.Close()

	var id int
	err = r.Db.QueryRow(`
        INSERT INTO users (phone_number, full_name, password_hash)
        VALUES ($1, $2, $3)
        RETURNING id
    `, phoneNumber, fullName, passwordHash).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *Repository) LoginUser(ctx context.Context, phoneNumber, passwordHash string) (out int, err error) {
	// db, err := sql.Open("postgres", "user=username dbname=mydb sslmode=disable")
	// if err != nil {
	//     return 0, err
	// }
	defer r.Db.Close()

	var id int
	err = r.Db.QueryRowContext(ctx, `
        SELECT id FROM users WHERE phone_number = $1 AND password_hash = $2
    `, phoneNumber, passwordHash).Scan(&id)

	if err != nil {
		return 0, err
	}

	// Insert a successful login attempt
	_, err = r.Db.Exec(`
        INSERT INTO login_attempts (user_id, successful)
        VALUES ($1, true)
    `, id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *Repository) GetMyProfile(ctx context.Context, userID int) (phoneNumber string, fullName string, err error) {
	// db, err := sql.Open("postgres", "user=username dbname=mydb sslmode=disable")
	// if err != nil {
	//     return "", "", err
	// }
	defer r.Db.Close()

	err = r.Db.QueryRowContext(ctx, `
        SELECT phone_number, full_name FROM users WHERE id = $1
    `, userID).Scan(&phoneNumber, &fullName)

	if err != nil {
		return "", "", err
	}

	return phoneNumber, fullName, nil
}

func (r *Repository) UpdateMyProfile(ctx context.Context, userID int, phoneNumber, fullName string) (err error) {
	// db, err := sql.Open("postgres", "user=username dbname=mydb sslmode=disable")
	// if err != nil {
	//     return err
	// }
	defer r.Db.Close()

	// Check if the new phone number is already in use
	var existingUserID int
	err = r.Db.QueryRowContext(ctx, `
        SELECT id FROM users WHERE phone_number = $1 AND id != $2
    `, phoneNumber, userID).Scan(&existingUserID)

	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if existingUserID != 0 {
		return utils.ErrPhoneNumberAlreadyExists
	}

	// Update the user's profile with the provided data
	_, err = r.Db.Exec(`
        UPDATE users
        SET phone_number = COALESCE(NULLIF($1, ''), phone_number),
            full_name = COALESCE(NULLIF($2, ''), full_name)
        WHERE id = $3
    `, phoneNumber, fullName, userID)

	if err != nil {
		return err
	}

	return nil
}

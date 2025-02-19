// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: 00002_pre_go_acc_user_base_9999.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const addUserBase = `-- name: AddUserBase :one
INSERT INTO pre_go_acc_user_base_9999 (
    user_account, user_password, user_salt, user_created_at, user_updated_at 
) 
VALUES (
    $1, $2, $3, NOW(), NOW()
)
RETURNING user_id
`

type AddUserBaseParams struct {
	UserAccount  string
	UserPassword string
	UserSalt     string
}

func (q *Queries) AddUserBase(ctx context.Context, arg AddUserBaseParams) (int32, error) {
	row := q.db.QueryRow(ctx, addUserBase, arg.UserAccount, arg.UserPassword, arg.UserSalt)
	var user_id int32
	err := row.Scan(&user_id)
	return user_id, err
}

const checkUserBaseExists = `-- name: CheckUserBaseExists :one
SELECT COUNT(*) 
FROM pre_go_acc_user_base_9999 
WHERE user_account = $1
`

func (q *Queries) CheckUserBaseExists(ctx context.Context, userAccount string) (int64, error) {
	row := q.db.QueryRow(ctx, checkUserBaseExists, userAccount)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getOneUserInfo = `-- name: GetOneUserInfo :one
SELECT user_id, user_account, user_password, user_salt
FROM pre_go_acc_user_base_9999 
WHERE user_account = $1
`

type GetOneUserInfoRow struct {
	UserID       int32
	UserAccount  string
	UserPassword string
	UserSalt     string
}

func (q *Queries) GetOneUserInfo(ctx context.Context, userAccount string) (GetOneUserInfoRow, error) {
	row := q.db.QueryRow(ctx, getOneUserInfo, userAccount)
	var i GetOneUserInfoRow
	err := row.Scan(
		&i.UserID,
		&i.UserAccount,
		&i.UserPassword,
		&i.UserSalt,
	)
	return i, err
}

const getOneUserInfoAdmin = `-- name: GetOneUserInfoAdmin :one
SELECT 
    user_id, 
    user_account,
    user_password, 
    user_salt,
    user_login_time, 
    user_logout_time, 
    user_login_ip, 
    user_created_at, 
    user_updated_at 
FROM pre_go_acc_user_base_9999 
WHERE user_account = $1
`

type GetOneUserInfoAdminRow struct {
	UserID         int32
	UserAccount    string
	UserPassword   string
	UserSalt       string
	UserLoginTime  pgtype.Timestamp
	UserLogoutTime pgtype.Timestamp
	UserLoginIp    pgtype.Text
	UserCreatedAt  pgtype.Timestamp
	UserUpdatedAt  pgtype.Timestamp
}

func (q *Queries) GetOneUserInfoAdmin(ctx context.Context, userAccount string) (GetOneUserInfoAdminRow, error) {
	row := q.db.QueryRow(ctx, getOneUserInfoAdmin, userAccount)
	var i GetOneUserInfoAdminRow
	err := row.Scan(
		&i.UserID,
		&i.UserAccount,
		&i.UserPassword,
		&i.UserSalt,
		&i.UserLoginTime,
		&i.UserLogoutTime,
		&i.UserLoginIp,
		&i.UserCreatedAt,
		&i.UserUpdatedAt,
	)
	return i, err
}

const loginUserBase = `-- name: LoginUserBase :exec
UPDATE pre_go_acc_user_base_9999 
    SET user_login_time = NOW(), 
    user_login_ip = $3
WHERE user_account = $1 AND user_password = $2
RETURNING 
    user_id
`

type LoginUserBaseParams struct {
	UserAccount  string
	UserPassword string
	UserLoginIp  pgtype.Text
}

func (q *Queries) LoginUserBase(ctx context.Context, arg LoginUserBaseParams) error {
	_, err := q.db.Exec(ctx, loginUserBase, arg.UserAccount, arg.UserPassword, arg.UserLoginIp)
	return err
}

const logoutUserBase = `-- name: LogoutUserBase :exec
UPDATE 
    pre_go_acc_user_base_9999 
SET user_logout_time = NOW() 
WHERE user_account = $1
`

func (q *Queries) LogoutUserBase(ctx context.Context, userAccount string) error {
	_, err := q.db.Exec(ctx, logoutUserBase, userAccount)
	return err
}

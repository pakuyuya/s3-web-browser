package user

import (
	"database/sql"
	"fmt"
	// Postgresql Driver
    _ "github.com/lib/pq"
)

// User is a modle of record.
type User struct {
	ID int `json:"id"`
	Username string `json:"username"`
	Loginid string `json:"loginid"`
	Password string `json:"password"`
}

func (m *User) String() string {
	return fmt.Sprintf("{Id:%d, Username:%s, Loginid:%s, Password:%s}", m.ID, m.Username, m.Loginid, m.Password)
}

// SelectAll is a function that get all users from repositoy.
func SelectAll(conn *sql.Tx) ([]User, error) {
	rows, err := conn.Query("SELECT id, username, loginid, '********' AS password FROM s3web.users ORDER BY id FOR READ ONLY;");

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		user := User{}
		rows.Scan(&user.ID, &user.Username, &user.Loginid, &user.Password)
		users = append(users, user)
	}

	return users, nil
}

// SelectByID is a function that get all users from repositoy.
func SelectByID(conn *sql.Tx, id int) (*User, error) {
	row := conn.QueryRow("SELECT id, username, loginid, '********' AS password FROM s3web.users WHERE id = $1 FOR READ ONLY;", id);

	user := User{}
	err := row.Scan(&user.ID, &user.Username, &user.Loginid, &user.Password)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// SelectForAuth is a function that try get a record using login infomation.
func SelectForAuth(conn *sql.Tx, loginid string, password string) (*User, error) {
	row := conn.QueryRow("SELECT id, username, loginid, '********' AS password FROM s3web.users WHERE loginid = $1 AND password_sha256 = digest($2, 'sha256')::varchar(256) FOR READ ONLY;", loginid, password);

	user := User{}
	err := row.Scan(&user.ID, &user.Username, &user.Loginid, &user.Password)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Insert is a function that insert a record to repositoy.
func Insert(conn *sql.Tx, m *User) (int, error) {
	query := "INSERT INTO s3web.users(username, loginid, password_sha256, create_at, update_at) VALUES($1, $2, digest($3, 'sha256'), CURRENT_TIMESTAMP, CURRENT_TIMESTAMP) RETURNING id;"
	args := []interface{}{&m.Username, &m.Loginid, &m.Password}

	row := conn.QueryRow(query, args...);

	id := 0
	err := row.Scan(&id)
	return id, err
}

// UpdateByID is a function that update a record in repositoy.
// This function do not update password.
func UpdateByID(conn *sql.Tx, m *User) (int64, error) {
	query := "UPDATE s3web.users SET username=$2, loginid=$3, update_at=CURRENT_TIMESTAMP WHERE id=$1;"
	args := []interface{}{m.ID, m.Username, m.Loginid}

	r, err := conn.Exec(query, args...);
	if err != nil {
		return 0, err
	}

	return r.RowsAffected()
}

// UpdatePasswordByID is a function that update the pasword of record in repositoy.
func UpdatePasswordByID(conn *sql.Tx, m *User) (int64, error) {
	query := "UPDATE s3web.users SET password_sha256=digest($2, 'sha256'), update_at=CURRENT_TIMESTAMP WHERE id=$1;"
	args := []interface{}{m.ID, m.Password}

	r, err := conn.Exec(query, args...);
	if err != nil {
		return 0, err
	}

	return r.RowsAffected()
}

// DeleteByID is a function that delete a record from repositoy.
func DeleteByID(conn *sql.Tx, id int) (int64, error) {
	query := "DELETE FROM s3web.users WHERE id=$1;"

	r, err := conn.Exec(query, id);
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

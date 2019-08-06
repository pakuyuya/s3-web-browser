package profile

import (
    "database/sql"

    _ "github.com/lib/pq"
)

// Profile is a modle of record.
type Profile struct {
	Profileid string
	Profilename string
	Connjson string
	Bucket string
	Basepath string
}

// GetAll is a function that get all profiles from repositoy.
func SelectAll(conn *sql.Tx) ([]Profile, error) {
	rows, err := conn.Query("SELECT profileid, profilename, connjson, bucket, basepath FROM s3web.profiles");

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	profiles := make([]Profile, 0)
	for rows.Next() {
		profile := Profile{}
		rows.Scan(&profile.Profileid, &profile.Profilename, &profile.Connjson, &profile.Bucket, &profile.Basepath)
		profiles = append(profiles, profile)
	}

	return profiles, nil
}

// GetAll is a function that get all profiles from repositoy.
func SelectById(conn *sql.Tx, profileid string) (*Profile, error) {
	row := conn.QueryRow("SELECT profileid, profilename, connjson, bucket, basepath FROM s3web.profiles WHERE profileid = $1;", profileid);

	profile := Profile{}
	err := row.Scan(&profile.Profileid, &profile.Profilename, &profile.Connjson, &profile.Bucket, &profile.Basepath)

	if err != nil {
		return nil, err
	}

	return &profile, nil
}

// Insert is a function that insert a record to repositoy.
func Insert(conn *sql.Tx, m *Profile) (error) {
	query := "INSERT INTO s3web.profiles(profileid, profilename, connjson, bucket, basepath, create_at, update_at) VALUES($1, $2, $3, $4, $5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);"
	args := []interface{}{m.Profileid, m.Profilename, m.Connjson, m.Connjson, m.Bucket, m.Basepath}

	_, err := conn.Exec(query, args...);
	if err != nil {
		return err
	}
	return nil
}

// Insert is a function that insert a record to repositoy.
func UpdateById(conn *sql.Tx, m *Profile) (error) {
	query := "UPDATE s3web.profiles SET profilename=$2, connjson=$3, bucket=$4, basepath=$5, update_at=CURRENT_TIMESTAMP WHERE profileid=$1;"
	args := []interface{}{m.Profileid, m.Profilename, m.Connjson, m.Connjson, m.Bucket, m.Basepath}

	_, err := conn.Exec(query, args...);
	if err != nil {
		return err
	}
	return nil
}

// Insert is a function that insert a record to repositoy.
func DeleteById(conn *sql.Tx, profileid string) (error) {
	query := "DELETE FROM s3web.profiles WHERE profileid=$1;"

	_, err := conn.Query(query, profileid);
	if err != nil {
		return err
	}
	return nil
}

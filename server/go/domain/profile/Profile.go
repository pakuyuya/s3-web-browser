package profile

import (
    "database/sql"
    "fmt"

	"s3-web-brawser/server/go/domain/db"

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
func GetAll(conn: *sql.Tx) ([]Profile, error) {
	rows, err := conn.Query("SELECT profileid, profilename, connjson, bucket, basepath FROM s3web.profiles");

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	profiles := make([]Profile)
	for rows.Next() {
		profile := Profile{}
		db.scan(&profile.profileid, &profile.poriflename, &profile.connjson, &profile.bucket, &profile.basepath)
		profiles = append(profiles, profile)
	}

	return profiles, nil
}

// Insert is a function that insert a record to repositoy.
func Insert(conn: *sql.Tx, m:*Profile) (error) {
	query := "INSERT INTO s3web.profiles(profileid, profilename, connjson, bucket, basepath, create_at, update_at) VALUES($1, $2, $3, $4, $5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);"
	args := []interface{m.Profileid, m.Profilename, m.Connjson, m.Connjson, m.Bucket, m.Basepath}

	_, err := conn.Exec(query, args...);
	if err != nil {
		return err
	}
	return nil
}

// Insert is a function that insert a record to repositoy.
func Update(conn: *sql.Tx, m:*Profile) (int64, error) {
	query := "UPDATE s3web.profiles SET profilename=$2, connjson=$3, bucket=$4, basepath=$5, update_at=CURRENT_TIMESTAMP WHERE profileid=$1;"
	args := []interface{m.Profileid, m.Profilename, m.Connjson, m.Connjson, m.Bucket, m.Basepath}

	result, err := conn.Exec(query, args...);
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

// Insert is a function that insert a record to repositoy.
func DeleteById(conn: *sql.Tx, profileid: string) (int64, error) {
	query := "DELETE FROM s3web.profiles WHERE profileid=$1;"

	result, err := conn.Query(query, profileid);
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

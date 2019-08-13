package profile

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
    // PostgreSQL driver
    _ "github.com/lib/pq"
)

// Profile is a modle of record.
type Profile struct {
	Profileid string
	Profilename string
	Connjson string
	Region string
	Bucket string
	Basepath string
}

func (m *Profile) String() string {
	return fmt.Sprintf("Profileid:%s, Profilename:%s, Connjson:%s, Region:%s, Bucket:%s, Basepath:%s", m.Profileid, m.Profilename, m.Region, m.Connjson, m.Bucket, m.Basepath)
}

// FormatBasepath is a function that normalize string as basepath
func FormatBasepath(basepath string) string {
	basepath = strings.TrimPrefix(basepath, "/")
	if basepath == "" {
		return basepath
	}
	if !strings.HasSuffix(basepath, "/") {
		basepath = basepath + "/"
	}
	return basepath
}

// SelectAll is a function that get all profiles from repositoy.
func SelectAll(conn *sql.Tx) ([]Profile, error) {
	rows, err := conn.Query("SELECT profileid, profilename, connjson, region, bucket, basepath FROM s3web.profiles ORDER BY profileid FOR READ ONLY;");

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	profiles := make([]Profile, 0)
	for rows.Next() {
		profile := Profile{}
		var profileid int
		rows.Scan(&profileid, &profile.Profilename, &profile.Connjson, &profile.Region, &profile.Bucket, &profile.Basepath)
		profile.Profileid = strconv.Itoa(profileid)
		profiles = append(profiles, profile)
	}

	return profiles, nil
}

// SelectByID is a function that get all profiles from repositoy.
func SelectByID(conn *sql.Tx, profileid string) (*Profile, error) {
	row := conn.QueryRow("SELECT profileid, profilename, connjson, region, bucket, basepath FROM s3web.profiles WHERE profileid = $1;", profileid);

	profile := Profile{}
	err := row.Scan(&profile.Profileid, &profile.Profilename, &profile.Connjson, &profile.Bucket, &profile.Bucket, &profile.Basepath)

	if err != nil {
		return nil, err
	}

	return &profile, nil
}

// Insert is a function that insert a record to repositoy.
func Insert(conn *sql.Tx, m *Profile) (int64, error) {
	query := "INSERT INTO s3web.profiles(profilename, connjson, bucket, region, basepath, create_at, update_at) VALUES($1, $2, $3, $4, $5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);"
	args := []interface{}{m.Profilename, m.Connjson, m.Region, m.Bucket, m.Basepath}

	r, err := conn.Exec(query, args...);
	if err != nil {
		return 0, err
	}

	return r.RowsAffected()
}

// UpdateByID is a function that insert a record to repositoy.
func UpdateByID(conn *sql.Tx, m *Profile) (int64, error) {
	query := "UPDATE s3web.profiles SET profilename=$2, connjson=$3, region=$4, bucket=$5, basepath=$6, update_at=CURRENT_TIMESTAMP WHERE profileid=$1;"
	profileid, _ := strconv.Atoi(m.Profileid)
	args := []interface{}{profileid, m.Profilename, m.Connjson, m.Region, m.Bucket, m.Basepath}

	r, err := conn.Exec(query, args...);
	if err != nil {
		return 0, err
	}

	return r.RowsAffected()
}

// DeleteByID is a function that insert a record to repositoy.
func DeleteByID(conn *sql.Tx, profileid string) (int64, error) {
	query := "DELETE FROM s3web.profiles WHERE profileid=$1;"

	r, err := conn.Exec(query, profileid);
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

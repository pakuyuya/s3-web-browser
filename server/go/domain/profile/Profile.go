package profile

import (
    "database/sql"
    "fmt"

	"s3-web-brawser/server/go/domain/db"

    _ "github.com/lib/pq"
)

// Profile is a modle of record.
type Profile struct {
	profileid string
	profilename string
	connjson string
	bucket string
	basepath string
}

// GetAll is a function that get all profiles from repositoy.
func GetAll() ([]Profile, error) {
	conn, err := db.Connection()
	if err != nil {
		return nil, err
	}

	rows, err := conn.Query("SELECT profileid, profilename, connjson, bucket, basepath FROM s3web.profiles");

	if err != nil {
		db.Close()
		return nil, err
	}

	profiles := make([]Profile)
	for rows.Next() {
		profile := Profile{}
		db.scan(&profile.profileid, &profile.poriflename, &profile.connjson, &profile.bucket, &profile.basepath)
		profiles = append(profiles, profile)
	}

	return profiles, nil
}

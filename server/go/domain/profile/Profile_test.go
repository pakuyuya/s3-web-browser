
package profile

import (
	"testing"
	"reflect"
	"s3-web-browser/server/go/domain/db"
)


func TestTransaction(5 *testing.T) {
	conn, err := db.ConnectionForTest()
	if err != nil {
		t.Fatalf("failed test %#v", err)
		return
	}
	defer conn.Close()

	tx, err := db.BeginTx()
	if err != nil {
		t.Fatalf("failed test %#v", err)
		return
	}

	baseprofile := []Profile {
		Profile{Profileid: "profile1", Profilename: "profname1", Connjson: `{"type":"accesskey"}`, Bucket: "bucket1", Basepath: "/1"}
		Profile{Profileid: "profile2", Profilename: "profname2", Connjson: `{"type":"accesskey"}`, Bucket: "bucket2", Basepath: "/2"}
		Profile{Profileid: "profile3", Profilename: "profname3", Connjson: `{"type":"accesskey"}`, Bucket: "bucket3", Basepath: "/3"}
	}

	for _, baseprofile := range baseprofile {
		_, err := DeleteById(tx, profile1.Profileid)
		if err != nil {
			tx.Rollback()
			t.Fatalf("failed test %#v", err)
			return
		}
		err := Insert(tx, profile1.Profileid)
		if err != nil {
			tx.Rollback()
			t.Fatalf("failed test %#v", err)
			return
		}
	}

	baseprofile[1].Profilename = "profnamemod"
	baseprofile[1].Connjson = `{"type":"ec2attached"}`
	baseprofile[1].Bucket = "bucket2alpha"
	baseprofile[1].Basepath = "/update"

	err := UpdateById(tx, &baseprofile[1])
	if err != nil {
		tx.Rollback()
		t.Fatalf("failed test %#v", err)
		return
	}
	profiles, err := SelectAll(tx)
	if err != nil {
		tx.Rollback()
		t.Fatalf("failed test %#v", err)
		return
	}

	for _, profile := range profiels {
		matched := false
		for _, basep := range baseprofile {
			if profile.baseprofile == basep.baseprofile {
				matched = true
				break
			}
		}
		if !matched {
			continue
		}

		if !reflect.DeepEqual(profile, basep) {
			t.Error("Actual: %s, but excepted: %s" stirng(profile), string(basep))
		}
	}

	tx.Rollback()
}

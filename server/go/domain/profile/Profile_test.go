
package profile

import (
	"testing"
	"reflect"
	"s3-web-browser/server/go/domain/db"
)


func TestTransaction(t *testing.T) {
	conn, err := db.ConnectionForTest()
	if err != nil {
		t.Fatalf("failed test %#v", err)
		return
	}
	defer conn.Close()

	tx, err := conn.Begin()
	if err != nil {
		t.Fatalf("failed test %#v", err)
		return
	}

	baseprofile := []Profile {
		Profile{Profileid: "profile1", Profilename: "profname1", Connjson: `{"type":"accesskey"}`, Bucket: "bucket1", Basepath: "/1"},
		Profile{Profileid: "profile2", Profilename: "profname2", Connjson: `{"type":"accesskey"}`, Bucket: "bucket2", Basepath: "/2"},
		Profile{Profileid: "profile3", Profilename: "profname3", Connjson: `{"type":"accesskey"}`, Bucket: "bucket3", Basepath: "/3"},
	}

	for _, baseprofile := range baseprofile {
		_, err := DeleteByID(tx, baseprofile.Profileid)
		if err != nil {
			tx.Rollback()
			t.Fatalf("failed test %#v", err)
			return
		}
		_, err = Insert(tx, &baseprofile)
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

	_, err = UpdateByID(tx, &baseprofile[1])
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

	for _, profile := range profiles {
		var foundProfile *Profile = nil 
		for _, basep := range baseprofile {
			if profile.Profileid == basep.Profileid {
				foundProfile = &basep
				break
			}
		}
		if foundProfile == nil {
			continue
		}

		if !reflect.DeepEqual(profile, *foundProfile) {
			t.Errorf("Actual: %s, but excepted: %s", profile.String(), foundProfile.String())
		}
	}

	
	p1, err := SelectByID(tx, baseprofile[1].Profileid)
	if !reflect.DeepEqual(p1, &baseprofile[1]) {
		t.Errorf("Actual: %s, but excepted: %s", p1.String(), baseprofile[1].String())
	}

	tx.Rollback()
}

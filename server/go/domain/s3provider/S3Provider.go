package s3provider

import (
    "encoding/json"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/service/s3/s3manager"
    "fmt"
	"strings"
)

// ConnParameter is a parameter to connect S3
type ConnParameter struct {
	Type string `json:"type"`
	Region string `json:"region"`
	Profile  string `json:"profile"`
	Accesskey string `json:"accesskey"`
	Secretkey string `json:"secretkey"`
}

// S3Item is a struct for infomation of S3 object / S3 prefix
type S3Item struct {
    Type: string
    Name: string
	Fullpath: string
	Size: string
	LastModified string
}

// CreateSession is a function to create session to AWS
func CreateSession(connJson string): (*session.Session, error) {
	param := ConnParameter{}
	if err := json.Unmarshal(param, connJson); err != nil {
		return nil, err
	}

	if param.Region == "" {
		return nil, fmt.Errorf("Invalid s3 profile: region is empty")
	}

	var c *credentials.Credentials
	switch param.Type {
	case "inherit":
		c, err = credentialWithInherit(param)
	case "accesskey":
		c, err = credentialWithAccesskey(param)
	}
	if err != nil {
		return nil, err
	}

	return session.NewSession(&aws.Config{
		Region: aws.String(param.Region),
		Credentials: c,
	})
}

func credentialWithInherit(*ConnParameter param) (*credentials.Credentials, error) {
	profile := param.Profile
	return session.NewSessionWithOptions(session.Options{Profile:profile})
}

func credentialWithAccesskey(*ConnParameter param) (*credentials.Credentials, error) {
	aws_access_key_id = param.Accesskey
	aws_secret_access_key = param.Secretkey
	return credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, "")
}


// List is a function that get list of object / folder in S3
func List(*session.Session sess, string bucket, string prefix) (S3Item[], err) {
	svc := s3.New(sess)

	resp, err := svc.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(bucketName),
		Prefix: aws.String(prefix),
		Delimiter: "/",
	})
	if err != nil {
		return nil, err
	}

	items := make(S3Item[], *)
	
	for _, prefix := range resp.CommonPrefixes {
		cutprefix = prefix.Prefix[0:len(prefix.Prefix)-1]
		idxDelimiter = strings.LastIndex(prefix.Prefix, "/")
		if idxDelimiter < 0 {
			idxDelimiter = 0
		}
		items = append(items, S3Item{
			Type: "directory",
			Name: cutprefix[idxDelimiter:],
			Fullpath: prefix.Prefix,
			Size: sprintSize(0),
			LastModified: "",
		})
	}
	for _, content := range resp.Contents {
		idxDelimiter = strings.LastIndex(content.Key, "/")
		if idxDelimiter < 0 {
			idxDelimiter = 0
		}
		items = append(items, S3Item{
			Type: "file",
			Name: content.Key[idxDelimiter:],
			Fullpath: content.Key,
			Size: sprintSize(content.Size),
			LastModified: content.LastModified,
		})
	}
}

// List is a function that download a file in S3 and write body to parametered writer
func DownloadStream(*session.Session sess, string bucket, string key, io.Writer w) err {
	svc := s3.New(sess)

	resp, err := svc.GetObject(&s3.ListObjectsInput{
		Bucket: aws.String(bucketName),
		Key: aws.String(key),
	})
	if err != nil {
		return err
	}

	_, err = io.Copy(w, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func sprintSize(n int64) string {
	if n < 1024 {
		return n + " Bytes"
	}
	if n < 1024*1024 {
		return n + " KB"
	}
	if n < 1024*1024*1024 {
		return n + " MB"
	}
	return n + " GB"
}
package s3provider

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// ConnParameter is a parameter to connect S3
type ConnParameter struct {
	Type      string `json:"type"`
	Region    string `json:"region"`
	Profile   string `json:"profile"`
	Accesskey string `json:"accesskey"`
	Secretkey string `json:"secretkey"`
}

// S3Item is a struct for infomation of S3 object / S3 prefix
type S3Item struct {
	Type         string
	Name         string
	Fullpath     string
	Size         string
	LastModified string
}

// CreateSession is a function to create session to AWS
func CreateSession(connJSON string) (*session.Session, error) {
	var err error
	param := ConnParameter{}
	if err = json.Unmarshal([]byte(connJSON), &param); err != nil {
		return nil, err
	}

	if param.Region == "" {
		return nil, fmt.Errorf("Invalid s3 profile: region is empty")
	}

	var c *credentials.Credentials
	switch param.Type {
	case "credentialfile":
		c, err = credentialWithInherit(&param)
	case "accesskey":
		c, err = credentialWithAccesskey(&param)
	}
	if err != nil {
		return nil, err
	}

	return session.NewSession(&aws.Config{
		Region:      aws.String(param.Region),
		Credentials: c,
	})
}

func credentialWithInherit(param *ConnParameter) (*credentials.Credentials, error) {
	profile := param.Profile
	return credentials.NewSharedCredentials("", profile), nil
}

func credentialWithAccesskey(param *ConnParameter) (*credentials.Credentials, error) {
	awsAccessKeyID := param.Accesskey
	awsSecretAccessKey := param.Secretkey
	return credentials.NewStaticCredentials(awsAccessKeyID, awsSecretAccessKey, ""), nil
}

// List is a function that get list of object / folder in S3
func List(sess *session.Session, bucket string, prefix string) ([]S3Item, error) {
	svc := s3.New(sess)

	delimiter := "/"

	resp, err := svc.ListObjects(&s3.ListObjectsInput{
		Bucket:    aws.String(bucket),
		Prefix:    aws.String(prefix),
		Delimiter: &delimiter,
	})
	if err != nil {
		return nil, err
	}

	items := make([]S3Item, 0)

	for _, p := range resp.CommonPrefixes {
		if strings.TrimSuffix(*(p.Prefix), "/") == strings.TrimSuffix(prefix, "/") {
			continue
		}
		pstr := *(p.Prefix)
		cutprefix := pstr[0 : len(pstr)-1]
		idxDelimiter := strings.LastIndex(cutprefix, "/")
		if idxDelimiter < 0 {
			idxDelimiter = -1
		}
		items = append(items, S3Item{
			Type:         "directory",
			Name:         cutprefix[idxDelimiter+1:],
			Fullpath:     *(p.Prefix),
			Size:         sprintSize(0),
			LastModified: "",
		})
	}
	for _, content := range resp.Contents {
		if strings.TrimSuffix(*(content.Key), "/") == strings.TrimSuffix(prefix, "/") {
			continue
		}
		key := *(content.Key)
		idxDelimiter := strings.LastIndex(key, "/")
		if idxDelimiter < 0 {
			idxDelimiter = -1
		}
		items = append(items, S3Item{
			Type:         "file",
			Name:         key[idxDelimiter+1:],
			Fullpath:     key,
			Size:         sprintSize(*(content.Size)),
			LastModified: content.LastModified.Format("2006-01-02 15:04:05"),
		})
	}

	return items, nil
}

// DownloadStream is a function that download a file in S3 and write body to parametered writer
func DownloadStream(sess *session.Session, bucket string, key string, w io.Writer) error {
	svc := s3.New(sess)

	fmt.Println(key)

	resp, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
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
		return fmt.Sprintf("%d Bytes", n)
	}
	if n < 1024*1024 {
		return fmt.Sprintf("%d KB", n)
	}
	if n < 1024*1024*1024 {
		return fmt.Sprintf("%d MB", n)
	}
	return fmt.Sprintf("%d GB", n)
}

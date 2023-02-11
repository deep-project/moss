package drivers

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"moss/infrastructure/persistent/storage"
)

type S3 struct {
	Endpoint  string `json:"endpoint"` // s3-us-west-1.amazonaws.com
	Region    string `json:"region"`   // us-west-1
	Bucket    string `json:"bucket"`
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
	Client    *s3.S3 `json:"-"`
}

func (s *S3) Init() (err error) {
	if s.Endpoint == "" {
		return errors.New("endpoint undefined")
	}
	if s.Region == "" {
		return errors.New("region undefined")
	}
	if s.Bucket == "" {
		return errors.New("bucket undefined")
	}
	if s.AccessKey == "" {
		return errors.New("accessKey undefined")
	}
	if s.SecretKey == "" {
		return errors.New("secretKey undefined")
	}
	conf := aws.Config{
		Credentials:      credentials.NewStaticCredentials(s.AccessKey, s.SecretKey, ""),
		Endpoint:         aws.String(s.Endpoint),
		Region:           aws.String(s.Region),
		S3ForcePathStyle: aws.Bool(true),
	}
	cliSession, err := session.NewSessionWithOptions(session.Options{Config: conf})
	if err != nil {
		return
	}
	s.Client = s3.New(cliSession)
	return
}

func (s *S3) Close() error {
	return nil
}

func (s *S3) Get(key string) (*storage.GetValue, error) {
	if s.Client == nil {
		return nil, errors.New("client undefined")
	}
	get := &s3.GetObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(key),
	}
	resp, err := s.Client.GetObjectWithContext(context.Background(), get)
	return storage.NewGetValue(resp.Body), err
}

func (s *S3) Set(key string, val *storage.SetValue) (err error) {
	if s.Client == nil {
		return errors.New("client undefined")
	}
	put := &s3.PutObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(key),
		Body:   val.Reader,
	}
	_, err = s.Client.PutObjectWithContext(context.Background(), put, request.WithGetResponseHeader("Content-Type", &val.ContentType))
	return
}

func (s *S3) Delete(key string) (err error) {
	if s.Client == nil {
		return errors.New("client undefined")
	}
	del := &s3.DeleteObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(key),
	}
	_, err = s.Client.DeleteObjectWithContext(context.Background(), del)
	return
}

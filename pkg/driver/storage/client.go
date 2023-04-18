package storage

import (
	"bytes"
	"cloud.google.com/go/storage"
	"context"
	"io"
)

type Storage struct {
	client *storage.Client
	ctx    context.Context
}

type Option struct {
	ContentEncoding string
}

func NewStorage(ctx context.Context) (*Storage, func() error, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, nil, err
	}
	return &Storage{
			client: client,
			ctx:    ctx,
		}, func() error {
			if err := client.Close(); err != nil {
				return err
			}
			return nil
		}, nil
}

func (s *Storage) UploadObject(bucketName string, objectName string, b []byte, option *Option) error {
	writer := s.client.Bucket(bucketName).Object(objectName).NewWriter(s.ctx)
	if option != nil {
		if option.ContentEncoding != "" {
			writer.ContentEncoding = option.ContentEncoding
		}
	}
	if _, err := io.Copy(writer, bytes.NewBuffer(b)); err != nil {
		return err
	}
	return nil
}

func (s *Storage) GetObject(bucketName string, objectName string) ([]byte, error) {
	reader, err := s.client.Bucket(bucketName).Object(objectName).NewReader(s.ctx)
	if err != nil {
		return nil, err
	}
	b, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	if err := reader.Close(); err != nil {
		return nil, err
	}
	return b, nil
}

package secret_manager

import (
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"context"
	"fmt"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
	"hash/crc32"
)

type SecretManager struct {
	*secretmanager.Client
	context.Context
}

func NewSecretManager(ctx context.Context) (*SecretManager, error) {
	client, err := createClient(ctx)
	if err != nil {
		return nil, err
	}
	return &SecretManager{Context: ctx, Client: client}, nil
}
func createClient(ctx context.Context) (*secretmanager.Client, error) {
	c, err := secretmanager.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	return c, nil
}
func (m *SecretManager) GetVersion(name string) (string, error) {
	req := &secretmanagerpb.AccessSecretVersionRequest{Name: name}
	result, err := m.Client.AccessSecretVersion(m.Context, req)
	if err != nil {
		return "", err
	}
	crc32c := crc32.MakeTable(crc32.Castagnoli)
	checkSum := int64(crc32.Checksum(result.Payload.Data, crc32c))
	if checkSum != *result.Payload.DataCrc32C {
		return "", fmt.Errorf(
			"data corruption detected. payload=%s payloadCRC32=%d, checkSum=%d",
			string(result.Payload.Data), *result.Payload.DataCrc32C, checkSum,
		)
	}
	return string(result.Payload.Data), nil
}

package secret_manager

import (
	"context"
	"fmt"
	"hash/crc32"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
)

type SecretManager struct {
	*secretmanager.Client
	context.Context
	projectID string
}

func NewSecretManager(ctx context.Context, projectIDNumber string) (*SecretManager, error) {
	client, err := createClient(ctx)
	if err != nil {
		return nil, err
	}
	return &SecretManager{Context: ctx, Client: client, projectID: projectIDNumber}, nil
}
func createClient(ctx context.Context) (*secretmanager.Client, error) {
	c, err := secretmanager.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	return c, nil
}
func (m *SecretManager) GetVersion(name string) (string, error) {
	req := &secretmanagerpb.AccessSecretVersionRequest{Name: fmt.Sprintf("projects/%s/secrets/%s/versions/latest", m.projectID, name)}
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

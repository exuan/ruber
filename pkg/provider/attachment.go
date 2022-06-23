package provider

import (
	"fmt"
)

const (
	TypeCos = "cos"
)

type (
	AttachmentConfig struct {
		Name      string `json:"name"`
		Type      string `json:"type"`
		SecretId  string `json:"secret_id"`
		SecretKey string `json:"secret_key"`
		BucketUrl string `json:"bucket_url"`
		Domain    string `json:"domain"`
	}

	Attachment interface {
		PresignedURL(name, ext string) (string, string, error)
	}
)

func NewAttachment(c *AttachmentConfig) (Attachment, error) {
	switch c.Type {
	case TypeCos:
		return NewCos(c)
	}

	return nil, fmt.Errorf("unknown type: %v", c.Type)
}

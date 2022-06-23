package provider

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/exuan/ruber/pkg/util"
	"github.com/google/uuid"
	client "github.com/tencentyun/cos-go-sdk-v5"
)

type cos struct {
	client    *client.Client
	SecretId  string
	SecretKey string
	Domain    string
}

func NewCos(c *AttachmentConfig) (Attachment, error) {
	u, _ := url.Parse(c.BucketUrl)
	b := &client.BaseURL{BucketURL: u}
	client := client.NewClient(b, &http.Client{
		Transport: &client.AuthorizationTransport{
			SecretID:  c.SecretId,
			SecretKey: c.SecretKey,
			Expire:    time.Minute,
		},
	})

	return &cos{
		client:    client,
		SecretId:  c.SecretId,
		SecretKey: c.SecretKey,
		Domain:    c.Domain,
	}, nil
}

func (c *cos) PresignedURL(name, ext string) (string, string, error) {
	if len(name) == 0 {
		name = util.Md5(uuid.NewString())
	}
	ctx := context.Background()
	f := fmt.Sprintf("%s.%s", name, ext)
	presignedURL, err := c.client.Object.GetPresignedURL(ctx, http.MethodPut, f, c.SecretId, c.SecretKey, time.Hour, nil)
	if err != nil {
		return "", "", err
	}

	return presignedURL.String(), c.Domain + "/" + f, nil
}

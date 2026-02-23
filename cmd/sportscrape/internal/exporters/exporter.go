package exporters

import (
	"context"
	"fmt"
	"net/url"
)

type Exporter[E any] interface {
	WriteParquet(ctx context.Context, record []E) error
	WriteJSONL(ctx context.Context, record []E) error
}

func SupportedDestination(destination *url.URL) error {
	scheme := destination.Scheme
	if scheme == "file" || scheme == "" || scheme == "s3" {
		return nil
	}
	return fmt.Errorf("unsupported destination type: %s. Only local and S3 supported", scheme)
}

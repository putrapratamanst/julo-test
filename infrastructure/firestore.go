package infrastructure

import (
	"context"
	"os"

	"cloud.google.com/go/firestore"
)

func CreateClient(ctx context.Context) (*firestore.Client, error) {
	// Sets your Google Cloud Platform project ID.
	projectID := os.Getenv("GOOGLE_PROJECT_ID")
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return &firestore.Client{}, err
	}

	return client, nil
}

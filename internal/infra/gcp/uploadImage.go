package gcp

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"time"
)

func UploadFile(object string, file multipart.File) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Cannot instantiate a client: %s", err)
	}

	defer client.Close()
	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	o := client.Bucket("si_images_unb").Object(object)
	o = o.If(storage.Conditions{DoesNotExist: true})

	wc := o.NewWriter(ctx)

	if _, err = io.Copy(wc, file); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %v", err)
	}

	return nil
}

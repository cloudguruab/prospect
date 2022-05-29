package controllers

import (
        "context"
        "fmt"
        "log"
        "time"

        "github.com/cloudguruab/prospect/common"
        "cloud.google.com/go/functions/metadata"
)


func TriggerGCS(ctx context.Context, e common.GCSEvent) error {
        meta, err := metadata.FromContext(ctx)

        if err != nil {
                return fmt.Errorf("metadata.FromContext: %v", err)
        }

        log.Printf("Event ID: %v\n", meta.EventID)
        log.Printf("Event type: %v\n", meta.EventType)
        log.Printf("Bucket: %v\n", e.Bucket)
        log.Printf("File: %v\n", e.Name)
        log.Printf("Metageneration: %v\n", e.Metageneration)
        log.Printf("Created: %v\n", e.TimeCreated)
        log.Printf("Updated: %v\n", e.Updated)

        return nil
}

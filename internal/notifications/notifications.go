package notifications

import (
	"context"
	"log"
	"server/internal/database"
	"time"
)

type NotificationsManager struct {
	dbManager *database.DBManager
}

func NewNotificationManager(dbManager *database.DBManager) *NotificationsManager {
	return &NotificationsManager{dbManager: dbManager}
}

func (nm *NotificationsManager) CheckEventsAndNotify(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			rows, err := nm.dbManager.GetOccuredEvents()
			if err != nil {
				log.Fatal(err)
			}
			for _, val := range rows {
				log.Printf("Email: %s\nEvent: %s\n", val.UserEmail, val.EventName)
			}
		case <-ctx.Done():
			return
		}
	}
}

package handlers

import (
	db "server/internal/database"
	not "server/internal/notifications"
)

type Server struct {
	dbManager     *db.DBManager
	NotifyManager *not.NotificationsManager
}

func NewServer() (*Server, error) {
	dbManager, err := db.NewManager()
	if err != nil {
		return nil, err
	}
	return &Server{dbManager: dbManager, NotifyManager: not.NewNotificationManager(dbManager)}, nil
}

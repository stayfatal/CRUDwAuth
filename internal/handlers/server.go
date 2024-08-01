package handlers

import (
	db "server/internal/database"
)

type Server struct {
	manager *db.DBManager
}

func NewServer() (*Server, error) {
	man, err := db.NewManager()
	if err != nil {
		return nil, err
	}
	return &Server{manager: man}, nil
}

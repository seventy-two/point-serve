package server

import (
	"context"

	"github.com/seventy-two/point-serve/database"
	points "github.com/seventy-two/point-serve/gen/go/point_server/v1"
)

type Server struct {
	db *database.Database
}

func NewServer(db *database.Database) *Server {
	return &Server{db: db}
}

func (s *Server) AddPoints(ctx context.Context, req *points.AddPointsRequest) (*points.AddPointsResponse, error) {
	u, err := s.db.GetUser(req.UserId)
	if err != nil {
		return nil, err
	}
	u.Points += req.Points
	err = s.db.SetUser(u)
	if err != nil {
		return nil, err
	}
	return &points.AddPointsResponse{
		UpdatedPoints: u.Points,
	}, nil
}

func (s *Server) RemovePoints(ctx context.Context, req *points.RemovePointsRequest) (*points.RemovePointsResponse, error) {
	u, err := s.db.GetUser(req.UserId)
	if err != nil {
		return nil, err
	}
	u.Points -= req.Points
	if u.Points < 0 {
		u.Points = 0
	}
	err = s.db.SetUser(u)
	if err != nil {
		return nil, err
	}
	return &points.RemovePointsResponse{
		UpdatedPoints: u.Points,
	}, nil
}

func (s *Server) GetPoints(ctx context.Context, req *points.GetPointsRequest) (*points.GetPointsResponse, error) {
	u, err := s.db.GetUser(req.UserId)
	if err != nil {
		return nil, err
	}
	return &points.GetPointsResponse{
		Points: u.Points,
	}, nil
}

func (s *Server) GetLeaderboard(ctx context.Context, req *points.GetLeaderboardRequest) (*points.GetLeaderboardResponse, error) {
	users, err := s.db.GetTopUsers(int(req.TopN))
	if err != nil {
		return nil, err
	}
	var leaderboard []*points.LeaderboardEntry
	for _, u := range users {
		leaderboard = append(leaderboard, &points.LeaderboardEntry{
			UserId: u.ID,
			Points: u.Points,
		})
	}
	return &points.GetLeaderboardResponse{
		Entries: leaderboard,
	}, nil
}

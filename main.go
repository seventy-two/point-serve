package main

import (
	"log"
	"net"

	cli "github.com/jawher/mow.cli"
	"github.com/seventy-two/point-serve/database"
	points "github.com/seventy-two/point-serve/gen/go/point_server/v1"
	"github.com/seventy-two/point-serve/server"
	"google.golang.org/grpc"
)

type appLink struct {
	name string
	url  string
}

var appMeta = struct {
	name        string
	description string
	discord     string
	maintainers string
	links       []appLink
}{
	name:        "Points Server",
	description: "A simple gRPC server that allows you to add and remove points from users",
	discord:     "https://discord.gg/F2cD4cN",
	maintainers: "github.com/seventy-two",
	links: []appLink{
		{name: "vcs", url: "https://github.com/seventy-two/points-server"},
	},
}

func main() {
	app := cli.App(appMeta.name, appMeta.description)

	dbPath := *app.String(cli.StringOpt{
		Name:   "DBPath",
		Value:  "/root/pointsDB",
		EnvVar: "DB_PATH",
	})

	if dbPath == "" {
		log.Fatalf("DB_PATH environment variable is not set")
	}

	db := database.NewDatabase(dbPath)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	points.RegisterPointServiceServer(s, server.NewServer(db))

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

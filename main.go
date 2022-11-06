package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math"
	"net"

	pb "github.com/yhjeon92/gpx-analyzer/gpx"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 8080, "port number")
)

type server struct {
	pb.UnsafeAnalyzerServer
}

func haversine(lat1 float32, lon1 float32, ele1 float32, lat2 float32, lon2 float32, ele2 float32) float64 {
	const R = 6371000
	phi1, phi2 := float64(lat1*math.Pi/180), float64(lat2*math.Pi/180)
	phiDelta, lambdaDelta := float64((lat2-lat1)*math.Pi/180), float64((lon2-lon1)*math.Pi/180)
	a := math.Pow(math.Sin(phiDelta/2), 2) + math.Cos(phi1)*math.Cos(phi2)*math.Pow(math.Sin(lambdaDelta/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	baseline := R * c
	elevation := float64(ele2 - ele1)

	return math.Sqrt(baseline*baseline + elevation*elevation)
}

func (s *server) AnalyzeWorkout(ctx context.Context, workout *pb.Workout) (*pb.WorkoutResultView, error) {

	return &pb.WorkoutResultView{Name: "Hello!"}, nil
	//return nil, nil
}

func main() {
	fmt.Println("Hello, World!")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Println(port)
	fmt.Println(*port)
	fmt.Println(&port)

	s := grpc.NewServer()
	analyzerServer := server{}
	pb.RegisterAnalyzerServer(s, &analyzerServer)
	log.Printf("server up and running at: %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

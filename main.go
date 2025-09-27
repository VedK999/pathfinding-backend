package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	. "pathfinding-backend/algorithms"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// type Node struct {
// 	Row      int     `json:"row"`
// 	Col      int     `json:"col"`
// 	IsWall   bool    `json:"isWall"`
// 	Distance float64 `json:"distance,omitempty"`
// 	Visited  bool    `json:"visited,omitempty"`
// 	Previous *Node   `json:"previous,omitempty"`
// 	FScore   float64 `json:"fscore,omitempty"`
// 	GScore   float64 `json:"gscore,omitempty"`
// }

type PathFindRequest struct {
	Grid      [][]Node `json:"grid"`
	Start     Node     `json:"start"`
	End       Node     `json:"end"`
	Algorithm string   `json:"algorithm"`
	GridSize  int      `json:"gridsize"`
}

type PathFindResponse struct {
	Success       bool    `json:"success"`
	Path          []Node  `json:"path"`
	VisitedNodes  []Node  `json:"visitedNodes"`
	ExecutionTime float64 `json:"executionTime"`
	Message       string  `json:"message,omitempty"`
}

// func enableCors(w http.ResponseWriter) {
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-control-Allow-Methods", "POST,GET,OPTIONS,PUT,DELETE")
// 	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
// }

func pathfinderHandler(w http.ResponseWriter, r *http.Request) {
	//enableCors(w)

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req PathFindRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// if len(req.Start) != 2 || len(req.End) != 2 {
	// 	http.Error(w, "Start and end points must have exactly two coordinates", http.StatusBadRequest)
	// 	return
	// }

	var path []Node
	var visited []Node
	var success bool
	//var executionTime float64

	startTime := time.Now()

	switch req.Algorithm {
	case "astar":
		path, visited, success = AStar(req.Grid, req.Start, req.End)
	case "bfs":
		path, visited, success = BFS(req.Grid, req.Start, req.End)
	case "dijkstra":
		path, visited, success = Dijkstra(req.Grid, req.Start, req.End)
	default:
		http.Error(w, "Undefied Algorithm", http.StatusBadRequest)
	}

	executionTime := float64(time.Since(startTime).Nanoseconds()) / 1000000.0

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }

	response := PathFindResponse{
		Path:          path,
		VisitedNodes:  visited,
		Success:       success,
		ExecutionTime: executionTime,
	}

	if !success {
		response.Message = "NO PATH FOUND"
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func main() {
	// http.HandleFunc("/pathfind", pathfinderHandler)
	// log.Println("Server starting on port 8080...")
	// log.Fatal(http.ListenAndServe(":8080", nil))

	r := mux.NewRouter()

	r.PathPrefix("/").Subrouter().HandleFunc("/pathfind", pathfinderHandler).Methods("POST")

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{
		"http://localhost:3000",
		"http://localhost:5173",
		"https://*.netlify.app",
		"https://netlify.app"})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server starting on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(headers, methods, origins)(r)))
}

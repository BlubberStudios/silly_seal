package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"seal-ascii/animations"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

const (
	clearScreen = "\033[2J\033[H"
	port        = ":8081"
)

func main() {
	r := mux.NewRouter()
	
	// Animation streaming endpoints
	r.HandleFunc("/{animation}", streamAnimation).Methods("GET")
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		streamAnimation(w, r) // Default to seal animation
	}).Methods("GET")
	
	// API endpoint for listing available animations
	r.HandleFunc("/list", listAnimations).Methods("GET")
	
	fmt.Printf("🦭 Seal ASCII Animation Server running on http://localhost%s\n", port)
	fmt.Println("Try: curl http://localhost:8080")
	fmt.Println("Or visit: http://localhost:8080/list for available animations")
	
	log.Fatal(http.ListenAndServe(port, r))
}

func streamAnimation(w http.ResponseWriter, r *http.Request) {
	// Detect if request is from curl or similar terminal client
	userAgent := r.Header.Get("User-Agent")
	isCurl := strings.Contains(strings.ToLower(userAgent), "curl") || 
			  strings.Contains(strings.ToLower(userAgent), "wget") ||
			  strings.Contains(strings.ToLower(userAgent), "httpie")
	
	if !isCurl {
		// Serve HTML page for browsers
		w.Header().Set("Content-Type", "text/html")
		html := `<!DOCTYPE html>
<html>
<head>
    <title>Seal ASCII Animation</title>
    <style>
        body { background: #000; color: #0f0; font-family: monospace; margin: 0; padding: 20px; }
        .container { text-align: center; }
        pre { font-size: 12px; line-height: 1; white-space: pre; }
    </style>
</head>
<body>
    <div class="container">
        <h1>🦭 Seal ASCII Animation</h1>
        <p>Use curl to view the animation:</p>
        <code>curl http://` + r.Host + `</code>
        <p><a href="/list">View available animations</a></p>
    </div>
</body>
</html>`
		fmt.Fprint(w, html)
		return
	}
	
	// Get animation name from URL path
	vars := mux.Vars(r)
	animationName := vars["animation"]
	if animationName == "" {
		animationName = "seal" // Default animation
	}
	
	// Get animation from frame map
	animation, exists := animations.FrameMap[animationName]
	if !exists {
		http.Error(w, fmt.Sprintf("Animation '%s' not found", animationName), http.StatusNotFound)
		return
	}
	
	// Set headers for streaming
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Transfer-Encoding", "chunked")
	w.Header().Set("Cache-Control", "no-cache")
	
	// Get flusher for immediate response
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming not supported", http.StatusInternalServerError)
		return
	}
	
	// Get close notifier to detect client disconnect
	closeNotifier, ok := w.(http.CloseNotifier)
	if !ok {
		log.Println("Warning: Close notification not supported")
	}
	
	// Animation loop
	frameIndex := 0
	ticker := time.NewTicker(animation.GetSleep())
	defer ticker.Stop()
	
	for {
		select {
		case <-ticker.C:
			// Clear screen and display current frame
			fmt.Fprint(w, clearScreen)
			fmt.Fprint(w, animation.GetFrame(frameIndex))
			fmt.Fprint(w, "\n")
			
			flusher.Flush()
			
			// Move to next frame
			frameIndex = (frameIndex + 1) % animation.GetLength()
			
		case <-closeNotifier.CloseNotify():
			// Client disconnected
			return
		}
	}
}

func listAnimations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	animationList := make([]string, 0, len(animations.FrameMap))
	for name := range animations.FrameMap {
		animationList = append(animationList, name)
	}
	
	response := map[string]interface{}{
		"animations": animationList,
		"usage":      "curl http://" + r.Host + "/{animation}",
		"example":    "curl http://" + r.Host + "/seal",
	}
	
	json.NewEncoder(w).Encode(response)
}
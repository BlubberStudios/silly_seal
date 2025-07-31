package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const (
	clearScreen = "\033[2J\033[H"
)

// Selected 24 frames out of 187 total frames
var SealWiggleFrames = []string{
	`⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⢀⢀⡀⣀⢀⡀⣀⢀⡀⣀⢀⡀⣀⢀⡀⣀⢀⡀⣀⢀⡀⣀⢀⡀⣀⢀⡀⣀⢀⡀⣀⢀⡀⣀⢀⡀⣀⢀⡀⣀⢀⡀⣀⢀⡀⣀⢀⡀⣀⢀⡀⣀⢀⡀⣀⢀⡀⣀⢀⡀⣀⢀⢀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⢰⢕⢗⢕⢗⢝⢎⢗⢝⢎⢗⢝⢎⢗⢝⢎⢗⢝⢎⢗⢝⢎⢗⢝⢎⢗⢝⢎⢗⢝⢎⢗⢝⢎⢗⢝⢎⢗⢝⢎⢗⢝⢎⢗⢝⢎⢗⢝⢎⢗⢝⢎⢗⢝⢎⢗⢝⢎⢗⢝⢎⢗⢝⢎⢗⡳⡕⣕⡄⠄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`,
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// Get the path from query parameter (set by Vercel rewrite) or URL path
	path := r.URL.Query().Get("path")
	if path == "" {
		path = strings.TrimPrefix(r.URL.Path, "/silly_seal")
		if path == "" {
			path = "/"
		}
	}

	// Handle list endpoint
	if path == "/list" || path == "list" {
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{
			"animations": []string{"seal"},
			"usage":      "curl https://blubberstudios.com/silly_seal",
			"example":    "curl https://blubberstudios.com/silly_seal",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

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
    <title>🦭 Seal ASCII Animation</title>
    <style>
        body { 
            background: #000; 
            color: #0f0; 
            font-family: 'Courier New', monospace; 
            margin: 0; 
            padding: 20px; 
            text-align: center;
        }
        .container { max-width: 800px; margin: 0 auto; }
        pre { 
            font-size: 14px; 
            line-height: 1.2; 
            white-space: pre; 
            background: #111;
            padding: 20px;
            border-radius: 8px;
            margin: 20px 0;
        }
        code {
            background: #222;
            padding: 4px 8px;
            border-radius: 4px;
            color: #ff6;
        }
        h1 { color: #4af; }
        a { color: #4af; }
    </style>
</head>
<body>
    <div class="container">
        <h1>🦭 Wiggling Seal Animation</h1>
        <p>Use curl to view the animated seal:</p>
        <pre><code>curl https://blubberstudios.com/silly_seal</code></pre>
        <p>Or try:</p>
        <pre><code>curl https://blubberstudios.com/silly_seal/list</code></pre>
        <p>Made with ❤️ inspired by parrot.live</p>
    </div>
</body>
</html>`
		fmt.Fprint(w, html)
		return
	}

	// Set headers for streaming
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")

	// Get flusher for immediate response (if available)
	flusher, ok := w.(http.Flusher)
	if !ok {
		// Fallback for environments that don't support flushing
		// Show a few frames in sequence
		for i := 0; i < 3; i++ {
			for _, frame := range SealWiggleFrames {
				fmt.Fprint(w, clearScreen)
				fmt.Fprint(w, frame)
				fmt.Fprint(w, "\n\n")
				time.Sleep(200 * time.Millisecond)
			}
		}
		return
	}

	// Animation loop with streaming
	frameIndex := 0
	maxDuration := 30 * time.Second // Limit animation time for serverless
	startTime := time.Now()
	ticker := time.NewTicker(150 * time.Millisecond) // Slightly slower for web
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if time.Since(startTime) > maxDuration {
				fmt.Fprint(w, "\n\n🦭 Thanks for watching! Run the command again for more seal wiggling!")
				return
			}

			// Clear screen and display current frame
			fmt.Fprint(w, clearScreen)
			fmt.Fprint(w, SealWiggleFrames[frameIndex])
			fmt.Fprint(w, "\n\n🦭 Wiggling seal! Press Ctrl+C to stop\n")

			flusher.Flush()

			// Move to next frame
			frameIndex = (frameIndex + 1) % len(SealWiggleFrames)
		}
	}
}
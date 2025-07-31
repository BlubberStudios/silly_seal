package main

import (
	"fmt"
	"seal-ascii/animations"
)

func main() {
	sealAnimation := animations.FrameMap["seal"]
	fmt.Printf("Total frames: %d\n", sealAnimation.GetLength())
	fmt.Printf("Sleep duration: %v\n", sealAnimation.GetSleep())
	
	// Show first few frames to verify content
	for i := 0; i < min(3, sealAnimation.GetLength()); i++ {
		frame := sealAnimation.GetFrame(i)
		fmt.Printf("\n--- Frame %d (first 3 lines) ---\n", i+1)
		lines := splitLines(frame)
		for j := 0; j < min(3, len(lines)); j++ {
			fmt.Printf("%s\n", lines[j])
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func splitLines(s string) []string {
	lines := []string{}
	current := ""
	for _, char := range s {
		if char == '\n' {
			lines = append(lines, current)
			current = ""
		} else {
			current += string(char)
		}
	}
	if current != "" {
		lines = append(lines, current)
	}
	return lines
}
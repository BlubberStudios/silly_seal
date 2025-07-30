package animations

import "time"

// FrameType defines the interface for animation frames
type FrameType struct {
	frames []string
	sleep  time.Duration
}

// GetFrame returns the frame at the specified index with bounds checking
func (f *FrameType) GetFrame(index int) string {
	if index < 0 || index >= len(f.frames) {
		return f.frames[0] // Return first frame if index out of bounds
	}
	return f.frames[index]
}

// GetLength returns the total number of frames
func (f *FrameType) GetLength() int {
	return len(f.frames)
}

// GetSleep returns the duration to sleep between frames
func (f *FrameType) GetSleep() time.Duration {
	return f.sleep
}

// DefaultFrameType creates a new FrameType with default 70ms timing
func DefaultFrameType(frames []string) *FrameType {
	return &FrameType{
		frames: frames,
		sleep:  70 * time.Millisecond,
	}
}

// NewFrameType creates a new FrameType with custom timing
func NewFrameType(frames []string, sleep time.Duration) *FrameType {
	return &FrameType{
		frames: frames,
		sleep:  sleep,
	}
}

// FrameMap registry for all available animations
var FrameMap = map[string]*FrameType{
	"seal": DefaultFrameType(SealWiggleFrames),
}
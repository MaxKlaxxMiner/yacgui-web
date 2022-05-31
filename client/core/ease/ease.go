package ease

import "math"

// --- based on JQuery ---

// Cap Capping the Value between 0 and 1
func Cap(t float64) float64 {
	if t < 0 {
		return 0
	}
	if t > 1 {
		return 1
	}
	return t
}

// Linear no easing, no acceleration
func Linear(t float64) float64 {
	return Cap(t)
}

// InQuad accelerating from zero velocity
func InQuad(t float64) float64 {
	t = Cap(t)
	return t * t
}

// OutQuad decelerating to zero velocity
func OutQuad(t float64) float64 {
	t = Cap(t)
	return t * (2 - t)
}

// InOutQuad acceleration until halfway, then deceleration
func InOutQuad(t float64) float64 {
	t = Cap(t)
	if t < 0.5 {
		return 2 * t * t
	} else {
		return -1 + (4-2*t)*t
	}
}

// InCubic accelerating from zero velocity
func InCubic(t float64) float64 {
	t = Cap(t)
	return t * t * t
}

// OutCubic decelerating to zero velocity
func OutCubic(t float64) float64 {
	t = Cap(t)
	t -= 1
	return t*t*t + 1
}

// InOutCubic acceleration until halfway, then deceleration
func InOutCubic(t float64) float64 {
	t = Cap(t)
	if t < 0.5 {
		return 4 * t * t * t
	} else {
		return (t-1)*(2*t-2)*(2*t-2) + 1
	}
}

// InQuart accelerating from zero velocity
func InQuart(t float64) float64 {
	t = Cap(t)
	return t * t * t * t
}

// OutQuart decelerating to zero velocity
func OutQuart(t float64) float64 {
	t = Cap(t)
	t -= 1
	return 1 - t*t*t*t
}

// InOutQuart acceleration until halfway, then deceleration
func InOutQuart(t float64) float64 {
	t = Cap(t)
	if t < 0.5 {
		return 8 * t * t * t * t
	} else {
		t -= 1
		return 1 - 8*t*t*t*t
	}
}

// InQuint accelerating from zero velocity
func InQuint(t float64) float64 {
	t = Cap(t)
	return t * t * t * t * t
}

// OutQuint decelerating to zero velocity
func OutQuint(t float64) float64 {
	t = Cap(t)
	t -= 1
	return 1 + t*t*t*t*t
}

// InSine accelerating from zero velocity
func InSine(t float64) float64 {
	t = Cap(t)
	return 1 - math.Cos(math.Pi/2*t)
}

// OutSine decelerating to zero velocity
func OutSine(t float64) float64 {
	t = Cap(t)
	return math.Sin(math.Pi / 2 * t)
}

// InOutSine JQuery default: acceleration until halfway, then deceleration
func InOutSine(t float64) float64 {
	t = Cap(t)
	return -(math.Cos(math.Pi*t) - 1) / 2
}

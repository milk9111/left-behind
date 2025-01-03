package engine

import (
	"math"
	"math/rand"
	"time"

	"github.com/yohamta/donburi"
	dmath "github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

func ScaledPosition(e *donburi.Entry) dmath.Vec2 {
	return transform.WorldPosition(e).Mul(transform.WorldScale(e))
}

func IndexToVec2(col, row int) dmath.Vec2 {
	return dmath.NewVec2(float64(col*32), float64(row*32))
}

func Vec2ToIndex(vec2 dmath.Vec2) (int, int) {
	return int(vec2.X / 32), int(vec2.Y / 32)
}

func Vec2InverseScalar(vec2 dmath.Vec2, scalar float64) dmath.Vec2 {
	return dmath.NewVec2(scalar/vec2.X, scalar/vec2.Y)
}

// PoissonDiskSampling generates points with a minimum distance r
func PoissonDiskSampling(width, height, r float64, k int) []dmath.Vec2 {
	// Grid size (cell size is r / sqrt(2))
	cellSize := r / math.Sqrt2
	cols := int(math.Ceil(width / cellSize))
	rows := int(math.Ceil(height / cellSize))

	// Grid to store points
	grid := make([][]*dmath.Vec2, rows)
	for i := range grid {
		grid[i] = make([]*dmath.Vec2, cols)
	}

	// Active list and output points
	active := []dmath.Vec2{}
	points := []dmath.Vec2{}

	// Helper to check if a point is valid
	isValid := func(p dmath.Vec2) bool {
		if p.X < 0 || p.X >= width || p.Y < 0 || p.Y >= height {
			return false
		}

		// Get the cell indices
		col := int(p.X / cellSize)
		row := int(p.Y / cellSize)

		// Check neighboring cells
		for i := -2; i <= 2; i++ {
			for j := -2; j <= 2; j++ {
				neighborCol := col + i
				neighborRow := row + j
				if neighborCol >= 0 && neighborCol < cols && neighborRow >= 0 && neighborRow < rows {
					neighbor := grid[neighborRow][neighborCol]
					if neighbor != nil && distance(p, *neighbor) < r {
						return false
					}
				}
			}
		}
		return true
	}

	// Seed random and initialize with a random point
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Helper to generate a random point around another point
	generateRandomPointAround := func(center dmath.Vec2) dmath.Vec2 {
		radius := r * (1 + seededRand.Float64())
		angle := 2 * math.Pi * seededRand.Float64()
		return dmath.Vec2{
			X: center.X + radius*math.Cos(angle),
			Y: center.Y + radius*math.Sin(angle),
		}
	}

	initialPoint := dmath.Vec2{X: seededRand.Float64() * width, Y: seededRand.Float64() * height}
	active = append(active, initialPoint)
	points = append(points, initialPoint)

	// Place the initial point in the grid
	grid[int(initialPoint.Y/cellSize)][int(initialPoint.X/cellSize)] = &initialPoint

	// Process the active list
	for len(active) > 0 {
		// Pick a random active point
		idx := seededRand.Intn(len(active))
		point := active[idx]

		// Generate up to k candidate points
		found := false
		for i := 0; i < k && !found; i++ {
			candidate := generateRandomPointAround(point)
			if isValid(candidate) {
				// Add the candidate to the grid, points, and active list
				points = append(points, candidate)
				active = append(active, candidate)
				grid[int(candidate.Y/cellSize)][int(candidate.X/cellSize)] = &candidate
				found = true
			}
		}

		// If no valid candidates, remove the point from the active list
		if !found {
			active = append(active[:idx], active[idx+1:]...)
		}
	}

	return points
}

// Helper function to calculate the Euclidean distance
func distance(a, b dmath.Vec2) float64 {
	dx := a.X - b.X
	dy := a.Y - b.Y
	return math.Sqrt(dx*dx + dy*dy)
}

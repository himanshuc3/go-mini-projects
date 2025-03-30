package solver

import (
	"fmt"
	"image"
	"log"
)

// NOTE:
// 1. Any packages inside internal would not be exposed
// outside the module

type Solver struct {
	maze    *image.RGBA
	palette palette
}

func (s *Solver) Solve() error {
	entrance, err := s.findEntrance()
	if err != nil {
		return fmt.Errorf("unable to find entrance: %w", err)
	}

	log.Printf("starting at %v", entrance)
	return nil
}

func New(imagePath string) (*Solver, error) {
	img, err := openMaze(imagePath)

	if err != nil {
		return nil, fmt.Errorf("cannot open maze image: %w", err)
	}

	return &Solver{
		maze:    img,
		palette: defaultPallete(),
	}, nil
}

func (s *Solver) findEntrance() (image.Point, error) {
	// NOTE:
	// 1. Bounds, RGBAAt - methods attached to image.RGBA grid
	for row := s.maze.Bounds().Min.Y; row < s.maze.Bounds().Max.Y; row++ {
		for col := s.maze.Bounds().Min.X; col < s.maze.Bounds().Max.X; col++ {
			if s.maze.RGBAAt(col, row) == s.palette.entrance {
				return image.Point{X: col, Y: row}, nil
			}
		}
	}
	return image.Point{}, fmt.Errorf("entrace position not found")
}

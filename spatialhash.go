package main

type SpatialHash struct {
	CellSize   float32
	Cols, Rows int
	Buckets    [][]int
}

func NewSpatialHash(w, h int, cellSize float32) *SpatialHash {
	cols := int(w/int(cellSize)) + 1
	rows := int(h/int(cellSize)) + 1
	buckets := make([][]int, cols*rows)
	for i := range buckets {
		buckets[i] = make([]int, 0, 64)
	}
	return &SpatialHash{cellSize, cols, rows, buckets}
}

func (s *SpatialHash) index(x, y float32) int {
	cx := int(x / s.CellSize)
	cy := int(y / s.CellSize)
	if cx < 0 {
		cx = 0
	} else if cx >= s.Cols {
		cx = s.Cols - 1
	}
	if cy < 0 {
		cy = 0
	} else if cy >= s.Rows {
		cy = s.Rows - 1
	}
	return cy*s.Cols + cx
}

func (s *SpatialHash) RebuildHash(agents []Agent) error {
	for i := range s.Buckets {
		s.Buckets[i] = s.Buckets[i][:0]
	}
	for i, a := range agents {
		idx := s.index(a.X, a.Y)
		s.Buckets[idx] = append(s.Buckets[idx], i)
	}
	return nil
}

func (s *SpatialHash) Neighbors(x, y float32) []int {
	cx := int(x / s.CellSize)
	cy := int(y / s.CellSize)
	out := make([]int, 0, 64)
	for dy := -1; dy <= 1; dy++ {
		yy := cy + dy
		if yy < 0 || yy >= s.Rows {
			continue
		}
		for dx := -1; dx <= 1; dx++ {
			xx := cx + dx
			if xx < 0 || xx >= s.Cols {
				continue
			}
			out = append(out, s.Buckets[yy*s.Cols+xx]...)
		}
	}
	return out
}


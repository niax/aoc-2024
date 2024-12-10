package collections

type Point2D struct {
	X int
	Y int
}

var (
	Point2D_North = Point2D{X: 0, Y: -1}
	Point2D_East  = Point2D{X: 1, Y: 0}
	Point2D_South = Point2D{X: 0, Y: 1}
	Point2D_West  = Point2D{X: -1, Y: 0}

	Point2D_CardinalDirections = []Point2D{
		Point2D_North,
		Point2D_East,
		Point2D_South,
		Point2D_West,
	}
)

func (p Point2D) Add(other Point2D) Point2D {
	return Point2D{
		X: p.X + other.X,
		Y: p.Y + other.Y,
	}
}

package main

type CollisionType int

const (
	None  CollisionType = 0
	Above               = 1
	Below               = 2
	Left                = 4
	Right               = 8
	All                 = 15
)

type Collider interface {
	ColliderType() CollisionType
	CollisionBox() Rect
	PrevCollisionBox() Rect
}

func Collision(a, b Collider) CollisionType {
	a1 := a.PrevCollisionBox()
	a2 := a.CollisionBox()
	b1 := b.CollisionBox()

	// there is a collision, find the type
	if !contains(a1, b1) && contains(a2, b1) {
		if (a1.X+a1.W) < b1.X && (a2.X+a2.W) >= b1.X {
			return Left & b.ColliderType()
		} else if a1.X > (b1.X+b1.W) && a2.X <= (b1.X+b1.W) {
			return Right & b.ColliderType()
		} else if (a1.Y+a1.H) < b1.Y && (a2.Y+a2.H) >= b1.Y {
			return Above & b.ColliderType()
		} else if a1.Y > (b1.Y+b1.H) && a2.Y <= (b1.Y+b1.H) {
			return Below & b.ColliderType()
		}
	}

	return None
}

func contains(a, b Rect) bool {
	outsideX := (a.X+a.W) < b.X || a.X > (b.X+b.W)
	outsideY := (a.Y+a.H) < b.Y || a.Y > (b.Y+b.H)

	return !outsideX && !outsideY
}

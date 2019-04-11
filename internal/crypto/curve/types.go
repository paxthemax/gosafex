package curve

// PointLength is the length of the point (in bytes).
const PointLength = 32

// Point is a point on the elliptic curve. A point is analogous to a key.
type Point = [PointLength]byte

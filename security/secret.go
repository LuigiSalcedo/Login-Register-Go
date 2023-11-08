package security

// The Secret is in plain text just for Academic Purpose. In the practice
// this is a security problem.
var jwtSecret = "Academic-Purpose"

func Secret() string {
	return jwtSecret
}

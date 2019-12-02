package bo

// UserQueryBO defines table `user` query conditions
type UserQueryBO struct {
	Limit   *int
	Offset  *int
	Name    *string
	PhoneNo *string
}

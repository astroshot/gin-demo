package bo

// UserQueryBO defines table `user` query conditions
type UserQueryBO struct {
	PageNo   *int
	PageSize *int
	Name     *string
	PhoneNo  *string
}

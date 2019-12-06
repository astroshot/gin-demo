package dao

var (
	UserDao UserDAO
)

// InitDAO init package dao
func InitDAO() {
	once.Do(initDB)
	UserDao = &UserDAOImpl{}
}

package dao

var (
	UserDao *UserDAOImpl
)

func InitDAO() {
	once.Do(initDB)
	UserDao = &UserDAOImpl{}
	db := GetDB()
	UserDao.SetDB(db)
}

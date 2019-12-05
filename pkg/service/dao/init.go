package dao

var (
	UserDao *UserDAOImpl
)

func InitDAO() {
	Init()
	UserDao = &UserDAOImpl{}
	db := GetDB()
	UserDao.SetDB(db)
}

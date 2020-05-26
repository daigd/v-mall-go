package datasoure

// DataEngine 数据引擎
type DataEngine uint8

const (
	// Memory 数据从内存加载
	Memory DataEngine = iota
	// Mysql 数据从Mysql加载
	Mysql
)

// DS 数据源
type DS struct {
	// db *gorm.DB
}

// LoadMemoryData 从内存中加载数据
// func LoadMemoryData() (ds *DS, err error) {
// 	// 这会告诉 sqlite 使用内存作为一个临时数据
// 	db, err := gorm.Open("sqlite3", ":memory:")
// 	fmt.Println("生成的数据库", db)
// 	ds = &DS{db: db}
// 	return
// }

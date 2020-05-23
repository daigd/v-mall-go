package repository

// Query represents the visitor and action queries.
// Query 用来查询数据,充当数据的访问者和操作者
type Query func(model interface{}) bool

const (
	// ReadOnlyMode 数据访问使用只读模式
	ReadOnlyMode = iota
	// ReadWriteMode 数据访问使用读写模式
	ReadWriteMode
)

// BaseRepository 提供基础函数定义
type BaseRepository interface {
	// 按给定的条件来执行数据操作
	Exec(query Query, action Query, limit int, mode int) (ok bool)
	// 按给定的查询条件选择数据
	Select(query Query) (datamodel interface{}, found bool)
}

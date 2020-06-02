package service

// BaseService 基础服务
type BaseService interface {
	Create(value interface{}) (err error)
}

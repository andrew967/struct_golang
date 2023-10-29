package services

type HashTable interface {
	Add(key string, value interface{})
	Get(key string) (value interface{}, err error)
	Remove(key string) error
}

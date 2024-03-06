package hashtable

import (
	"fmt"
	"sync"
)

type HashTableStore struct {
	data sync.Map
}

// NewHashTableStore - конструктор для создания нового объекта HashTableStore
func NewHashTableStore() *HashTableStore {
	return &HashTableStore{
		data: sync.Map{},
	}
}

// AddKeyValue - метод для добавления ключа и значения
func (ht *HashTableStore) AddKeyValue(key, value string) {
	ht.data.Store(key, value)
}

// GetValueByKey - метод для получения значения по ключу
func (ht *HashTableStore) GetValueByKey(key string) (string, bool) {
	value, ok := ht.data.Load(key)
	if ok {
		return value.(string), true
	}
	return "", false
}

// GetKeyByValue - метод для получения ключа по значению
func (ht *HashTableStore) GetKeyByValue(value string) (string, bool) {
	var foundKey string
	found := false

	ht.data.Range(func(key, val interface{}) bool {
		if val.(string) == value {
			foundKey = key.(string)
			found = true
			return false // Прерываем итерацию
		}
		return true
	})

	return foundKey, found
}

// PrintHashTableStore - метод для вывода в консоль всех ключей
// и значений в хранилище
func (ht *HashTableStore) PrintHashTableStore() {
	ht.data.Range(func(key, value interface{}) bool {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
		return true
	})
}
package trymap


import (
	"fmt"
	"sync"
)

type TMap[K comparable, V any] struct{ m sync.Map }

func (t *TMap[K, V]) Store(k K, v V) { t.m.Store(k, v) }
func (t *TMap[K, V]) Load(k K) (V, bool) {
	if v, ok := t.m.Load(k); ok {
		return v.(V), true
	}
	var zero V
	return zero, false
}

type User struct {
	Name string
	Age  int
}

func TryTMap() {
	var users TMap[string, *User]
	users.Store("alice", &User{Name: "Alice", Age: 20})
	// we cannot use int because var users TMap[string, *User] has already assert the type
	// users.Store(20, 3)
	if u, ok := users.Load("alice"); ok {
		fmt.Printf("%+v\n", u)
	}
}

func TrySyncMap() {
	var m sync.Map
	m.Store("id", 123)
	var id int
	// 讀取：斷言失敗會 panic，所以通常先檢查 ok
	if v, ok := m.Load("id"); ok {
		id = v.(int) // 型別斷言
		fmt.Println(id)
	}
}


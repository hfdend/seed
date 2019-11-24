package cli

import (
	"seed/conf"
	"sync"
)

var once sync.Once

func Init() {
	once.Do(func() {
		conf.Init()
		InitializeLogger()
		InitializeDB()
	})
}

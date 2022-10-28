package simpledb

import (
	"fmt"
)

func Find[T any](entity T) T {
	// reflect.ValueOf(entity).Elem()
	fmt.Println(entity)

	return entity
}

func FindByPrimaryKey(entity any) any {
	fmt.Println(entity)

	return nil
}

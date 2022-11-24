package simpledb

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

func Find[T any](db string, entity T, key string) (T, bool) {
	data, ok := FindByPrimaryKey(db, key)
	if !ok {
		return entity, false
	}

	err := json.Unmarshal(data, &entity)
	if err != nil {
		fmt.Printf("unmarshal json failed: %s", err)
		return entity, false
	}

	return entity, true
}

func FindByPrimaryKey(db string, key string) ([]byte, bool) {
	f, err := os.Open(db)
	if err != nil {
		fmt.Println("err to open file")
	}
	defer f.Close()

	r := bufio.NewReader(f)

	var line []byte
	var rs []byte
	var ok bool

	for !ok {
		line, _, err = r.ReadLine()
		if err != nil {
			break
		}

		rs, ok = findBytesByPrimaryKey(line, []byte(key), rs, ok)
	}

	return rs, ok
}

func findBytesByPrimaryKey(line []byte, key []byte, rs []byte, ok bool) ([]byte, bool) {
	isContainKey := bytes.Contains(line[:bytes.Index(line, []byte(":"))], []byte(key))
	if isContainKey {
		rs, ok = splitBytesBySubbytes(line, []byte(":"))
	}
	return rs, ok
}

func splitBytesBySubbytes(line []byte, subbytes []byte) ([]byte, bool) {
	i := bytes.Index(line, subbytes)
	if i == -1 {
		return []byte(""), false
	}

	substr := bytes.TrimSpace(line[i+1:])
	return substr[:len(substr)-1], true
}

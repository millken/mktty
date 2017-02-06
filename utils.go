package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func randByte(n int) []byte {
	b := make([]byte, n)
	f, err := os.OpenFile("/dev/urandom", os.O_RDONLY, 0)
	if err != nil {
		_, err := rand.Read(b)
		if err != nil {
			log.Printf("[ERROR] rand.Read : %s", err)
		}
	} else {
		f.Read(b)
		f.Close()
	}
	return b
}
func guid() string {
	b := randByte(16)
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}

func randToken() string {
	b := randByte(16)
	return fmt.Sprintf("%x", b)
}

func randUserName() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := r.Intn(100000)
	return fmt.Sprintf("用户%d", n)
}

func strToInt(str string) int {
	var i int
	i, err = strconv.Atoi(str)
	if err != nil {
		i = 0
	}
	return i
}

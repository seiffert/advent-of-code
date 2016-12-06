package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func GetPassword(input string) string {
	pwd := [8]byte{'_', '_', '_', '_', '_', '_', '_', '_'}
	var chars int
	for i := 0; chars < 8; i++ {
		m := md5.New()
		io.WriteString(m, input+strconv.Itoa(i))
		hash := fmt.Sprintf("%x", m.Sum(nil))

		if string(hash)[:5] == "00000" {
			log.Println(string(hash))
			pos, err := strconv.Atoi(string(hash[5]))
			if err == nil && pos < 8 && pwd[pos] == '_' {
				pwd[pos] = hash[6]
				chars++
				log.Println(string(pwd[:8]))
			}
		}
	}

	return string(pwd[:8])
}

func main() {
	fmt.Println(GetPassword(os.Args[1]))
}

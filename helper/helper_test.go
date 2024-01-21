package helper

import (
	"fmt"
	"testing"
)

func TestEncrypt(t *testing.T) {
	encrypt, err := Encrypt("2", "abc&1*~#^2^#s0^=)^^7%b34")
	fmt.Println(err, encrypt)

	decrypt, err1 := Decrypt("WQ==", "abc&1*~#^2^#s0^=)^^7%b34")
	fmt.Println(err1, decrypt)

}

func TestName(t *testing.T) {
	list := []int{1, 2, 3, 4}

	f := func(item int) bool {
		return item == 3
	}
	result := Filter3(list, f)

	fmt.Printf("%v", result)
}

func Filter3(collection []int, sdfdsd func(int2 int) bool) []int {
	result := make([]int, 0, len(collection))

	for _, item := range collection {
		if sdfdsd(item) {
			result = append(result, item)
		}
	}

	return result
}

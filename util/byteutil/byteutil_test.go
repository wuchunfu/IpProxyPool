package byteutil

import (
	"fmt"
	"testing"
)

func TestByteSize(t *testing.T) {
	size := ByteSize(9999999999999999999)
	fmt.Println(size)
}

package pkg

import (
	"fmt"
	"math/rand"
	"time"
)

func GenID() string {
	return fmt.Sprintf(`%04x%04x-%04x-%04x-%04x-%04x%04x%04x`,
		// 32 bits for "time_low"
		MtRand(0, 0xffff), MtRand(0, 0xffff),
		// 16 bits for "time_mid"
		MtRand(0, 0xffff),
		// 16 bits for "time_hi_and_version",
		// four most significant bits holds version number 4
		MtRand(0, 0x0fff)|0x4000,
		// 16 bits, 8 bits for "clk_seq_hi_res",
		// 8 bits for "clk_seq_low",
		// two most significant bits holds zero and one for variant DCE1.1
		MtRand(0, 0x3fff)|0x8000,
		// 48 bits for "node"
		MtRand(0, 0xffff), MtRand(0, 0xffff), MtRand(0, 0xffff))
}

func MtRand(min, max int64) int64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int63n(max-min+1) + min
}

package config

import "time"

var JWT_KEY []byte
var EXP_TIME = time.Now().Add(5 * time.Hour)

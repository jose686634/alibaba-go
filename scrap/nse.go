package scrap

import (
	"time"

	"github.com/jose686634/alibaba-go/utils"
)

var k = utils.NewHTTPClient(2*time.Second, map[string]string{"a": "b"})

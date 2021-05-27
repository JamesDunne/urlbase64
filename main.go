package main

import (
	"encoding/base64"
	"io"
	"os"
)

func main() {
	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	// don't use padding chars since they are '=', horrible for URLs; there is no ambiguity when using URL-base64
	// encoded values without padding as query-string params since the values are already delimited
	enc := base64.RawURLEncoding

	o := make([]byte, enc.EncodedLen(len(b)))
	enc.Encode(o, b)

	os.Stdout.Write(o)
}

package main

import (
	"bufio"
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

type BpParse struct {
	BP Blueprint `json:"blueprint"`
}

func parseBPString(s string) (Blueprint, error) {
	s = s[1:]
	var out BpParse

	// base64 to bytes
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return out.BP, errors.New("Couldn't decode blueprint")
	}

	// decompress bytes to json string
	r, err := zlib.NewReader(bytes.NewReader(b))
	if err != nil {
		return out.BP, errors.New("Couldn't uncompress blueprint")
	}
	var dat bytes.Buffer
	io.Copy(bufio.NewWriter(&dat), r)
	r.Close()

	fmt.Println(string(dat.Bytes()))

	// parsing json
	err = json.Unmarshal(dat.Bytes(), &out)
	if err != nil {
		return out.BP, errors.New("Couldn't parse json")
	}

	return out.BP, nil
}

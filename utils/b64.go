// Package utils
// DateTime: 2023-02-10 20:30
// Author: CN
// Mail: Nanshao@n-s.fun
// Description:

package utils

import (
	"encoding/base64"

	"golang.org/x/text/encoding/simplifiedchinese"
)

func B64DecodeString(i string) []byte {
	decodeString, err := base64.StdEncoding.DecodeString(i)
	if err != nil {
		return nil
	}
	return decodeString
}
func B64DecodeStringGBK(i string) []byte {
	bs, err := base64.StdEncoding.DecodeString(i)
	if err != nil {
		return nil
	}

	b, err := simplifiedchinese.GB18030.NewDecoder().Bytes(bs)
	if err != nil {
		return nil
	}
	return b
}
func B64Encode(i []byte) string {
	return base64.StdEncoding.EncodeToString(i)
}

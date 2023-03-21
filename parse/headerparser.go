// Package parse
// DateTime: 2023-03-08 20:07
// Author: CN
// Mail: Nanshao@n-s.fun
// Description:

package parse

import (
	"strings"
	"time"

	"github.com/nanshaobit/nsmail/mail"
)

type headerParser struct {
	header *mail.Header
	err    error
}

func (hp headerParser) parseAddress(s string) (ma *mail.Address) {
	if hp.err != nil {
		return nil
	}
	if strings.Trim(s, " \n") != "" {
		ma, hp.err = mail.ParseAddress(s)

		return ma
	}
	return nil
}

func (hp headerParser) parseAddressList(s string) (ma []*mail.Address) {
	if hp.err != nil {
		return
	}
	if strings.Trim(s, " \n") != "" {
		ma, hp.err = mail.ParseAddressList(s)
		return
	}
	return
}

func (hp headerParser) parseTime(s string) (t time.Time) {
	if hp.err != nil || s == "" {
		return
	}

	formats := []string{
		time.RFC1123Z,
		"Mon, 2 Jan 2006 15:04:05 -0700",
		time.RFC1123Z + " (MST)",
		"Mon, 2 Jan 2006 15:04:05 -0700 (MST)",
	}

	for _, format := range formats {
		t, hp.err = time.Parse(format, s)
		if hp.err == nil {
			return
		}
	}

	return
}

func (hp headerParser) parseMessageId(s string) string {
	if hp.err != nil {
		return ""
	}
	return strings.Trim(s, "<> ")
}

func (hp headerParser) parseMessageIdList(s string) (result []string) {
	if hp.err != nil {
		return
	}
	for _, p := range strings.Split(s, " ") {
		if strings.Trim(p, " \n") != "" {
			result = append(result, hp.parseMessageId(p))
		}
	}
	return
}

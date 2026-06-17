package main

import (
	"bytes"
	"testing"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
)

func TestInitInputEncoding(t *testing.T) {
	t.Cleanup(func() {
		inputEncoding = nil
		inputEncName = ""
	})

	if err := initInputEncoding(""); err != nil {
		t.Fatal(err)
	}
	if inputEncoding != nil {
		t.Fatal("empty label should not set encoding")
	}

	if err := initInputEncoding("gbk"); err != nil {
		t.Fatal(err)
	}
	if inputEncoding == nil || inputEncName != "gbk" {
		t.Fatalf("gbk: encoding=%v name=%q", inputEncoding, inputEncName)
	}

	if err := initInputEncoding("not-a-real-encoding"); err == nil {
		t.Fatal("expected error for unknown encoding")
	}
}

func TestGBKDecode(t *testing.T) {
	e, _ := charset.Lookup("gbk")
	if e == nil {
		t.Fatal("gbk not recognized")
	}
	utf8, _, err := transform.Bytes(e.NewDecoder(), []byte("\xb2\xe2\xca\xd4")) // 测试
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(utf8, []byte("测试")) {
		t.Fatalf("got %q", utf8)
	}
}

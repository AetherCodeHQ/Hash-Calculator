package main

// Hash-Calculator: SHA-256, SHA-512, MD5 hash hesaplama
import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: hash-calculator <file> [md5|sha256|sha512]")
		os.Exit(1)
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer f.Close()

	algo := "all"
	if len(os.Args) > 2 {
		algo = os.Args[2]
	}

	switch algo {
	case "md5":
		h := md5.New()
		io.Copy(h, f)
		fmt.Printf("md5:    %s\n", hex.EncodeToString(h.Sum(nil)))
	case "sha256":
		h := sha256.New()
		io.Copy(h, f)
		fmt.Printf("sha256: %s\n", hex.EncodeToString(h.Sum(nil)))
	case "sha512":
		h := sha512.New()
		io.Copy(h, f)
		fmt.Printf("sha512: %s\n", hex.EncodeToString(h.Sum(nil)))
	default:
		f.Seek(0, 0)
		h256 := sha256.New()
		io.Copy(h256, f)
		f.Seek(0, 0)
		h512 := sha512.New()
		io.Copy(h512, f)
		f.Seek(0, 0)
		hmd5 := md5.New()
		io.Copy(hmd5, f)
		fmt.Printf("md5:    %s\n", hex.EncodeToString(hmd5.Sum(nil)))
		fmt.Printf("sha256: %s\n", hex.EncodeToString(h256.Sum(nil)))
		fmt.Printf("sha512: %s\n", hex.EncodeToString(h512.Sum(nil)))
	}
}

// Package crypto implements golang package crypto functionality for lua.
package crypto

import (
	"crypto/md5"
	"fmt"
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"

	lua "github.com/yuin/gopher-lua"
)

// MD5 lua crypto.md5(string) return string
func MD5(L *lua.LState) int {
	str := L.CheckString(1)
	hash := md5.Sum([]byte(str))
	L.Push(lua.LString(fmt.Sprintf("%x", hash)))
	return 1
}

// SHA256 lua crypto.sha256(string) return string
func SHA256(L *lua.LState) int {
	str := L.CheckString(1)
	hash := sha256.Sum256([]byte(str))
	L.Push(lua.LString(fmt.Sprintf("%x", hash)))
	return 1
}
func HMAC(L *lua.LState) int {
	algorithm := lua.LVAsString(L.Get(1))
	s := lua.LVAsString(L.Get(2))
	key := lua.LVAsString(L.Get(3))
	raw := lua.LVAsBool(L.Get(4))

	var h hash.Hash
	switch algorithm {
	case "md5":
		h = hmac.New(md5.New, []byte(key))
	case "sha1":
		h = hmac.New(sha1.New, []byte(key))
	case "sha256":
		h = hmac.New(sha256.New, []byte(key))
	case "sha512":
		h = hmac.New(sha512.New, []byte(key))
	default:
		L.Push(lua.LNil)
		L.Push(lua.LString("unsupported algorithm"))
		return 2
	}

	_, err := h.Write([]byte(s))
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	var result string
	if !raw {
		result = hex.EncodeToString(h.Sum(nil))
	} else {
		result = string(h.Sum(nil))
	}
	L.Push(lua.LString(result))
	return 1
}

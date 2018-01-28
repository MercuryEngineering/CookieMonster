# CookieMonster

[![GoDoc](https://godoc.org/github.com/MercuryEngineering/CookieMonster?status.svg)](https://godoc.org/github.com/MercuryEngineering/CookieMonster)

A simple package for parsing [Netscape Cookie File](http://curl.haxx.se/rfc/cookie_spec.html) format files into Go [Cookies](https://golang.org/pkg/net/http/#Cookie)

### Install

`go get -u github.com/MercuryEngineering/CookieMonster`

### Example

```
cookies, err := cookiemonster.ParseFile("cookies.txt")
if err != nil {
  panic(err)
}

for _, cookie := range cookies {
  fmt.Printf("%s=%s\n", cookie.Name, cookie.Value)
}
```

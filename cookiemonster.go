// Copyright (c) 2018 Henry Slawniak <https://datacenterscumbags.com/>
// Copyright (c) 2018 Mercury Engineering <https://mercury.engineering/>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cookiemonster

import (
	"bufio"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func ParseFile(path string) ([]*http.Cookie, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	cookies := []*http.Cookie{}

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") || line == "" {
			// Ignore comments and blank lines
			continue
		}

		split := strings.Split(line, "\t")
		expires, err := strconv.Atoi(strings.Split(split[4], ".")[0])
		if err != nil {
			return nil, err
		}

		cookie := &http.Cookie{
			Name:    split[5],
			Value:   split[6],
			Path:    split[2],
			Domain:  split[0],
			Expires: time.Unix(int64(expires), 0),
			Secure:  strings.ToLower(split[3]) == "true",
		}
		cookies = append(cookies, cookie)
	}

	return cookies, nil
}

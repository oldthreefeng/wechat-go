/*
Copyright 2017 wechat-go Authors. All Rights Reserved.
MIT License

Copyright (c) 2017

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package laosj

import (
	"math/rand"
	"testing"
	"time"

	"github.com/songtianyi/laosj/spider"
	"github.com/oldthreefeng/wechat-go/logs"
)

func Test_listenCmd(t *testing.T) {
	uri := "https://www.mzitu.com/zipai"
	s, err := spider.CreateSpiderFromUrl(uri)
	if err != nil {
		logs.Error(err)
		return
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	logs.Info(s.Url)
	srcs, _ := s.GetAttr("div.main>div.main-content>div.postlist>div>ul>li>div.comment-body>p>img", "src")
	if len(srcs) < 1 {
		logs.Error(srcs)
		logs.Error("cannot get mzitu images")
		return
	}
	for _, k := range srcs {
		logs.Info(k)
	}
	img := srcs[r.Intn(len(srcs))]
	logs.Info(img)
}
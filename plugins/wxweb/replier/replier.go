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

package replier

import (	
	"strings"

	"github.com/oldthreefeng/wechat-go/logs"
	"github.com/oldthreefeng/wechat-go/wxweb"
)

// register plugin
func Register(session *wxweb.Session) {
	session.HandlerRegister.Add(wxweb.MSG_TEXT, wxweb.Handler(autoReply), "text-replier")
	if err := session.HandlerRegister.Add(wxweb.MSG_IMG, wxweb.Handler(autoReply), "img-replier"); err != nil {
		logs.Error(err)
	}

	if err := session.HandlerRegister.EnableByName("text-replier"); err != nil {
		logs.Error(err)
	}

	if err := session.HandlerRegister.EnableByName("img-replier"); err != nil {
		logs.Error(err)
	}

}
func autoReply(session *wxweb.Session, msg *wxweb.ReceivedMessage) {
	logs.Info( "isgourp: %v, MsgId: %v, content: %v, from: %v, to: %v, who: %v, msgType: %v, subtype: %v, originContent: %v,at: %v, Url: %v", msg.IsGroup, msg.MsgId, msg.Content, msg.FromUserName, msg.ToUserName, msg.Who, msg.MsgType, msg.SubType,
	msg.OriginContent, msg.At, msg.Url)
	if !msg.IsGroup  {
		// session.SendText("暂时不在，稍后回复", session.Bot.UserName, msg.FromUserName)
		if msg.MsgType == wxweb.MSG_TEXT {
			str := msg.Content
			logs.Info(str)
			if strings.Contains(str, "经销商无法登陆") {
				logs.Info("auto send text")
				session.SendText("1、确认是否开了经销商账号\n2、经销商是否被禁用\n3、经销商是否登错账号", session.Bot.UserName, msg.FromUserName)
			
			}
		}
	}


}

// Package chatai is the "Chat AI (Claude)" DRIVER — Claude-backed chat responder driver: answers site-visitor questions with a real model instead of keyword matching.
//
// It self-registers into the chat base plugin on import. Install by blank-import:
//
//	import _ "github.com/togo-framework/chat-ai"
package chatai

import "github.com/togo-framework/chat"

func init() {
	chat.Register(chat.Descriptor{
		Key:         "ai",
		Title:       "Chat AI (Claude)",
		DefaultZone: "",
		Asset:       "/chats/ai.js",
	})
}

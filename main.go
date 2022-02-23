package main

import (
	"akita/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"log"
)

func main() {
	input, err := antlr.NewFileStream("/opt/3code/java/activity/kkcloud-activity-admin/kkcloud-activity-admin-biz/src/main/java/com/kkcloud/activity/admin/service/impl/CarnivalVoteVotingServiceImpl.java")

	if err != nil {
		log.Fatalf("err=%v", err)
		return
	}
	lexer := parser.NewJava9Lexer(input)
	for _, v := range lexer.GetAllTokens() {
		log.Println(v.GetLine(), v.GetText())
	}

	//stream := antlr.NewCommonTokenStream(lexer, 0)
	//stream.GetAllText()

}

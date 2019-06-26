package main

import (
	"flag"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/awaduharatk/go-batch-starter/db"
	errorh "github.com/awaduharatk/go-batch-starter/error"
	"github.com/awaduharatk/go-batch-starter/logic"
)

func main() {
	defer func() {
		// panicのキャッチ処理
		fmt.Println("予期せぬError")
		err := recover()
		if err != nil {
			// logging処理
			fmt.Println("Recover!:", err)
			os.Exit(errorh.ExitCodeError)
		}
	}()

	// 引数の取得
	flag.Parse()
	args := flag.Args()
	for _, s := range args {
		fmt.Println(s)
	}

	// 初期化
	// config.Init()
	db.Init()

	// 業務ロジックのキック処理
	err := run(args)

	// 終了コード判定を行い、処理終了
	os.Exit(errorh.HandleExit(err))
}

func run(args []string) error {
	fmt.Println("running")
	// sub := logic.NewSublogic(db.GetDB())
	logic := &logic.Mainlogic{}
	return logic.Logic(args)
}

package main

import (
	"flag"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/awaduharatk/go-batch-starter/configuration"
	errorh "github.com/awaduharatk/go-batch-starter/error"
	"github.com/awaduharatk/go-batch-starter/logic"
)

var db *gorm.DB

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
	db = configuration.NewDB()

	// 業務ロジックをキックし、終了コード判定
	os.Exit(errorh.HandleExit(run(args)))
}

func run(args []string) error {
	fmt.Println("running")
	logic := logic.NewMainlogic(
		logic.NewSublogic(db),
	)
	return logic.Logic(args)
}

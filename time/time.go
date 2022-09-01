package time

import (
	"fmt"
	"github.com/kingcanfish/keyu/utils"
	"github.com/urfave/cli/v2"
	"strconv"
	"time"
)

func timeStampToString(timeStamp string) (string, error) {
	format := "2006.01.02 15:04:05"
	ts, err := strconv.ParseInt(timeStamp, 10, 64)
	if err != nil {
		return "", err
	}
	// 秒
	if ts < 1_000_000_000_000 {
		return time.Unix(ts, 0).Format(format), nil
	}
	// 毫秒
	if ts < 1_000_000_000_000_000 {
		return time.UnixMilli(ts).Format(format), nil
	}
	// 微秒
	return time.UnixMicro(ts).Format(format), nil
}

func StampAction(context *cli.Context) error {
	str := context.Args().Get(0)
	if str == "" {
		utils.EmptyArg(context)
	}
	content, err := timeStampToString(str)
	if err != nil {
		utils.ErrArg(context)
	}
	fmt.Println(content)
	return nil
}

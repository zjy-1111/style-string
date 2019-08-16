package style

import "fmt"

type Color int

const (
	Black   Color = iota
	Red
	Green
	Yellow
	Blue
	Puple
	Cyan
	White
	Default Color = -1
)

func toFgColor(c Color) int {
	if c == Default {
		return 0
	}
	return int(c) + 30
}

func toBgColor(c Color) int {
	if c == Default {
		return 0
	}
	return int(c) + 40
}

//// 前景颜色
//type FColor int
//
//const (
//	FBlack FColor = iota + 30
//	FRed
//	FGreen
//	FYellow
//	FBlue
//	FPuple
//	FCyan
//	FWhite
//)
//
//// 背景颜色
//type BgColor int
//
//const (
//	BgBlack BgColor = iota + 40
//	BgRed
//	BgGreen
//	BgYellow
//	BgBlue
//	BgPuple
//	BgCyan
//	BgWhite
//)

// 字体属性
type Attr int

const (
	Reset        Attr = iota // 默认
	Bold                     // 粗体
	Faint                    // 微弱
	Italic                   // 斜体
	Underline                // 下划线
	BlinkSlow                // 缓慢闪烁
	BlinkRapid               // 快速闪烁
	ReverseVideo             // 反向显示
	Concealed                // 隐藏
	CrossedOut               // 删除线
)

func StyleString(s string, a Attr, b Color, f Color) string {
	return fmt.Sprintf("\033[%d;%d;%dm%s\033[0m", a, toBgColor(b), toFgColor(f), s)
}

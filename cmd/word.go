package cmd

import (
	"github.com/ice-waves/tour/internal/word"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

const (
	ModeUpper = iota + 1    		// 全部单词转大写
	ModeLower					    // 全部单词转小写
	ModeUnderscoreToUpperCamelCase  // 下划线单词转大骆驼单词
	ModeUnderscoreToLowerCamelCase  // 下划线单词转小驼峰单词
	ModeCamelCaseToUnderscore		// 驼峰单词转下划线
)

var desc = strings.Join([]string{
	"1：全部单词转大写",
	"2：全部单词转小写",
	"3：下划线单词转大骆驼单词",
	"4：下划线单词转小驼峰单词",
	"5：驼峰单词转下划线",
}, "\n")

var str string
var mode int8

var wordCmd = &cobra.Command{
	Use: "word",
	Short: "单词格式转换",
	Long: desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToLower(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderscoreToLowerCamelCase:
			content = word.UnderscoreToLowerCamelCase(str)
		case ModeUnderscoreToUpperCamelCase:
			content = word.UnderscoreToUpperCamelCase(str)
		case ModeCamelCaseToUnderscore:
			content = word.CamelCaseToUnderscore(str)
		default:
			log.Fatalf("暂不支持该种转换格式，请使用 help word 查看帮助文档")
		}

		log.Printf("转换结果：%s", content)
	},
}

func init()  {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换模式")
}

package cmd

import (
	"github.com/ice-waves/tour/internal/sql2struct"
	"github.com/spf13/cobra"
	"log"
)

var (
	username  string
	password  string
	host      string
	charset   string
	dbType    string
	dbName    string
	tableName string
)

var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "sql转换和处理",
	Long: "sql转换和处理",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var sql2structCmd = &cobra.Command{
	Use: "struct",
	Short: "转换",
	Long: "转换",
	Run: func(cmd *cobra.Command, args []string) {
		dbInfo := &sql2struct.DBInfo{
			DBType:   dbType,
			Host:     host,
			UserName: username,
			Password: password,
			Charset:  charset,
		}

		dbModel := sql2struct.NewDBModel(dbInfo)
		if err := dbModel.Connect(); err != nil {
			log.Fatalf("dbModel.Connect err: %v", err)
		}
		columns, err := dbModel.GetColumns(dbName, tableName)
		if err != nil {
			log.Fatalf("dbModel.GetColumns err: %v", columns)
		}
		template := sql2struct.NewStructTemplate()
		templateColumns := template.AssemblyColumns(columns)
		if err := template.Generate(tableName, templateColumns); err != nil {
			log.Fatalf("template.Generate err: %v", err)
		}
	},
}

func init()  {
	sqlCmd.AddCommand(sql2structCmd)
	sql2structCmd.Flags().StringVarP(&username, "username", "", "", "请输入数据库账号")
	sql2structCmd.Flags().StringVarP(&password, "password", "", "", "请输入数据库密码")
	sql2structCmd.Flags().StringVarP(&host, "host", "", "127.0.0.1:3306", "请输入数据库host")
	sql2structCmd.Flags().StringVarP(&charset, "charset", "", "utf8mb4", "请输入数据库编码")
	sql2structCmd.Flags().StringVarP(&dbType, "dbType", "", "mysql", "请输入数据库类型")
	sql2structCmd.Flags().StringVarP(&dbName, "dbName", "", "", "请输入数据库名称")
	sql2structCmd.Flags().StringVarP(&tableName, "tableName", "", "", "请输入表名称")
}
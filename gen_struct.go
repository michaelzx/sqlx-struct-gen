package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jimsmart/schema"
	"github.com/jinzhu/inflection"
	"github.com/michaelzx/zky-sqlx-gen/dbmeta"
	"github.com/michaelzx/zky-sqlx-gen/tpl"
	"github.com/michaelzx/zky-sqlx-gen/utils"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func genStruct() {
	currPath, _ := os.Getwd()
	var tableRegion string
	prompt := &survey.Select{
		Message: "数据表范围",
		Options: []string{"全部表", "部分表"},
	}
	if err := survey.AskOne(prompt, &tableRegion); err != nil {
		log.Fatal(err)
		return
	}
	sqlType := "mysql"
	// root@tcp(127.0.0.1:3306)/employees?&parseTime=True
	sqlConnStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?&parseTime=True",
		DbCfg.Usr,
		DbCfg.Psw,
		DbCfg.Host,
		DbCfg.Port,
		DbCfg.DbName,
	)
	fmt.Printf("type:%s\n", sqlType)
	fmt.Printf("conn:%s\n", sqlConnStr)
	// 创建数据库链接
	var db, err = sql.Open(sqlType, sqlConnStr)
	if err != nil {
		fmt.Println("数据库链接创建失败: " + err.Error())
		return
	}
	defer db.Close()

	// parse or read tables
	var tables []string
	if tables, err = schema.TableNames(db); err != nil {
		fmt.Println("Error in fetching tables information from mysql information schema")
		return
	}
	var genTables []string
	if tableRegion == "部分表" {
		prompt := &survey.MultiSelect{
			Message: "请选择表:",
			Options: tables,
		}
		if err := survey.AskOne(prompt, &genTables); err != nil {
			log.Fatal(err)
			return
		}
	} else {
		genTables = tables
	}
	// 检查model文件夹是否存在，不存在则创建
	if exist, err := utils.PathExists("model"); err == nil && !exist {
		if err := os.Mkdir("model", 0777); err != nil {
			log.Fatal(err)
		}
	}

	modelTpl, err := utils.GetTemplate(tpl.ModelTmpl)
	if err != nil {
		fmt.Println("模板加载失败: " + err.Error())
		return
	}

	// 为每张表生成model文件
	for _, tableName := range genTables {
		structName := dbmeta.FmtFieldName(tableName)

		modelInfo := dbmeta.GenerateStruct(db, tableName, structName, "model")

		var buf bytes.Buffer
		err = modelTpl.Execute(&buf, modelInfo)
		if err != nil {
			fmt.Println("Error in rendering model: " + err.Error())
			return
		}
		fmt.Print(buf.String())
		data, err := format.Source(buf.Bytes())
		if err != nil {
			fmt.Println("Error in formating source: " + err.Error())
			return
		}

		relativePath := filepath.Join("model", inflection.Singular(tableName)+".go")
		if exist, err := utils.PathExists(relativePath); err == nil && exist {
			var existAction string
			prompt := &survey.Select{
				Message: fmt.Sprintf("文件已存在，是否覆盖？[%s]", relativePath),
				Options: []string{"否", "是"},
			}
			if err := survey.AskOne(prompt, &existAction); err != nil {
				log.Fatal(err)
				return
			}
			if existAction == "否" {
				utils.PrintRed(fmt.Sprintf("已存在，跳过 : %s", filepath.Join(currPath, relativePath)))
				continue
			}
		}
		ioutil.WriteFile(filepath.Join(relativePath), data, 0777)
		utils.PrintGreen(fmt.Sprintf("成功生成 : %s", filepath.Join(currPath, relativePath)))
	}
}

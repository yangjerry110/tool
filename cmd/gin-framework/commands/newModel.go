/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-25 10:45:19
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-11 11:25:54
 * @Description: new model
 */
package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golib/cli"
	"github.com/yangjerry110/tool/cmd/gin-framework/errors"
	"github.com/yangjerry110/tool/cmd/gin-framework/templates"
	"github.com/yangjerry110/tool/conf"
	"github.com/yangjerry110/tool/db"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

type NewModelCommands interface {
	NewModel(ctx *cli.Context) error
	CreateModel() error
	CreateWd() error
	AskModelConfigPath() error
}

type NewModel struct {
	ModelPath string
}

/**
 * @description: NewModelParams
 * @author: Jerry.Yang
 * @date: 2023-04-25 14:14:13
 * @return {*}
 */
var NewModelParams = &NewModel{}

/**
 * @description: NewModel
 * @param {*cli.Context} ctx
 * @author: Jerry.Yang
 * @date: 2023-05-11 11:25:53
 * @return {*}
 */
func (n *NewModel) NewModel(ctx *cli.Context) error {

	/**
	 * @step
	 * @获取第一个参数的名称
	 **/
	modelName := ctx.Args().First()
	if modelName == "" {
		return errors.ErrModelNameIsEmpty
	}

	/**
	 * @step
	 * @进行赋值
	 **/
	InitParms.ModelName = modelName

	/**
	 * @step
	 * @问一下modelConfig的path
	 **/
	err := n.AskModelConfigPath()
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @设置projectPath
	 **/
	err = CreateInitCommands().SetProjectPath()
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @设置projectImportPath
	 **/
	err = CreateInitCommands().SetImportProjectPath()
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @执行actionModelGen
	 **/
	err = n.ActionModelGen()
	if err != nil {
		return err
	}
	return nil
}

/**
 * @description: CreateModel
 * @author: Jerry.Yang
 * @date: 2023-04-25 15:00:39
 * @return {*}
 */
func (n *NewModel) CreateModel() error {

	/**
	 * @step
	 * @创建目录
	 **/
	err := n.CreateWd()
	if err != nil {
		return err
	}
	return nil
}

/**
 * @description: CreateWd
 * @author: Jerry.Yang
 * @date: 2023-04-25 14:15:53
 * @return {*}
 */
func (n *NewModel) CreateWd() error {

	/**
	 * @step
	 * @获取config的path
	 **/
	path := fmt.Sprintf("%s/%s", InitParms.ProjectPath, "model")

	/**
	 * @step
	 * @创建configPath
	 **/
	err := os.MkdirAll(path, 0777)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @赋值
	 **/
	NewModelParams.ModelPath = path
	return nil
}

/**
 * @description: AskModelConfig
 * @author: Jerry.Yang
 * @date: 2023-05-10 11:00:26
 * @return {*}
 */
func (n *NewModel) AskModelConfigPath() error {

	/**
	 * @step
	 * @初始化reader
	 **/
	reader := bufio.NewReader(os.Stdin)

	/**
	 * @step
	 * @定义输入的提示
	 **/
	fmt.Println("请输入你的model配置文件路径，按回车结束")
	fmt.Print("\r\n")
	fmt.Print("need username; password; port; ip; database;  ")
	fmt.Print("=> ")

	/**
	 * @step
	 * @获取输入的内容
	 **/
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("发生错误 : %+v \r\n", err)
		return err
	}

	/**
	 * @step
	 * @假如输入内容为空，报错直接
	 **/
	if len(text) == 1 {
		return errors.ErrModelConfigIsEmpty
	}

	/**
	 * @step
	 * @删除换行
	 **/
	text = strings.ReplaceAll(text, "\n", "")

	/**
	 * @step
	 * @赋值
	 **/
	InitParms.ModelConfigPath = text
	return nil
}

/**
 * @description: ActionModelGen
 * @author: Jerry.Yang
 * @date: 2023-05-10 14:42:20
 * @return {*}
 */
func (n *NewModel) ActionModelGen() error {

	/**
	 * @step
	 * @获取gormDb
	 **/
	gormDb, err := n.GetGormDb()
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @初始化gen
	 **/
	modelPath := fmt.Sprintf("%s%s", InitParms.ProjectPath, "model")
	queryPath := fmt.Sprintf("%s%s", InitParms.ProjectPath, "query")
	genObj := gen.NewGenerator(gen.Config{
		// 相对执行`go run`时的路径, 会自动创建目录
		OutPath: queryPath,
		// outFile
		// OutFile: "go",
		// WithDefaultQuery 生成默认查询结构体(作为全局变量使用), 即`Q`结构体和其字段(各表模型)
		// WithoutContext 生成没有context调用限制的代码供查询
		// WithQueryInterface 生成interface形式的查询代码(可导出), 如`Where()`方法返回的就是一个可导出的接口类型
		Mode: gen.WithDefaultQuery | gen.WithQueryInterface,
		// 表字段可为 null 值时, 对应结体字段使用指针类型
		FieldNullable: true,
		// 表字段默认值与模型结构体字段零值不一致的字段, 在插入数据时需要赋值该字段值为零值的, 结构体字段须是指针类型才能成功, 即`FieldCoverable:true`配置下生成的结构体字段.
		// 因为在插入时遇到字段为零值的会被GORM赋予默认值. 如字段`age`表默认值为10, 即使你显式设置为0最后也会被GORM设为10提交.
		// 如果该字段没有上面提到的插入时赋零值的特殊需要, 则字段为非指针类型使用起来会比较方便.
		FieldCoverable: false,
		// 模型结构体字段的数字类型的符号表示是否与表字段的一致, `false`指示都用有符号类型
		FieldSignable: false,
		// 生成 gorm 标签的字段索引属性
		FieldWithIndexTag: true,
		// 生成 gorm 标签的字段类型属性
		FieldWithTypeTag: true,
	})

	/**
	 * @step
	 * @设置目标db
	 **/
	genObj.UseDB(gormDb)

	/**
	 * @step
	 * @统一int类型
	 **/
	dataIntMap := map[string]func(columnType gorm.ColumnType) (dataType string){
		"tinyint":   func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		"smallint":  func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		"mediumint": func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		"bigint":    func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		"int":       func(columnType gorm.ColumnType) (dataType string) { return "int64" },
	}

	/**
	 * @step
	 * @使用类型映射
	 **/
	genObj.WithDataTypeMap(dataIntMap)

	/**
	 * @step
	 * @生成数据库内所有表的结构体
	 **/
	// genObj.GenerateAllTable()

	/**
	 * @step
	 * @生成某张表的结构体
	 **/
	modelName := templates.CreateCommonTemplate().FirstUpper(InitParms.ModelName)
	genObj.ApplyBasic(genObj.GenerateModelAs(InitParms.ModelName, modelName))
	// genObj.ApplyBasic(genObj.GenerateAllTable()...)

	/**
	 * @step
	 * @执行
	 **/
	genObj.Execute()

	/**
	 * @step
	 * @判断文件是否存在
	 **/
	oldModelFileName := fmt.Sprintf("%s/%s.gen.go", modelPath, InitParms.ModelName)
	_, err = os.Stat(oldModelFileName)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @重命名model文件
	 **/
	NewAppParams.AppModelFileName = fmt.Sprintf("%s/%s.go", modelPath, InitParms.ModelName)
	err = os.Rename(oldModelFileName, NewAppParams.AppModelFileName)
	if err != nil {
		return err
	}
	return nil
}

/**
 * @description: GetGormDb
 * @author: Jerry.Yang
 * @date: 2023-05-10 14:18:25
 * @return {*}
 */
func (n *NewModel) GetGormDb() (*gorm.DB, error) {

	/**
	 * @step
	 * @获取配置
	 **/
	dataBase := db.BaseDb{}
	yamlConf := conf.YamlConf{FilePath: InitParms.ModelConfigPath, FileName: "database.yaml", FileType: "yaml", Intervals: 10 * time.Minute, Conf: dataBase}
	err := yamlConf.GetConf(&dataBase)
	if err != nil {
		return nil, err
	}

	/**
	 * @step
	 * @初始化配置
	 **/
	db, err := gorm.Open(mysql.Open(dataBase.Dsn))
	if err != nil {
		return nil, err
	}
	return db, nil
}

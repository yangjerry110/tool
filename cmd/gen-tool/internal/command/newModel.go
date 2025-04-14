package command

import (
	"context"
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/golib/cli"
	"github.com/yangjerry110/tool/cmd/gen-tool/internal/config"
	"github.com/yangjerry110/tool/cmd/gen-tool/internal/errors"
	"github.com/yangjerry110/tool/cmd/gen-tool/internal/template"
	"github.com/yangjerry110/tool/conf"
	"github.com/yangjerry110/tool/db"
	"gorm.io/gen"
	"gorm.io/gorm"
)

type NewModel struct {
	CliContext *cli.Context
}

/**
 * @description: New
 * @author: Jerry.Yang
 * @date: 2023-12-21 14:23:31
 * @return {*}
 */
func (n *NewModel) New() error {

	// If CliContext == nil
	if n.CliContext == nil {
		return errors.ErrCommandNoCliContext
	}

	// Set ModelName
	if err := conf.CreateConf(&config.Model{CliContext: n.CliContext}).SetConfig(); err != nil {
		return err
	}

	// Set ProjectPath
	if err := conf.CreateConf(&config.ProjectPath{}).SetConfig(); err != nil {
		return err
	}

	// askDatabaseConfPath
	databaseConfPath, err := n.askDatabaseConfPath()
	if err != nil {
		return err
	}

	// Set ConfigPath
	if err := conf.CreatePathConf(databaseConfPath).SetConfig(); err != nil {
		return err
	}

	// Set databaseConf
	if err := db.SetGormConf().SetConfig(); err != nil {
		return err
	}

	// Ask databaseConfName
	databaseConfName, err := n.askDatabaseConfName()
	if err != nil {
		return err
	}

	// Action Gen
	if err := n.actionGen(databaseConfName); err != nil {
		return err
	}
	return nil
}

/**
 * @description: actionGen
 * @param {string} databaseConfName
 * @author: Jerry.Yang
 * @date: 2023-12-21 14:23:06
 * @return {*}
 */
func (n *NewModel) actionGen(databaseConfName string) error {

	// Define
	modelPath := fmt.Sprintf("%s/%s/%s", config.ProjectPathConf.Path, "internal", "model")
	queryPath := fmt.Sprintf("%s/%s/%s", config.ProjectPathConf.Path, "internal", "query")

	// Get GormDb
	gormDb, err := n.getGormDb(databaseConfName)
	if err != nil {
		return err
	}

	// Init gen
	genGenerate := gen.NewGenerator(gen.Config{
		// Relative path when executing `go run`, the directory will be created automatically
		OutPath: queryPath,

		// outFile
		// OutFile: "go",

		// WithDefaultQuery generates a default query struct (used as a global variable),
		// including the `Q` struct and its fields (models of various tables).
		// WithoutContext generates code without context calling restrictions for queries.
		// WithQueryInterface generates the query code in interface form (exportable),
		// such as the `Where()` method returning an exportable interface type.
		Mode: gen.WithDefaultQuery | gen.WithQueryInterface,

		// When table fields can be null, the corresponding struct fields use pointer types
		FieldNullable: true,

		// For fields where the table default value is inconsistent with the zero value of the model struct field,
		// during data insertion, the field value must be assigned to zero value.
		// The struct field must be a pointer type to be successful, i.e., under the `FieldCoverable:true` configuration,
		// the generated struct fields.
		// Because, when inserting, if a field encounters a zero value, GORM will assign the default value.
		// For example, if the table default value for the `age` field is 10, even if you explicitly set it to 0,
		// it will be set to 10 by GORM in the end.
		// If this field does not have the special need mentioned above to assign zero value during insertion,
		// then using non-pointer types for fields will be more convenient.
		FieldCoverable: false,

		// Represents whether the sign of the numeric type of the model struct field is consistent with the table field.
		// `false` indicates that signed types are used consistently for both.
		FieldSignable: false,

		// Generates the GORM tag's field index attribute
		FieldWithIndexTag: true,

		// Generates the GORM tag's field type attribute
		FieldWithTypeTag: true,
	})

	// Use gormDb
	genGenerate.UseDB(gormDb)

	// cover int => int64
	dataIntMap := map[string]func(columnType gorm.ColumnType) (dataType string){
		"tinyint":   func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		"smallint":  func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		"mediumint": func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		"bigint":    func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		"int":       func(columnType gorm.ColumnType) (dataType string) { return "int64" },
	}

	// WithDataTypeMap
	genGenerate.WithDataTypeMap(dataIntMap)

	// Generate All Table
	// genGenerate.GenerateAllTable()

	// Generate one table
	modelNameUp := template.FirstUpper(config.ModelConf.ModelName)
	genGenerate.ApplyBasic(genGenerate.GenerateModelAs(config.ModelConf.TableName, modelNameUp))
	// genObj.ApplyBasic(genObj.GenerateAllTable()...)

	/**
	 * @step
	 * @执行
	 **/
	genGenerate.Execute()

	/**
	 * @step
	 * @判断文件是否存在
	 **/
	oldModelFileName := fmt.Sprintf("%s/%s.gen.go", modelPath, config.ModelConf.TableName)
	_, err = os.Stat(oldModelFileName)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @重命名model文件
	 **/
	newModelFileName := fmt.Sprintf("%s/%s.go", modelPath, config.ModelConf.ModelName)
	err = os.Rename(oldModelFileName, newModelFileName)
	if err != nil {
		return err
	}
	return nil
}

/**
 * @description: getGormDb
 * @param {string} databaseConfName
 * @author: Jerry.Yang
 * @date: 2023-12-21 14:02:14
 * @return {*}
 */
func (n *NewModel) getGormDb(databaseConfName string) (*gorm.DB, error) {

	// CreateClient
	if err := db.CreateGormDb().CreateClient(databaseConfName); err != nil {
		return nil, err
	}

	// Return GetClient
	ctx := context.Background()
	return db.CreateGormDb().GetClient(ctx, databaseConfName)
}

/**
 * @description: askDatabaseConfPath
 * @author: Jerry.Yang
 * @date: 2023-12-21 14:38:06
 * @return {*}
 */
func (n *NewModel) askDatabaseConfPath() (string, error) {

	// define question
	questions := []*survey.Question{
		{
			Name: "databaseConfPath",
			Prompt: &survey.Input{
				Message: "please input databaseConfPath ? ",
				Default: "",
			},
		},
	}

	// set answer
	type Answer struct {
		DatabaseConfPath string `survey:"databaseConfPath"`
	}

	// action ask
	// Set the answer data to Answer
	answer := &Answer{}
	err := survey.Ask(questions, answer)
	if err != nil {
		return "", err
	}
	return answer.DatabaseConfPath, nil
}

/**
 * @description: askDatabaseConfName
 * @author: Jerry.Yang
 * @date: 2023-12-21 13:59:12
 * @return {*}
 */
func (n *NewModel) askDatabaseConfName() (string, error) {

	// define question
	questions := []*survey.Question{
		{
			Name: "databaseConfName",
			Prompt: &survey.Input{
				Message: "please input databaseConfName ? ",
				Default: "",
			},
		},
	}

	// set answer
	type Answer struct {
		DatabaseConfName string `survey:"databaseConfName"`
	}

	// action ask
	// Set the answer data to Answer
	answer := &Answer{}
	err := survey.Ask(questions, answer)
	if err != nil {
		return "", err
	}
	return answer.DatabaseConfName, nil

}

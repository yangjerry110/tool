package dao

import (
	"fmt"

	"github.com/yangjerry110/tool/internal/cmd/config"
	"github.com/yangjerry110/tool/internal/cmd/template"
)

type NewDao struct {
	ProjectImportPath string
	DaoName           string
	DaoNameUp         string
	FirstDaoName      string
	DbName            string
	Time              string
}

/**
 * @description: New
 * @author: Jerry.Yang
 * @date: 2023-12-21 15:40:51
 * @return {*}
 */
func (n *NewDao) New() error {

	// filePath
	filePath := fmt.Sprintf("%s/internal/dao", config.ProjectPathConf.Path)
	return template.SaveTemplate(filePath, fmt.Sprintf("%sDao.go", n.DaoName), n.getTemplate(), n)
}

/**
 * @description: getTemplate
 * @author: Jerry.Yang
 * @date: 2023-12-21 15:40:41
 * @return {*}
 */
func (n *NewDao) getTemplate() string {
	return `/*
	* @Author: Jerry.Yang
	* @Date: {{.Time}}
	* @LastEditors: Jerry.Yang
	* @LastEditTime: {{.Time}}
	* @Description: {{.DaoName}} dao
	*/
   package dao
   
   import (
	   "context"
   
	   "{{.ProjectImportPath}}/logger"
	   "{{.ProjectImportPath}}/model"
   )
   
   type {{.DaoNameUp}}Dao interface {
		Get{{.DaoNameUp}}List(ctx context.Context, {{.DaoName}}Model *model.{{.DaoNameUp}}) ([]*model.{{.DaoNameUp}}, error)
	   Get{{.DaoNameUp}}Info(ctx context.Context, {{.DaoName}}Model *model.{{.DaoNameUp}}) (*model.{{.DaoNameUp}}, error)
	   Save{{.DaoNameUp}}(ctx context.Context, {{.DaoName}}Model *model.{{.DaoNameUp}}) (int64, error)
	   Delete{{.DaoNameUp}}(ctx context.Context, {{.DaoName}}Model *model.{{.DaoNameUp}}) (bool, error)
   }
   
   type {{.DaoNameUp}} struct{}
   
   /**
	* @description: Get{{.DaoNameUp}}List
	* @param {context.Context} ctx
	* @param {*model.{{.DaoNameUp}}} {{.DaoName}}Model
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/
   func ({{.FirstDaoName}} *{{.DaoNameUp}}) Get{{.DaoNameUp}}List(ctx context.Context, {{.DaoName}}Model *model.{{.DaoNameUp}}) ([]*model.{{.DaoNameUp}}, error) {
   
	   /**
		* @step
		* @result
		**/
	   result := []*model.{{.DaoNameUp}}{} /**
	   * @step
	   * @执行结果
	   **/
	  if err := CreateCommonDao().DbClient({{.DbName}}).Where({{.DaoName}}Model).Find(result).Error; err != nil {
		  logger.Logger().Errorf("{{.DaoName}}Dao Get{{.DaoNameUp}}List Err : %+v", err)
		  return nil, err
	  }
	  return result, nil
  }
  
  /**
   * @description: Get{{.DaoNameUp}}Info
   * @param {context.Context} ctx
   * @param {*model.{{.DaoNameUp}}} {{.DaoName}}Model
   * @author: Jerry.Yang
   * @date: {{.Time}}
   * @return {*}
   */
  func ({{.FirstDaoName}} *{{.DaoNameUp}}) Get{{.DaoNameUp}}Info(ctx context.Context,{{.DaoName}}Model *model.{{.DaoNameUp}}) (*model.{{.DaoNameUp}}, error) {
  
	  /**
	   * @step
	   * @result
	   **/
	  result := &model.{{.DaoNameUp}}{}
  
	  /**
	   * @step
	   * @执行结果
	   **/
	  if err := CreateCommonDao().DbClient({{.DbName}}).Where({{.DaoName}}Model).First(result).Error; err != nil {
		  logger.Logger().Errorf("{{.DaoName}}Dao Get{{.DaoNameUp}}Info Err : %+v", err)
		  return nil, err
	  }
	  return result, nil
  } 
  
  /**
  * @description: Save{{.DaoNameUp}}
  * @param {context.Context} ctx
  * @param {*model.{{.DaoNameUp}}} {{.DaoName}}Model
  * @author: Jerry.Yang
  * @date: {{.Time}}
  * @return {*}
  */
 func ({{.FirstDaoName}} *{{.DaoNameUp}}) Save{{.DaoNameUp}} (ctx context.Context, {{.DaoName}}Model *model.{{.DaoNameUp}}) (int64, error) {
 
	 /**
	  * @step
	  * @执行结果
	  **/
	 if err := CreateCommonDao().DbClient({{.DbName}}).Save({{.DaoName}}Model).Error; err != nil {
		 logger.Logger().Errorf("{{.DaoName}}Dao Save{{.DaoNameUp}} Err : %+v", err)
		 return 0, err
	 }
	 return {{.DaoName}}Model.ID, nil
 }

 /**
  * @description: Delete{{.DaoNameUp}}
  * @param {context.Context} ctx
  * @param {*model.{{.DaoNameUp}}} {{.DaoName}}Model
  * @author: Jerry.Yang
  * @date: {{.Time}}
  * @return {*}
  */
 func ({{.FirstDaoName}} *{{.DaoNameUp}}) Delete{{.DaoNameUp}} (ctx context.Context, {{.DaoName}}Model *model.{{.DaoNameUp}}) (bool, error) {
 
	 /**
	  * @step
	  * @执行结果
	  **/
	 if err := CreateCommonDao().DbClient({{.DbName}}).Where({{.DaoName}}Model).Update("is_deleted = ?", model.Is_Deleted).Error; err != nil {
		 logger.Logger().Errorf("{{.DaoName}}Dao Delete{{.DaoNameUp}} Err : %+v", err)
		 return false, err
	 }
	 return true, nil
 }
 `
}

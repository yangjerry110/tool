/*
 * @Author: Jerry.Yang
 * @Date: 2023-05-10 17:55:34
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-11 14:51:06
 * @Description:new dao
 */
package dao

import (
	"github.com/yangjerry110/tool/cmd/gin-framework/templates"
)

type NewDao interface {
	SaveTemplate(path string, projectImportPath string, daoName string, daoFileName string) error
	GetTemplate() string
}

type New struct{}

/**
 * @description: SaveTemplate
 * @param {string} path
 * @param {string} projectImportPath
 * @param {string} daoName
 * @param {string} daoFileName
 * @author: Jerry.Yang
 * @date: 2023-05-11 11:22:28
 * @return {*}
 */
func (n *New) SaveTemplate(path string, projectImportPath string, daoName string, daoFileName string) error {
	/**
	 * @step
	 * @定义渲染的数据
	 **/
	type Data struct {
		ProjectImportPath string
		DaoName           string
		DaoNameUp         string
		FirstDaoName      string
	}

	/**
	 * @step
	 * @appName进行大写字母的转换
	 **/
	daoNameUp := templates.CreateCommonTemplate().FirstUpper(daoName)

	/**
	 * @step
	 * @对第一个字母进行大写
	 **/
	FirstDaoName := daoName[:1]

	/**
	 * @step
	 * @进行赋值
	 **/
	data := &Data{ProjectImportPath: projectImportPath, DaoName: daoName, DaoNameUp: daoNameUp, FirstDaoName: FirstDaoName}
	return templates.CreateCommonTemplate().SaveTemplate(path, daoFileName, n.GetTemplate(), data)
}

/**
 * @description: GetTemplate
 * @author: Jerry.Yang
 * @date: 2023-05-10 18:01:53
 * @return {*}
 */
func (n *New) GetTemplate() string {
	return `/*
	* @Author: Jerry.Yang
	* @Date: 2023-05-10 17:02:13
	* @LastEditors: Jerry.Yang
	* @LastEditTime: 2023-05-10 17:54:17
	* @Description: {{.DaoName}} dao
	*/
   package dao
   
   import (
	   "context"
   
	   "{{.ProjectImportPath}}/logger"
	   "{{.ProjectImportPath}}/model"
   )
   
   type {{.DaoNameUp}}Dao interface {
	   GetList(ctx context.Context, {{.DaoName}}Model *model.{{.DaoNameUp}}) ([]*model.{{.DaoNameUp}}, error)
	   GetInfo(ctx context.Context, {{.DaoName}}Model *model.{{.DaoNameUp}}) (*model.{{.DaoNameUp}}, error)
	   Save(ctx context.Context, {{.DaoName}}Model *model.{{.DaoNameUp}}) (int64, error)
	   Delete(ctx context.Context, {{.DaoName}}Model *model.{{.DaoNameUp}}) (bool, error)
   }
   
   type {{.DaoNameUp}} struct{}
   
   /**
	* @description: GetList
	* @param {context.Context} ctx
	* @param {*model.{{.DaoNameUp}}} {{.DaoName}}Model
	* @author: Jerry.Yang
	* @date: 2023-05-10 17:49:19
	* @return {*}
	*/
   func ({{.FirstDaoName}} *{{.DaoNameUp}}) GetList(ctx context.Context, {{.DaoName}}Model *model.{{.DaoNameUp}}) ([]*model.{{.DaoNameUp}}, error) {
   
	   /**
		* @step
		* @result
		**/
	   result := []*model.{{.DaoNameUp}}{}
   
	   /**
		* @step
		* @执行结果
		**/
	   if err := CreateCommonDao().DbClient().Where({{.DaoName}}Model).Find(result).Error; err != nil {
		   logger.Logger().Errorf("{{.DaoName}}Dao GetList Err : %+v", err)
		   return nil, err
	   }
	   return result, nil
   }
   
   /**
	* @description: GetInfo
	* @param {context.Context} ctx
	* @param {*model.{{.DaoNameUp}}} {{.DaoName}}Model
	* @author: Jerry.Yang
	* @date: 2023-05-10 17:50:21
	* @return {*}
	*/
   func ({{.FirstDaoName}} *{{.DaoNameUp}}) GetInfo(ctx context.Context,{{.DaoName}}Model *model.{{.DaoNameUp}}) (*model.{{.DaoNameUp}}, error) {
   
	   /**
		* @step
		* @result
		**/
	   result := &model.{{.DaoNameUp}}{}
   
	   /**
		* @step
		* @执行结果
		**/
	   if err := CreateCommonDao().DbClient().Where({{.DaoName}}Model).First(result).Error; err != nil {
		   logger.Logger().Errorf("{{.DaoName}}Dao GetList Err : %+v", err)
		   return nil, err
	   }
	   return result, nil
   }
   
   /**
	* @description: Save
	* @param {context.Context} ctx
	* @param {*model.{{.DaoNameUp}}} {{.DaoName}}Model
	* @author: Jerry.Yang
	* @date: 2023-05-10 17:51:41
	* @return {*}
	*/
   func ({{.FirstDaoName}} *{{.DaoNameUp}}) Save(ctx context.Context, {{.DaoName}}Model *model.{{.DaoNameUp}}) (int64, error) {
   
	   /**
		* @step
		* @执行结果
		**/
	   if err := CreateCommonDao().DbClient().Save({{.DaoName}}Model).Error; err != nil {
		   logger.Logger().Errorf("{{.DaoName}}Dao GetList Err : %+v", err)
		   return 0, err
	   }
	   return {{.DaoName}}Model.UID, nil
   }
   
   /**
	* @description: Delete
	* @param {context.Context} ctx
	* @param {*model.{{.DaoNameUp}}} {{.DaoName}}Model
	* @author: Jerry.Yang
	* @date: 2023-05-10 17:53:14
	* @return {*}
	*/
   func ({{.FirstDaoName}} *{{.DaoNameUp}}) Delete(ctx context.Context, {{.DaoName}}Model *model.{{.DaoNameUp}}) (bool, error) {
   
	   /**
		* @step
		* @执行结果
		**/
	   if err := CreateCommonDao().DbClient().Where({{.DaoName}}Model).Update("is_deleted = ?", 0).Error; err != nil {
		   logger.Logger().Errorf("{{.DaoName}}Dao GetList Err : %+v", err)
		   return false, err
	   }
	   return true, nil
   }
   `
}

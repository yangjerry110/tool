/*
 * @Author: Jerry.Yang
 * @Date: 2023-05-10 17:55:34
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-07-05 16:09:38
 * @Description:new dao
 */
package dao

import (
	"fmt"

	"github.com/yangjerry110/tool/cmd/gen-tool/templates"
)

type NewDao interface {
	SaveTemplate(path string, projectImportPath string, daoName string, daoFileName string) error
	GetTemplate() string
	AppendFuncTemplate(path string, daoName string, baseDaoName string) error
	GetAppendTemplate() string
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
		Time              string
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
	data := &Data{ProjectImportPath: projectImportPath, DaoName: daoName, DaoNameUp: daoNameUp, FirstDaoName: FirstDaoName, Time: templates.CreateCommonTemplate().GetFormatNowTime()}
	return templates.CreateCommonTemplate().SaveTemplate(path, daoFileName, n.GetTemplate(), data)
}

/**
 * @description: AppendFuncTemplate
 * @param {string} path
 * @param {string} daoName
 * @param {string} baseDaoName
 * @author: Jerry.Yang
 * @date: 2023-05-16 11:22:40
 * @return {*}
 */
func (n *New) AppendFuncTemplate(path string, daoName string, baseDaoName string) error {

	/**
	 * @step
	 * @解析需要添加的文件
	 **/
	basePath := fmt.Sprintf("%s/%sDao.go", path, baseDaoName)

	/**
	 * @step
	 * @定义数据结构
	 **/
	type Data struct {
		BaseDaoName        string
		FirstBaseDaoNameUp string
		DaoName            string
		DaoNameUp          string
		Time               string
	}

	/**
	 * @step
	 * @定义大写的参数
	 **/
	firstBaseDaoNameUp := templates.CreateCommonTemplate().FirstUpper(baseDaoName)
	DaoNameUp := templates.CreateCommonTemplate().FirstUpper(daoName)

	/**
	 * @step
	 * @进行赋值
	 **/
	data := &Data{FirstBaseDaoNameUp: firstBaseDaoNameUp, BaseDaoName: baseDaoName, DaoNameUp: DaoNameUp, DaoName: daoName, Time: templates.CreateCommonTemplate().GetFormatNowTime()}
	return templates.CreateCommonTemplate().AppendTemplate(basePath, n.GetAppendTemplate(), data)

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
	   result := []*model.{{.DaoNameUp}}{}
   
	   /**
		* @step
		* @执行结果
		**/
	   if err := CreateCommonDao().DbClient().Where({{.DaoName}}Model).Where("is_deleted = ?",model.No_Deleted).Find(result).Error; err != nil {
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
	   if err := CreateCommonDao().DbClient().Where({{.DaoName}}Model).Where("is_deleted = ?",model.No_Deleted).First(result).Error; err != nil {
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
	   if err := CreateCommonDao().DbClient().Save({{.DaoName}}Model).Error; err != nil {
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
	   if err := CreateCommonDao().DbClient().Where({{.DaoName}}Model).Update("is_deleted = ?", model.Is_Deleted).Error; err != nil {
		   logger.Logger().Errorf("{{.DaoName}}Dao Delete{{.DaoNameUp}} Err : %+v", err)
		   return false, err
	   }
	   return true, nil
   }
   `
}

/**
 * @description: GetAppendTemplate
 * @author: Jerry.Yang
 * @date: 2023-05-16 11:16:14
 * @return {*}
 */
func (n *New) GetAppendTemplate() string {
	return `/**
	* @description: Get{{.DaoNameUp}}List
	* @param {context.Context} ctx
	* @param {*model.{{.DaoNameUp}}} {{.DaoName}}Model
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/
   func ({{.FirstBaseDaoName}} *{{.BaseDaoNameUp}}) Get{{.DaoNameUp}}List(ctx context.Context, {{.DaoName}}Model *model.{{.DaoNameUp}}) ([]*model.{{.DaoNameUp}}, error) {
   
	   /**
		* @step
		* @result
		**/
	   result := []*model.{{.DaoNameUp}}{}
   
	   /**
		* @step
		* @执行结果
		**/
	   if err := CreateCommonDao().DbClient().Where({{.DaoName}}Model).Where("is_deleted = ?",model.No_Deleted).Find(result).Error; err != nil {
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
   func ({{.FirstBaseDaoName}} *{{.BaseDaoNameUp}}) Get{{.DaoNameUp}}Info(ctx context.Context,{{.DaoName}}Model *model.{{.DaoNameUp}}) (*model.{{.DaoNameUp}}, error) {
   
	   /**
		* @step
		* @result
		**/
	   result := &model.{{.DaoNameUp}}{}
   
	   /**
		* @step
		* @执行结果
		**/
	   if err := CreateCommonDao().DbClient().Where({{.DaoName}}Model).Where("is_deleted = ?",model.No_Deleted).First(result).Error; err != nil {
		   logger.Logger().Errorf("{{.DaoName}}Dao  Get{{.DaoNameUp}}Info Err : %+v", err)
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
   func ({{.FirstBaseDaoName}} *{{.BaseDaoNameUp}}) Save{{.DaoNameUp}} (ctx context.Context, {{.DaoName}}Model *model.{{.DaoNameUp}}) (int64, error) {
   
	   /**
		* @step
		* @执行结果
		**/
	   if err := CreateCommonDao().DbClient().Save({{.DaoName}}Model).Error; err != nil {
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
   func ({{.FirstBaseDaoName}} *{{.BaseDaoNameUp}}) Delete{{.DaoNameUp}} (ctx context.Context, {{.DaoName}}Model *model.{{.DaoNameUp}}) (bool, error) {
   
	   /**
		* @step
		* @执行结果
		**/
	   if err := CreateCommonDao().DbClient().Where({{.DaoName}}Model).Update("is_deleted = ?", model.Is_Deleted).Error; err != nil {
		   logger.Logger().Errorf("{{.DaoName}}Dao Delete{{.DaoNameUp}} Err : %+v", err)
		   return false, err
	   }
	   return true, nil
   }`
}

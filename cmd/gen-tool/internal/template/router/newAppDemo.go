/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-19 14:53:51
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-28 16:29:52
 * @Description: newApp demo router
 */
package router

import (
	"fmt"

	"github.com/yangjerry110/tool/cmd/gen-tool/internal/config"
	"github.com/yangjerry110/tool/cmd/gen-tool/internal/template"
)

type NewAppDemoRouter struct{}

/**
 * @description: NewAppDemoRouter
 * @author: Jerry.Yang
 * @date: 2023-12-19 14:55:24
 * @return {*}
 */
func (n *NewAppDemoRouter) New() error {

	// The structure that needs to be rendered
	type Data struct {
		Time       string
		ImportPath string
	}

	// Set Data
	data := &Data{}
	data.Time = template.GetFormatNowTime()
	data.ImportPath = config.ProjectImportPathConf.ImportPath

	filePath := fmt.Sprintf("%s/router", config.ProjectPathConf.Path)
	return template.SaveTemplate(filePath, "demoRouter.go", n.getTemplate(), data)
}

/**
 * @description: getTemplate
 * @author: Jerry.Yang
 * @date: 2023-12-19 14:55:14
 * @return {*}
 */
func (n *NewAppDemoRouter) getTemplate() string {
	return `
/*
 * @Author: Jerry.Yang
 * @Date: {{.Time}}
 * @LastEditors: Jerry.Yang
 * @LastEditTime: {{.Time}}
 * @Description: demo
 */
package router

import (
	"git.qutoutiao.net/gopher/qms/pkg/qlog"
	"{{.ImportPath}}/vo/protobuf"
	"github.com/gin-gonic/gin"
	toolRouter "github.com/yangjerry110/tool/router"
	"net/http"
)

type Demo struct{
	HttpServer protobuf.DemoApiServer
}

/**
 * @description: RouterName
 * @author: Jerry.Yang
 * @date: {{.Time}}
 * @return {*}
 */
func (*Demo) RouterName() string {
	return "demo"
}

/**
 * @description: CreateRouter
 * @author: Jerry.Yang
 * @date: {{.Time}}
 * @return {*}
 */
func (d *Demo) RegisterHTTP(router gin.IRouter) {

	/**
	* @step
	* @POST /api/demo
	* @增加Demo的详细描述
	**/
	router.POST("/api/demo", d.AddDemo)

	/**
	* @step
	* @DELETE /api/demo/:id
	* @删除Demo的详细描述
	**/
	router.DELETE("/api/demo/:id", d.DeleteDemo)

	/**
	* @step
	* @PATCH /api/demo/:id
	* @更新Demo的详细描述
	**/
	router.PATCH("/api/demo/:id", d.UpdateDemo)

	/**
	* @step
	* @GET /api/demo
	* @查找Demo的详细描述
	**/
	router.GET("/api/demo", d.GetDemo)
}

/**
 * @description: RegisterService
 * @param {router.RouterHTTPService} service
 * @author: Jerry.Yang
 * @date: {{.Time}}
 * @return {*}
 */
func (d *Demo) RegisterHTTPService(service toolRouter.RouterHTTPService) {
	d.HttpServer = service.(protobuf.DemoApiServer)
}

// AddDemo 增加Demo的详细描述
// @ID AddDemo
// @Summary 增加Demo的详细描述
// @Tags Demo
// @Param input body protobuf.AddDemoReq false " - "
// @Success 200 {object} protobuf.Empty
// @Router /api/demo [POST]
func (d *Demo) AddDemo(ctx *gin.Context) {

	/**
	 * @step
	 * @inputVo
	 **/
	inputVo := &protobuf.AddDemoReq{}

	/**
	 * @step
	 * @should bind
	 **/
	if err := ctx.ShouldBind(inputVo); err != nil {
		qlog.Errorf("AddDemo shouldBind Err : %+v", err)
		return
	}

	/**
	 * @step
	 * @调用service
	 **/
	outputVo, err := d.HttpServer.AddDemo(ctx, inputVo)
	if err != nil {
		qlog.Errorf("DemoService AddDemo Err : %+v", err)
		ctx.JSON(http.StatusOK, &protobuf.Empty{})
		return
	}

	/**
	 * @step
	 * @return
	 **/
	ctx.JSON(http.StatusOK, outputVo)
}

// DeleteDemo 删除Demo的详细描述
// @ID DeleteDemo
// @Summary 删除Demo的详细描述
// @Tags Demo
// @Success 200 {object} protobuf.Empty
// @Router /api/demo/:id [DELETE]
func (d *Demo) DeleteDemo(ctx *gin.Context) {

	/**
	 * @step
	 * @inputVo
	 **/
	inputVo := &protobuf.DeleteDemoReq{}

	/**
	 * @step
	 * @should bind
	 **/
	if err := ctx.ShouldBind(inputVo); err != nil {
		qlog.Errorf("DeleteDemo shouldBind Err : %+v", err)
		return
	}

	/**
	 * @step
	 * @调用service
	 **/
	outputVo, err := d.HttpServer.DeleteDemo(ctx, inputVo)
	if err != nil {
		qlog.Errorf("DemoService DeleteDemo Err : %+v", err)
		ctx.JSON(http.StatusOK, &protobuf.Empty{})
		return
	}

	/**
	 * @step
	 * @return
	 **/
	ctx.JSON(http.StatusOK, outputVo)
}

// UpdateDemo 更新Demo的详细描述
// @ID UpdateDemo
// @Summary 更新Demo的详细描述
// @Tags Demo
// @Success 200 {object} protobuf.Empty
// @Router /api/demo/:id [PATCH]
func (d *Demo) UpdateDemo(ctx *gin.Context) {

	/**
	 * @step
	 * @inputVo
	 **/
	inputVo := &protobuf.UpdateDemoReq{}

	/**
	 * @step
	 * @should bind
	 **/
	if err := ctx.ShouldBind(inputVo); err != nil {
		qlog.Errorf("UpdateDemo shouldBind Err : %+v", err)
		return
	}

	/**
	 * @step
	 * @调用service
	 **/
	outputVo, err := d.HttpServer.UpdateDemo(ctx, inputVo)
	if err != nil {
		qlog.Errorf("DemoService UpdateDemo Err : %+v", err)
		ctx.JSON(http.StatusOK, &protobuf.Empty{})
		return
	}

	/**
	 * @step
	 * @return
	 **/
	ctx.JSON(http.StatusOK, outputVo)
}

// GetDemo 查找Demo的详细描述
// @ID GetDemo
// @Summary 查找Demo的详细描述
// @Tags Demo
// @Param id query int32 false "-"
// @Success 200 {object} protobuf.GetDemoResp
// @Router /api/demo [GET]
func (d *Demo) GetDemo(ctx *gin.Context) {

	/**
	 * @step
	 * @inputVo
	 **/
	inputVo := &protobuf.GetDemoReq{}

	/**
	 * @step
	 * @should bind
	 **/
	if err := ctx.ShouldBind(inputVo); err != nil {
		qlog.Errorf("GetDemo shouldBind Err : %+v", err)
		return
	}

	/**
	 * @step
	 * @调用service
	 **/
	outputVo, err := d.HttpServer.GetDemo(ctx, inputVo)
	if err != nil {
		qlog.Errorf("DemoService GetDemo Err : %+v", err)
		ctx.JSON(http.StatusOK, &protobuf.GetDemoResp{})
		return
	}

	/**
	 * @step
	 * @return
	 **/
	ctx.JSON(http.StatusOK, outputVo)
}
`
}

/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-25 16:49:20
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-04-25 17:25:24
 * @Description: baseVo
 */
package vo

/**
 * @description: CreateTestInputVo
 * @param {...TestInputVo} TestInputVos
 * @author: Jerry.Yang
 * @date: 2023-04-25 17:24:35
 * @return {*}
 */
func CreateTestInputVo(TestInputVos ...TestInputVo) TestInputVo {
	if len(TestInputVos) == 0 {
		return &TestInput{}
	}
	return TestInputVos[0]
}

/**
 * @description: CreateTestOutputVo
 * @param {...TestOutputVo} TestOutputVos
 * @author: Jerry.Yang
 * @date: 2023-04-25 17:25:22
 * @return {*}
 */
func CreateTestOutputVo(TestOutputVos ...TestOutputVo) TestOutputVo {
	if len(TestOutputVos) == 0 {
		return &TestOutput{}
	}
	return TestOutputVos[0]
}

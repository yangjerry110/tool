/*
 * @Author: Jerry.Yang
 * @Date: 2023-05-18 16:01:59
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-22 15:50:02
 * @Description: baseProtobuf
 */
package protobuf

/**
 * @description: CreateDemoProtobuf
 * @param {...DemoProtobuf} DemoProtobufs
 * @author: Jerry.Yang
 * @date: 2023-05-18 16:02:47
 * @return {*}
 */
func CreateDemoProtobuf(DemoProtobufs ...DemoProtobuf) DemoProtobuf {
	if len(DemoProtobufs) == 0 {
		return &Demo{}
	}
	return DemoProtobufs[0]
}

/**
 * @description: CreateHttpProtobuf
 * @param {...HttpProtobuf} HttpProtobufs
 * @author: Jerry.Yang
 * @date: 2023-05-22 15:50:09
 * @return {*}
 */
func CreateHttpProtobuf(HttpProtobufs ...HttpProtobuf) HttpProtobuf {
	if len(HttpProtobufs) == 0 {
		return &Http{}
	}
	return HttpProtobufs[0]
}

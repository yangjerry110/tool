/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-15 14:26:09
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-15 14:26:13
 * @Description: 
 */
package testvo


type GetAccountReq struct {
	AccountId int32  `protobuf:"varint,1,opt,name=account_id,json=accountId,proto3" json:"account_id" form:"account_id" uri:"account_id"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" form:"name" uri:"name"`
	Type      string `protobuf:"bytes,3,opt,name=type,proto3" json:"type" form:"type" uri:"type"`
}

type GetAccountResp struct {
	RetCode   int32                 `protobuf:"varint,1,opt,name=retCode,proto3" json:"retCode" form:"retCode" uri:"retCode"`
	RetMsg    string                `protobuf:"bytes,2,opt,name=retMsg,proto3" json:"retMsg" form:"retMsg" uri:"retMsg"`
	RetResult *GetAccountResultResp `protobuf:"bytes,3,opt,name=retResult,proto3" json:"retResult" form:"retResult" uri:"retResult"`
}

type GetAccountResultResp struct {
	DisplayName   string `protobuf:"bytes,1,opt,name=DisplayName,proto3" json:"DisplayName" form:"DisplayName" uri:"DisplayName"`
	Email         string `protobuf:"bytes,2,opt,name=Email,proto3" json:"Email" form:"Email" uri:"Email"`
	UpdateDate    string `protobuf:"bytes,3,opt,name=UpdateDate,proto3" json:"UpdateDate" form:"UpdateDate" uri:"UpdateDate"`
	MobilePhone   string `protobuf:"bytes,4,opt,name=MobilePhone,proto3" json:"MobilePhone" form:"MobilePhone" uri:"MobilePhone"`
	UserId        string `protobuf:"bytes,5,opt,name=UserId,proto3" json:"UserId" form:"UserId" uri:"UserId"`
	Comments      string `protobuf:"bytes,6,opt,name=Comments,proto3" json:"Comments" form:"Comments" uri:"Comments"`
	LastLoginDate string `protobuf:"bytes,7,opt,name=LastLoginDate,proto3" json:"LastLoginDate" form:"LastLoginDate" uri:"LastLoginDate"`
	CreateDate    string `protobuf:"bytes,8,opt,name=CreateDate,proto3" json:"CreateDate" form:"CreateDate" uri:"CreateDate"`
	UserName      string `protobuf:"bytes,9,opt,name=UserName,proto3" json:"UserName" form:"UserName" uri:"UserName"`
}
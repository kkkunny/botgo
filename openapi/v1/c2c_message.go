package v1

import (
	"context"

	"github.com/tencent-connect/botgo/dto"
)

// PostC2CMessage 发送单聊消息
func (o *openAPI) PostC2CMessage(ctx context.Context, userID string, msg *dto.C2CMessageToCreate) (*dto.C2CMessageToReply, error) {
	resp, err := o.request(ctx).
		SetResult(dto.C2CMessageToReply{}).
		SetPathParam("user_id", userID).
		SetBody(msg).
		Post(o.getURL(userMessageURI))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*dto.C2CMessageToReply), nil
}

// PostGroupMessage 发送群组消息
func (o *openAPI) PostGroupMessage(ctx context.Context, groupID string, msg *dto.C2CMessageToCreate) (*dto.C2CMessageToReply, error) {
	resp, err := o.request(ctx).
		SetResult(dto.C2CMessageToReply{}).
		SetPathParam("group_id", groupID).
		SetBody(msg).
		Post(o.getURL(groupMessageURI))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*dto.C2CMessageToReply), nil
}

// PostC2CFile 上传单聊富文本
func (o *openAPI) PostC2CFile(ctx context.Context, userID string, msg *dto.C2CFileToCreate) (*dto.C2CFileToReply, error) {
	resp, err := o.request(ctx).
		SetResult(dto.C2CFileToReply{}).
		SetPathParam("user_id", userID).
		SetBody(msg).
		Post(o.getURL(userFileURI))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*dto.C2CFileToReply), nil
}

// PostGroupFile 上传群组富文本
func (o *openAPI) PostGroupFile(ctx context.Context, groupID string, msg *dto.C2CFileToCreate) (*dto.C2CFileToReply, error) {
	resp, err := o.request(ctx).
		SetResult(dto.C2CFileToReply{}).
		SetPathParam("group_id", groupID).
		SetBody(msg).
		Post(o.getURL(groupFileURI))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*dto.C2CFileToReply), nil
}

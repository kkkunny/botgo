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

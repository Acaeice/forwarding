package pkg

import (
	"errors"
)

// MessageCreateParam Message创建参数
type MessageCreateParam struct {
	MessageBody string `json:"messageBody" form:"messageBody"` //短信正文
	TimeReceipt string `json:"timeReceipt" form:"timeReceipt"` //收件时间
	Sender      string `json:"sender" form:"sender"`           //发件人

}

// IsValid 校验参数
func (param MessageCreateParam) IsValid() error {

	if len(param.MessageBody) == 0 {
		return errors.New("MessageCreateParam.MessageBody不合法")
	}

	if len(param.TimeReceipt) == 0 {
		return errors.New("MessageCreateParam.TimeReceipt不合法")
	}

	if len(param.Sender) == 0 {
		return errors.New("MessageCreateParam.Sender不合法")
	}

	return nil
}

// MessageModifyParam Message修改参数
type MessageModifyParam struct {
	ID          uint   `json:"id" form:"id"`
	MessageBody string `json:"messageBody" form:"messageBody"` //短信正文
	TimeReceipt string `json:"timeReceipt" form:"timeReceipt"` //收件时间
	Sender      string `json:"sender" form:"sender"`           //发件人

}

// IsValid 校验参数
func (param MessageModifyParam) IsValid() error {
	if param.ID == 0 {
		return errors.New("MessageModifyParam.id不合法")
	}

	if len(param.MessageBody) == 0 {
		return errors.New("MessageModifyParam.MessageBody不合法")
	}

	if len(param.TimeReceipt) == 0 {
		return errors.New("MessageModifyParam.TimeReceipt不合法")
	}

	if len(param.Sender) == 0 {
		return errors.New("MessageModifyParam.Sender不合法")
	}

	return nil
}

//MessageUpdateParam Message单独修改参数
type MessageUpdateParam struct {
	ID          uint   `json:"id" form:"id"`
	MessageBody string `json:"messageBody" form:"messageBody"` //短信正文
	TimeReceipt string `json:"timeReceipt" form:"timeReceipt"` //收件时间
	Sender      string `json:"sender" form:"sender"`           //发件人

}

// IsValid 校验参数
func (param MessageUpdateParam) IsValid() error {
	if param.ID == 0 {
		return errors.New("MessageUpdateParam.id不合法")
	}
	return nil
}

// MessageSearchParam Message查询参数
type MessageSearchParam struct {
	MessageBody string `json:"messageBody" form:"messageBody"` //短信正文
	TimeReceipt string `json:"timeReceipt" form:"timeReceipt"` //收件时间
	Sender      string `json:"sender" form:"sender"`           //发件人

	PageSize uint `json:"pageSize" form:"pageSize"`
	Page     uint `json:"page" form:"page"`
}

// MessageListParam Message查询参数
type MessageListParam struct {
	MessageBody string `json:"messageBody" form:"messageBody"` //短信正文
	TimeReceipt string `json:"timeReceipt" form:"timeReceipt"` //收件时间
	Sender      string `json:"sender" form:"sender"`           //发件人

}

// MessageDeleteParam Message删除参数
type MessageDeleteParam struct {
	ID uint `json:"id" form:"id"`
}

// IsValid 校验参数
func (param MessageDeleteParam) IsValid() error {
	if param.ID == 0 {
		return errors.New("MessageDeleteParam.id不合法")
	}
	return nil
}

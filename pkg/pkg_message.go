package pkg

//Message 短信转发参数
type Message struct {
	MessageBody string `json:"messageBody" form:"messageBody"` //短信正文
	TimeReceipt string `json:"timeReceipt" form:"timeReceipt"` //收件时间
	Sender      string `json:"sender" form:"sender"`           //发件人
}

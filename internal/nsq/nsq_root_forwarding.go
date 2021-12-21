package nsq

import (
	"encoding/json"

	"code.meikeland.com/me/forwarding/pkg"
	"github.com/meikeland/logger"
)

//收到短信进行投递
func MessageForward(message *pkg.Message) {
	body, _ := json.Marshal(message)
	if err := producer.Publish(messageForwardTopic, body); err != nil {
		logger.Infof("publish nsq by order complete error : %s", err.Error())
	}
}

package forwarding

import (
	"code.meikeland.com/me/forwarding/internal/nsq"
	"code.meikeland.com/me/forwarding/pkg"
	"github.com/gin-gonic/gin"
	"github.com/meikeland/errkit"
)

//获取消息
func getMsg(c *gin.Context) {
	param := pkg.Message{}
	if err := c.ShouldBind(&param); err != nil {
		fail(c, errkit.Wrap(err, "参数错误"))
		return
	}

	go nsq.MessageForward(&param)

	ok(c, resp{
		"msg": "发送成功!",
	})
}

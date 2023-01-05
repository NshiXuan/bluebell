package models

// 定义请求的参数结构体

// 请求postList列表的参数
const (
	OrderTime  = "time"
	OrderScore = "score"
)

// 注册请求参数
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`                     // 用户名
	Password   string `json:"password" binding:"required"`                     // 密码
	RePassword string `json:"re_password" binding:"required,eqfield=Password"` // 确认密码器
}

// 登录请求参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"` // 用户名
	Password string `json:"password" binding:"required"` // 密码
}

// 投票数据参数
type ParamVoteDate struct {
	// UserID 从请求中获取当前用户的ID
	PostID    string `json:"post_id" binding:"required"`              // 帖子ID
	Direction int8   `json:"direction,string" binding:"oneof=1 0 -1"` // 赞成票(1) 反对票(-1) 取消投票(0)
}

// ParamPostList 获取帖子列表的query string参数 使用到form来指定获取query参数
type ParamPostList struct {
	CommunityID int64  `json:"community_id" form:"community_id"` // 社区id
	Page        int64  `json:"page" form:"page"`                 // 页码
	Size        int64  `json:"size" form:"size"`                 // 每页的数据量
	Order       string `json:"order" form:"order"`               // 排序依据
}

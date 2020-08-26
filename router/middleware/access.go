package middleware

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var (
	NoAccessErr = errors.New("您当前没有权限访问此API")
)

func AccessAuthority(items ...string) gin.HandlerFunc {
	if len(items) == 0 {
		panic("func AccessAuthority args (items == 0)")
	}

	return func(c *gin.Context) {
		// member, ok := c.Keys["member"].(model.Member)
		// if !ok {
		// 	return
		// }

		// authorityItems, err := dao.Current.FindAuthorityItemsByMemberID(member.ID)
		// if err != nil || len(authorityItems) == 0 {
		// 	return
		// }

		// for _, item := range items {
		// 	for _, authority := range authorityItems {
		// 		if item == authority.Name {
		// 			goto END
		// 		}
		// 	}

		// 	return
		// END:
		// }

	}
}

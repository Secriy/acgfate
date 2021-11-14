package cache

import (
	"fmt"

	"acgfate/util"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type WordDao struct{}

func (w *WordDao) Like(c *gin.Context, wid, uid interface{}) {
	keyLike := fmt.Sprintf("%s:%s", KeyWordLike, util.ToString(wid))
	rdb.SAdd(c, keyLike, uid)
	// for ranking
	keyRank := fmt.Sprintf("%s:%s", KeyWordRank, util.DateToday())
	rdb.ZIncrBy(c, keyRank, 1, fmt.Sprintf("%s", wid))
}

func (w *WordDao) Unlike(c *gin.Context, wid, uid interface{}) {
	key := fmt.Sprintf("%s:%s", KeyWordLike, util.ToString(wid))
	rdb.SRem(c, key, uid)
	// for ranking
	keyRank := fmt.Sprintf("%s:%s", KeyWordRank, util.DateToday())
	rdb.ZIncrBy(c, keyRank, -1, fmt.Sprintf("%s", wid))
}

// IsLiked check if user like that word.
func (w *WordDao) IsLiked(c *gin.Context, wid, uid interface{}) bool {
	keyLike := fmt.Sprintf("%s:%s", KeyWordLike, util.ToString(wid))
	cmd := rdb.SIsMember(c, keyLike, uid)
	return cmd.Val()
}

// Likes return the number of likes.
func (w *WordDao) Likes(c *gin.Context, wid interface{}) int64 {
	keyLike := fmt.Sprintf("%s:%s", KeyWordLike, util.ToString(wid))
	cmd := rdb.SCard(c, keyLike)
	return cmd.Val()
}

// Trend return the string slice of wid order by score.
func (w *WordDao) Trend(c *gin.Context) []string {
	keyRank := fmt.Sprintf("%s:%s", KeyWordRank, util.DateToday())
	cmd := rdb.ZRevRangeByScore(c, keyRank, &redis.ZRangeBy{
		Min:    "1",
		Max:    "+inf",
		Offset: 0,
		Count:  100,
	})
	return cmd.Val()
}

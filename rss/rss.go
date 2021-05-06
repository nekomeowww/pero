package rss

import (
	"github.com/nekomeowww/pero/config"
	"github.com/nekomeowww/pero/logger"
	"github.com/robfig/cron/v3"
)

// Feed 新的订阅源
func Feed(conf config.Config) *cron.Cron {
	c := cron.New()
	c.AddFunc("@every 10s", func() {
		for i := range conf.RSS {
			newFeedFromURI(conf.RSS[i])
		}
	})
	c.Start()
	return c
}

func newFeedFromURI(rss config.RSS) {
	logger.Infof("更新订阅源：%s %s", rss.Name, rss.URI)
	// fp := gofeed.NewParser()
	// feed, err := fp.ParseURL(rss.URI)
	// if err != nil {
	// 	logger.Error(err)
	// 	// PASS
	// }

	// for _, v := range feed.Items {
	// t, err := time.Parse(time.RFC1123Z, v.Published)
	// if err != nil {
	// logger.Error(err)
	// PASS
	// }
	// }
}

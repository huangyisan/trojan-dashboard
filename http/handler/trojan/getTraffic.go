package trojan

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"trojan-dashboard/pkg/common"
	"trojan-dashboard/pkg/configs"
	"trojan-dashboard/pkg/meInfluxdb"
	"trojan-dashboard/pkg/meTrojan"
)

func GetAllTraffic(c *gin.Context) {
	i := meInfluxdb.NewInfluxDBClient(
		configs.InfluxDBConf.Url,
		configs.InfluxDBConf.AuthToken,
		configs.InfluxDBConf.Org,
		configs.InfluxDBConf.Bucket)
	startTime := c.DefaultQuery("start", common.GetMonthFirstDay())
	endTime := c.DefaultQuery("end", common.GetCurrentDay())
	res := meTrojan.QueryAllTraffic(i, startTime, endTime)
	c.JSON(http.StatusOK, res)
}

func GetAllTrafficByGroup(c *gin.Context) {
	i := meInfluxdb.NewInfluxDBClient(
		configs.InfluxDBConf.Url,
		configs.InfluxDBConf.AuthToken,
		configs.InfluxDBConf.Org,
		configs.InfluxDBConf.Bucket)
	startTime := c.DefaultQuery("start", common.GetMonthFirstDay())
	endTime := c.DefaultQuery("end", common.GetCurrentDay())
	res := meTrojan.QueryAllTrafficByGroup(i, startTime, endTime)
	c.JSON(http.StatusOK, res)

}

func GetTrafficByTag(c *gin.Context) {
	i := meInfluxdb.NewInfluxDBClient(
		configs.InfluxDBConf.Url,
		configs.InfluxDBConf.AuthToken,
		configs.InfluxDBConf.Org,
		configs.InfluxDBConf.Bucket)
	startTime := c.DefaultQuery("start", common.GetMonthFirstDay())
	endTime := c.DefaultQuery("end", common.GetCurrentDay())
	res := meTrojan.QueryTrafficByTag(i, startTime, endTime, c.Param("tag"))
	c.JSON(http.StatusOK, res)
}

func GetTrafficByGroup(c *gin.Context) {
	i := meInfluxdb.NewInfluxDBClient(
		configs.InfluxDBConf.Url,
		configs.InfluxDBConf.AuthToken,
		configs.InfluxDBConf.Org,
		configs.InfluxDBConf.Bucket)
	startTime := c.DefaultQuery("start", common.GetMonthFirstDay())
	endTime := c.DefaultQuery("end", common.GetCurrentDay())
	res := meTrojan.QueryTrafficByGroup(i, startTime, endTime, c.Param("group"))
	c.JSON(http.StatusOK, res)
}

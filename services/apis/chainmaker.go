package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"zeusnet.com/zhf/zeusnet/modules/config"
	"zeusnet.com/zhf/zeusnet/modules/defence/info"
)

func StartAttack(c *gin.Context) {
	targetPeerId, err := info.InformationInstance.TbftConsensusImpl.StartAttack()
	if err != nil {
		fmt.Printf("error: %v", err)
		c.JSON(500, gin.H{
			"message": fmt.Errorf("start attack failed: %v", err),
		})
	} else {
		c.JSON(200, gin.H{
			"target":  config.EnvLoaderInstance.PeerIdToContainerNameMapping[targetPeerId],
			"message": "successfully launch insider attack",
		})
	}
}

func StopAttack(c *gin.Context) {
	err := info.InformationInstance.TbftConsensusImpl.StopAttack()
	if err != nil {
		fmt.Printf("error: %v", err)
		c.JSON(500, gin.H{
			"message": fmt.Errorf("stop attack failed: %v", err),
		})
	} else {
		c.JSON(200, gin.H{
			"message": "successfully stop insider attack",
		})
	}
}

package network

import (
	"time"

	"github.com/pagu-project/Pagu/internal/engine/command"
	"github.com/pagu-project/Pagu/internal/entity"
	utils2 "github.com/pagu-project/Pagu/pkg/utils"
)

func (n *Network) networkHealthHandler(
	_ *entity.User,
	cmd *command.Command,
	_ map[string]string,
) command.CommandResult {
	lastBlockTime, lastBlockHeight := n.clientMgr.GetLastBlockTime()
	lastBlockTimeFormatted := time.Unix(int64(lastBlockTime), 0).Format("02/01/2006, 15:04:05")
	currentTime := time.Now()
	timeDiff := currentTime.Unix() - int64(lastBlockTime)

	healthStatus := true
	if timeDiff > 15 {
		healthStatus = false
	}

	var status string
	if healthStatus {
		status = "Healthy✅"
	} else {
		status = "UnHealthy❌"
	}

	return cmd.SuccessfulResultF("Network is %s\nCurrentTime: %v\n"+
		"LastBlockTime: %v\nTime Diff: %v\nLast Block Height: %v",
		status, currentTime.Format("02/01/2006, 15:04:05"), lastBlockTimeFormatted, timeDiff,
		utils2.FormatNumber(int64(lastBlockHeight)))
}

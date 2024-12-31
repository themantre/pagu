package zealy

import (
	"fmt"

	"github.com/pagu-project/pagu/internal/engine/command"
	"github.com/pagu-project/pagu/internal/entity"
	"github.com/pagu-project/pagu/pkg/log"
)

func (z *Zealy) claimHandler(caller *entity.User, cmd *command.Command, args map[string]string) command.CommandResult {
	user, err := z.db.GetZealyUser(caller.UserID)
	if err != nil {
		return cmd.ErrorResult(err)
	}

	if user.IsClaimed() {
		return cmd.FailedResultF("You already claimed your reward: https://pacviewer.com/transaction/%s",
			user.TxHash)
	}

	address := args["address"]
	txHash, err := z.wallet.TransferTransaction(address, "Pagu Zealy reward distribution", user.Amount)
	if err != nil {
		log.Error("error in transfer zealy reward", "err", err)
		transferErr := fmt.Errorf("failed to transfer zealy reward. Please make sure the address is valid")

		return cmd.ErrorResult(transferErr)
	}

	if err = z.db.UpdateZealyUser(caller.UserID, txHash); err != nil {
		return cmd.ErrorResult(err)
	}

	return cmd.SuccessfulResultF("Zealy reward claimed successfully: https://pacviewer.com/transaction/%s",
		txHash)
}

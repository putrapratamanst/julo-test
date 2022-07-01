package wallet

import (
	"julo-test/model/wallet"
	"julo-test/pkg"
	"julo-test/presenter"
	walletResponse "julo-test/presenter/wallet"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (handler *WalletController) Enable(ctx *gin.Context) {
	cid, errCid := ctx.Get("customer_xid")
	if !errCid {
		result := presenter.Response{
			Status: pkg.HTTP_STATUS_FAIL,
			Data: presenter.ErrorResponseMessage{
				Error: pkg.ErrCustomerXID.Error(),
			},
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	//check wallet whether enabled or disabled
	statusWallet, balance := handler.iws.CheckWalletService(cid.(string))
	if !statusWallet {
		result := presenter.Response{
			Status: pkg.HTTP_STATUS_FAIL,
			Data: presenter.ErrorResponseMessage{
				Error: pkg.ErrWalletAlreadyEnabled.Error(),
			},
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	timeNow := time.Now()
	dataWallet := wallet.WalletModel{
		ID:          pkg.GenID(),
		OwnedBy:     cid.(string),
		EnableAt:    &timeNow,
		Status:      pkg.WALLET_ENABLED,
		Balance:     balance,
		CustomerXID: cid.(string),
	}

	handler.iws.EnableWalletService(dataWallet)
	result := presenter.Response{
		Status: pkg.HTTP_STATUS_SUCCESS,
		Data: walletResponse.EnableResponse{
			Wallet: dataWallet,
		},
	}
	ctx.JSON(http.StatusCreated, result)
}

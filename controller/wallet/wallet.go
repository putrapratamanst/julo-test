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
		ID:       pkg.GenID(),
		OwnedBy:  cid.(string),
		EnableAt: &timeNow,
		Status:   pkg.WALLET_ENABLED,
		Balance:  balance,
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

func (handler *WalletController) View(ctx *gin.Context) {
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

	detail, errDetail := handler.iws.ViewWalletService(cid.(string))
	if errDetail != nil {
		result := presenter.Response{
			Status: pkg.HTTP_STATUS_FAIL,
			Data: presenter.ErrorResponseMessage{
				Error: errDetail.Message,
			},
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	if detail.Status == pkg.WALLET_DISABLED {
		result := presenter.Response{
			Status: pkg.HTTP_STATUS_FAIL,
			Data: presenter.ErrorResponseMessage{
				Error: pkg.ErrWalletAlreadyDisabled.Error(),
			},
		}
		ctx.JSON(http.StatusNotFound, result)
		return
	}

	result := presenter.Response{
		Status: pkg.HTTP_STATUS_SUCCESS,
		Data: walletResponse.EnableResponse{
			Wallet: detail,
		},
	}
	ctx.JSON(http.StatusOK, result)
}

func (handler *WalletController) Disable(ctx *gin.Context) {
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

	var input wallet.DisableModel
	err := ValidateRequest(pkg.BIND_TYPE_JSON, ctx, &input)
	if err != nil {
		result := presenter.Response{
			Status:  pkg.HTTP_STATUS_FAIL,
			Message: err.Message,
			Data:    nil,
		}
		ctx.JSON(err.Code, result)
		return
	}

	detail, errDetail := handler.iws.ViewWalletService(cid.(string))
	if errDetail != nil {
		result := presenter.Response{
			Status: pkg.HTTP_STATUS_FAIL,
			Data: presenter.ErrorResponseMessage{
				Error: errDetail.Message,
			},
		}
		ctx.JSON(http.StatusBadRequest, result)
		return
	}

	timeNow := time.Now()
	detail.Status = pkg.WALLET_DISABLED
	detail.DisabledAt = &timeNow

	handler.iws.DisableWalletService(detail)
	result := presenter.Response{
		Status: pkg.HTTP_STATUS_SUCCESS,
		Data: walletResponse.EnableResponse{
			Wallet: wallet.WalletModel{
				ID:         detail.ID,
				OwnedBy:    detail.OwnedBy,
				Status:     detail.Status,
				DisabledAt: detail.DisabledAt,
				Balance:    detail.Balance,
			},
		},
	}
	ctx.JSON(http.StatusOK, result)
}

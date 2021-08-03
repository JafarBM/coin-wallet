package core

import (
	"errors"
	"fmt"
	"request_test/tamrin1/pkg"
	"time"
)

const (
	timeFormat = "2006-01-02 15:04"
)

type WalletData struct {
	wallets []*pkg.Wallet
}

func GetNewWalletService() pkg.WalletService {
	return &WalletData{
		wallets: make([]*pkg.Wallet, 0),
	}
}

func (wd *WalletData) CreateWallet(req *pkg.CreateWalletRequest) (*pkg.CreateWalletResponse, error) {
	if wd.isWalletNameExists(req.Name) {
		return nil, errors.New("wallet name already taken")
	}
	wallet := wd.createNewWallet(req.Name)
	return &pkg.CreateWalletResponse{
		BaseResponse: pkg.BaseResponse{
			Code:    200,
			Message: "Wallet added successfully!",
		},
		Wallet: *wallet,
	}, nil

}

func (wd *WalletData) createNewWallet(name string) *pkg.Wallet {
	wallet := &pkg.Wallet{
		Name:        name,
		Balance:     0.0,
		Coins:       make([]*pkg.Coin, 0),
		LastUpdated: time.Now().Format(timeFormat),
	}
	wd.wallets = append(wd.wallets, wallet)
	return wallet
}

func (wd *WalletData) isWalletNameExists(name string) bool {
	for _, wallet := range wd.wallets {
		if wallet.Name == name {
			return true
		}
	}
	return false
}

func (wd *WalletData) GetWallets() (*pkg.GetWalletsResponse, error) {
	return &pkg.GetWalletsResponse{
		Size:    len(wd.wallets),
		Wallets: wd.wallets,
		BaseResponse: pkg.BaseResponse{
			Code:    200,
			Message: "‫‪All‬‬ ‫‪wallets‬‬ ‫‪received‬‬ ‫!‪successfully‬‬",
		},
	}, nil
}

func (wd *WalletData) UpdateWallet(req *pkg.UpdateWalletRequest) (*pkg.UpdateWalletResponse, error) {
	if !wd.isWalletNameExists(req.OldName) {
		return nil, errors.New("wallet name does not exists")
	}
	var responseWallet *pkg.Wallet
	for _, wallet := range wd.wallets {
		if wallet.Name == req.OldName {
			wallet.Name = req.Name
			wallet.LastUpdated = time.Now().Format(timeFormat)
			responseWallet = wallet
		}
	}
	return &pkg.UpdateWalletResponse{
		BaseResponse: pkg.BaseResponse{
			Code:    200,
			Message: "‫‪Wallet‬‬ ‫‪name‬‬ ‫‪changed‬‬ ‫!‪successfully‬‬",
		},
		Wallet: *responseWallet,
	}, nil
}

func (wd *WalletData) DeleteWallet(req *pkg.DeleteWalletRequest) (*pkg.DeleteWalletResponse, error) {
	if !wd.isWalletNameExists(req.Name) {
		return nil, errors.New("wallet name does not exists")
	}
	var responseWallet *pkg.Wallet
	for index, wallet := range wd.wallets {
		if wallet.Name == req.Name {
			wd.wallets = removeWalletFromWallets(wd.wallets, index)
			wallet.LastUpdated = time.Now().Format(timeFormat)
			responseWallet = wallet
		}
	}
	return &pkg.DeleteWalletResponse{
		BaseResponse: pkg.BaseResponse{
			Code:    200,
			Message: "Wallet deleted (logged out) successfully!",
		},
		Wallet: *responseWallet,
	}, nil
}

func removeWalletFromWallets(mySlice []*pkg.Wallet, index int) []*pkg.Wallet {
	copy(mySlice[index:], mySlice[index+1:])
	mySlice[len(mySlice)-1] = &pkg.Wallet{}
	mySlice = mySlice[:len(mySlice)-1]
	return mySlice
}

func (wd *WalletData) CreateCoinForWallet(
	req *pkg.CreateCoinForWalletRequest) (*pkg.CreateCoinForWalletResponse, error) {

	if !wd.isWalletNameExists(req.WalletName) {
		return nil, errors.New("wallet does not exists")
	}
	if wd.isCoinExistsInWallet(req.WalletName, req.Name, req.Symbol) {
		return nil, errors.New("coin or symbol already exists")
	}
	var coin *pkg.Coin
	for _, wallet := range wd.wallets {
		if wallet.Name == req.WalletName {
			coin = &pkg.Coin{
				Name: req.Name,
				Symbol: req.Symbol,
				Amount: req.Amount,
				Rate: req.Rate,
			}
			wallet.Coins = append(wallet.Coins, coin)
			wallet.LastUpdated = time.Now().Format(timeFormat)
			wallet.Balance += req.Rate * req.Amount
		}
	}

	return &pkg.CreateCoinForWalletResponse{
		Coin: *coin,
		BaseResponse: pkg.BaseResponse{
			Code: 200,
			Message: "Coin added successfully!",
		},
	}, nil
}

func (wd *WalletData) isCoinExistsInWallet(walletName string, coinName string, symbol string) bool {
	for _, wallet := range wd.wallets {
		if wallet.Name == wallet.Name {
			for _, coin := range wallet.Coins {
				if coin.Name == coinName || coin.Symbol == symbol {
					return true
				}
			}
		}
	}
	return false
}

func (wd *WalletData) GetCoinsForWallet(req *pkg.GetCoinsForWalletRequest) (*pkg.GetCoinsForWalletResponse, error) {
	if !wd.isWalletNameExists(req.WalletName) {
		return nil, errors.New("wallet does not exists")
	}
	var responseWallet *pkg.Wallet
	for _, wallet := range wd.wallets {
		if wallet.Name == req.WalletName {
			responseWallet = wallet
		}
	}
	return &pkg.GetCoinsForWalletResponse{
		BaseResponse: pkg.BaseResponse{
			Code: 200,
			Message: "All coins received successfully",
		},
		Wallet: *responseWallet,
	}, nil
}

func (wd *WalletData) UpdateCoinForWallet(
	req *pkg.UpdateCoinForWalletRequest) (*pkg.UpdateCoinForWalletResponse, error) {

	if !wd.isWalletNameExists(req.WalletName) {
		return nil, errors.New("wallet does not exists")
	}
	if !wd.isCoinExistsInWallet(req.WalletName, "", req.OldSymbol) {
		return nil, errors.New("coin does not exists")
	}
	if !wd.isUpdateCoinValid(req.WalletName, req.OldSymbol, req.Name, req.Symbol) {
		return nil, errors.New("can not change coin new name or symbol already taken")
	}

	var responseCoin *pkg.Coin
	for _, wallet := range wd.wallets {
		if wallet.Name == req.WalletName {
			for _, coin := range wallet.Coins {
				if coin.Symbol == req.OldSymbol {
					wallet.Balance += req.Rate * req.Amount
					wallet.Balance -= coin.Amount * coin.Rate
					coin.Name = req.Name
					coin.Symbol = req.Symbol
					coin.Amount = req.Amount
					coin.Rate = req.Rate
					wallet.LastUpdated = time.Now().Format(timeFormat)
					responseCoin = coin
				}
			}
		}
	}

	return &pkg.UpdateCoinForWalletResponse{
		Coin: *responseCoin,
		BaseResponse: pkg.BaseResponse{
			Code: 200,
			Message: "Coin updated successfully!",
		},
	}, nil
}

func (wd *WalletData) isUpdateCoinValid(
	walletName string,
	oldSymbol string,
	newName string,
	newSymbol string) bool {

	fmt.Println(oldSymbol, newName, newSymbol, walletName)

	for _, wallet := range wd.wallets {
		if wallet.Name == walletName {
			for _, coin := range wallet.Coins {
				if (coin.Symbol == newSymbol || coin.Name == newName) && coin.Symbol != oldSymbol {
					return false
				}
			}
		}
	}
	return true
}

func (wd *WalletData) DeleteCoinForWallet(
	req *pkg.DeleteCoinForWalletRequest) (*pkg.DeleteCoinForWalletResponse, error) {

	if !wd.isWalletNameExists(req.WalletName) {
		return nil, errors.New("wallet does not exists")
	}
	if !wd.isCoinExistsInWallet(req.WalletName, "", req.CoinSymbol) {
		return nil, errors.New("coin does not exists")
	}

	var responseCoin *pkg.Coin
	for _, wallet := range wd.wallets {
		if wallet.Name == req.WalletName {
			for index, coin := range wallet.Coins {
				if coin.Symbol == req.CoinSymbol {
					responseCoin = coin
					wallet.Balance -= coin.Amount * coin.Rate
					wallet.Coins = removeCoinFromWallet(wallet.Coins, index)
					wallet.LastUpdated = time.Now().Format(timeFormat)
				}
			}
		}
	}

	return &pkg.DeleteCoinForWalletResponse{
		Coin: *responseCoin,
		BaseResponse: pkg.BaseResponse{
			Code: 200,
			Message: "Coin deleted successfully!",
		},
	}, nil
}

func removeCoinFromWallet(mySlice []*pkg.Coin, index int) []*pkg.Coin {
	copy(mySlice[index:], mySlice[index+1:])
	mySlice[len(mySlice)-1] = &pkg.Coin{}
	mySlice = mySlice[:len(mySlice)-1]
	return mySlice
}


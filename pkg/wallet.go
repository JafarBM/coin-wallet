package pkg

type WalletService interface {
	CreateWallet(req *CreateWalletRequest) (*CreateWalletResponse, error)
	GetWallets() (*GetWalletsResponse, error)
	UpdateWallet(req *UpdateWalletRequest) (*UpdateWalletResponse, error)
	DeleteWallet(req *DeleteWalletRequest) (*DeleteWalletResponse, error)
	CreateCoinForWallet(req *CreateCoinForWalletRequest) (*CreateCoinForWalletResponse, error)
	GetCoinsForWallet(req *GetCoinsForWalletRequest) (*GetCoinsForWalletResponse, error)
	UpdateCoinForWallet(req *UpdateCoinForWalletRequest) (*UpdateCoinForWalletResponse, error)
	DeleteCoinForWallet(req *DeleteCoinForWalletRequest) (*DeleteCoinForWalletResponse, error)
}

type BaseResponse struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}
type Wallet struct {
	Name        string  `json:"name"`
	Balance     float64 `json:"balance"`
	Coins       []*Coin `json:"coins"`
	LastUpdated string  `json:"last_updated"`
}
type Coin struct {
	Name   string  `json:"name"`
	Symbol string  `json:"symbol"`
	Amount float64 `json:"amount"`
	Rate   float64 `json:"rate"`
}

type CreateWalletRequest struct {
	Name string `json:"name"`
}
type CreateWalletResponse struct {
	Wallet
	BaseResponse
}

type GetWalletsResponse struct {
	Size    int       `json:"size"`
	Wallets []*Wallet `json:"wallets"`
	BaseResponse
}

type UpdateWalletRequest struct {
	OldName string
	Name    string `json:"name"`
}
type UpdateWalletResponse struct {
	Wallet
	BaseResponse
}

type DeleteWalletRequest struct {
	Name string
}
type DeleteWalletResponse struct {
	Wallet
	BaseResponse
}

type CreateCoinForWalletRequest struct {
	WalletName string
	Coin
}
type CreateCoinForWalletResponse struct {
	Coin
	BaseResponse
}

type GetCoinsForWalletRequest struct {
	WalletName string
}
type GetCoinsForWalletResponse struct {
	Wallet
	BaseResponse
}

type UpdateCoinForWalletRequest struct {
	WalletName string
	OldSymbol  string
	Coin
}
type UpdateCoinForWalletResponse struct {
	Coin
	BaseResponse
}

type DeleteCoinForWalletRequest struct {
	WalletName string
	CoinSymbol string
}
type DeleteCoinForWalletResponse struct {
	Coin
	BaseResponse
}

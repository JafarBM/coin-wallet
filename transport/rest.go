package transport

import (
	"encoding/json"
	"fmt"
	"net/http"
	"request_test/tamrin1/pkg"
	"strings"
)

type rESTService struct {
	walletService pkg.WalletService
}

func GetNewRestService(walletService pkg.WalletService) *rESTService {
	service := &rESTService{
		walletService: walletService,
	}

	http.HandleFunc("/", service.handleWallets)

	return service
}

func (rs rESTService) handleWallets(w http.ResponseWriter, r *http.Request) {
	urlParams := strings.Split(r.URL.Path, "/")[1:]

	if len(urlParams) > 2 {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if urlParams[0] == "wallets" {
		rs.handleWalletsAPIs(w, r, urlParams)
	} else {
		rs.handleCoinAPIs(w, r, urlParams)
	}
}

func (rs rESTService) handleWalletsAPIs(w http.ResponseWriter, r *http.Request, urlParams []string) {

	var response interface{}
	var err error

	switch r.Method {
	case "GET":
		if len(urlParams) != 1 {
			http.Error(w, "Method not supported!", http.StatusMethodNotAllowed)
			return
		}
		response, err = rs.walletService.GetWallets()
	case "POST":
		if len(urlParams) != 1 {
			http.Error(w, "Method not supported!", http.StatusMethodNotAllowed)
			return
		}
		validatedRequest := getValidRequestForCreateWallet(r)
		response, err = rs.walletService.CreateWallet(validatedRequest)
	case "PUT":
		if len(urlParams) != 2 {
			http.Error(w, "Method not supported!", http.StatusMethodNotAllowed)
			return
		}
		validatedRequest := getValidRequestForUpdateWallet(r, urlParams)
		response, err = rs.walletService.UpdateWallet(validatedRequest)
	case "DELETE":
		if len(urlParams) != 2 {
			http.Error(w, "Method not supported!", http.StatusMethodNotAllowed)
			return
		}
		validatedRequest := getValidRequestForDeleteWallet(urlParams)
		response, err = rs.walletService.DeleteWallet(validatedRequest)
	default:
		http.Error(w, "Method not supported!", http.StatusMethodNotAllowed)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Fprintf(w, string(jsonResponse))
}

func (rs rESTService) handleCoinAPIs(w http.ResponseWriter, r *http.Request, urlParams []string) {

	var response interface{}
	var err error

	switch r.Method {
	case "GET":
		if len(urlParams) != 1 {
			http.Error(w, "Method not supported!", http.StatusMethodNotAllowed)
			return
		}
		validatedRequest := getValidRequestForGetCoinForWallet(urlParams)
		response, err = rs.walletService.GetCoinsForWallet(validatedRequest)
	case "POST":
		if len(urlParams) != 2 || urlParams[1] != "coins" {
			http.Error(w, "Method not supported!", http.StatusMethodNotAllowed)
			return
		}
		validatedRequest := getValidRequestForCreateCoinForWallet(r, urlParams)
		response, err = rs.walletService.CreateCoinForWallet(validatedRequest)
	case "PUT":
		if len(urlParams) != 2 {
			http.Error(w, "Method not supported!", http.StatusMethodNotAllowed)
			return
		}
		validatedRequest := getValidRequestForUpdateCoinForWallet(r, urlParams)
		response, err = rs.walletService.UpdateCoinForWallet(validatedRequest)
	case "DELETE":
		if len(urlParams) != 2 {
			http.Error(w, "Method not supported!", http.StatusMethodNotAllowed)
			return
		}
		validatedRequest := getValidRequestForDeleteCoinForWallet(urlParams)
		response, err = rs.walletService.DeleteCoinForWallet(validatedRequest)
	default:
		http.Error(w, "Method not supported!", http.StatusMethodNotAllowed)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Fprintf(w, string(jsonResponse))

}

func getValidRequestForCreateWallet(r *http.Request) *pkg.CreateWalletRequest {
	req := &pkg.CreateWalletRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		panic(err)
	}
	return req
}

func getValidRequestForUpdateWallet(r *http.Request, urlParams []string) *pkg.UpdateWalletRequest {
	req := &pkg.UpdateWalletRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		panic(err)
	}
	req.OldName = urlParams[1]
	return req
}

func getValidRequestForDeleteWallet(urlParams []string) *pkg.DeleteWalletRequest {
	return &pkg.DeleteWalletRequest{
		Name: urlParams[1],
	}
}

func getValidRequestForGetCoinForWallet(urlParams []string) *pkg.GetCoinsForWalletRequest {
	return &pkg.GetCoinsForWalletRequest{
		WalletName: urlParams[0],
	}
}

func getValidRequestForCreateCoinForWallet(r *http.Request, urlParams []string) *pkg.CreateCoinForWalletRequest {
	req := &pkg.CreateCoinForWalletRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		panic(err)
	}
	req.WalletName = urlParams[0]
	return req
}

func getValidRequestForUpdateCoinForWallet(r *http.Request, urlParams []string) *pkg.UpdateCoinForWalletRequest {
	req := &pkg.UpdateCoinForWalletRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		panic(err)
	}
	req.WalletName = urlParams[0]
	req.OldSymbol = urlParams[1]
	fmt.Println("aaaaaaaaaaaaaaaaa", len(urlParams), urlParams)
	return req
}

func getValidRequestForDeleteCoinForWallet(urlParams []string) *pkg.DeleteCoinForWalletRequest {
	return &pkg.DeleteCoinForWalletRequest{
		WalletName: urlParams[0],
		CoinSymbol: urlParams[1],
	}
}

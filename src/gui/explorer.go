package gui

import (
	"net/http"
	"strconv"

	"github.com/skycoin/skycoin/src/cipher"
	"github.com/skycoin/skycoin/src/daemon"
	wh "github.com/skycoin/skycoin/src/util/http" //http,json helpers
	"github.com/skycoin/skycoin/src/visor"
	"github.com/skycoin/skycoin/src/wallet"
)

func RegisterExploerHandlers(mux *http.ServeMux, gateway *daemon.Gateway) {
	// get set of pending transactions
	mux.HandleFunc("/explorer/address", getTransactionsForAddress(gateway))

	mux.HandleFunc("/explorer/getEffectiveOutputs", getCoinSupply(gateway))
}

func getCoinSupply(gateway *daemon.Gateway) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			var AddrList []string = []string{
				"R6aHqKWSQfvpdo2fGSrq4F1RYXkBWR9HHJ",
				"2EYM4WFHe4Dgz6kjAdUkM6Etep7ruz2ia6h",
				"25aGyzypSA3T9K6rgPUv1ouR13efNPtWP5m",
				"ix44h3cojvN6nqGcdpy62X7Rw6Ahnr3Thk",
				"AYV8KEBEAPCg8a59cHgqHMqYHP9nVgQDyW",
				"2Nu5Jv5Wp3RYGJU1EkjWFFHnebxMx1GjfkF",
				"2THDupTBEo7UqB6dsVizkYUvkKq82Qn4gjf",
				"tWZ11Nvor9parjg4FkwxNVcby59WVTw2iL",
				"m2joQiJRZnj3jN6NsoKNxaxzUTijkdRoSR",
				"8yf8PAQqU2cDj8Yzgz3LgBEyDqjvCh2xR7",
				"sgB3n11ZPUYHToju6TWMpUZTUcKvQnoFMJ",
				"2UYPbDBnHUEc67e7qD4eXtQQ6zfU2cyvAvk",
				"wybwGC9rhm8ZssBuzpy5goXrAdE31MPdsj",
				"JbM25o7kY7hqJZt3WGYu9pHZFCpA9TCR6t",
				"2efrft5Lnwjtk7F1p9d7BnPd72zko2hQWNi",
				"Syzmb3MiMoiNVpqFdQ38hWgffHg86D2J4e",
				"2g3GUmTQooLrNHaRDhKtLU8rWLz36Beow7F",
				"D3phtGr9iv6238b3zYXq6VgwrzwvfRzWZQ",
				"gpqsFSuMCZmsjPc6Rtgy1FmLx424tH86My",
				"2EUF3GPEUmfocnUc1w6YPtqXVCy3UZA4rAq",
				"TtAaxB3qGz5zEAhhiGkBY9VPV7cekhvRYS",
				"2fM5gVpi7XaiMPm4i29zddTNkmrKe6TzhVZ",
				"ix3NDKgxfYYANKAb5kbmwBYXPrkAsha7uG",
				"2RkPshpFFrkuaP98GprLtgHFTGvPY5e6wCK",
				"Ak1qCDNudRxZVvcW6YDAdD9jpYNNStAVqm",
				"2eZYSbzBKJ7QCL4kd5LSqV478rJQGb4UNkf",
				"KPfqM6S96WtRLMuSy4XLfVwymVqivdcDoM",
				"5B98bU1nsedGJBdRD5wLtq7Z8t8ZXio8u5",
				"2iZWk5tmBynWxj2PpAFyiZzEws9qSnG3a6n",
				"XUGdPaVnMh7jtzPe3zkrf9FKh5nztFnQU5",
				"hSNgHgewJme8uaHrEuKubHYtYSDckD6hpf",
				"2DeK765jLgnMweYrMp1NaYHfzxumfR1PaQN",
				"orrAssY5V2HuQAbW9K6WktFrGieq2m23pr",
				"4Ebf4PkG9QEnQTm4MVvaZvJV6Y9av3jhgb",
				"7Uf5xJ3GkiEKaLxC2WmJ1t6SeekJeBdJfu",
				"oz4ytDKbCqpgjW3LPc52pW2CaK2gxCcWmL",
				"2ex5Z7TufQ5Z8xv5mXe53fSQRfUr35SSo7Q",
				"WV2ap7ZubTxeDdmEZ1Xo7ufGMkekLWikJu",
				"ckCTV4r1pNuz6j2VBRHhaJN9HsCLY7muLV",
				"MXJx96ZJVSjktgeYZpVK8vn1H3xWP8ooq5",
				"wyQVmno9aBJZmQ99nDSLoYWwp7YDJCWsrH",
				"2cc9wKxCsFNRkoAQDAoHke3ZoyL1mSV14cj",
				"29k9g3F5AYfVaa1joE1PpZjBED6hQXes8Mm",
				"2XPLzz4ZLf1A9ykyTCjW5gEmVjnWa8CuatH",
				"iH7DqqojTgUn2JxmY9hgFp165Nk7wKfan9",
				"RJzzwUs3c9C8Y7NFYzNfFoqiUKeBhBfPki",
				"2W2cGyiCRM4nwmmiGPgMuGaPGeBzEm7VZPn",
				"ALJVNKYL7WGxFBSriiZuwZKWD4b7fbV1od",
				"tBaeg9zE2sgmw5ZQENaPPYd6jfwpVpGTzS",
				"2hdTw5Hk3rsgpZjvk8TyKcCZoRVXU5QVrUt",
				"A1QU6jKq8YgTP79M8fwZNHUZc7hConFKmy",
				"q9RkXoty3X1fuaypDDRUi78rWgJWYJMmpJ",
				"2Xvm6is5cAPA85xnSYXDuAqiRyoXiky5RaD",
				"4CW2CPJEzxhn2PS4JoSLoWGL5QQ7dL2eji",
				"24EG6uTzL7DHNzcwsygYGRR1nfu5kco7AZ1",
				"KghGnWw5fppTrqHSERXZf61yf7GkuQdCnV",
				"2WojewRA3LbpyXTP9ANy8CZqJMgmyNm3MDr",
				"2BsMfywmGV3M2CoDA112Rs7ZBkiMHfy9X11",
				"kK1Q4gPyYfVVMzQtAPRzL8qXMqJ67Y7tKs",
				"28J4mx8xfUtM92DbQ6i2Jmqw5J7dNivfroN",
				"gQvgyG1djgtftoCVrSZmsRxr7okD4LheKw",
				"3iFGBKapAWWzbiGFSr5ScbhrEPm6Esyvia",
				"NFW2akQH2vu7AqkQXxFz2P5vkXTWkSqrSm",
				"2MQJjLnWRp9eHh6MpCwpiUeshhtmri12mci",
				"2QjRQUMyL6iodtHP9zKmxCNYZ7k3jxtk49C",
				"USdfKy7B6oFNoauHWMmoCA7ND9rHqYw2Mf",
				"cA49et9WtptYHf6wA1F8qqVgH3kS5jJ9vK",
				"qaJT9TjcMi46sTKcgwRQU8o5Lw2Ea1gC4N",
				"22pyn5RyhqtTQu4obYjuWYRNNw4i54L8xVr",
				"22dkmukC6iH4FFLBmHne6modJZZQ3MC9BAT",
				"z6CJZfYLvmd41GRVE8HASjRcy5hqbpHZvE",
				"GEBWJ2KpRQDBTCCtvnaAJV2cYurgXS8pta",
				"oS8fbEm82cprmAeineBeDkaKd7QownDZQh",
				"rQpAs1LVQdphyj9ipEAuukAoj9kNpSP8cM",
				"6NSJKsPxmqipGAfFFhUKbkopjrvEESTX3j",
				"cuC68ycVXmD2EBzYFNYQ6akhKGrh3FGjSf",
				"bw4wtYU8toepomrhWP2p8UFYfHBbvEV425",
				"HvgNmDz5jD39Gwmi9VfDY1iYMhZUpZ8GKz",
				"SbApuZAYquWP3Q6iD51BcMBQjuApYEkRVf",
				"2Ugii5yxJgLzC59jV1vF8GK7UBZdvxwobeJ",
				"21N2iJ1qnQRiJWcEqNRxXwfNp8QcmiyhtPy",
				"9TC4RGs6AtFUsbcVWnSoCdoCpSfM66ALAc",
				"oQzn55UWG4iMcY9bTNb27aTnRdfiGHAwbD",
				"2GCdwsRpQhcf8SQcynFrMVDM26Bbj6sgv9M",
				"2NRFe7REtSmaM2qAgZeG45hC8EtVGV2QjeB",
				"25RGnhN7VojHUTvQBJA9nBT5y1qTQGULMzR",
				"26uCBDfF8E2PJU2Dzz2ysgKwv9m4BhodTz9",
				"Wkvima5cF7DDFdmJQqcdq8Syaq9DuAJJRD",
				"286hSoJYxvENFSHwG51ZbmKaochLJyq4ERQ",
				"FEGxF3HPoM2HCWHn82tyeh9o7vEQq5ySGE",
				"h38DxNxGhWGTq9p5tJnN5r4Fwnn85Krrb6",
				"2c1UU8J6Y3kL4cmQh21Tj8wkzidCiZxwdwd",
				"2bJ32KuGmjmwKyAtzWdLFpXNM6t83CCPLq5",
				"2fi8oLC9zfVVGnzzQtu3Y3rffS65Hiz6QHo",
				"TKD93RxFr2Am44TntLiJQus4qcEwTtvEEQ",
				"zMDywYdGEDtTSvWnCyc3qsYHWwj9ogws74",
				"25NbotTka7TwtbXUpSCQD8RMgHKspyDubXJ",
				"2ayCELBERubQWH5QxUr3cTxrYpidvUAzsSw",
				"RMTCwLiYDKEAiJu5ekHL1NQ8UKHi5ozCPg",
				"ejJjiCwp86ykmFr5iTJ8LxQXJ2wJPTYmkm",
			}

			filters := []daemon.OutputsFilter{}

			filters = append(filters, daemon.FbyAddressesNotIncluded(AddrList))
			outs := gateway.GetUnspentOutputs(filters...)
			totalSupply := 0
			for _, u := range outs.HeadOutputs {
				coin, err := strconv.Atoi(u.Coins)
				if err == nil {
					totalSupply = totalSupply + coin
				}

			}

			wh.SendOr404(w, wallet.CoinSupply{
				CurrentSupply: totalSupply,
				CoinCap:       100000000,
			})
		}
	}

}

func getTransactionsForAddress(gateway *daemon.Gateway) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			wh.Error405(w, "")
			return
		}
		addr := r.FormValue("address")
		if addr == "" {
			wh.Error400(w, "address is empty")
			return
		}

		cipherAddr, err := cipher.DecodeBase58Address(addr)
		if err != nil {
			wh.Error400(w, err.Error())
			return
		}

		uxs, err := gateway.GetAddressUxOuts(cipherAddr)
		if err != nil {
			wh.Error400(w, err.Error())
			return
		}

		resTxs := make([]visor.ReadableAddressTransaction, len(uxs))

		for i, ux := range uxs {
			sourceTxnNumber, err := cipher.SHA256FromHex(ux.Out.Body.SrcTransaction.Hex())
			if err != nil {
				wh.Error400(w, "Transaction id is not good")
				return
			}
			sourceTransaction, err := gateway.GetTransaction(sourceTxnNumber)
			in := make([]visor.ReadableTransactionInput, len(sourceTransaction.Txn.In))
			for i := range sourceTransaction.Txn.In {
				id, err := cipher.SHA256FromHex(sourceTransaction.Txn.In[i].Hex())
				if err != nil {
					wh.Error400(w, err.Error())
					return
				}
				uxout, err := gateway.GetUxOutByID(id)
				in[i] = visor.NewReadableTransactionInput(sourceTransaction.Txn.In[i].Hex(), uxout.Out.Body.Address.String())
			}

			resTxs[i] = visor.NewReadableAddressTransaction(sourceTransaction, in)
		}
		wh.SendOr404(w, &resTxs)
	}
}

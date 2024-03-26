package main

import (
	"fmt"
	"strings"
)

type Currency int

const (
	BITCOIN Currency = iota
	ETHEREUM
	TRON
	USDT
)

func (c Currency) String() string {
	switch c {
	case BITCOIN:
		return "BTC"
	case ETHEREUM:
		return "ETH"
	case TRON:
		return "TRX"
	case USDT:
		return "USDT"
	}
	return "UNKNOWN"
}

type Wallet struct {
	Amount   int64
	Currency Currency
	name     string
}

func (w Wallet) WalletMessage() string {
	balance, err := GetBalance(w.name)
	if err != nil {
		return "Error getting balance"
	}
	return strings.Replace(fmt.Sprintf("💰My Wallet \n\n*%s*: %s", w.Currency.String(), balance), ".", "\\.", -1)
}

func (w Wallet) Receive() string {
	address := GetAddress(w.name)
	return fmt.Sprintf("*Receive*\n\nUse the address below to send BTC to the CryptoOwl bot wallet address\\.\nNetwork: *Bitcoin \\- BTC*\\.\n\n*Address:* `%s`\n\n Funds will be credited within 30\\-60 minutes\\.", address)
}

func (w Wallet) Send(amount int64, address string) string {
	txid, err := Send(w.name, address, amount)
	if err != nil {
		return "Error: `" + err.Error() + "`"
	}
	return fmt.Sprintf("Transaction ID: %s", txid)
}

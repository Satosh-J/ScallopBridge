// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package subtest

import (
	"math/big"
	"testing"

	msg "github.com/ChainSafe/ChainBridge/message"
	utils "github.com/ChainSafe/ChainBridge/shared/substrate"
	"github.com/centrifuge/go-substrate-rpc-client/signature"
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

func CreateClient(t *testing.T, key *signature.KeyringPair, endpoint string) *utils.Client {
	client, err := utils.CreateClient(key, endpoint)
	if err != nil {
		t.Fatal(err)
	}
	return client
}

func AddRelayer(t *testing.T, client *utils.Client, relayer types.AccountID) {
	err := client.AddRelayer(relayer)
	if err != nil {
		t.Fatal(err)
	}
}

func WhitelistChain(t *testing.T, client *utils.Client, id msg.ChainId) {
	err := client.WhitelistChain(id)
	if err != nil {
		t.Fatal(err)
	}
}

func InitiateHashTransfer(t *testing.T, client *utils.Client, hash types.Hash, destId msg.ChainId) {
	err := client.InitiateHashTransfer(hash, destId)
	if err != nil {
		t.Fatal(err)
	}
}

func InitiateSubstrateNativeTransfer(t *testing.T, client *utils.Client, amount types.U32, recipient []byte, destId msg.ChainId) {
	err := client.InitiateSubstrateNativeTransfer(amount, recipient, destId)
	if err != nil {
		t.Fatal(err)
	}
}

func RegisterResource(t *testing.T, client *utils.Client, id msg.ResourceId, method string) {
	err := client.RegisterResource(id, method)
	if err != nil {
		t.Fatal(err)
	}
}

func BalanceOf(t *testing.T, client *utils.Client, publicKey []byte) *big.Int {
	balance, err := utils.BalanceOf(client, publicKey)
	if err != nil {
		t.Fatal(err)
	}
	return balance
}

func AssertBalanceOf(t *testing.T, client *utils.Client, publicKey []byte, expected *big.Int) {
	current, err := utils.BalanceOf(client, publicKey)
	if err != nil {
		t.Fatal(err)
	}

	if expected.Cmp(current) != 0 {
		t.Fatalf("Balance does not match expected. Expected: %s Got: %s (change %s) \n", expected.String(), current.String(), big.NewInt(0).Sub(current, expected).String())
	}
}
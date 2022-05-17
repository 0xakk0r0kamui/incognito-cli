package main

import (
	"fmt"
	"github.com/incognitochain/go-incognito-sdk-v2/incclient"
	"github.com/urfave/cli/v2"
	"log"
)

// getPortalDepositAddress generates the portal depositing (i.e, shielding) address for a payment address and a tokenID.
func getPortalDepositAddress(c *cli.Context) error {
	address := c.String(addressFlag)
	if !isValidAddress(address) {
		return fmt.Errorf("%v is invalid", addressFlag)
	}

	tokenIDStr := c.String(tokenIDFlag)
	if !isValidTokenID(tokenIDStr) {
		return fmt.Errorf("%v is invalid", tokenIDFlag)
	}

	shieldAddress, err := cfg.incClient.GeneratePortalShieldingAddress(address, tokenIDStr)
	if err != nil {
		return err
	}

	log.Printf("ShieldingAddress: %v\n", shieldAddress)
	return nil
}

// portalShield deposits a portal token (e.g, BTC) into the Incognito chain.
func portalShield(c *cli.Context) error {
	if cfg.btcClient == nil {
		return fmt.Errorf("portal shielding is not supported by this CLI configuration")
	}

	privateKey := c.String(privateKeyFlag)
	if !isValidPrivateKey(privateKey) {
		return fmt.Errorf("%v is invalid", privateKeyFlag)
	}

	address := c.String(addressFlag)
	if address == "" {
		address = incclient.PrivateKeyToPaymentAddress(privateKey, -1)
	}
	if !isValidAddress(address) {
		return fmt.Errorf("%v is invalid", addressFlag)
	}

	portalTxHashStr := c.String(externalTxIDFlag)

	tokenIDStr := c.String(tokenIDFlag)
	if !isValidTokenID(tokenIDStr) {
		return fmt.Errorf("%v is invalid", tokenIDFlag)
	}

	// check if the transaction has enough confirmations.
	isConfirmed, blkHeight, err := cfg.btcClient.IsConfirmedTx(portalTxHashStr)
	if err != nil {
		return err
	}
	if !isConfirmed {
		return fmt.Errorf("tx %v does not have enough 6 confirmations", portalTxHashStr)
	}

	// generate the shielding proof.
	shieldingProof, err := cfg.btcClient.BuildProof(portalTxHashStr, blkHeight)
	if err != nil {
		return err
	}

	// create an Incognito transaction to submit the proof.
	txHash, err := cfg.incClient.CreateAndSendPortalShieldTransaction(privateKey, tokenIDStr, address, shieldingProof, nil, nil)
	if err != nil {
		return err
	}
	log.Printf("TxHash: %v\n", txHash)
	return nil
}

// getPortalShieldStatus returns the status of a portal shielding request.
func getPortalShieldStatus(c *cli.Context) error {
	txHash := c.String(txHashFlag)
	if txHash == "" {
		return fmt.Errorf("%v is invalid", txHashFlag)
	}

	status, err := cfg.incClient.GetPortalShieldingRequestStatus(txHash)
	if err != nil {
		return err
	}

	return jsonPrint(status)
}

// portalUnShield creates and sends a port un-shielding transaction.
func portalUnShield(c *cli.Context) error {
	privateKey := c.String(privateKeyFlag)
	if !isValidPrivateKey(privateKey) {
		return fmt.Errorf("%v is invalid", privateKeyFlag)
	}

	tokenIDStr := c.String(tokenIDFlag)
	if !isValidTokenID(tokenIDStr) {
		return fmt.Errorf("%v is invalid", tokenIDFlag)
	}

	unShieldAmount := c.Uint64(amountFlag)
	if unShieldAmount == 0 {
		return fmt.Errorf("%v cannot be zero", amountFlag)
	}

	remoteAddress := c.String(externalAddressFlag)
	if remoteAddress == "" {
		return fmt.Errorf("%v is invalid", externalAddressFlag)
	}

	// create a transaction to burn the Incognito token.
	txHash, err := cfg.incClient.CreateAndSendPortalUnShieldTransaction(
		privateKey,
		tokenIDStr,
		remoteAddress,
		unShieldAmount,
		nil,
		nil,
	)
	if err != nil {
		return err
	}

	log.Printf("TxHash: %v\n", txHash)
	log.Println("Please wait for ~ 30-60 minutes for the fund to be released!!")
	log.Println("Use command `portalunshieldstatus` to check the status of the request.")

	return nil
}

// getPortalUnShieldStatus returns the status of a portal un-shielding request.
func getPortalUnShieldStatus(c *cli.Context) error {
	txHash := c.String(txHashFlag)
	if txHash == "" {
		return fmt.Errorf("%v is invalid", txHashFlag)
	}

	status, err := cfg.incClient.GetPortalUnShieldingRequestStatus(txHash)
	if err != nil {
		return err
	}

	return jsonPrint(status)
}

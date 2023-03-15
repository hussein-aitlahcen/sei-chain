package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/crypto/bn256"
)

// AddGenesisAccountCmd returns add-genesis-account cobra Command.
func GenBn256() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gen-bn256",
		Short: "Generate a BN256 key pair",
		Long: `Boring`,
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			privKey := bn256.GenPrivKey();
			pubKey := privKey.PubKey();
			privKeyRepr, err := json.Marshal(privKey);
			if err != nil {
				return err;
			}
			pubKeyRepr, err := json.Marshal(pubKey);
			if err != nil {
				return err;
			}
			fmt.Println("PubKey: ", string(privKeyRepr));
			fmt.Println("PubKey:", string(pubKeyRepr));
			return nil;
		},
	}
	return cmd
}

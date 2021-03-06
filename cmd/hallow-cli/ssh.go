package main

import (
	"fmt"
	"strings"

	"github.com/urfave/cli/v2"

	"github.com/hallowauth/hallow/client"
)

var (
	SSHCommand = &cli.Command{
		Name:   "ssh",
		Usage:  "Prints the command to SSH into a server. Generally used `eval $(hallow-cli ssh myserver.com)`",
		Action: SSH,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "key-id",
				Value: client.DefaultComment(),
				Usage: "KeyID to use for the SSH Certificate",
			},
		},
	}
)

func SSH(c *cli.Context) error {
	if c.NArg() != 1 {
		return fmt.Errorf("ssh takes exactly one argument")
	}

	hallow := hallowClientFromCLI(c)

	signer, sshCert, err := hallow.GenerateAndRequestCertificate(
		c.Context,
		client.KeyTypeED25519,
		c.String("key-id"),
	)
	if err != nil {
		return err
	}
	args, err := client.SSHCLI(signer, sshCert, c.Args().Get(0))
	if err != nil {
		return err
	}
	fmt.Println(strings.Join(args, " "))

	return nil
}

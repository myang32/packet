package cmd

import (
	"io/ioutil"

	"github.com/spf13/cobra"
)

// sshCmd represents the ssh command
var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "Manage your ssh keys for secure login to your packet devices.",
	// Long: ``,
}

var listSSHKeysCmd = &cobra.Command{
	Use:   "listall",
	Short: "View all configured SSH keys",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := ListSSHKeys()
		return err
	},
}

var listSSHKeyCmd = &cobra.Command{
	Use:   "list",
	Short: "View configured SSH key associated with the given ID.",
	RunE: func(cmd *cobra.Command, args []string) error {
		sshKeyID := cmd.Flag("key-id").Value.String()
		err := ListSSHKey(sshKeyID)
		return err
	},
}

var createSSHKeyCmd = &cobra.Command{
	Use:   "create",
	Short: "Configure a new SSH key",
	RunE: func(cmd *cobra.Command, args []string) error {
		keyFile := cmd.Flag("file").Value.String()
		key, err := ioutil.ReadFile(keyFile)
		if err != nil {
			return err
		}
		label := cmd.Flag("label").Value.String()
		err = CreateSSHKey(label, string(key))
		return err
	},
}

var deleteSSHKeyCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete SSH key associated with the given ID.",
	RunE: func(cmd *cobra.Command, args []string) error {
		sshKeyID := cmd.Flag("key-id").Value.String()
		err := DeleteSSHKey(sshKeyID)
		return err
	},
}

var updateSSHKeyCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a SSH key: change the key or its label",
	RunE: func(cmd *cobra.Command, args []string) error {
		sshKeyID := cmd.Flag("key-id").Value.String()
		keyFile := cmd.Flag("file").Value.String()
		key, err := ioutil.ReadFile(keyFile)
		if err != nil {
			return err
		}
		label := cmd.Flag("label").Value.String()
		err = UpdateSSHKey(sshKeyID, label, string(key))
		return err
	},
}

func init() {
	// Subcommands
	sshCmd.AddCommand(listSSHKeysCmd, listSSHKeyCmd, createSSHKeyCmd, deleteSSHKeyCmd, updateSSHKeyCmd)
	RootCmd.AddCommand(sshCmd)

	// Flags for command: packet ssh list
	listSSHKeyCmd.Flags().String("key-id", "", "SSH key ID")

	//Flags for command: packet ssh create
	createSSHKeyCmd.Flags().String("label", "", "Label to assign to the key")
	createSSHKeyCmd.Flags().StringP("file", "f", "", "Read public key from file.")

	// Flags for command: packet ssh delete
	deleteSSHKeyCmd.Flags().String("key-id", "", "SSH key ID")

	// Flags for command: packet ssh update
	updateSSHKeyCmd.Flags().String("key-id", "", "SSH key ID")
	updateSSHKeyCmd.Flags().String("label", "", "Label to assign to the key")
	updateSSHKeyCmd.Flags().StringP("file", "f", "", "Read public key from file.")
}

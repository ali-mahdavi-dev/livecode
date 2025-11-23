package command

import (
	"fmt"
	"livecode/scripts"

	"github.com/spf13/cobra"
)

func readUserFromJsonCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "readUserFromJson",
		Short: "handle database migration actions",
	}

	readUserFronJson := &cobra.Command{
		Use:   "",
		Short: "insert json user to db",
		RunE: func(_ *cobra.Command, args []string) error {
			return readUserFromJson()
		},
	}

	cmd.AddCommand(readUserFronJson)

	return cmd
}

func readUserFromJson() error {
	err := scripts.RunScriptReadUserFromJSON("./users_data.json")
	if err != nil {
		return fmt.Errorf("fail to run job: %w", err)
	}

	return nil
}

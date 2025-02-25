package secret

import (
	"github.com/MakeNowJust/heredoc"
	cmdDelete "github.com/cli/cli/v2/pkg/cmd/secret/delete"
	cmdList "github.com/cli/cli/v2/pkg/cmd/secret/list"
	cmdSet "github.com/cli/cli/v2/pkg/cmd/secret/set"
	"github.com/cli/cli/v2/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdSecret(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "secret <command>",
		Short: "Manage GitHub secrets",
		Long: heredoc.Doc(`
			Secrets can be set at the repository, or organization level for use in
			GitHub Actions or Dependabot. User secrets can be set for use in GitHub Codespaces.
			Environment secrets can be set for use in GitHub Actions.
			Run "gh help secret set" to learn how to get started.
`),
	}

	cmdutil.EnableRepoOverride(cmd, f)

	cmd.AddCommand(cmdList.NewCmdList(f, nil))
	cmd.AddCommand(cmdSet.NewCmdSet(f, nil))
	cmd.AddCommand(cmdDelete.NewCmdDelete(f, nil))

	return cmd
}

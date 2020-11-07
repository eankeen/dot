package cmd

import (
	"github.com/eankeen/dotty/actions"
	"github.com/eankeen/dotty/config"
	"github.com/eankeen/dotty/internal/util"
	"github.com/spf13/cobra"
)

var userUnapplyCmd = &cobra.Command{
	Use:   "unapply",
	Short: "Unapply a",
	Long:  "This unapplies all user dotfiles, unlinking them from the destination (user) directory",
	Run: func(cmd *cobra.Command, args []string) {
		dotfilesDir := cmd.Flag("dotfiles-dir").Value.String()
		dottyCfg := config.DottyCfg(dotfilesDir)

		srcDir := util.Src(dotfilesDir, dottyCfg, "user")
		destDir := util.Dest(dotfilesDir, dottyCfg, "user")

		actions.Unapply(dotfilesDir, srcDir, destDir)
	},
}

func init() {
	userCmd.AddCommand(userUnapplyCmd)
}
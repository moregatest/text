package facade

import (
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/text"
	"github.com/spiegel-im-spiegel/text/cli/gonkf/list"
	"github.com/spiegel-im-spiegel/text/cli/gonkf/norm"
	"github.com/spiegel-im-spiegel/text/normalize"
)

//newNormCmd returns cobra.Command instance for norm sub-command
func newNormCmd() *cobra.Command {
	normCmd := &cobra.Command{
		Use:   "norm [flags] [text file]",
		Short: "Unicode normalization",
		Long:  "Unicode normalization (UTF-8 text only)",
		RunE: func(cmd *cobra.Command, args []string) error {
			str, _ := cmd.Flags().GetString("form")
			form := normalize.FormofNormalize(str)
			if form == normalize.Unknown {
				return text.ErrNoImplement
			}
			outPath, _ := cmd.Flags().GetString("output")

			reader := cui.Reader()
			if len(args) > 0 {
				file, err := os.OpenFile(args[0], os.O_RDONLY, 0400) //args[0] is maybe file path
				if err != nil {
					return err
				}
				defer file.Close()
				reader = file
			}
			dst := norm.Run(reader, form)

			if len(outPath) > 0 {
				file, err := os.OpenFile(outPath, os.O_WRONLY|os.O_CREATE, 0666)
				if err != nil {
					return err
				}
				defer file.Close()
				io.Copy(file, dst)
			} else {
				cui.WriteFrom(dst)
			}
			return nil
		},
	}

	normCmd.Flags().StringP("form", "f", "nfc", "normalization form ["+list.NormOptionsList("|")+"]")
	normCmd.Flags().StringP("output", "o", "", "output file path")

	return normCmd
}

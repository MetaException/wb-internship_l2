package data

import "github.com/spf13/cobra"

type Flags struct {
	F int
	S bool
	D string
}

func NewFlags(F int, S bool, D string) *Flags {
	return &Flags{
		F: F,
		S: S,
		D: D,
	}
}

func NewFlagsParse(cmd *cobra.Command) (*Flags, error) {

	//rootCmd.Flags().IntP("fields", "f", 1, "выбрать поля (колонки)")
	//rootCmd.Flags().StringP("delimiter", "d", `(\t)`, "использовать другой разделитель")
	//Flags().BoolP("separated", "s", false, "только строки с разделителем")

	F, err := cmd.Flags().GetInt("fields")
	if err != nil {
		return nil, err
	}

	S, err := cmd.Flags().GetBool("separated")
	if err != nil {
		return nil, err
	}

	D, err := cmd.Flags().GetString("delimiter")
	if err != nil {
		return nil, err
	}

	return &Flags{
		F: F,
		S: S,
		D: D,
	}, nil
}

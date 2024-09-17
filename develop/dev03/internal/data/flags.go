package data

import "github.com/spf13/cobra"

//	rootCmd.Flags().StringP("column", "k", " ", "sort by key")
//	rootCmd.Flags().BoolP("numeric", "n", false, "sort by numeric")
//	rootCmd.Flags().BoolP("reverse", "r", false, "reverse order sorting")
//	rootCmd.Flags().BoolP("distinct", "u", false, "unique sort")
//  rootCmd.Flags().StringP("output", "o", "", "output path")

type Flags struct {
	OutputPath string
	N, R, U    bool
	K          int
}

func NewFlags(k int, outputPath string, n, r, u bool) *Flags {
	return &Flags{
		K:          k,
		R:          r,
		U:          u,
		N:          n,
		OutputPath: outputPath,
	}
}

func NewFlagsParse(cmd *cobra.Command) (*Flags, error) {

	k, err := cmd.Flags().GetInt("column")
	if err != nil {
		return nil, err
	}

	n, err := cmd.Flags().GetBool("numeric")
	if err != nil {
		return nil, err
	}

	r, err := cmd.Flags().GetBool("reverse")
	if err != nil {
		return nil, err
	}

	u, err := cmd.Flags().GetBool("distinct")
	if err != nil {
		return nil, err
	}

	outputPath, err := cmd.Flags().GetString("output")
	if err != nil {
		return nil, err
	}

	return &Flags{
		K:          k,
		N:          n,
		R:          r,
		U:          u,
		OutputPath: outputPath,
	}, nil
}

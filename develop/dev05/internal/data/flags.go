package data

import "github.com/spf13/cobra"

type Flags struct {
	//OutputPath string
	I, V, F, CC, N bool
	A, B, C        int
}

func NewFlags(A, B, C int, i, c, n, v, F bool) *Flags {
	return &Flags{
		A:  A,
		B:  B,
		C:  C,
		CC: c,
		N:  n,
		I:  i,
		V:  v,
		F:  F,
	}
}

func NewFlagsParse(cmd *cobra.Command) (*Flags, error) {

	//	rootCmd.Flags().BoolP("ignore-case", "i", false, "игнорировать регистр")
	//	rootCmd.Flags().BoolP("invert", "v", false, "инвертированный поиск")
	//	rootCmd.Flags().BoolP("fixed", "F", false, "точное совпадение со строкой, не паттерн")
	//	rootCmd.Flags().IntP("after", "A", -1, "печать +N строк после совпадения")
	//	rootCmd.Flags().IntP("before", "B", -1, "печать +N строк до совпадения")
	//	rootCmd.Flags().IntP("context", "C", -1, "печать +-строк вокруг совпадения")
	//	rootCmd.Flags().IntP("count", "c", -1, "количество строк")
	//	rootCmd.Flags().BoolP("line-num", "n", false, "напечать номер строки")

	i, err := cmd.Flags().GetBool("ignore-case")
	if err != nil {
		return nil, err
	}

	v, err := cmd.Flags().GetBool("invert")
	if err != nil {
		return nil, err
	}

	F, err := cmd.Flags().GetBool("fixed")
	if err != nil {
		return nil, err
	}

	n, err := cmd.Flags().GetBool("line-num")
	if err != nil {
		return nil, err
	}

	A, err := cmd.Flags().GetInt("after")
	if err != nil {
		return nil, err
	}

	B, err := cmd.Flags().GetInt("before")
	if err != nil {
		return nil, err
	}

	C, err := cmd.Flags().GetInt("context")
	if err != nil {
		return nil, err
	}

	c, err := cmd.Flags().GetBool("count")
	if err != nil {
		return nil, err
	}

	return &Flags{
		A:  A,
		B:  B,
		C:  C,
		CC: c,
		N:  n,
		I:  i,
		V:  v,
		F:  F,
	}, nil
}

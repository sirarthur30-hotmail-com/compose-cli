/*
	Copyright (c) 2020 Docker Inc.

	Permission is hereby granted, free of charge, to any person
	obtaining a copy of this software and associated documentation
	files (the "Software"), to deal in the Software without
	restriction, including without limitation the rights to use, copy,
	modify, merge, publish, distribute, sublicense, and/or sell copies
	of the Software, and to permit persons to whom the Software is
	furnished to do so, subject to the following conditions:

	The above copyright notice and this permission notice shall be
	included in all copies or substantial portions of the Software.

	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
	EXPRESS OR IMPLIED,
	INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
	FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
	IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
	HOLDERS BE LIABLE FOR ANY CLAIM,
	DAMAGES OR OTHER LIABILITY,
	WHETHER IN AN ACTION OF CONTRACT,
	TORT OR OTHERWISE,
	ARISING FROM, OUT OF OR IN CONNECTION WITH
	THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package context

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	cliopts "github.com/docker/api/cli/options"
	apicontext "github.com/docker/api/context"
	"github.com/docker/api/context/store"
)

func showCommand(opts *cliopts.GlobalOpts) *cobra.Command {
	return &cobra.Command{
		Use:   "show",
		Short: "Print the current context",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runShow(cmd.Context(), opts)
		},
	}
}

func runShow(ctx context.Context, opts *cliopts.GlobalOpts) error {
	s := store.ContextStore(ctx)
	name := apicontext.CurrentContext(ctx)
	// Match behavior of existing CLI
	if name != store.DefaultContextName {
		if _, err := s.Get(name, nil); err != nil {
			return err
		}
	}
	fmt.Println(name)
	return nil
}
/*
Copyright 2023 The Vitess Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreedto in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package command

import (
	"github.com/spf13/cobra"
)

var AddAuth = &cobra.Command{
	Use:  "addAuth <digest> <user:pass>",
	Args: cobra.ExactArgs(2),
	RunE: commandAddAuth,
}

func commandAddAuth(cmd *cobra.Command, args []string) error {
	scheme, auth := cmd.Flags().Arg(0), cmd.Flags().Arg(1)
	return fs.Conn.AddAuth(cmd.Context(), scheme, []byte(auth))
}

func init() {
	Root.AddCommand(AddAuth)
}

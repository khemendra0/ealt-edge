/*
Copyright 2020 Huawei Technologies Co., Ltd.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package applcmpkg

import (
	"ealt/cmd/adapter"

	"github.com/spf13/cobra"
)

// allCmd represents the all command
func NewApplcmDeleteCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "delete",
		Short: "Install Complete EALT Deployment Environment",
		Long:  `Install Complete EALT Deployment Environment`,
		RunE: func(cmd *cobra.Command, args []string) error {
			applicationIdArg := cmd.Flag("appid")
			applicationId := applicationIdArg.Value.String()
			err := adapter.BuilderRequest(applicationId, "NewApplcmDeleteCommand")
			if err != nil {
				return err
			}
			return nil
		},
	}

	cmd.Flags().StringP("appid", "i", "", "Application Package ID to be deleted from MEP")

	return cmd
}
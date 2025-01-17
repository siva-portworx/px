// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	api "github.com/libopenstorage/openstorage-sdk-clients/sdk/golang"
	"github.com/portworx/px/pkg/portworx"
	"github.com/portworx/px/pkg/util"
	"github.com/spf13/cobra"
)

var (
	dvReq = &api.SdkVolumeDeleteRequest{}
)

// deleteVolumeCmd represents the deleteVolume command
var deleteVolumeCmd = &cobra.Command{
	Use:   "volume",
	Short: "Delete a volume in Portworx",
	RunE: func(cmd *cobra.Command, args []string) error {
		return deleteVolumeExec(cmd, args)
	},
}

func init() {
	deleteCmd.AddCommand(deleteVolumeCmd)
	deleteVolumeCmd.Flags().StringVar(&dvReq.VolumeId, "name", "", "Name/Id of volume (required)")
	deleteVolumeCmd.Flags().SortFlags = false

	// TODO bring the flags from rootCmd
}

func deleteVolumeExec(cmd *cobra.Command, args []string) error {
	ctx, conn, err := portworx.PxConnectCurrent(GetConfigFile())
	if err != nil {
		return err
	}
	defer conn.Close()

	// Send request
	volumes := api.NewOpenStorageVolumeClient(conn)
	_, err = volumes.Delete(ctx, dvReq)
	if err != nil {
		return util.PxErrorMessage(err, "Failed to delete volume")
	}

	util.Printf("Volume %s deleted\n", dvReq.GetVolumeId())

	return nil
}

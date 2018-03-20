package services

import (
	"gp_upgrade/hub/upgradestatus"
	pb "gp_upgrade/idl"

	"github.com/greenplum-db/gp-common-go-libs/gplog"
	"golang.org/x/net/context"
)

func (h *HubClient) UpgradeValidateStartCluster(ctx context.Context,
	in *pb.UpgradeValidateStartClusterRequest) (*pb.UpgradeValidateStartClusterReply, error) {
	gplog.Info("Started processing validate-start-cluster request")

	err := h.startNewCluster(in.NewBinDir, in.NewDataDir)
	return &pb.UpgradeValidateStartClusterReply{}, err
}

func (h *HubClient) startNewCluster(newBinDir string, newDataDir string) error {
	c := upgradestatus.NewChecklistManager(h.conf.StateDir)
	c.ResetStateDir("validate-start-cluster")
	c.MarkInProgress("validate-start-cluster")

	_, err := h.commandExecer("bash", "gpstart -a").Output()
	if err != nil {
		gplog.Error(err.Error())
		c.MarkFailed("validate-start-cluster")

		return err
	}

	c.MarkComplete("validate-start-cluster")

	return nil
}

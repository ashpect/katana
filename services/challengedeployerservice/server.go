package challengedeployerservice

import (
	"fmt"
	"log"

	g "github.com/sdslabs/katana/configs"
	"github.com/sdslabs/katana/lib/utils"
	v1 "k8s.io/api/core/v1"
)

func DeployToAll(localFilePath string, pathInPod string) error {

	if err := GetClient(g.KatanaConfig.KubeConfig); err != nil {
		return err
	}

	// Get pods from different namespaces
	var pods []v1.Pod
	numberOfTeams := utils.GetTeamNumber()
	for i := 0; i < numberOfTeams; i++ {
		podsInTeam, err := getPods(map[string]string{
			"app": g.ClusterConfig.TeamLabel,
		}, "katana-team-"+fmt.Sprint(i)+"-ns")
		if err != nil {
			log.Println(err)
			return err
		}
		pods = append(pods, podsInTeam...)
	}
	// Loop over pods
	for _, pod := range pods {
		// Copy file into pod
		if err := utils.CopyIntoPod(pod.Name, g.TeamVmConfig.ContainerName, pathInPod, localFilePath, pod.Namespace); err != nil {
			log.Println(err)
			return err
		}
	}

	return nil
}

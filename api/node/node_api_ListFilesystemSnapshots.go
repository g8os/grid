package node

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/zero-os/0-orchestrator/api/tools"
)

// ListFilesystemSnapshots is the handler for GET /nodes/{nodeid}/storagepools/{storagepoolname}/filesystem/{filesystemname}/snapshot
// List snapshots of this filesystem
func (api *NodeAPI) ListFilesystemSnapshots(w http.ResponseWriter, r *http.Request) {
	aysClient, err := tools.GetAysConnection(api)
	if err != nil {
		tools.WriteError(w, http.StatusUnauthorized, err, "")
		return
	}
	fileSystemName := mux.Vars(r)["filesystemname"]

	// only show the snapshots under the filesystem specified in the URL
	queryParams := map[string]interface{}{
		"parent": fmt.Sprintf("filesystem!%s", fileSystemName),
	}

	services, res, err := aysClient.Ays.ListServicesByRole("fssnapshot", api.AysRepo, nil, queryParams)
	if !tools.HandleAYSResponse(err, res, w, "listing snapshots") {
		return
	}

	respBody := make([]string, len(services))
	for i, serv := range services {
		respBody[i] = serv.Name
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&respBody)
}

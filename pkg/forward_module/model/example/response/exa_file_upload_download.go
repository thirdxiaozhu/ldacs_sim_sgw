package response

import "ldacs_sim_sgw/pkg/forward_module/model/example"

type ExaFileResponse struct {
	File example.ExaFileUploadAndDownload `json:"file"`
}

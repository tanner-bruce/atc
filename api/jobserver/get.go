package jobserver

import (
	"encoding/json"
	"net/http"

	"github.com/concourse/atc/api/present"
)

func (s *Server) GetJob(w http.ResponseWriter, r *http.Request) {
	jobName := r.FormValue(":job_name")

	finished, next, err := s.db.GetJobFinishedAndNextBuild(jobName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(present.Job(jobName, finished, next))
}

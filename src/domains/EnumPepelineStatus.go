package domains

type PipelineStatus string

const (
	SIGNUP_ANALISIS   PipelineStatus = "SIGNUP_ANALISIS"
	SIGNUP_COMPLETED  PipelineStatus = "SIGNUP_COMPLETED"
	SIGNUP_DENIED     PipelineStatus = "SIGNUP_DENIED"
	PROPOSAL_ANALISIS PipelineStatus = "PROPOSAL_ANALISIS"
	PROPOSAL_CREATED  PipelineStatus = "PROPOSAL_CREATED"
	PROPOSAL_DENIED   PipelineStatus = "PROPOSAL_DENIED"
)

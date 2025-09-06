package enum

type ContractStatus string

const (
	ContractStatusPending   ContractStatus = "PENDING"
	ContractStatusActive    ContractStatus = "ACTIVE"
	ContractStatusCompleted ContractStatus = "COMPLETED"
	ContractStatusCancelled ContractStatus = "CANCELLED"
)

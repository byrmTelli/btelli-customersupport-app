package models

type ComplaintStatus string

const (
	Pending    ComplaintStatus = "Pending"
	Resolved   ComplaintStatus = "Resolved"
	Cancelled  ComplaintStatus = "Cancelled"
	InProgress ComplaintStatus = "InProgress"
)

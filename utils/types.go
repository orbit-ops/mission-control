package utils

type AuditAction string

const (
	Approve      AuditAction = "approve"
	Reject       AuditAction = "reject"
	CreateAccess AuditAction = "createAccess"
	RemoveAccess AuditAction = "removeAccess"
)

package models

import (
	"database/sql"
	"time"
)

type FMRStatus string

const (
	StatusDraft      FMRStatus = "Draft"
	StatusSubmitted  FMRStatus = "Submitted"
	StatusUnresolved FMRStatus = "Unresolved"
	StatusInProgress FMRStatus = "In-Progress"
	StatusResolved   FMRStatus = "Resolved"
	StatusArchived   FMRStatus = "Archived"
)

type PointOfFailure string

const (
	FailureReceipt PointOfFailure = "Receipt"
	FailureTesting PointOfFailure = "Testing"
	FailureMission PointOfFailure = "Mission"
	FailureOther   PointOfFailure = "Other"
)

type RepairAction string

const (
	ActionRMA         RepairAction = "RMA"
	ActionRework      RepairAction = "Rework"
	ActionScrapNewPart RepairAction = "Scrap/New Part"
	ActionOther       RepairAction = "Other"
)

type Signed1149Status string

const (
	Signed1149NA              Signed1149Status = "N/A"
	Signed1149WaitingSignature Signed1149Status = "Waiting for Signature"
	Signed1149SignDate        Signed1149Status = "Sign Date"
)

type TierLevel string

const (
	TierLevelTier1 TierLevel = "tier1"
	TierLevelTier2 TierLevel = "tier2"
	TierLevelTier3 TierLevel = "tier3"
	TierLevelNone  TierLevel = "none"
)

type FMRRequest struct {
	ID             int64          `json:"id"`
	Title          string         `json:"title"`
	Description    *string        `json:"description,omitempty"`
	Status         FMRStatus      `json:"status"`
	CreationDate   time.Time      `json:"creationDate"`
	ResolutionDate *time.Time     `json:"resolutionDate,omitempty"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`

	// Report section
	ControlNumber           *string         `json:"controlNumber,omitempty"`
	Program                 *string         `json:"program,omitempty"`
	Unit                    *string         `json:"unit,omitempty"`
	OmmodSn                 *string         `json:"ommodSn,omitempty"`
	OmmodSystemMeterHrs     *string         `json:"ommodSystemMeterHrs,omitempty"`
	TechnicianName          *string         `json:"technicianName,omitempty"`
	Email                   *string         `json:"email,omitempty"`
	CommNumber              *string         `json:"commNumber,omitempty"`
	DateReported            *time.Time      `json:"dateReported,omitempty"`
	FailedDate              *time.Time      `json:"failedDate,omitempty"`
	PointOfFailure          *PointOfFailure `json:"pointOfFailure,omitempty"`
	FailedPartNomenclature  *string         `json:"failedPartNomenclature,omitempty"`
	FailedPartNumber        *string         `json:"failedPartNumber,omitempty"`
	Qty                     *string         `json:"qty,omitempty"`
	PartSerialNumber        *string         `json:"partSerialNumber,omitempty"`
	DescribeFailure         *string         `json:"describeFailure,omitempty"`
	TroubleShootingPerformed *string        `json:"troubleShootingPerformed,omitempty"`
	RepairTechnician        *string         `json:"repairTechnician,omitempty"`
	RepairStartDate         *time.Time      `json:"repairStartDate,omitempty"`
	RepairActivityDescription *string       `json:"repairActivityDescription,omitempty"`

	// Processing section
	RepairAction      *RepairAction `json:"repairAction,omitempty"`
	RepairActionOther *string       `json:"repairActionOther,omitempty"`

	// Supply Chain section
	Signed1149             *Signed1149Status `json:"signed1149,omitempty"`
	Signed1149Date         *time.Time        `json:"signed1149Date,omitempty"`
	FailedPartReceivedDate *time.Time        `json:"failedPartReceivedDate,omitempty"`
	NewPartReceivedDate    *time.Time        `json:"newPartReceivedDate,omitempty"`
	NewPartIssuedDate      *time.Time        `json:"newPartIssuedDate,omitempty"`
	NewPartSerialNumber    *string           `json:"newPartSerialNumber,omitempty"`
	SupplyComments         *string           `json:"supplyComments,omitempty"`

	// Procurement section
	PoNumber              *string    `json:"poNumber,omitempty"`
	WarrantyStartDate     *time.Time `json:"warrantyStartDate,omitempty"`
	NewPartOrderedDate    *time.Time `json:"newPartOrderedDate,omitempty"`
	VendorContactAddress  *string    `json:"vendorContactAddress,omitempty"`
	ProcurementComments   *string    `json:"procurementComments,omitempty"`

	// Logistics section
	RequestDate          *time.Time `json:"requestDate,omitempty"`
	RequestToVendorDate  *time.Time `json:"requestToVendorDate,omitempty"`
	RmaReceivedDate      *time.Time `json:"rmaReceivedDate,omitempty"`
	ShippedToVendorDate  *time.Time `json:"shippedToVendorDate,omitempty"`
	LogisticsComment     *string    `json:"logisticsComment,omitempty"`

	// Contractor section
	ProblemAssistanceRequested *string `json:"problemAssistanceRequested,omitempty"`
	Tier1Date                  *time.Time `json:"tier1Date,omitempty"`
	Tier2Date                  *time.Time `json:"tier2Date,omitempty"`
	Tier3Date                  *time.Time `json:"tier3Date,omitempty"`
	ProblemResolved            *bool      `json:"problemResolved,omitempty"`
	ContractorComments         *string    `json:"contractorComments,omitempty"`

	// Additional tracking fields
	FormSubmissionType          *string    `json:"formSubmissionType,omitempty"`
	FmrCreatedDate              *time.Time `json:"fmrCreatedDate,omitempty"`
	WarrantyStopDate            *time.Time `json:"warrantyStopDate,omitempty"`
	CompletedDate               *time.Time `json:"completedDate,omitempty"`
	FmrReceivedDateByFailures   *time.Time `json:"fmrReceivedDateByFailures,omitempty"`
	TrackingNoType              *string    `json:"trackingNoType,omitempty"`
	TrackingNo                  *string    `json:"trackingNo,omitempty"`
	RmaNoType                   *string    `json:"rmaNoType,omitempty"`
	RmaNo                       *string    `json:"rmaNo,omitempty"`
	EngineerDeployedType        *string    `json:"engineerDeployedType,omitempty"`
	EngineerName                *string    `json:"engineerName,omitempty"`
	DeploymentDate              *time.Time `json:"deploymentDate,omitempty"`
	RecordUpdated               *time.Time `json:"recordUpdated,omitempty"`
	RecordEnteredOn             *time.Time `json:"recordEnteredOn,omitempty"`
	RecordEnteredBy             *string    `json:"recordEnteredBy,omitempty"`
}

type CreateFMRRequest struct {
	Title       string     `json:"title" binding:"required"`
	Description *string    `json:"description,omitempty"`
	Status      *FMRStatus `json:"status,omitempty"`

	// All optional fields from FMRRequest
	ControlNumber                *string           `json:"controlNumber,omitempty"`
	Program                      *string           `json:"program,omitempty"`
	Unit                         *string           `json:"unit,omitempty"`
	OmmodSn                      *string           `json:"ommodSn,omitempty"`
	OmmodSystemMeterHrs          *string           `json:"ommodSystemMeterHrs,omitempty"`
	TechnicianName               *string           `json:"technicianName,omitempty"`
	Email                        *string           `json:"email,omitempty"`
	CommNumber                   *string           `json:"commNumber,omitempty"`
	DateReported                 *time.Time        `json:"dateReported,omitempty"`
	FailedDate                   *time.Time        `json:"failedDate,omitempty"`
	PointOfFailure               *PointOfFailure   `json:"pointOfFailure,omitempty"`
	FailedPartNomenclature       *string           `json:"failedPartNomenclature,omitempty"`
	FailedPartNumber             *string           `json:"failedPartNumber,omitempty"`
	Qty                          *string           `json:"qty,omitempty"`
	PartSerialNumber             *string           `json:"partSerialNumber,omitempty"`
	DescribeFailure              *string           `json:"describeFailure,omitempty"`
	TroubleShootingPerformed     *string           `json:"troubleShootingPerformed,omitempty"`
	RepairTechnician             *string           `json:"repairTechnician,omitempty"`
	RepairStartDate              *time.Time        `json:"repairStartDate,omitempty"`
	RepairActivityDescription    *string           `json:"repairActivityDescription,omitempty"`
	RepairAction                 *RepairAction     `json:"repairAction,omitempty"`
	RepairActionOther            *string           `json:"repairActionOther,omitempty"`
	Signed1149                   *Signed1149Status `json:"signed1149,omitempty"`
	Signed1149Date               *time.Time        `json:"signed1149Date,omitempty"`
	FailedPartReceivedDate       *time.Time        `json:"failedPartReceivedDate,omitempty"`
	NewPartReceivedDate          *time.Time        `json:"newPartReceivedDate,omitempty"`
	NewPartIssuedDate            *time.Time        `json:"newPartIssuedDate,omitempty"`
	NewPartSerialNumber          *string           `json:"newPartSerialNumber,omitempty"`
	SupplyComments               *string           `json:"supplyComments,omitempty"`
	PoNumber                     *string           `json:"poNumber,omitempty"`
	WarrantyStartDate            *time.Time        `json:"warrantyStartDate,omitempty"`
	NewPartOrderedDate           *time.Time        `json:"newPartOrderedDate,omitempty"`
	VendorContactAddress         *string           `json:"vendorContactAddress,omitempty"`
	ProcurementComments          *string           `json:"procurementComments,omitempty"`
	RequestDate                  *time.Time        `json:"requestDate,omitempty"`
	RequestToVendorDate          *time.Time        `json:"requestToVendorDate,omitempty"`
	RmaReceivedDate              *time.Time        `json:"rmaReceivedDate,omitempty"`
	ShippedToVendorDate          *time.Time        `json:"shippedToVendorDate,omitempty"`
	LogisticsComment             *string           `json:"logisticsComment,omitempty"`
	ProblemAssistanceRequested   *string           `json:"problemAssistanceRequested,omitempty"`
	Tier1Date                    *time.Time        `json:"tier1Date,omitempty"`
	Tier2Date                    *time.Time        `json:"tier2Date,omitempty"`
	Tier3Date                    *time.Time        `json:"tier3Date,omitempty"`
	ProblemResolved              *bool             `json:"problemResolved,omitempty"`
	ContractorComments           *string           `json:"contractorComments,omitempty"`
	FormSubmissionType           *string           `json:"formSubmissionType,omitempty"`
	FmrCreatedDate               *time.Time        `json:"fmrCreatedDate,omitempty"`
	WarrantyStopDate             *time.Time        `json:"warrantyStopDate,omitempty"`
	CompletedDate                *time.Time        `json:"completedDate,omitempty"`
	FmrReceivedDateByFailures    *time.Time        `json:"fmrReceivedDateByFailures,omitempty"`
	TrackingNoType               *string           `json:"trackingNoType,omitempty"`
	TrackingNo                   *string           `json:"trackingNo,omitempty"`
	RmaNoType                    *string           `json:"rmaNoType,omitempty"`
	RmaNo                        *string           `json:"rmaNo,omitempty"`
	EngineerDeployedType         *string           `json:"engineerDeployedType,omitempty"`
	EngineerName                 *string           `json:"engineerName,omitempty"`
	DeploymentDate               *time.Time        `json:"deploymentDate,omitempty"`
	RecordEnteredBy              *string           `json:"recordEnteredBy,omitempty"`
}

type UpdateFMRRequest struct {
	ID              int64      `json:"id" binding:"required"`
	Title           *string    `json:"title,omitempty"`
	Description     *string    `json:"description,omitempty"`
	Status          *FMRStatus `json:"status,omitempty"`
	ResolutionDate  *time.Time `json:"resolutionDate,omitempty"`

	// All the same optional fields as CreateFMRRequest
	ControlNumber                *string           `json:"controlNumber,omitempty"`
	Program                      *string           `json:"program,omitempty"`
	Unit                         *string           `json:"unit,omitempty"`
	OmmodSn                      *string           `json:"ommodSn,omitempty"`
	OmmodSystemMeterHrs          *string           `json:"ommodSystemMeterHrs,omitempty"`
	TechnicianName               *string           `json:"technicianName,omitempty"`
	Email                        *string           `json:"email,omitempty"`
	CommNumber                   *string           `json:"commNumber,omitempty"`
	DateReported                 *time.Time        `json:"dateReported,omitempty"`
	FailedDate                   *time.Time        `json:"failedDate,omitempty"`
	PointOfFailure               *PointOfFailure   `json:"pointOfFailure,omitempty"`
	FailedPartNomenclature       *string           `json:"failedPartNomenclature,omitempty"`
	FailedPartNumber             *string           `json:"failedPartNumber,omitempty"`
	Qty                          *string           `json:"qty,omitempty"`
	PartSerialNumber             *string           `json:"partSerialNumber,omitempty"`
	DescribeFailure              *string           `json:"describeFailure,omitempty"`
	TroubleShootingPerformed     *string           `json:"troubleShootingPerformed,omitempty"`
	RepairTechnician             *string           `json:"repairTechnician,omitempty"`
	RepairStartDate              *time.Time        `json:"repairStartDate,omitempty"`
	RepairActivityDescription    *string           `json:"repairActivityDescription,omitempty"`
	RepairAction                 *RepairAction     `json:"repairAction,omitempty"`
	RepairActionOther            *string           `json:"repairActionOther,omitempty"`
	Signed1149                   *Signed1149Status `json:"signed1149,omitempty"`
	Signed1149Date               *time.Time        `json:"signed1149Date,omitempty"`
	FailedPartReceivedDate       *time.Time        `json:"failedPartReceivedDate,omitempty"`
	NewPartReceivedDate          *time.Time        `json:"newPartReceivedDate,omitempty"`
	NewPartIssuedDate            *time.Time        `json:"newPartIssuedDate,omitempty"`
	NewPartSerialNumber          *string           `json:"newPartSerialNumber,omitempty"`
	SupplyComments               *string           `json:"supplyComments,omitempty"`
	PoNumber                     *string           `json:"poNumber,omitempty"`
	WarrantyStartDate            *time.Time        `json:"warrantyStartDate,omitempty"`
	NewPartOrderedDate           *time.Time        `json:"newPartOrderedDate,omitempty"`
	VendorContactAddress         *string           `json:"vendorContactAddress,omitempty"`
	ProcurementComments          *string           `json:"procurementComments,omitempty"`
	RequestDate                  *time.Time        `json:"requestDate,omitempty"`
	RequestToVendorDate          *time.Time        `json:"requestToVendorDate,omitempty"`
	RmaReceivedDate              *time.Time        `json:"rmaReceivedDate,omitempty"`
	ShippedToVendorDate          *time.Time        `json:"shippedToVendorDate,omitempty"`
	LogisticsComment             *string           `json:"logisticsComment,omitempty"`
	ProblemAssistanceRequested   *string           `json:"problemAssistanceRequested,omitempty"`
	Tier1Date                    *time.Time        `json:"tier1Date,omitempty"`
	Tier2Date                    *time.Time        `json:"tier2Date,omitempty"`
	Tier3Date                    *time.Time        `json:"tier3Date,omitempty"`
	ProblemResolved              *bool             `json:"problemResolved,omitempty"`
	ContractorComments           *string           `json:"contractorComments,omitempty"`
	FormSubmissionType           *string           `json:"formSubmissionType,omitempty"`
	FmrCreatedDate               *time.Time        `json:"fmrCreatedDate,omitempty"`
	WarrantyStopDate             *time.Time        `json:"warrantyStopDate,omitempty"`
	CompletedDate                *time.Time        `json:"completedDate,omitempty"`
	FmrReceivedDateByFailures    *time.Time        `json:"fmrReceivedDateByFailures,omitempty"`
	TrackingNoType               *string           `json:"trackingNoType,omitempty"`
	TrackingNo                   *string           `json:"trackingNo,omitempty"`
	RmaNoType                    *string           `json:"rmaNoType,omitempty"`
	RmaNo                        *string           `json:"rmaNo,omitempty"`
	EngineerDeployedType         *string           `json:"engineerDeployedType,omitempty"`
	EngineerName                 *string           `json:"engineerName,omitempty"`
	DeploymentDate               *time.Time        `json:"deploymentDate,omitempty"`
	RecordEnteredBy              *string           `json:"recordEnteredBy,omitempty"`
}

type FMRListResponse struct {
	FMRs  []FMRRequest `json:"fmrs"`
	Total int          `json:"total"`
}

type FMRFilters struct {
	Status    *FMRStatus `json:"status,omitempty"`
	Search    *string    `json:"search,omitempty"`
	SortBy    *string    `json:"sortBy,omitempty"`
	SortOrder *string    `json:"sortOrder,omitempty"` // "asc" or "desc"
	Limit     int        `json:"limit,omitempty"`
	Offset    int        `json:"offset,omitempty"`
	TierLevel *TierLevel `json:"tierLevel,omitempty"`
}

// Database row type for scanning from SQL
type FMRRow struct {
	ID                        int64
	Title                     string
	Description               sql.NullString
	Status                    string
	CreationDate              time.Time
	ResolutionDate            sql.NullTime
	CreatedAt                 time.Time
	UpdatedAt                 time.Time
	ControlNumber             sql.NullString
	Program                   sql.NullString
	Unit                      sql.NullString
	OmmodSn                   sql.NullString
	OmmodSystemMeterHrs       sql.NullString
	TechnicianName            sql.NullString
	Email                     sql.NullString
	CommNumber                sql.NullString
	DateReported              sql.NullTime
	FailedDate                sql.NullTime
	PointOfFailure            sql.NullString
	FailedPartNomenclature    sql.NullString
	FailedPartNumber          sql.NullString
	Qty                       sql.NullString
	PartSerialNumber          sql.NullString
	DescribeFailure           sql.NullString
	TroubleShootingPerformed  sql.NullString
	RepairTechnician          sql.NullString
	RepairStartDate           sql.NullTime
	RepairActivityDescription sql.NullString
	RepairAction              sql.NullString
	RepairActionOther         sql.NullString
	Signed1149                sql.NullString
	Signed1149Date            sql.NullTime
	FailedPartReceivedDate    sql.NullTime
	NewPartReceivedDate       sql.NullTime
	NewPartIssuedDate         sql.NullTime
	NewPartSerialNumber       sql.NullString
	SupplyComments            sql.NullString
	PoNumber                  sql.NullString
	WarrantyStartDate         sql.NullTime
	NewPartOrderedDate        sql.NullTime
	VendorContactAddress      sql.NullString
	ProcurementComments       sql.NullString
	RequestDate               sql.NullTime
	RequestToVendorDate       sql.NullTime
	RmaReceivedDate           sql.NullTime
	ShippedToVendorDate       sql.NullTime
	LogisticsComment          sql.NullString
	ProblemAssistanceRequested sql.NullString
	Tier1Date                 sql.NullTime
	Tier2Date                 sql.NullTime
	Tier3Date                 sql.NullTime
	ProblemResolved           sql.NullBool
	ContractorComments        sql.NullString
	FormSubmissionType        sql.NullString
	FmrCreatedDate            sql.NullTime
	WarrantyStopDate          sql.NullTime
	CompletedDate             sql.NullTime
	FmrReceivedDateByFailures sql.NullTime
	TrackingNoType            sql.NullString
	TrackingNo                sql.NullString
	RmaNoType                 sql.NullString
	RmaNo                     sql.NullString
	EngineerDeployedType      sql.NullString
	EngineerName              sql.NullString
	DeploymentDate            sql.NullTime
	RecordUpdated             sql.NullTime
	RecordEnteredOn           sql.NullTime
	RecordEnteredBy           sql.NullString
}

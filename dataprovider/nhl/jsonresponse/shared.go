package jsonresponse

// LocalizedString is a localized text value e.g. {"default": "Penguins"}
type LocalizedString struct {
	Default string `json:"default"`
}

// PeriodDescriptor describes a game period
type PeriodDescriptor struct {
	Number int32 `json:"number"`
	// PeriodType e.g. REG, OT, SO
	PeriodType           string `json:"periodType"`
	MaxRegulationPeriods int32  `json:"maxRegulationPeriods"`
}

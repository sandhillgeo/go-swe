package swe

type ResultTemplate struct {
  Type string `json:"type"`
  Name string `json:"name,omitempty"`
  Definition string `json:"definition"`
  Description string `json:"description"`
  ReferenceFrame string `json:"referenceFrame,omitempty"`
  Label string `json:"label,omitempty"`
  UnitsOfMeasurement map[string]string `json:"uom,omitempty"`
  ElementCount *ElementCount `json:"elementCount,omitempty"`
  ElementType *ResultTemplate `json:"elementType,omitempty"`
  Fields []ResultTemplate `json:"field,omitempty"`
}

type ElementCount struct {
  Type string `json:"type"`
  Value int `json:"value"`
}

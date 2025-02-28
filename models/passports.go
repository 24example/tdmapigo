package models

type PassportTrace struct {
	File     string `json:"file"`
	Line     int    `json:"line"`
	Function string `json:"function"`
	Class    string `json:"class"`
	Type     string `json:"type"`
}

type Passport struct {
	Message   string          `json:"message"`
	Exception string          `json:"exception"`
	File      string          `json:"file"`
	Line      int             `json:"line"`
	Trace     []PassportTrace `json:"trace"`
}

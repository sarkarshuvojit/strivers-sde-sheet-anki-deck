// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    striverQuestions, err := UnmarshalStriverQuestions(bytes)
//    bytes, err = striverQuestions.Marshal()

package types

import "encoding/json"

func UnmarshalStriverQuestions(data []byte) (StriverQuestions, error) {
	var r StriverQuestions
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *StriverQuestions) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type StriverQuestions struct {
	SheetData []SheetDatum `json:"sheetData"`
}

type SheetDatum struct {
	StepNo     int64   `json:"step_no"`
	HeadStepNo string  `json:"head_step_no"`
	Topics     []Topic `json:"topics"`
}

type Topic struct {
	ID          string      `json:"id"`
	StepNo      int64       `json:"step_no"`
	SlNoInStep  int64       `json:"sl_no_in_step"`
	HeadStepNo  string      `json:"head_step_no"`
	Title       string      `json:"title"`
	PostLink    *string     `json:"post_link"`
	YtLink      *string     `json:"yt_link"`
	CSLink      string      `json:"cs_link"`
	GfgLink     *string     `json:"gfg_link"`
	LcLink      *string     `json:"lc_link"`
	CompanyTags interface{} `json:"company_tags"`
	Difficulty  *int64      `json:"difficulty"`
	QuesTopic   *string     `json:"ques_topic"`
}


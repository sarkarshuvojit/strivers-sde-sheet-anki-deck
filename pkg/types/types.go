// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    striverQuestions, err := UnmarshalStriverQuestions(bytes)
//    bytes, err = striverQuestions.Marshal()

package types

import "encoding/json"

type StriverQuestions []StriverQuestion

func UnmarshalStriverQuestions(data []byte) (StriverQuestions, error) {
	var r StriverQuestions
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *StriverQuestions) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type StriverQuestion struct {
	StepNo    int64     `json:"step_no"`
	StepTitle string    `json:"step_title"`
	SubSteps  []SubStep `json:"sub_steps"`
}

type SubStep struct {
	SubStepNo    int64   `json:"sub_step_no"`
	SubStepTitle string  `json:"sub_step_title"`
	Topics       []Topic `json:"topics"`
}

type Topic struct {
	ID            string      `json:"id"`
	StepNo        int64       `json:"step_no"`
	SubStepNo     int64       `json:"sub_step_no"`
	SlNo          int64       `json:"sl_no"`
	StepTitle     string      `json:"step_title"`
	SubStepTitle  string      `json:"sub_step_title"`
	QuestionTitle string      `json:"question_title"`
	PostLink      *string     `json:"post_link"`
	YtLink        *string     `json:"yt_link"`
	GfgLink       *string     `json:"gfg_link"`
	CSLink        *string     `json:"cs_link"`
	LcLink        *string     `json:"lc_link"`
	CompanyTags   interface{} `json:"company_tags"`
	Difficulty    *int64      `json:"difficulty"`
	QuesTopic     *string     `json:"ques_topic"`
}




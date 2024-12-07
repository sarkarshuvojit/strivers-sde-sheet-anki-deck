// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    questionTopics, err := UnmarshalQuestionTopics(bytes)
//    bytes, err = questionTopics.Marshal()

package types

import "encoding/json"

type QuestionTopics []QuestionTopic

func UnmarshalQuestionTopics(data []byte) (QuestionTopics, error) {
	var r QuestionTopics
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *QuestionTopics) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type QuestionTopic struct {
	Value string `json:"value"`
	Label string `json:"label"`
}


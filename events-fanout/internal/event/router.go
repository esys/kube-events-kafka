package event

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Topic   string
	Message []byte
}

func parseEvent(value []byte) (*KubeEvent, error) {
	var data KubeEvent
	err := json.Unmarshal(value, &data)
	if err != nil {
		return nil, fmt.Errorf("failed to decode event: %w", err)
	}
	return &data, nil
}

func CreateDestinationMessage(event []byte) (*Message, error) {
	e, err := parseEvent(event)
	if err != nil {
		return nil, err
	}
	return &Message{
		Topic:   e.Metadata.Namespace,
		Message: event,
	}, nil
}

type KubeEvent struct {
	Verb  string `json:"verb"`
	Event `json:"event"`
}

type Event struct {
	Metadata           `json:"metadata"`
	InvolvedObject     `json:"involvedObject"`
	Reason             string `json:"reason"`
	Message            string `json:"message"`
	Source             `json:"source"`
	FirstTimestamp     string `json:"firstTimestamp"`
	LastTimestamp      string `json:"lastTimestamp"`
	Count              int    `json:"count"`
	Type               string `json:"type"`
	EventTime          string `json:"eventTime"`
	ReportingComponent string `json:"reportingComponent"`
	ReportingInstance  string `json:"reportingInstance"`
}

type Metadata struct {
	Name              string          `json:"name"`
	Namespace         string          `json:"namespace"`
	SelfLink          string          `json:"selfLink"`
	UID               string          `json:"uid"`
	ResourceVersion   string          `json:"resourceVersion"`
	CreationTimestamp string          `json:"creationTimestamp"`
	ManagedFields     []ManagedFields `json:"managedFields"`
}

type ManagedFields struct {
	Manager    string `json:"manager"`
	Operation  string `json:"operation"`
	APIVersion string `json:"apiVersion"`
	Time       string `json:"time"`
}

type InvolvedObject struct {
	Kind            string `json:"kind"`
	Namespace       string `json:"namespace"`
	Name            string `json:"name"`
	UID             string `json:"uid"`
	APIVersion      string `json:"apiVersion"`
	ResourceVersion string `json:"resourceVersion"`
	FieldPath       string `json:"fieldPath"`
}

type Source struct {
	Component string `json:"component"`
	Host      string `json:"host"`
}

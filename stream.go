package tesla

import (
	"bufio"
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// StreamingURL is the base URL for vehicle data streaming
var StreamingURL = "https://streaming.vn.teslamotors.com"

// StreamEventResponse represents an event response returned by vehicle data stream API
type StreamEventResponse struct {
	MsgType           string `json:"msg_type"`
	ConnectionTimeout int    `json:"connection_timeout"`
	PingFrequency     int    `json:"ping_frequency"`
	Autopark          struct {
		HeartbeatFrequency   int `json:"heartbeat_frequency"`
		AutoparkPauseTimeout int `json:"autopark_pause_timeout"`
		AutoparkStopTimeout  int `json:"autopark_stop_timeout"`
	} `json:"autopark"`
	IceServers                              interface{}   `json:"ice_servers"`
	Heading                                 float64       `json:"heading"`
	Latitude                                float64       `json:"latitude"`
	Longitude                               float64       `json:"longitude"`
	ShiftState                              string        `json:"shift_state"`
	Speed                                   float64       `json:"speed"`
	AutoparkState                           string        `json:"autopark_state"`
	AutoparkStateReason                     string        `json:"autopark_state_reason"`
	SmartSummonLastAbortReason              string        `json:"smart_summon_last_abort_reason"`
	SmartSummonLastSafetyMonitorAbortReason string        `json:"smart_summon_last_safety_monitor_abort_reason"`
	SmartSummonLastUnavailableReason        string        `json:"smart_summon_last_unavailable_reason"`
	SmartSummonLastWaitReason               string        `json:"smart_summon_last_wait_reason"`
	SmartSummonLeashProximityMeters         float64       `json:"smart_summon_leash_proximity_meters"`
	SmartSummonLeashTravelMeters            float64       `json:"smart_summon_leash_travel_meters"`
	SmartSummonState                        string        `json:"smart_summon_state"`
	SmartSummonTimeUntilReady               float64       `json:"smart_summon_time_until_ready"`
	HomelinkNearby                          bool          `json:"homelink_nearby"`
	Path                                    []interface{} `json:"path"`
}

// Stream requests a stream from the vehicle and returns a Go channel
func (v Vehicle) Stream() (chan *StreamEventResponse, chan error, error) {
	url := StreamingURL + "/connect/" + strconv.Itoa(v.VehicleID)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Upgrade", "websocket")
	req.Header.Set("Sec-Websocket-Version", "13")
	req.Header.Set("Sec-WebSocket-Key", "SGVsbG8sIHevcmxkIQ==")
	req.SetBasicAuth(ActiveClient.Auth.Email, v.Tokens[0])
	resp, err := ActiveClient.HTTP.Do(req)

	if err != nil {
		return nil, nil, err
	}

	eventChan := make(chan *StreamEventResponse)
	errChan := make(chan error)
	go readStream(resp, eventChan, errChan)

	return eventChan, errChan, nil
}

// readStream reads the stream itself from the vehicle
func readStream(resp *http.Response, eventChan chan *StreamEventResponse, errChan chan error) {
	reader := bufio.NewReader(resp.Body)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	defer resp.Body.Close()

	for scanner.Scan() {
		streamEvent, err := parseStreamEvent(scanner.Text())
		if err == nil {
			eventChan <- streamEvent
		} else {
			errChan <- err
		}
	}
	errChan <- errors.New("HTTP stream closed")
}

// parseStreamEvent parses the stream event, setting all of the appropriate data types
func parseStreamEvent(rawResponse string) (*StreamEventResponse, error) {
	// Cleanup response into something more readable.
	// TODO: make this less hacky
	response := regexp.MustCompile(`[[:^ascii:]]`).ReplaceAllLiteralString(rawResponse, "")
	response = strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) {
			return r
		}
		return -1
	}, response)

	// Remove non-valid characters
	response = regexp.MustCompile(`}[0-9]{`).ReplaceAllString(response, "}~{")
	response = regexp.MustCompile(`}\|{`).ReplaceAllString(response, "}~{")
	response = regexp.MustCompile(`}{`).ReplaceAllString(response, "~")

	// Split the response into a slice of json object strings
	groups := strings.Split(response, "~")

	// Merge into a single json object string
	var groupsMerged string
	for _, v := range groups {
		if len(v) < 3 {
			continue
		}
		if []rune(v)[0] == rune('{') {
			v = "," + v[1:]
		}
		if []rune(v)[len(v)-1] == rune('}') {
			v = v[:len(v)-1] + ","
		}
		groupsMerged += v
	}
	groupsMerged = "{" + groupsMerged[1:len(groupsMerged)-1] + "}"
	groupsMerged = strings.Replace(groupsMerged, ",,", ",", -1)

	// Unmarshal the json object
	// TODO: This object has multiple msg_types
	// TODO: Consider changing StreamEventResponse.MsgType to []string
	streamEvent := &StreamEventResponse{}
	json.Unmarshal([]byte(groupsMerged), streamEvent)

	return streamEvent, nil
}

package tesla

import (
	"bufio"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

var (
	// StreamParams are the params sent to the event stream URL
	StreamParams = "speed,odometer,soc,elevation,est_heading,est_lat,est_lng,power,shift_state,range,est_range,heading"
	// StreamingURL is the base URL for pulling streaming events
	StreamingURL = "https://streaming.vn.teslamotors.com"
)

// StreamEvent is the event returned by the vehicle by the Tesla API
type StreamEvent struct {
	Timestamp  time.Time `json:"timestamp"`
	Speed      int       `json:"speed"`
	Odometer   float64   `json:"odometer"`
	Soc        int       `json:"soc"`
	Elevation  int       `json:"elevation"`
	EstHeading int       `json:"est_heading"`
	EstLat     float64   `json:"est_lat"`
	EstLng     float64   `json:"est_lng"`
	Power      int       `json:"power"`
	ShiftState string    `json:"shift_state"`
	Range      int       `json:"range"`
	EstRange   int       `json:"est_range"`
	Heading    int       `json:"heading"`
}

// Stream requests a stream from the vehicle and returns a Go channel
func (v Vehicle) Stream() (chan *StreamEvent, chan error, error) {
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

	eventChan := make(chan *StreamEvent)
	errChan := make(chan error)
	go readStream(resp, eventChan, errChan)

	return eventChan, errChan, nil
}

// readStream reads the stream itself from the vehicle
func readStream(resp *http.Response, eventChan chan *StreamEvent, errChan chan error) {
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
func parseStreamEvent(event string) (*StreamEvent, error) {
	// TODO: rewrite parser to handle latest response format
	/*
		�~�{"msg_type":"control:hello","connection_timeout":10000,"ping_frequency":5000,"autopark":{"heartbeat_frequency":100,"autopark_pause_timeout":2000,"autopark_stop_timeout":10000},"ice_servers":null}�}{"heading":
		0.0,"latitude":123.123456,"longitude":-123.123456,"msg_type":"vehicle_data:location","shift_state":"P","speed":0.0}�9{"autopark_style":"dead_man","msg_type":"autopark:style"}�~�{"autopark_state":"standby","autopa
		rk_state_reason":"plugged_in","msg_type":"autopark:status","smart_summon_last_abort_reason":"unknown","smart_summon_last_safety_monitor_abort_reason":"none","smart_summon_last_unavailable_reason":"unknown","smar
		t_summon_last_wait_reason":"unknown","smart_summon_leash_proximity_meters":0.0,"smart_summon_leash_travel_meters":0.0,"smart_summon_state":"initializing","smart_summon_time_until_ready":0.0}�6{"homelink_nearby":
		false,"msg_type":"homelink:status"}�2{"msg_type":"autopark:smart_summon_viz","path":[]}�~�{"autopark_state":"standby","autopark_state_reason":"plugged_in","msg_type":"autopark:status","smart_summon_last_abort_re
		ason":"","smart_summon_last_safety_monitor_abort_reason":"none","smart_summon_last_unavailable_reason":"no_phone_location","smart_summon_last_wait_reason":"none","smart_summon_leash_proximity_meters":65.0,"smart
		_summon_leash_travel_meters":145.0,"smart_summon_state":"unavailable","smart_summon_time_until_ready":0.0}�~{"heading":310.5,"latitude":123.123456,"longitude":-123.123456,"msg_type":"vehicle_data:location","shift
		_state":"P","speed":0.0}������������
	*/

	fmt.Println(event)
	return nil, errors.New("todo: rewrite parser")
	/*
		data := strings.Split(event, ",")
		if len(data) != 13 {
			return nil, errors.New("Bad message from Tesla API stream")
		}

		streamEvent := &StreamEvent{}
		timestamp, _ := strconv.ParseInt(data[0], 10, 64)
		streamEvent.Timestamp = time.Unix(0, timestamp*int64(time.Millisecond))
		streamEvent.Speed, _ = strconv.Atoi(data[1])
		streamEvent.Odometer, _ = strconv.ParseFloat(data[2], 64)
		streamEvent.Soc, _ = strconv.Atoi(data[3])
		streamEvent.Elevation, _ = strconv.Atoi(data[4])
		streamEvent.EstHeading, _ = strconv.Atoi(data[5])
		streamEvent.EstLat, _ = strconv.ParseFloat(data[6], 64)
		streamEvent.EstLng, _ = strconv.ParseFloat(data[7], 64)
		streamEvent.Power, _ = strconv.Atoi(data[8])
		streamEvent.ShiftState = data[9]
		streamEvent.Range, _ = strconv.Atoi(data[10])
		streamEvent.EstRange, _ = strconv.Atoi(data[11])
		streamEvent.Heading, _ = strconv.Atoi(data[12])
		return streamEvent, nil
	*/
}

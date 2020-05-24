package tesla

import (
	"encoding/json"
	"errors"
	"strconv"
)

// CommandResponse represents a response from the Tesla API after POSTing a command
type CommandResponse struct {
	Response struct {
		Reason string `json:"reason"`
		Result bool   `json:"result"`
	} `json:"response"`
}

// AutoParkRequest represents parameters to POST an Autopark/Summon request
type AutoParkRequest struct {
	VehicleID int     `json:"vehicle_id,omitempty"`
	Lat       float64 `json:"lat"`
	Lon       float64 `json:"lon"`
	Action    string  `json:"action,omitempty"`
}

// AutoparkAbort aborts an autopark request
func (v Vehicle) AutoparkAbort() error {
	return v.autoPark("abort")
}

// AutoparkForward commands the vehicle to pull forward
func (v Vehicle) AutoparkForward() error {
	return v.autoPark("start_forward")
}

// AutoparkReverse commands the vehicle to go in reverse
func (v Vehicle) AutoparkReverse() error {
	return v.autoPark("start_reverse")
}

// autoPark performs the auto park/summon request for the vehicle
func (v Vehicle) autoPark(action string) error {
	apiURL := BaseURL + "/vehicles/" + strconv.FormatInt(v.ID, 10) + "/command/autopark_request"
	driveState, _ := v.DriveState()
	autoParkRequest := &AutoParkRequest{
		VehicleID: v.VehicleID,
		Lat:       driveState.Latitude,
		Lon:       driveState.Longitude,
		Action:    action,
	}
	body, _ := json.Marshal(autoParkRequest)

	_, err := sendCommand(apiURL, body)
	return err
}

// TBD based on Github issue #7
// Toggles defrost on and off, locations values are 'front' or 'rear'
// func (v Vehicle) Defrost(location string, state bool) error {
// 	command := location + "_defrost_"
// 	if state {
// 		command += "on"
// 	} else {
// 		command += "off"
// 	}
// 	apiURL := BaseURL + "/vehicles/" + strconv.FormatInt(v.ID, 10) + "/command/" + command
// 	fmt.Println(apiURL)
// 	_, err := sendCommand(apiURL, nil)
// 	return err
// }

// TriggerHomelink opens and closes the configured Homelink garage door of the vehicle
// This is a toggle and the garage door state is unknown
func (v Vehicle) TriggerHomelink() error {
	apiURL := BaseURL + "/vehicles/" + strconv.FormatInt(v.ID, 10) + "/command/trigger_homelink"
	driveState, _ := v.DriveState()
	autoParkRequest := &AutoParkRequest{
		Lat: driveState.Latitude,
		Lon: driveState.Longitude,
	}
	body, _ := json.Marshal(autoParkRequest)

	_, err := sendCommand(apiURL, body)
	return err
}

// Wakeup wakes up the vehicle when it is powered off
func (v Vehicle) Wakeup() (*Vehicle, error) {
	apiURL := BaseURL + "/vehicles/" + strconv.FormatInt(v.ID, 10) + "/wake_up"
	body, err := sendCommand(apiURL, nil)
	if err != nil {
		return nil, err
	}
	vehicleResponse := &VehicleResponse{}
	err = json.Unmarshal(body, vehicleResponse)
	if err != nil {
		return nil, err
	}
	return vehicleResponse.Response, nil
}

// OpenChargePort opens the vehicle's charge port
func (v Vehicle) OpenChargePort() error {
	apiURL := BaseURL + "/vehicles/" + strconv.FormatInt(v.ID, 10) + "/command/charge_port_door_open"
	_, err := sendCommand(apiURL, nil)
	return err
}

// CloseChargePort closes the vehicle's charge port
func (v Vehicle) CloseChargePort() error {
	apiURL := BaseURL + "/vehicles/" + strconv.FormatInt(v.ID, 10) + "/command/charge_port_door_close"
	_, err := sendCommand(apiURL, nil)
	return err
}

// ResetValetPIN resets the valet mode PIN, if set
func (v Vehicle) ResetValetPIN() error {
	apiURL := BaseURL + "/vehicles/" + strconv.FormatInt(v.ID, 10) + "/command/reset_valet_pin"
	_, err := sendCommand(apiURL, nil)
	return err
}

// SetChargeLimitStandard sets the charge limit to the default setting
func (v Vehicle) SetChargeLimitStandard() error {
	apiURL := BaseURL + "/vehicles/" + strconv.FormatInt(v.ID, 10) + "/command/charge_standard"
	_, err := sendCommand(apiURL, nil)
	return err
}

// SetChargeLimitMax sets the charge limit to the maximum value
func (v Vehicle) SetChargeLimitMax() error {
	apiURL := BaseURL + "/vehicles/" + strconv.FormatInt(v.ID, 10) + "/command/charge_max_range"
	_, err := sendCommand(apiURL, nil)
	return err
}

// SetChargeLimit sets the charge limit to a supplied percent value
func (v Vehicle) SetChargeLimit(percent int) error {
	apiURL := BaseURL + "/vehicles/" + strconv.FormatInt(v.ID, 10) + "/command/set_charge_limit"
	postJSON := `{"percent": ` + strconv.Itoa(percent) + `}`
	_, err := ActiveClient.post(apiURL, []byte(postJSON))
	return err
}

// StartCharging starts the charging of the vehicle if charging cable is inserted
func (v Vehicle) StartCharging() error {
	apiURL := BaseURL + "/vehicles/" + strconv.FormatInt(v.ID, 10) + "/command/charge_start"
	_, err := sendCommand(apiURL, nil)
	return err
}

// StopCharging stops a vehicle's charge session
func (v Vehicle) StopCharging() error {
	apiURL := BaseURL + "/vehicles/" + strconv.FormatInt(v.ID, 10) + "/command/charge_stop"
	_, err := sendCommand(apiURL, nil)
	return err
}

// FlashLights flashes the lights of the vehicle
func (v Vehicle) FlashLights() error {
	apiURL := BaseURL + "/vehicles/" + strconv.FormatInt(v.ID, 10) + "/command/flash_lights"
	_, err := sendCommand(apiURL, nil)
	return err
}

// HonkHorn honks the vehicle's horn
func (v *Vehicle) HonkHorn() error {
	apiURL := BaseURL + "/vehicles/" + strconv.FormatInt(v.ID, 10) + "/command/honk_horn"
	_, err := sendCommand(apiURL, nil)
	return err
}

// UnlockDoors unlocks the vehicle's doors
func (v Vehicle) UnlockDoors() error {
	apiURL := BaseURL + "/vehicles/" + strconv.FormatInt(v.ID, 10) + "/command/door_unlock"
	_, err := sendCommand(apiURL, nil)
	return err
}

// LockDoors locks the vehicle's doors
func (v Vehicle) LockDoors() error {
	apiURL := BaseURL + "/vehicles/" + strconv.FormatInt(v.ID, 10) + "/command/door_lock"
	_, err := sendCommand(apiURL, nil)
	return err
}

// SetTemperature sets the temperature of the vehicle
// Driver and passenger zones are controlled individually
func (v Vehicle) SetTemperature(driver float64, passenger float64) error {
	driveTemp := strconv.FormatFloat(driver, 'f', -1, 32)
	passengerTemp := strconv.FormatFloat(passenger, 'f', -1, 32)
	apiURL := BaseURL + "/vehicles/" + strconv.FormatInt(v.ID, 10) + "/command/set_temps"
	postJSON := `{"driver_temp": "` + driveTemp + `", "passenger_temp":` + passengerTemp + `}`
	_, err := ActiveClient.post(apiURL, []byte(postJSON))

	return err
}

// StartAirConditioning starts the vehicle's air conditioner
func (v Vehicle) StartAirConditioning() error {
	url := BaseURL + "/vehicles/" + strconv.FormatInt(v.ID, 10) + "/command/auto_conditioning_start"
	_, err := sendCommand(url, nil)
	return err
}

// StopAirConditioning stops the vehicle's air conditioner
func (v Vehicle) StopAirConditioning() error {
	apiURL := BaseURL + "/vehicles/" + strconv.FormatInt(v.ID, 10) + "/command/auto_conditioning_stop"
	_, err := sendCommand(apiURL, nil)
	return err
}

// MovePanoRoof controls the state of the panoramic roof. The approximate percent open
// values for each state are open = 100%, close = 0%, comfort = 80%, vent = %15, move = set %
func (v Vehicle) MovePanoRoof(state string, percent int) error {
	apiURL := BaseURL + "/vehicles/" + strconv.FormatInt(v.ID, 10) + "/command/sun_roof_control"
	postJSON := `{"state": "` + state + `", "percent":` + strconv.Itoa(percent) + `}`
	_, err := ActiveClient.post(apiURL, []byte(postJSON))
	return err
}

// Start starts the car by turning it on. Requires the Tesla account password
func (v Vehicle) Start(password string) error {
	apiURL := BaseURL + "/vehicles/" + strconv.FormatInt(v.ID, 10) + "/command/remote_start_drive?password=" + password
	_, err := sendCommand(apiURL, nil)
	return err
}

// OpenTrunk opens the trunk. Valid trunk values are 'front' and 'rear'
func (v Vehicle) OpenTrunk(trunk string) error {
	apiURL := BaseURL + "/vehicles/" + strconv.FormatInt(v.ID, 10) + "/command/trunk_open" // ?which_trunk=" + trunk
	postJSON := `{"which_trunk": "` + trunk + `"}`
	_, err := ActiveClient.post(apiURL, []byte(postJSON))
	return err
}

// VentWindows vents the vehicle's windows
func (v Vehicle) VentWindows() error {
	return v.windows("vent")
}

// CloseWindows closes the vehicle's windows (model 3 only?)
func (v Vehicle) CloseWindows() error {
	return v.windows("close")
}

// windows vents or closes the windows
func (v Vehicle) windows(action string) error {
	apiURL := BaseURL + "/vehicles/" + strconv.FormatInt(v.ID, 10) + "/command/window_control"
	windowRequest := struct {
		VehicleID string `json:"command"`
		Lat       int    `json:"lat"`
		Lon       int    `json:"lon"`
	}{
		action, 0, 0,
	}
	body, _ := json.Marshal(windowRequest)

	_, err := sendCommand(apiURL, body)
	return err
}

// SetSentryMode controls Sentry Mode's active state (true/false)
func (v Vehicle) SetSentryMode(on bool) error {
	apiURL := BaseURL + "/vehicles/" + strconv.FormatInt(v.ID, 10) + "/command/set_sentry_mode"
	postJSON := `{"on": "` + strconv.FormatBool(on) + `"}`
	_, err := sendCommand(apiURL, []byte(postJSON))
	return err
}

// HeatSeat sets heating for the supplied seat number (0=driver, 1=passenger, 2=rear-left...)
func (v Vehicle) HeatSeat(seat, level int) error {
	//requires climate to be set first
	err := v.StartAirConditioning()
	if err != nil {
		panic(err)
	}
	apiURL := BaseURL + "/vehicles/" + strconv.FormatInt(v.ID, 10) + "/command/remote_seat_heater_request"
	seatRequest := struct {
		Heater int `json:"heater"`
		Level  int `json:"level"`
	}{
		seat, level,
	}
	body, _ := json.Marshal(seatRequest)
	_, err = sendCommand(apiURL, body)
	return err
}

// HeatWheel turns steering wheel heat on or off
func (v Vehicle) HeatWheel(on bool) error {
	//requires climate to be set first
	err := v.StartAirConditioning()

	apiURL := BaseURL + "/vehicles/" + strconv.FormatInt(v.ID, 10) + "/command/remote_steering_wheel_heater_request"
	postJSON := `{"on": "` + strconv.FormatBool(on) + `"}`
	_, err = sendCommand(apiURL, []byte(postJSON))
	return err
}

// ScheduleSoftwareUpdate schedules the installation of the available software update.
// An update must already be available for this command to work
func (v Vehicle) ScheduleSoftwareUpdate(offset int64) error {
	apiURL := BaseURL + "/vehicles/" + strconv.FormatInt(v.ID, 10) + "/command/schedule_software_update"
	theJSON := `{"offset_sec": ` + strconv.FormatInt(offset, 10) + `}`
	_, err := ActiveClient.post(apiURL, []byte(theJSON))
	return err
}

// CancelSoftwareUpdate cancels a previously-scheduled software update that has not yet started
func (v Vehicle) CancelSoftwareUpdate() error {
	apiURL := BaseURL + "/vehicles/" + strconv.FormatInt(v.ID, 10) + "/command/cancel_software_update"
	_, err := sendCommand(apiURL, nil)
	return err
}

// sendCommand sends a command to the vehicle
func sendCommand(url string, reqBody []byte) ([]byte, error) {
	body, err := ActiveClient.post(url, reqBody)
	if err != nil {
		return nil, err
	}
	if len(body) > 0 {
		response := &CommandResponse{}
		err = json.Unmarshal(body, response)
		if err != nil {
			return nil, err
		}
		if response.Response.Result != true && response.Response.Reason != "" {
			return nil, errors.New(response.Response.Reason)
		}
	}
	return body, nil
}

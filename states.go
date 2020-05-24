package tesla

import (
	"encoding/json"
	"strconv"
)

// ChargeState contains the current charge states that exist within the vehicle
type ChargeState struct {
	BatteryHeaterOn             bool        `json:"battery_heater_on"`
	BatteryLevel                int         `json:"battery_level"`
	BatteryRange                float64     `json:"battery_range"`
	ChargeCurrentRequest        int         `json:"charge_current_request"`
	ChargeCurrentRequestMax     int         `json:"charge_current_request_max"`
	ChargeEnableRequest         bool        `json:"charge_enable_request"`
	ChargeEnergyAdded           float64     `json:"charge_energy_added"`
	ChargeLimitSoc              int         `json:"charge_limit_soc"`
	ChargeLimitSocMax           int         `json:"charge_limit_soc_max"`
	ChargeLimitSocMin           int         `json:"charge_limit_soc_min"`
	ChargeLimitSocStd           int         `json:"charge_limit_soc_std"`
	ChargeMilesAddedIdeal       float64     `json:"charge_miles_added_ideal"`
	ChargeMilesAddedRated       float64     `json:"charge_miles_added_rated"`
	ChargePortColdWeatherMode   interface{} `json:"charge_port_cold_weather_mode"`
	ChargePortDoorOpen          bool        `json:"charge_port_door_open"`
	ChargePortLatch             string      `json:"charge_port_latch"`
	ChargeRate                  float64     `json:"charge_rate"`
	ChargeToMaxRange            bool        `json:"charge_to_max_range"`
	ChargerActualCurrent        int         `json:"charger_actual_current"`
	ChargerPhases               interface{} `json:"charger_phases"`
	ChargerPilotCurrent         int         `json:"charger_pilot_current"`
	ChargerPower                int         `json:"charger_power"`
	ChargerVoltage              int         `json:"charger_voltage"`
	ChargingState               string      `json:"charging_state"`
	ConnChargeCable             string      `json:"conn_charge_cable"`
	EstBatteryRange             float64     `json:"est_battery_range"`
	FastChargerBrand            string      `json:"fast_charger_brand"`
	FastChargerPresent          bool        `json:"fast_charger_present"`
	FastChargerType             string      `json:"fast_charger_type"`
	IdealBatteryRange           float64     `json:"ideal_battery_range"`
	ManagedChargingActive       bool        `json:"managed_charging_active"`
	ManagedChargingStartTime    interface{} `json:"managed_charging_start_time"`
	ManagedChargingUserCanceled bool        `json:"managed_charging_user_canceled"`
	MaxRangeChargeCounter       int         `json:"max_range_charge_counter"`
	MinutesToFullCharge         int         `json:"minutes_to_full_charge"`
	NotEnoughPowerToHeat        bool        `json:"not_enough_power_to_heat"`
	ScheduledChargingPending    bool        `json:"scheduled_charging_pending"`
	ScheduledChargingStartTime  interface{} `json:"scheduled_charging_start_time"`
	TimeToFullCharge            float64     `json:"time_to_full_charge"`
	Timestamp                   int64       `json:"timestamp"`
	TripCharging                bool        `json:"trip_charging"`
	UsableBatteryLevel          int         `json:"usable_battery_level"`
	UserChargeEnableRequest     bool        `json:"user_charge_enable_request"`
}

// ClimateState contains the current climate states availale from the vehicle
type ClimateState struct {
	BatteryHeater              bool    `json:"battery_heater"`
	BatteryHeaterNoPower       bool    `json:"battery_heater_no_power"`
	ClimateKeeperMode          string  `json:"climate_keeper_mode"`
	DefrostMode                int     `json:"defrost_mode"`
	DriverTempSetting          float64 `json:"driver_temp_setting"`
	FanStatus                  int     `json:"fan_status"`
	InsideTemp                 float64 `json:"inside_temp"`
	IsAutoConditioningOn       bool    `json:"is_auto_conditioning_on"`
	IsClimateOn                bool    `json:"is_climate_on"`
	IsFrontDefrosterOn         bool    `json:"is_front_defroster_on"`
	IsPreconditioning          bool    `json:"is_preconditioning"`
	IsRearDefrosterOn          bool    `json:"is_rear_defroster_on"`
	LeftTempDirection          int     `json:"left_temp_direction"`
	MaxAvailTemp               float64 `json:"max_avail_temp"`
	MinAvailTemp               float64 `json:"min_avail_temp"`
	OutsideTemp                float64 `json:"outside_temp"`
	PassengerTempSetting       float64 `json:"passenger_temp_setting"`
	RemoteHeaterControlEnabled bool    `json:"remote_heater_control_enabled"`
	RightTempDirection         int     `json:"right_temp_direction"`
	SeatHeaterLeft             int     `json:"seat_heater_left"`
	SeatHeaterRearCenter       int     `json:"seat_heater_rear_center"`
	SeatHeaterRearLeft         int     `json:"seat_heater_rear_left"`
	SeatHeaterRearRight        int     `json:"seat_heater_rear_right"`
	SeatHeaterRight            int     `json:"seat_heater_right"`
	SeatHeaterThirdRowLeft     int     `json:"seat_heater_third_row_left"`
	SeatHeaterThirdRowRight    int     `json:"seat_heater_third_row_right"`
	SideMirrorHeaters          bool    `json:"side_mirror_heaters"`
	SteeringWheelHeater        bool    `json:"steering_wheel_heater"`
	Timestamp                  int64   `json:"timestamp"`
	WiperBladeHeater           bool    `json:"wiper_blade_heater"`
}

// DriveState contains the current drive state of the vehicle
type DriveState struct {
	GpsAsOf                 int         `json:"gps_as_of"`
	Heading                 int         `json:"heading"`
	Latitude                float64     `json:"latitude"`
	Longitude               float64     `json:"longitude"`
	NativeLatitude          float64     `json:"native_latitude"`
	NativeLocationSupported int         `json:"native_location_supported"`
	NativeLongitude         float64     `json:"native_longitude"`
	NativeType              string      `json:"native_type"`
	Power                   int         `json:"power"`
	ShiftState              string      `json:"shift_state"`
	Speed                   interface{} `json:"speed"`
	Timestamp               int64       `json:"timestamp"`
}

// GuiSettings contains the current GUI settings of the vehicle
type GuiSettings struct {
	Gui24HourTime       bool   `json:"gui_24_hour_time"`
	GuiChargeRateUnits  string `json:"gui_charge_rate_units"`
	GuiDistanceUnits    string `json:"gui_distance_units"`
	GuiRangeDisplay     string `json:"gui_range_display"`
	GuiTemperatureUnits string `json:"gui_temperature_units"`
	ShowRangeUnits      bool   `json:"show_range_units"`
	Timestamp           int64  `json:"timestamp"`
}

// VehicleConfig describes the vehicle's configured hardware and software features
type VehicleConfig struct {
	CanAcceptNavigationRequests bool   `json:"can_accept_navigation_requests"`
	CanActuateTrunks            bool   `json:"can_actuate_trunks"`
	CarSpecialType              string `json:"car_special_type"`
	CarType                     string `json:"car_type"`
	ChargePortType              string `json:"charge_port_type"`
	EuVehicle                   bool   `json:"eu_vehicle"`
	ExteriorColor               string `json:"exterior_color"`
	HasAirSuspension            bool   `json:"has_air_suspension"`
	HasLudicrousMode            bool   `json:"has_ludicrous_mode"`
	KeyVersion                  int    `json:"key_version"`
	HasMotorizedChargePort      bool   `json:"motorized_charge_port"`
	PerfConfig                  string `json:"perf_config"`
	PLG                         bool   `json:"plg"`
	RearSeatHeaters             int    `json:"rear_seat_heaters"`
	RearSeatType                int    `json:"rear_seat_type"`
	RHD                         bool   `json:"rhd"`
	RoofColor                   string `json:"roof_color"`
	SeatType                    int    `json:"seat_type"`
	SpoilerType                 string `json:"spoiler_type"`
	SunRoofInstalled            int    `json:"sun_roof_installed"`
	ThirdRowSeats               string `json:"third_row_seats"`
	Timestamp                   uint64 `json:"timestamp"`
	TrimBadging                 string `json:"trim_badging"`
	WheelType                   string `json:"wheel_type"`
}

// VehicleState contains the current state of the vehicle
type VehicleState struct {
	APIVersion          int    `json:"api_version"`
	AutoparkStateV2     string `json:"autopark_state_v2"`
	AutoparkStyle       string `json:"autopark_style"`
	CalendarSupported   bool   `json:"calendar_supported"`
	CarVersion          string `json:"car_version"`
	CenterDisplayState  int    `json:"center_display_state"`
	Df                  int    `json:"df"`
	Dr                  int    `json:"dr"`
	FdWindow            int    `json:"fd_window"`
	FpWindow            int    `json:"fp_window"`
	Ft                  int    `json:"ft"`
	HomelinkDeviceCount int    `json:"homelink_device_count"`
	HomelinkNearby      bool   `json:"homelink_nearby"`
	IsUserPresent       bool   `json:"is_user_present"`
	LastAutoparkError   string `json:"last_autopark_error"`
	Locked              bool   `json:"locked"`
	MediaState          struct {
		RemoteControlEnabled bool `json:"remote_control_enabled"`
	} `json:"media_state"`
	NotificationsSupported  bool    `json:"notifications_supported"`
	Odometer                float64 `json:"odometer"`
	ParsedCalendarSupported bool    `json:"parsed_calendar_supported"`
	Pf                      int     `json:"pf"`
	Pr                      int     `json:"pr"`
	RdWindow                int     `json:"rd_window"`
	RemoteStart             bool    `json:"remote_start"`
	RemoteStartEnabled      bool    `json:"remote_start_enabled"`
	RemoteStartSupported    bool    `json:"remote_start_supported"`
	RpWindow                int     `json:"rp_window"`
	Rt                      int     `json:"rt"`
	SentryMode              bool    `json:"sentry_mode"`
	SentryModeAvailable     bool    `json:"sentry_mode_available"`
	SmartSummonAvailable    bool    `json:"smart_summon_available"`
	SoftwareUpdate          struct {
		DownloadPerc        int    `json:"download_perc"`
		ExpectedDurationSec int    `json:"expected_duration_sec"`
		InstallPerc         int    `json:"install_perc"`
		Status              string `json:"status"`
		Version             string `json:"version"`
	} `json:"software_update"`
	SpeedLimitMode struct {
		Active          bool    `json:"active"`
		CurrentLimitMph float64 `json:"current_limit_mph"`
		MaxLimitMph     int     `json:"max_limit_mph"`
		MinLimitMph     int     `json:"min_limit_mph"`
		PinCodeSet      bool    `json:"pin_code_set"`
	} `json:"speed_limit_mode"`
	SummonStandbyModeEnabled bool   `json:"summon_standby_mode_enabled"`
	Timestamp                int64  `json:"timestamp"`
	ValetMode                bool   `json:"valet_mode"`
	ValetPinNeeded           bool   `json:"valet_pin_needed"`
	VehicleName              string `json:"vehicle_name"`
}

// StateRequest represents the request to get the states of the vehicle
type StateRequest struct {
	Response struct {
		*ChargeState
		*ClimateState
		*DriveState
		*GuiSettings
		*VehicleConfig
		*VehicleState
	} `json:"response"`
}

// Response represents a state request response from the API server
type Response struct {
	Bool bool `json:"response"`
}

// MobileEnabled returns a flag indicating whether the vehicle is mobile enabled for Tesla API control
func (v *Vehicle) MobileEnabled() (bool, error) {
	body, err := ActiveClient.get(BaseURL + "/vehicles/" + strconv.FormatInt(v.ID, 10) + "/mobile_enabled")
	if err != nil {
		return false, err
	}
	response := &Response{}
	err = json.Unmarshal(body, response)
	if err != nil {
		return false, err
	}
	return response.Bool, nil
}

// ChargeState returns the charge state of the vehicle
func (v *Vehicle) ChargeState() (*ChargeState, error) {
	stateRequest, err := fetchState("/charge_state", v.ID)
	if err != nil {
		return nil, err
	}
	return stateRequest.Response.ChargeState, nil
}

// ClimateState returns the climate state of the vehicle
func (v Vehicle) ClimateState() (*ClimateState, error) {
	stateRequest, err := fetchState("/climate_state", v.ID)
	if err != nil {
		return nil, err
	}
	return stateRequest.Response.ClimateState, nil
}

// DriveState returns the drive state of the vehicle
func (v Vehicle) DriveState() (*DriveState, error) {
	stateRequest, err := fetchState("/drive_state", v.ID)
	if err != nil {
		return nil, err
	}
	return stateRequest.Response.DriveState, nil
}

// GuiSettings returns the GUI settings of the vehicle
func (v Vehicle) GuiSettings() (*GuiSettings, error) {
	stateRequest, err := fetchState("/gui_settings", v.ID)
	if err != nil {
		return nil, err
	}
	return stateRequest.Response.GuiSettings, nil
}

// VehicleConfig retrieves the vehicle's configured features
func (v Vehicle) VehicleConfig() (*VehicleConfig, error) {
	stateRequest, err := fetchState("/vehicle_config", v.ID)
	if err != nil {
		return nil, err
	}
	return stateRequest.Response.VehicleConfig, nil
}

// VehicleState returns the vehicle state
func (v Vehicle) VehicleState() (*VehicleState, error) {
	stateRequest, err := fetchState("/vehicle_state", v.ID)
	if err != nil {
		return nil, err
	}
	return stateRequest.Response.VehicleState, nil
}

// fetchState fetches the a given state of the vehicle
func fetchState(resource string, id int64) (*StateRequest, error) {
	stateRequest := &StateRequest{}
	body, err := ActiveClient.get(BaseURL + "/vehicles/" + strconv.FormatInt(id, 10) + "/data_request" + resource)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, stateRequest)
	if err != nil {
		return nil, err
	}
	return stateRequest, nil
}

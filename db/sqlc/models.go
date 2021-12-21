// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
	"time"
)

type Device struct {
	DeviceName        string         `json:"device_name"`
	LastUpdated       string         `json:"last_updated"`
	Expected          string         `json:"expected"`
	Price             int32          `json:"price"`
	ImgUrl            string         `json:"img_url"`
	SourceUrl         string         `json:"source_url"`
	SpecScore         int32          `json:"spec_score"`
	Ram               sql.NullString `json:"ram"`
	Processor         sql.NullString `json:"processor"`
	FrontCamera       sql.NullString `json:"front_camera"`
	RearCamera        sql.NullString `json:"rear_camera"`
	Battery           sql.NullString `json:"battery"`
	Display           sql.NullString `json:"display"`
	OperatingSystem   sql.NullString `json:"operating_system"`
	CustomUi          sql.NullString `json:"custom_ui"`
	Chipset           sql.NullString `json:"chipset"`
	Cpu               sql.NullString `json:"cpu"`
	Architecture      sql.NullString `json:"architecture"`
	Graphics          sql.NullString `json:"graphics"`
	DisplayType       sql.NullString `json:"display_type"`
	ScreenSize        sql.NullString `json:"screen_size"`
	Resolution        sql.NullString `json:"resolution"`
	PixelDensity      sql.NullString `json:"pixel_density"`
	Touchscreen       sql.NullString `json:"touchscreen"`
	InternalMemory    sql.NullString `json:"internal_memory"`
	ExpandableMemory  sql.NullString `json:"expandable_memory"`
	MCameraSetup      sql.NullString `json:"m_camera_setup"`
	MResolution       sql.NullString `json:"m_resolution"`
	MAutofocus        sql.NullString `json:"m_autofocus"`
	MOis              sql.NullString `json:"m_ois"`
	MSensors          sql.NullString `json:"m_sensors"`
	MFlash            sql.NullString `json:"m_flash"`
	MImageResolution  sql.NullString `json:"m_image_resolution"`
	MSettings         sql.NullString `json:"m_settings"`
	MShootingModes    sql.NullString `json:"m_shooting_modes"`
	MCameraFeatures   sql.NullString `json:"m_camera_features"`
	MVideoRecording   sql.NullString `json:"m_video_recording"`
	SCameraSetup      sql.NullString `json:"s_camera_setup"`
	SResolution       sql.NullString `json:"s_resolution"`
	SVideoRecording   sql.NullString `json:"s_video_recording"`
	Capacity          sql.NullString `json:"capacity"`
	RemovableBattery  sql.NullString `json:"removable_battery"`
	WirelessCharging  sql.NullString `json:"wireless_charging"`
	QuickCharging     sql.NullString `json:"quick_charging"`
	Usb               sql.NullString `json:"usb"`
	SimSlots          sql.NullString `json:"sim_slots"`
	NetworkSupport    sql.NullString `json:"network_support"`
	FingerprintSensor sql.NullString `json:"fingerprint_sensor"`
	OtherSensors      sql.NullString `json:"other_sensors"`
	ScrapeTimestamp   time.Time      `json:"scrape_timestamp"`
}

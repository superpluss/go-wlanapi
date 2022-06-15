package wlanapi

//The WFD_DISPLAY_SINK_NOTIFICATION_TYPE enumerated type defines the type of the notification passed to the WFD_DISPLAY_SINK_NOTIFICATION_CALLBACK function.
//https://docs.microsoft.com/en-us/windows/win32/nativewifi/wfd-display-sink-notification-type
type WFD_DISPLAY_SINK_NOTIFICATION_TYPE uint32

const (
	ProvisioningRequestNotification WFD_DISPLAY_SINK_NOTIFICATION_TYPE = 0
	ReconnectRequestNotification    WFD_DISPLAY_SINK_NOTIFICATION_TYPE = 1
	ConnectedNotification           WFD_DISPLAY_SINK_NOTIFICATION_TYPE = 2
	DisconnectedNotification        WFD_DISPLAY_SINK_NOTIFICATION_TYPE = 3
)

//The DOT11_AUTH_ALGORITHM enumerated type defines a wireless LAN authentication algorithm.
//https://docs.microsoft.com/en-us/windows/win32/nativewifi/dot11-auth-algorithm
type DOT11_AUTH_ALGORITHM uint32

const (
	DOT11_AUTH_ALGO_80211_OPEN       DOT11_AUTH_ALGORITHM = 1
	DOT11_AUTH_ALGO_80211_SHARED_KEY DOT11_AUTH_ALGORITHM = 2
	DOT11_AUTH_ALGO_WPA              DOT11_AUTH_ALGORITHM = 3
	DOT11_AUTH_ALGO_WPA_PSK          DOT11_AUTH_ALGORITHM = 4
	DOT11_AUTH_ALGO_WPA_NONE         DOT11_AUTH_ALGORITHM = 5
	DOT11_AUTH_ALGO_RSNA             DOT11_AUTH_ALGORITHM = 6
	DOT11_AUTH_ALGO_RSNA_PSK         DOT11_AUTH_ALGORITHM = 7
	DOT11_AUTH_ALGO_IHV_START        DOT11_AUTH_ALGORITHM = 0x80000000
	DOT11_AUTH_ALGO_IHV_END          DOT11_AUTH_ALGORITHM = 0xffffffff
)

//The DOT11_BSS_TYPE enumerated type defines a basic service set (BSS) network type.
//https://docs.microsoft.com/en-us/windows/win32/nativewifi/dot11-bss-type
type DOT11_BSS_TYPE uint32

const (
	dot11_BSS_type_infrastructure DOT11_BSS_TYPE = 1
	dot11_BSS_type_independent    DOT11_BSS_TYPE = 2
	dot11_BSS_type_any            DOT11_BSS_TYPE = 3
)

//The DOT11_CIPHER_ALGORITHM enumerated type defines a cipher algorithm for data encryption and decryption.
//https://docs.microsoft.com/en-us/windows/win32/nativewifi/dot11-cipher-algorithm
type DOT11_CIPHER_ALGORITHM uint32

const (
	DOT11_CIPHER_ALGO_NONE          DOT11_CIPHER_ALGORITHM = 0x00
	DOT11_CIPHER_ALGO_WEP40         DOT11_CIPHER_ALGORITHM = 0x01
	DOT11_CIPHER_ALGO_TKIP          DOT11_CIPHER_ALGORITHM = 0x02
	DOT11_CIPHER_ALGO_CCMP          DOT11_CIPHER_ALGORITHM = 0x04
	DOT11_CIPHER_ALGO_WEP104        DOT11_CIPHER_ALGORITHM = 0x05
	DOT11_CIPHER_ALGO_WPA_USE_GROUP DOT11_CIPHER_ALGORITHM = 0x100
	DOT11_CIPHER_ALGO_RSN_USE_GROUP DOT11_CIPHER_ALGORITHM = 0x100
	DOT11_CIPHER_ALGO_WEP           DOT11_CIPHER_ALGORITHM = 0x101
	DOT11_CIPHER_ALGO_IHV_START     DOT11_CIPHER_ALGORITHM = 0x80000000
	DOT11_CIPHER_ALGO_IHV_END       DOT11_CIPHER_ALGORITHM = 0xffffffff
)

//The DOT11_PHY_TYPE enumeration defines an 802.11 PHY and media type.
//https://docs.microsoft.com/en-us/windows/win32/nativewifi/dot11-phy-type
type DOT11_PHY_TYPE uint32

const (
	dot11_phy_type_unknownDOT11_PHY_TYPE DOT11_PHY_TYPE = 0
	dot11_phy_type_anyDOT11_PHY_TYPE     DOT11_PHY_TYPE = 0
	dot11_phy_type_fhss                  DOT11_PHY_TYPE = 1
	dot11_phy_type_dsss                  DOT11_PHY_TYPE = 2
	dot11_phy_type_irbaseband            DOT11_PHY_TYPE = 3
	dot11_phy_type_ofdm                  DOT11_PHY_TYPE = 4
	dot11_phy_type_hrdsss                DOT11_PHY_TYPE = 5
	dot11_phy_type_erp                   DOT11_PHY_TYPE = 6
	dot11_phy_type_ht                    DOT11_PHY_TYPE = 7
	dot11_phy_type_vht                   DOT11_PHY_TYPE = 8
	dot11_phy_type_IHV_start             DOT11_PHY_TYPE = 0x80000000
	dot11_phy_type_IHV_end               DOT11_PHY_TYPE = 0xffffffff
)

//The DOT11_RADIO_STATE enumeration specifies an 802.11 radio state.
//https://docs.microsoft.com/en-us/windows/win32/api/wlanapi/ne-wlanapi-dot11_radio_state-r1
type DOT11_RADIO_STATE uint32

const (
	dot11_radio_state_unknown DOT11_RADIO_STATE = iota
	dot11_radio_state_on
	dot11_radio_state_off
)

//Specifies the active tab when the wireless profile user interface dialog box appears.
//https://docs.microsoft.com/en-us/windows/win32/api/wlanapi/ne-wlanapi-wl_display_pages
type WL_DISPLAY_PAGES uint32

const (
	WLConnectionPage WL_DISPLAY_PAGES = iota
	WLSecurityPage
	WLAdvPage
)

//The WLAN_CONNECTION_MODE enumerated type defines the mode of connection.Windows XP with SP3 and Wireless LAN API for Windows XP with SP2:  Only the wlan_connection_mode_profile value is supported.
//https://docs.microsoft.com/en-us/windows/win32/api/wlanapi/ne-wlanapi-wlan_connection_mode
type WLAN_CONNECTION_MODE uint32

const (
	wlan_connection_mode_profile WLAN_CONNECTION_MODE = iota
	wlan_connection_mode_temporary_profile
	wlan_connection_mode_discovery_secure
	wlan_connection_mode_discovery_unsecure
	wlan_connection_mode_auto
	wlan_connection_mode_invalid
)

//The WLAN_FILTER_LIST_TYPE enumerated type indicates types of filter lists.
//https://docs.microsoft.com/en-us/windows/win32/api/wlanapi/ne-wlanapi-wlan_filter_list_type
type WLAN_FILTER_LIST_TYPE uint32

const (
	wlan_filter_list_type_gp_permit WLAN_FILTER_LIST_TYPE = iota
	wlan_filter_list_type_gp_deny
	wlan_filter_list_type_user_permit
	wlan_filter_list_type_user_deny
)

//The WLAN_HOSTED_NETWORK_NOTIFICATION_CODE enumerated type specifies the possible values of the NotificationCode parameter for received notifications on the wireless Hosted Network.
//https://docs.microsoft.com/en-us/windows/win32/api/wlanapi/ne-wlanapi-wlan_hosted_network_notification_code
type WLAN_HOSTED_NETWORK_NOTIFICATION_CODE uint32

const (
	wlan_hosted_network_state_change WLAN_HOSTED_NETWORK_NOTIFICATION_CODE = iota
	wlan_hosted_network_peer_state_change
	wlan_hosted_network_radio_state_change
)

//The WLAN_HOSTED_NETWORK_OPCODE enumerated type specifies the possible values of the operation code for the properties to query or set on the wireless Hosted Network.
//https://docs.microsoft.com/en-us/windows/win32/api/wlanapi/ne-wlanapi-wlan_hosted_network_opcode
type WLAN_HOSTED_NETWORK_OPCODE uint32

const (
	wlan_hosted_network_opcode_connection_settings WLAN_HOSTED_NETWORK_OPCODE = iota
	wlan_hosted_network_opcode_security_settings
	wlan_hosted_network_opcode_station_profile
	wlan_hosted_network_opcode_enable
)

//The WLAN_HOSTED_NETWORK_PEER_AUTH_STATE enumerated type specifies the possible values for the authentication state of a peer on the wireless Hosted Network.
//https://docs.microsoft.com/en-us/windows/win32/api/wlanapi/ne-wlanapi-wlan_hosted_network_peer_auth_state
type WLAN_HOSTED_NETWORK_PEER_AUTH_STATE uint32

const (
	wlan_hosted_network_peer_state_invalid WLAN_HOSTED_NETWORK_PEER_AUTH_STATE = iota
	wlan_hosted_network_peer_state_authenticated
)

//The WLAN_HOSTED_NETWORK_REASON enumerated type specifies the possible values for the result of a wireless Hosted Network function call.
//https://docs.microsoft.com/en-us/windows/win32/api/wlanapi/ne-wlanapi-wlan_hosted_network_reason
type WLAN_HOSTED_NETWORK_REASON uint32

const (
	wlan_hosted_network_reason_success WLAN_HOSTED_NETWORK_REASON = iota
	wlan_hosted_network_reason_unspecified
	wlan_hosted_network_reason_bad_parameters
	wlan_hosted_network_reason_service_shutting_down
	wlan_hosted_network_reason_insufficient_resources
	wlan_hosted_network_reason_elevation_required
	wlan_hosted_network_reason_read_only
	wlan_hosted_network_reason_persistence_failed
	wlan_hosted_network_reason_crypt_error
	wlan_hosted_network_reason_impersonation
	wlan_hosted_network_reason_stop_before_start
	wlan_hosted_network_reason_interface_available
	wlan_hosted_network_reason_interface_unavailable
	wlan_hosted_network_reason_miniport_stopped
	wlan_hosted_network_reason_miniport_started
	wlan_hosted_network_reason_incompatible_connection_started
	wlan_hosted_network_reason_incompatible_connection_stopped
	wlan_hosted_network_reason_user_action
	wlan_hosted_network_reason_client_abort
	wlan_hosted_network_reason_ap_start_failed
	wlan_hosted_network_reason_peer_arrived
	wlan_hosted_network_reason_peer_departed
	wlan_hosted_network_reason_peer_timeout
	wlan_hosted_network_reason_gp_denied
	wlan_hosted_network_reason_service_unavailable
	wlan_hosted_network_reason_device_change
	wlan_hosted_network_reason_properties_change
	wlan_hosted_network_reason_virtual_station_blocking_use
	wlan_hosted_network_reason_service_available_on_virtual_station
)

//The WLAN_HOSTED_NETWORK_STATE enumerated type specifies the possible values for the network state of the wireless Hosted Network.
//https://docs.microsoft.com/en-us/windows/win32/api/wlanapi/ne-wlanapi-wlan_hosted_network_state
type WLAN_HOSTED_NETWORK_STATE uint32

const (
	wlan_hosted_network_unavailable WLAN_HOSTED_NETWORK_STATE = iota
	wlan_hosted_network_idle
	wlan_hosted_network_active
)

//The WLAN_INTERFACE_TYPE enumeration specifies the wireless interface type.
//https://docs.microsoft.com/en-us/windows/win32/api/wlanapi/ne-wlanapi-wlan_interface_type
type WLAN_INTERFACE_TYPE uint32

const (
	wlan_interface_type_emulated_802_11 WLAN_INTERFACE_TYPE = iota
	wlan_interface_type_native_802_11
	wlan_interface_type_invalid
)

//The WLAN_SECURABLE_OBJECT enumerated type defines the securable objects used by Native Wifi Functions.
//https://docs.microsoft.com/en-us/windows/win32/api/wlanapi/ne-wlanapi-wlan_securable_object
type WLAN_SECURABLE_OBJECT uint32

const (
	wlan_secure_permit_list WLAN_SECURABLE_OBJECT = iota
	wlan_secure_deny_list
	wlan_secure_ac_enabled
	wlan_secure_bc_scan_enabled
	wlan_secure_bss_type
	wlan_secure_show_denied
	wlan_secure_interface_properties
	wlan_secure_ihv_control
	wlan_secure_all_user_profiles_order
	wlan_secure_add_new_all_user_profiles
	wlan_secure_add_new_per_user_profiles
	wlan_secure_media_streaming_mode_enabled
	wlan_secure_current_operation_mode
	wlan_secure_get_plaintext_key
	wlan_secure_hosted_network_elevated_access
	wlan_secure_virtual_station_extensibility
	wlan_secure_wfd_elevated_access
	WLAN_SECURABLE_OBJECT_COUNT
)

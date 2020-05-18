



From: "https://community.hubitat.com/t/zigbee-thermostat-mode-reporting/20492/3"


// send configure commad to the thermostat
def cmds = [
        //Set long poll interval to 2 qs
        "raw 0x0020 {11 00 02 02 00 00 00}",
        "send 0x${device.deviceNetworkId} 1 1",
        //Thermostat - Cluster 201
        "zdo bind 0x${device.deviceNetworkId} 1 1 0x201 {${device.zigbeeId}} {}", "delay 500",
        "zcl global send-me-a-report 0x201 0 0x29 5 300 {3200}", "delay 500",      // report temperature changes over 0.5°C (0x3200 in little endian)
        "send 0x${device.deviceNetworkId} 1 1", "delay 500",
        "zcl global send-me-a-report 0x201 0x0011 0x29 5 300 {3200}", "delay 500", // report cooling setpoint delta: 0.5°C
        "send 0x${device.deviceNetworkId} 1 1", "delay 500",
        "zcl global send-me-a-report 0x201 0x0012 0x29 5 300 {3200}", "delay 500", // report heating setpoint delta: 0.5°C
        "send 0x${device.deviceNetworkId} 1 1", "delay 500",
        //"zcl global send-me-a-report 0x201 0x001C 0x30 5 300 {}", "delay 500",     // report system mode
        //"send 0x${device.deviceNetworkId} 1 1", "delay 500",
        "zcl global send-me-a-report 0x201 0x0029 0x19 5 300 {}", "delay 500",     // report running state
        "send 0x${device.deviceNetworkId} 1 1", "delay 500",
        //Fan Control - Cluster 202
        "zdo bind 0x${device.deviceNetworkId} 1 1 0x202 {${device.zigbeeId}} {}", "delay 500",
        "zcl global send-me-a-report 0x202 0 0x30 5 300 {}", "delay 500",          // report fan mode
        "send 0x${device.deviceNetworkId} 1 1",
    ]
// System Mode
//cmds += zigbee.configureReporting(THERMOSTAT_CLUSTER, ATTRIBUTE_SYSTEM_MODE, TypeENUM8, 5, 300)
cmds += zigbee.configureReporting(0x0201, 0x001C, 0x30, 0, 0xFFFE, null, [:], 500)
//Power Control - Cluster 0001 (report battery status)
cmds += zigbee.batteryConfig()
//sendZigbeeCmds(cmds)
// Delay polling device attribute until the config is done
runIn(15, "pollDevice", [overwrite: true])

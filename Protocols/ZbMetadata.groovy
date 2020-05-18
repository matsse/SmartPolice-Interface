metadata {
    // Automatically generated. Make future change here.
    definition (name: "SmartPower Outlet", namespace: "smartthings", author: "SmartThings") {
        capability "Actuator"
        capability "Switch"
        capability "Power Meter"
        capability "Configuration"
        capability "Refresh"
        capability "Sensor"

    // indicates that device keeps track of heartbeat (in state.heartbeat)
    attribute "heartbeat", "string"

    fingerprint profileId: "0104", inClusters: "0000,0003,0004,0005,0006,0B04,0B05", outClusters: "0019", manufacturer: "CentraLite",  model: "3200", deviceJoinName: "Outlet"
    fingerprint profileId: "0104", inClusters: "0000,0003,0004,0005,0006,0B04,0B05", outClusters: "0019", manufacturer: "CentraLite",  model: "3200-Sgb", deviceJoinName: "Outlet"
    fingerprint profileId: "0104", inClusters: "0000,0003,0004,0005,0006,0B04,0B05", outClusters: "0019", manufacturer: "CentraLite",  model: "4257050-RZHAC", deviceJoinName: "Outlet"
    fingerprint profileId: "0104", inClusters: "0000,0003,0004,0005,0006,0B04,0B05", outClusters: "0019"
}
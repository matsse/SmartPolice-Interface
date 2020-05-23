package Utils



const (
	////////// Main screens //////////
	MainHome string = `
 Welcome to the smart policing interface for data aqcuisition from IoT devices.
 This program lets you define custom formats to extract useful information from smart
 devices and control what gets prioritized.

	Disclaimer: this is a Work in progress, so some issues might arise from using its features!

	Please select one of the options:

	device)    Enters the device management interface
	format)    Enters the format management interface
	listen)    Listen for the data that is gathered
	options)   Adjust settings
	exit)      Exit the program`
	
	
	
	MainHelp string = `
 Please select one of the options:

	device)    Enters the device management interface
	format)    Enters the format management interface
	listen)    Listen for the data that is gathered
	options)   Adjust settings
	exit)      Exit the program`
	
	
	
	////////// Device screens //////////
	
	DeviceHome string = `
 This interface lets you manage your devices. All devices are stored in json format for future use.

	Please select one of the options:
	new)       Create a new device - A prompt will ask for more info.
	edit)      Edits an existing device - Example: edit device1
	delete)    Deletes a device - Example: delete device1
	dat)       dump current data from all devices
	list)      List all devices
	exit)      goes back to precious prompt`
	
	
	////////// Format screens //////////
	
	
	TargetHelp string = `
 Specify the target of your data. It can be done with a key or a path
	
	## PATH ##
		To specify a path, type the key value of the the path which leads to the right key.
		Denote each corresponding joints in the path with a period (.)
		
		Example:   RootValue.Child1.Child2.Child3

	## KEY ##
		To specify the target with a key will be reduce efficiency, as the program will recursively look for the value.
		However, after the value is found the first time, a path will be created to reduce time/complexity.`
	
	
	TypeHelp string = `
 Specify the type of your targetted data
 To specify the data type of your data entry, you can use the following options:
	
	## Primitives ##
		int = type:int
		uint16 = type:uint16
		float32 = type:float32
		float64 = type:float64
		string = type:string
		date = type:date
		interface = type:interface

	## Arrays ##
		int = type:[]int
		uint = type:[]uint16
		float32 = type:[]float32
		float64 = type:[]float64
		string = type:[]string
		interface = type:[]interface`
	
	
	MainPS string = "\n\n Home#  "
	TypePS string = "\n\n Type#   "
	TargetPS string = "\n\n Target#   "
	DevicePS string = "\n\n Device#   "


)





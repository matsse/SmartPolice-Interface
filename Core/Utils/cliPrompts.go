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
	
	
	
	FormatName string = `
 What will the name of your format be?`
	
	
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
	
	
	ActionsScreen string = `
 Actions are functions that transforms data and manipulates if to the user's liking.
 Do you wish to add actions to this format (yes/no)?`
	
	//ActionsMain string = `
 //You can either add an action or a chain.
 //
	//Actions are signular funtion calls that requires the user to only specify the name of the function and args
	//	Type: action:{functionName(type1#arg1, ..., typen#argn)}
 //
	//Chains are targets that runs consecutively after one an other. So the order of the chain is important.
	//	Type: chain.{function1(type1#arg1, ..., typen#argn)}.{function2(type1#arg1, ..., typen#argn)}
 //
 //
 //help:
	//Type "#help <keyword>" to print a help screen.
	//
	//list)
	//Args)
	//Chains)
	//Actions)`
	//
	ActionsMain string = `
 You can either add an action or a chain.
 
	Actions are signular funtion calls that requires the user to only specify the name of the function and args
		Type: action:{functionName(type1#arg1, ..., typen#argn)}
  
	Chains are targets that runs consecutively after one an other. So the order of the chain is important.
		Type: chain.{function1(type1#arg1, ..., typen#argn)}.{function2(type1#arg1, ..., typen#argn)}
`
	
	
	
	
	
	ActionsHelp string = `
About:
	Actions are function calls that can be made on format targets once they have been recieved by this proram.
	Problems relating to other programs requiring different inputs or sensors recording data in different measuring
	sytstems (i.e. pounds/kilos, Fahrenheit/Celcius, etc)
	
Syntax:
	keyword:
		The keyword "action" is used to specify that the entry is an action.
    Characters:
		Actions are are formatted with special characters.
			action:functionName(type1#arg1, ..., )

		- colon ':' separates the keyword "action" from the rest
		- opening parenthesis '('  and closing parenthesis ')' are used to enclose arguments
		- number-sign '#' are used to separate argument types from argument values`
	
	
	
	ChainsHelp string = `
 About:
	Chains are a sequence of actions that runs in consequtive order.
	They are especially useful for when multiple actions are neededto be performed on
    data in an order.
	

Ordering:
	The way chains are typed is important. The program will read the chain from left to
	right and execute in the order that it reads values. For example the two cases are different:
	
	case1)
		chain:sumx(arg1, arg2).DecryptAES(arg1, ..., argn)

	case2)
		chain:DecryptAES(arg1, ..., argn).sumx(arg1, arg2)

	Both chains have the same functions, but the ordering of these are different. The
	first case will execute the addition, before decrypting. While the second chain
	will decrypt first and then sum two numbers
	

Syntax:
	keyword:
		The keyword "chain" is used to specify that the entry is an action.
   
    Characters:
		chains
			chain:function1(arg1, arg2).function2(arg1, ..., argn)

		- colon ':' separates the keyword "action" from the rest
		- opening parenthesis '('  and closing parenthesis ')' are used to enclose arguments
		- number-sign '#' are used to separate argument types from argument values
		- periods '.' are used to separate actions from each other`
		
	
	
	
	
	ActTypessHelp string = `
 Arguments are the input that a function takes when a chain or an action is to be executed.
 There are TWO components to an argument that must be satisfied, in order for the action or chain to compile.
 
 Type declaration is the first part of the argument and specifies what type the argument is.
	 Types includes:
	    sfloat64    - Float value
	    sint        - Integer value
	    sstring     - String value
		self        - self reference
		sref        - Reference to another target
		

Argument value can either be referenced by name or by value
    Examples of values includes:
		sing#10             - integer
		sfloat64#13.1       - float
		something           - string
		sref#Temperature    - Reference to Temperature target
		self#               - Self referral, (type is inferred from  collected value)`
	
	AddMore string = `
 Do you wish to add additional format targets (yes/no)?`
	
	MainPS string = "\n\n Home#  "
	TypePS string = "\n\n Type#   "
	TargetPS string = "\n\n Target#   "
	DevicePS string = "\n\n Device#   "
	FmtNamePS string = "\n\n NAME#   "
	AddActionPS string = "\n\n NewAction#   "
	AdditonalT string = "\n\n MoreTargets#   "
	ActionsPS string = "\n\n Action#   "
)





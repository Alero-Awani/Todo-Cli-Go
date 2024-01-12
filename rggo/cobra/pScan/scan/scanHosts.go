package scan

//PortState represents the state of a single TCP port


type PortState struct {
	Port int 
	Open state
}

type state bool 

//String converts the boolean value of state to a human readable string
func(s state) String() string {
	if s {
		return "open"
	}
	return "closed"
}

//scanPort performs a port scan on a single TCP port 
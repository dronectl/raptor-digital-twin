
package device


type Device struct {
    Name string // user configured name
    UUID string // serial id
    channel chan string // channel to send messages
}

func (d *Device) Start() {
    // start keep alive message stream (1 Hz)
    // fetch conditions data
}

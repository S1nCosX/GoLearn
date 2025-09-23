package reader

import "errors"

type PhoneNumber struct {
	str []rune
}

func (this *PhoneNumber) Read(buffer []rune) (int, error) {
	// lets read only phone numbers starts from num or '+' and with 0-9 num in country code
	iter := 0
	if this.str == nil {
		this.str = make([]rune, 0)
	}
	if buffer[0] == '+' {
		this.str = append(this.str, '+')
		iter++
	} else {
		if !(buffer[0] >= '0' && buffer[0] <= '9') {
			return 0, errors.New("Phone number format error: Got a non + or number as first char")
		}
	}
	for ; iter < 11; iter++ { // lets read num only format (for real this is just learning task this thing must be done by regexp, not hands)
		if !(buffer[iter] >= '0' && buffer[iter] <= '9') {
			return 0, errors.New("Phone number format error: Got a non + or number as first char")
		}
		this.str = append(this.str, buffer[iter])
	}
	return iter, nil
}

func Reader() {
	// lets talk about reading. Most read funcs in goalang uses interface Reader.
	// this interface means, that our class must have func: Read(buffer []rune) (int, error)
	// basicaly this means: we get buffer, where we trying to read our type, and return 2 values: int - amount of runes we readed, error - error
	// there was realy nice demonstration for this: Lets write phone number reader
	var phone PhoneNumber
	amount, err := phone.Read([]rune("88005553535"))
	println(amount, err, string(phone.str))
	amount, err = phone.Read([]rune("+78005553535"))
	println(amount, err, string(phone.str))
	amount, err = phone.Read([]rune("a88005553535"))
	println(amount, err, string(phone.str))
}

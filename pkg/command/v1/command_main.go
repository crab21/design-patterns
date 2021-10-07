package v1

func Mainest() {
	tv := &Tv{}
	on := &OnCommand{
		Device: tv,
	}
	bt := &button{
		Command: on,
	}
	bt.Press()

	off := &OffCommand{
		Device: tv,
	}
	btOff := &button{
		Command: off,
	}
	btOff.Press()
}

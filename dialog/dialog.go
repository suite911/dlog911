package dialog

import (
	"github.com/amy911/amy911/onfail"
	"github.com/amy911/amy911/str"
)

var (
	OverrideLogIn func(*Dialog, []interface{})
)

type Dialog struct {
	Embed, UserData interface{}
}

func New(kind string, args ...interface{}) *Dialog {
	return new(Dialog).Init(kind, args...)
}

func (d *Dialog) Init(kind string, args ...interface{}) *Dialog {
	if Library == nil {
		onfail.Fail("Dialog Library was not loaded or was unloaded", d, onfail.Panic, args)
		return d
	}
	var aps []*string
	var api8 []*int8
	for _, arg := range args {
		switch arg.(type) {
		case *string:
			aps = append(aps, arg.(*string))
		case *int8:
			api8 = append(api8, arg.(*int8))
		}
	}
	switch simp := str.Simp(kind); simp {
	case "login":
		if OverrideLogIn != nil {
			OverrideLogIn(d, args)
		} else {
			w := Library().NewWindow("Log in") // TODO: translate
			f := w.NewFrame("Account") // TODO: translate
			if len(aps) >= 1 {
				label := "Login:" // TODO: translate
				if len(aps) > 3 {
					label = "E-mail address:" // TODO: translate
				}
				f.NewLabel(label)
				f.NewEntry(aps[0], false)
			}
			if len(aps) >= 3 {
				f.NewLabel("Username:") // TODO: translate
				f.NewEntry(aps[2], false)
			}
			if len(aps) >= 2 {
				f.NewLabel("Password:") // TODO: translate
				f.NewEntry(aps[1], true)
			}
			var out *int8
			if len(api8) >= 1 {
				out = api8[0]
			}
			w.NewButtonGroup(out, NewButtonGroup("login"))
			w.Show(576, 324)
		}
	default:
		onfail.Fail("Unknown Dialog kind \""+simp+"\"", d, onfail.Panic, args)
	}
	return d
}

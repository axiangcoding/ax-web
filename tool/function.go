package tool

import "github.com/axiangcoding/ax-web/logging"

// GoWithRecover handle go routines with panic recover.
// FIXME: use ants instead
func GoWithRecover(f func()) {
	go func(handler func()) {
		defer func() {
			if r := recover(); r != nil {
				logging.Errorf("recover from go func error. %s", r)
			}
		}()
		handler()
	}(f)
}

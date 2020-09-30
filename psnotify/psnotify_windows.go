// windows stub function

package psnotify

import (
	"errors"
)

// Initialize linux implementation of the eventListener interface
func createListener() (eventListener, error) {
	return nil, errors.New("Not support windows yet!")
}

func (w *Watcher) readEvents() {
}

// Delete filter for given pid from the queue
func (w *Watcher) unregister(pid int) error {
	return nil
}

// noop on linux
func (w *Watcher) register(pid int, flags uint32) error {
	return nil
}

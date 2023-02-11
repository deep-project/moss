package osx

import (
	"github.com/fsnotify/fsnotify"
)

func Watch(paths []string, fn func(fsnotify.Event)) (watcher *fsnotify.Watcher, err error) {
	watcher, err = fsnotify.NewWatcher()
	if err != nil {
		return
	}
	for _, p := range paths {
		if err = watcher.Add(p); err != nil {
			return
		}
	}
	go func() {
		for {
			select {
			case e, ok := <-watcher.Events:
				if ok {
					fn(e)
				}
			}
		}
	}()
	return
}

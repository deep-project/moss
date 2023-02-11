package engine

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/CloudyKit/jet/v6"
)

// Modified from https://github.com/gofiber/template/tree/master/jet

type Jet struct {
	dir       string                 // views folder
	layout    string                 // layout variable name that incapsulates the template
	loaded    bool                   // determines if the engine parsed all templates
	reload    bool                   // reload on each render
	mutex     sync.RWMutex           // lock for funcMap and templates
	funcMap   map[string]interface{} // template funcMap
	skipLoad  func(path string) bool // skip add loader
	Templates *jet.Set               // templates
}

func NewJet(dir string, skipLoad func(path string) bool) *Jet {
	return &Jet{
		dir:       dir,
		layout:    "embed",
		funcMap:   make(map[string]interface{}),
		skipLoad:  skipLoad,
		Templates: new(jet.Set),
	}
}

func (e *Jet) Layout(key string) *Jet {
	e.layout = key
	return e
}

// AddFunc adds the function to the template's function map.
// It is legal to overwrite elements of the default actions
func (e *Jet) AddFunc(name string, fn interface{}) {
	e.mutex.Lock()
	e.funcMap[name] = fn
	e.mutex.Unlock()
	return
}

// Reload if set to true the templates are reloading on each render,
// use it when you're in development and you don't want to restart
// the application when you edit a template file.
func (e *Jet) Reload() {
	e.reload = true
	return
}

// Load parses the templates to the engine.
func (e *Jet) Load() (err error) {
	// race safe
	e.mutex.Lock()
	defer e.mutex.Unlock()

	var loader = jet.NewInMemLoader()
	e.Templates = jet.NewSet(loader)

	for name, fn := range e.funcMap {
		e.Templates.AddGlobal(name, fn)
	}
	walkFn := func(path string, info os.FileInfo, err error) error {
		// Return error if exist
		if err != nil {
			return err
		}
		// Skip file if it's a directory or has no file info
		if info == nil || info.IsDir() {
			return nil
		}

		// ./views/html/index.tmpl -> index.tmpl
		name, err := filepath.Rel(e.dir, path)
		if err != nil {
			return err
		}
		if e.skipLoad(name) == true {
			return nil
		}
		buf, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		loader.Set(name, string(buf))
		return err
	}
	e.loaded = true
	return filepath.Walk(e.dir, walkFn)
}

func (e *Jet) Render(out io.Writer, template string, binds map[string]any, layout ...string) error {
	if !e.loaded || e.reload {
		e.reload = false
		if err := e.Load(); err != nil {
			return err
		}
	}
	tmpl, err := e.Templates.GetTemplate(template)
	if err != nil || tmpl == nil {
		if tmpl == nil {
			return fmt.Errorf("template not found: template %s", template)
		}
		return fmt.Errorf("render: template %s could not be loaded: %v", template, err)
	}

	var bind = make(jet.VarMap)
	for key, value := range binds {
		bind.Set(key, value)
	}
	if len(layout) > 0 && layout[0] != "" {
		lay, err := e.Templates.GetTemplate(layout[0])
		if err != nil {
			return err
		}
		bind.Set(e.layout, func() {
			_ = tmpl.Execute(out, bind, nil)
		})
		return lay.Execute(out, bind, nil)
	}
	return tmpl.Execute(out, bind, nil)
}

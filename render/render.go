package render

import (
	"context"
	"html/template"
	"io"
	"sync"

	"github.com/ONSdigital/log.go/log"
	"github.com/rav-pradhan/test-modules/render/models"
	"github.com/unrolled/render"
	unrolled "github.com/unrolled/render"
)

type Render struct {
	client *unrolled.Render
	hMutex *sync.Mutex
	jMutex *sync.Mutex
}

// TODO: need to find a way of passing in assets from application
// when this is instantiated so that common assets can be built
// in the render package and custom assets can still be provided
func New(assetsPath, siteDomain string, assetFn func(name string) ([]byte, error), assetNameFn func() []string) *Render {
	return &Render{
		client: unrolled.New(render.Options{
			Asset:      assetFn,
			AssetNames: assetNameFn,
			Layout:     "main",
			Funcs:      []template.FuncMap{registeredFuncs},
		}),
		hMutex: &sync.Mutex{},
		jMutex: &sync.Mutex{},
	}
}

//Page resolves the rendering of a specific page with a given model and template name
func (r *Render) Page(w io.Writer, page interface{}, templateName string) {
	ctx := context.Background()
	if err := r.HTML(w, 200, templateName, page); err != nil {
		r.JSON(w, 500, models.ErrorResponse{
			Error: err.Error(),
		})
		log.Event(ctx, "failed to render template", log.Error(err), log.ERROR)
		return
	}
	log.Event(ctx, "rendered template", log.Data{"template": templateName}, log.INFO)
}

// HTML controls the rendering of an HTML template with a given name and template parameters to an io.Writer
func (r *Render) HTML(w io.Writer, status int, name string, binding interface{}, htmlOpt ...render.HTMLOptions) error {
	r.hMutex.Lock()
	defer r.hMutex.Unlock()
	return r.client.HTML(w, status, name, binding, htmlOpt...)
}

// JSON controls the rendering of a JSON template with a given name and template parameters to an io.Writer
func (r *Render) JSON(w io.Writer, status int, v interface{}) error {
	r.jMutex.Lock()
	defer r.jMutex.Unlock()
	return r.client.JSON(w, status, v)
}

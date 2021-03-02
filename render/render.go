package render

import (
	"context"
	"html/template"
	"io"
	"sync"

	"github.com/ONSdigital/dp-frontend-models/model"
	"github.com/ONSdigital/log.go/log"
	"github.com/unrolled/render"
	unrolled "github.com/unrolled/render"
)

type Render struct {
	client *unrolled.Render
	hMutex *sync.Mutex
	jMutex *sync.Mutex
}

func New(assetsPath, siteDomain string) *Render {
	return &Render{
		client: unrolled.New(render.Options{
			Layout: "main",
			Funcs:  []template.FuncMap{registeredFuncs},
		}),
		hMutex: &sync.Mutex{},
		jMutex: &sync.Mutex{},
	}
}

//Handler resolves the rendering of a specific pagem with a given model and template name
func (r *Render) Page(w io.Writer, page interface{}, templateName string) {
	ctx := context.Background()

	// // assert that the struct contains a model.Page common type/fields
	// if _, ok := page.(model.Page); !ok {
	// 	log.Event(ctx, "invalid template type - core fields not present", nil, log.ERROR)
	// 	return
	// }

	if err := r.HTML(w, 200, templateName, page); err != nil {
		r.JSON(w, 500, model.ErrorResponse{
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

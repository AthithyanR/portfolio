package routes

import (
	"html/template"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

const (
	templatesDir = "internal/templates"

	contentType     = "content-type"
	textHtml        = "text-html"
	applicationJson = "application/json"

	cacheControl = "cache-control"
	noCache      = "no-cache"
)

func root(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set(contentType, textHtml)
	ctx.Response.Header.Set(cacheControl, noCache)

	tmpl := template.Must(template.ParseFiles(templatesDir + "/index.html"))

	tmpl.Execute(ctx, nil)
}

func NewRouter() *router.Router {
	r := router.New()

	r.GET("/", root)

	fsHandler := &fasthttp.FS{
		Root: "static",
		PathRewrite: func(ctx *fasthttp.RequestCtx) []byte {
			return fasthttp.NewPathSlashesStripper(1)(ctx)
		},
		Compress: true,
	}

	r.GET("/static/{filepath:*}", fsHandler.NewRequestHandler())

	return r
}

package routes

import (
	"html/template"
	"log"
	"time"

	"github.com/AthithyanR/portfolio/internal/database"
	"github.com/fasthttp/router"
	realip "github.com/ferluci/fast-realip"
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

	tmpl.Execute(ctx, map[string]interface{}{
		"techIKnow":    GetTechIKnow(),
		"catFact":      GetCatFact(),
		"chatMessages": GetChatMessages(),
	})
}

func addNewMessage(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set(contentType, applicationJson)
	ctx.Response.Header.Set(cacheControl, noCache)

	name := string(ctx.FormValue("name"))
	message := string(ctx.FormValue("message"))

	if name == "" || message == "" {
		log.Println("Invalid message body")
		ctx.Response.SetStatusCode(500)
		return
	}

	chatMessage := database.ChatMessage{
		Ip:      realip.FromRequest(ctx),
		Name:    name,
		Message: message,
	}

	err := CreateChatMessage(&chatMessage)
	if err != nil {
		log.Println("Unable to create chat message", err)
		ctx.Response.SetStatusCode(500)
		return
	}

	tmpl := template.Must(template.ParseFiles(templatesDir + "/index.html"))

	tmpl.ExecuteTemplate(ctx, "chatMessages", map[string]interface{}{
		"chatMessages": []database.ChatMessage{chatMessage},
	})
}

func NewRouter() *router.Router {
	r := router.New()

	r.GET("/", root)
	r.POST("/new-message", addNewMessage)

	fsHandler := &fasthttp.FS{
		Root: "static",
		PathRewrite: func(ctx *fasthttp.RequestCtx) []byte {
			return fasthttp.NewPathSlashesStripper(1)(ctx)
		},
		Compress:      true,
		CacheDuration: time.Second * 1,
	}

	r.GET("/static/{filepath:*}", fsHandler.NewRequestHandler())

	return r
}

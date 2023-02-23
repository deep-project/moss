package router

import (
	"crypto/tls"
	"errors"
	"go.uber.org/zap"
	"moss/api/web/middleware"
	"moss/domain/config"
	"moss/infrastructure/general/conf"
	"moss/infrastructure/support/log"
	"net"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Router struct {
	app *fiber.App
}

func New() *Router {
	r := &Router{}
	r.app = r.newFiber()
	return r
}

func (r *Router) newFiber() *fiber.App {
	app := fiber.New(config.Config.Router.GetOptions())

	// pprof
	var pprofPrefix = ""
	if config.Config.Router.PprofSecret != "" {
		pprofPrefix = "/" + config.Config.Router.PprofSecret
	}
	app.Use(pprof.New(pprof.Config{Prefix: pprofPrefix}))

	// 捕捉堆栈错误
	app.Use(recover.New())
	app.Use(middleware.CatchPanicError)
	// http log
	app.Use(middleware.HttpLog)
	// ETag
	if config.Config.Router.ETag {
		app.Use(etag.New())
	}
	// 压缩
	app.Use(compress.New(compress.Config{Level: compress.Level(config.Config.Router.CompressLevel)}))

	// TLS
	app.Use(middleware.TLS)

	// admin
	app.Route(config.Config.Router.GetAdminPath(), r.RegisterAdmin)

	// home
	app.Route("/", r.RegisterHome)

	return app
}

func (r *Router) Run() error {
	log.Info("app starting...")
	go func() {
		if config.Config.TLS.Enable {
			err := r.listenerTLS()
			if err != nil {
				log.Error("tls listen error", zap.Error(err))
			}
		}
	}()
	return r.app.Listen(conf.Addr)
}

func (r *Router) Reload() {
	r.app.Server().Handler = r.newFiber().Handler()
}

func (r *Router) listenerTLS() error {
	ln, err := r.ln()
	if err != nil {
		return err
	}
	return r.app.Listener(ln)
}

func (r *Router) ln() (ln net.Listener, err error) {
	if config.Config.TLS.CertPEM == "" || config.Config.TLS.KeyPEM == "" {
		return ln, errors.New("tls Cert or KEY is undefined")
	}
	cert, err := tls.X509KeyPair([]byte(config.Config.TLS.CertPEM), []byte(config.Config.TLS.KeyPEM))
	if err != nil {
		return
	}
	tlsHandler := &fiber.TLSHandler{}
	c := &tls.Config{
		MinVersion:     tls.VersionTLS12,
		Certificates:   []tls.Certificate{cert},
		GetCertificate: tlsHandler.GetClientInfo,
	}
	netWork := config.Config.Router.Options.Network
	if netWork == "" {
		netWork = "tcp"
	}
	ln, err = net.Listen(netWork, config.Config.TLS.ListenAddr())
	ln = tls.NewListener(ln, c)
	r.app.SetTLSHandler(tlsHandler)
	return
}

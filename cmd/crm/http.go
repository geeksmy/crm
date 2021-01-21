package crm

import (
	"context"
	"net/http"
	"os"
	"sync"
	"time"

	"crm/gen/auth"
	"crm/gen/customer"
	"crm/gen/group"
	authsvr "crm/gen/http/auth/server"
	customersvr "crm/gen/http/customer/server"
	groupsvr "crm/gen/http/group/server"
	inventorysvr "crm/gen/http/inventory/server"
	procurementsvr "crm/gen/http/procurement/server"
	productsvr "crm/gen/http/product/server"
	salessvr "crm/gen/http/sales/server"
	suppliersvr "crm/gen/http/supplier/server"
	usersvr "crm/gen/http/user/server"
	warehousesvr "crm/gen/http/warehouse/server"
	"crm/gen/inventory"
	"crm/gen/log"
	"crm/gen/procurement"
	"crm/gen/product"
	"crm/gen/sales"
	"crm/gen/supplier"
	"crm/gen/user"
	"crm/gen/warehouse"

	"github.com/geeksmy/go-libs/goa-libs/middleware/metrics"
	"go.uber.org/zap"
	goahttp "goa.design/goa/v3/http"
	httpmdlwr "goa.design/goa/v3/http/middleware"
	"goa.design/goa/v3/middleware"

	mdlwr "crm/pkg/middleware"
)

// handleHTTPServer starts configures and starts a HTTP server on the given
// URL. It shuts down the server if any error is received in the error channel.
func handleHTTPServer(ctx context.Context, host string, userEndpoints *user.Endpoints, authEndpoints *auth.Endpoints,
	groupEndpoints *group.Endpoints, productEndpoints *product.Endpoints, supplierEndpoints *supplier.Endpoints,
	procurementEndpoints *procurement.Endpoints, salesEndpoints *sales.Endpoints, customerEndpoints *customer.Endpoints,
	warehouseEndpoints *warehouse.Endpoints, inventoryEndpoints *inventory.Endpoints,
	wg *sync.WaitGroup, errc chan error, logger *log.Logger, metrics *metrics.Prometheus, debug bool) {

	// Setup goa log adapter.
	var (
		adapter middleware.Logger
	)
	{
		adapter = logger
	}

	// Provide the transport specific request decoder and response encoder.
	// The goa http package has built-in support for JSON, XML and gob.
	// Other encodings can be used by providing the corresponding functions,
	// see goa.design/implement/encoding.
	var (
		dec = goahttp.RequestDecoder
		enc = goahttp.ResponseEncoder
	)

	// Build the service HTTP request multiplexer and configure it to serve
	// HTTP requests to the service endpoints.
	var mux goahttp.Muxer
	{
		mux = goahttp.NewMuxer()
	}

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to HTTP requests and
	// responses.
	var (
		authServer        *authsvr.Server
		userServer        *usersvr.Server
		groupServer       *groupsvr.Server
		productServer     *productsvr.Server
		supplierServer    *suppliersvr.Server
		procurementServer *procurementsvr.Server
		salesServer       *salessvr.Server
		customerServer    *customersvr.Server
		warehouseServer   *warehousesvr.Server
		inventoryServer   *inventorysvr.Server
	)
	{
		eh := errorHandler(logger)
		userServer = usersvr.New(userEndpoints, mux, dec, enc, eh, nil)
		authServer = authsvr.New(authEndpoints, mux, dec, enc, eh, nil)
		groupServer = groupsvr.New(groupEndpoints, mux, dec, enc, eh, nil)
		productServer = productsvr.New(productEndpoints, mux, dec, enc, eh, nil)
		supplierServer = suppliersvr.New(supplierEndpoints, mux, dec, enc, eh, nil)
		procurementServer = procurementsvr.New(procurementEndpoints, mux, dec, enc, eh, nil)
		salesServer = salessvr.New(salesEndpoints, mux, dec, enc, eh, nil)
		customerServer = customersvr.New(customerEndpoints, mux, dec, enc, eh, nil)
		warehouseServer = warehousesvr.New(warehouseEndpoints, mux, dec, enc, eh, nil)
		inventoryServer = inventorysvr.New(inventoryEndpoints, mux, dec, enc, eh, nil)
	}
	// Configure the mux.
	usersvr.Mount(mux, userServer)
	authsvr.Mount(mux, authServer)
	groupsvr.Mount(mux, groupServer)
	productsvr.Mount(mux, productServer)
	suppliersvr.Mount(mux, supplierServer)
	procurementsvr.Mount(mux, procurementServer)
	salessvr.Mount(mux, salesServer)
	customersvr.Mount(mux, customerServer)
	warehousesvr.Mount(mux, warehouseServer)
	inventorysvr.Mount(mux, inventoryServer)

	// Wrap the multiplexer with additional middlewares. Middlewares mounted
	// here apply to all the service endpoints.
	var handler http.Handler = mux
	{
		if debug {
			handler = httpmdlwr.Debug(mux, os.Stdout)(handler)
		}

		handler = mdlwr.PopulateRequestContext()(handler)
		handler = httpmdlwr.RequestID()(handler)

		if metrics != nil {
			handler = metrics.HandlerFunc(adapter)(handler)
		} else {
			handler = httpmdlwr.Log(adapter)(handler)
		}
	}

	// Start HTTP server using default configuration, change the code to
	// configure the server as required by your service.
	srv := &http.Server{Addr: host, Handler: handler}

	for _, m := range userServer.Mounts {
		logger.Infof("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	for _, m := range authServer.Mounts {
		logger.Infof("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	for _, m := range groupServer.Mounts {
		logger.Infof("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	for _, m := range productServer.Mounts {
		logger.Infof("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	for _, m := range supplierServer.Mounts {
		logger.Infof("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	for _, m := range salesServer.Mounts {
		logger.Infof("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	for _, m := range customerServer.Mounts {
		logger.Infof("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	for _, m := range warehouseServer.Mounts {
		logger.Infof("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	for _, m := range inventoryServer.Mounts {
		logger.Infof("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	for _, m := range procurementServer.Mounts {
		logger.Infof("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}

	(*wg).Add(1)
	go func() {
		defer (*wg).Done()

		// Start HTTP server in a separate goroutine.
		go func() {
			logger.Infof("HTTP server listening on %q", host)
			errc <- srv.ListenAndServe()
		}()

		<-ctx.Done()
		logger.Infof("shutting down HTTP server at %q", host)

		// Shutdown gracefully with a 30s timeout.
		ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()

		_ = srv.Shutdown(ctx)
	}()
}

// errorHandler returns a function that writes and logs the given error.
// The function also writes and logs the error unique ID so that it's possible
// to correlate.
func errorHandler(logger *log.Logger) func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		id := ctx.Value(middleware.RequestIDKey).(string)
		_, _ = w.Write([]byte("[" + id + "] encoding: " + err.Error()))
		logger.With(zap.String("id", id)).Error(err.Error())
	}
}

package crm

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"crm/config"
	handler "crm/controller"
	"crm/gen/auth"
	"crm/gen/customer"
	"crm/gen/group"
	"crm/gen/inventory"
	"crm/gen/log"
	"crm/gen/procurement"
	"crm/gen/product"
	"crm/gen/sales"
	"crm/gen/supplier"
	"crm/gen/user"
	"crm/gen/warehouse"

	metricsMlwr "github.com/geeksmy/go-libs/goa-libs/middleware/metrics"
	"github.com/geeksmy/go-libs/panichandler"
	"go.uber.org/zap"
)

func RunServer(cfg *config.Config, metrics *metricsMlwr.Prometheus) {

	// Setup logger. Replace logger with your own log package of choice.
	var (
		logger *log.Logger
	)
	{
		logger = log.New("crm", !cfg.Debug)
	}

	// Initialize the services.
	var (
		authSvc        auth.Service
		userSvc        user.Service
		groupSvc       group.Service
		productSvc     product.Service
		supplierSvc    supplier.Service
		procurementSvc procurement.Service
		salesSvc       sales.Service
		customerSvc    customer.Service
		warehouseSvc   warehouse.Service
		inventorySvc   inventory.Service
	)
	{
		authSvc = handler.NewAuth(logger)
		userSvc = handler.NewUser(logger)
		groupSvc = handler.NewGroup(logger)
		productSvc = handler.NewProduct(logger)
		supplierSvc = handler.NewSupplier(logger)
		procurementSvc = handler.NewProcurement(logger)
		salesSvc = handler.NewSales(logger)
		customerSvc = handler.NewCustomer(logger)
		warehouseSvc = handler.NewWarehouse(logger)
		inventorySvc = handler.NewInventory(logger)
	}

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	var (
		authEndpoints        *auth.Endpoints
		userEndpoints        *user.Endpoints
		groupEndpoints       *group.Endpoints
		productEndpoints     *product.Endpoints
		supplierEndpoints    *supplier.Endpoints
		procurementEndpoints *procurement.Endpoints
		salesEndpoints       *sales.Endpoints
		customerEndpoints    *customer.Endpoints
		warehouseEndpoints   *warehouse.Endpoints
		inventoryEndpoints   *inventory.Endpoints
	)
	{
		authEndpoints = auth.NewEndpoints(authSvc)
		userEndpoints = user.NewEndpoints(userSvc)
		groupEndpoints = group.NewEndpoints(groupSvc)
		productEndpoints = product.NewEndpoints(productSvc)
		supplierEndpoints = supplier.NewEndpoints(supplierSvc)
		procurementEndpoints = procurement.NewEndpoints(procurementSvc)
		salesEndpoints = sales.NewEndpoints(salesSvc)
		customerEndpoints = customer.NewEndpoints(customerSvc)
		warehouseEndpoints = warehouse.NewEndpoints(warehouseSvc)
		inventoryEndpoints = inventory.NewEndpoints(inventorySvc)
	}

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the services to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	addr := fmt.Sprintf("http://%s:%s", cfg.Server.Host, cfg.Server.HTTPPort)
	u, _ := url.Parse(addr)
	handleHTTPServer(ctx, u.Host, userEndpoints, authEndpoints, groupEndpoints, productEndpoints, supplierEndpoints,
		procurementEndpoints, salesEndpoints, customerEndpoints, warehouseEndpoints, inventoryEndpoints,
		&wg, errc, logger, metrics, cfg.Debug)

	// Wait for signal.
	logger.Infof("exiting (%v)", <-errc)

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	logger.Info("exited")
}

func RunDebugPprofServer(addr string) {
	defer panichandler.ZapHandler(zap.L()).Handle()
	zap.L().Sugar().Infof("启动 pprof 监听 %s.", addr)
	err := http.ListenAndServe(addr, nil)

	if err != nil {
		zap.L().Error("开启 pprof 监听失败 %s", zap.Error(err))
	}
}

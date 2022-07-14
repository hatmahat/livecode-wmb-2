package manager

import (
	"fmt"
	"go-api-with-gin2/config"
	"go-api-with-gin2/model"
	"go-api-with-gin2/service"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// infra -> sebagai database penyimpanan pengganti slice
type Infra interface {
	SqlDb() *gorm.DB
	LopeiClientConn() service.LopeiPaymentClient
}

type infra struct {
	db          *gorm.DB
	lopeiClient service.LopeiPaymentClient
	cfg         config.Config
}

func (i *infra) SqlDb() *gorm.DB {
	return i.db
}

func (i *infra) LopeiClientConn() service.LopeiPaymentClient {
	return i.lopeiClient
}

func NewInfra(config *config.Config) Infra {
	resource, err := initDbResource(config.DataSourceName)
	if err != nil {
		log.Fatal(err.Error())
	}
	infra := infra{
		db:  resource,
		cfg: *config,
	}
	infra.initGrpcClient()
	return &infra
}

func initDbResource(dataSourceName string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	db.AutoMigrate(
		&model.Menu{},
		&model.MenuPrice{},
		&model.BillDetail{},
		&model.Bill{},
		&model.Table{},
		&model.TransType{},
		&model.Customer{},
		&model.Discount{},
	)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (i *infra) initGrpcClient() {
	dial, err := grpc.Dial(i.cfg.GrpcConfig.Url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("did not connect...", err)
	}

	client := service.NewLopeiPaymentClient(dial)
	i.lopeiClient = client
	fmt.Println("GRPC client connected...")
}

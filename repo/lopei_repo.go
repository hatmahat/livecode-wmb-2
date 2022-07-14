package repo

import (
	"context"
	"encoding/json"
	"fmt"
	"go-api-with-gin2/service"
	"log"
)

type LopeiRepo interface {
	CheckBalance(lopeiId int32) (float32, error)
	DoPayment(lopeiId int32, amount float32) error
}

type lopeiRepo struct {
	client service.LopeiPaymentClient
}

func (c *lopeiRepo) CheckBalance(lopeiId int32) (float32, error) {
	result, err := c.client.CheckBalance(context.Background(), &service.CheckBalanceMessage{LopeiId: lopeiId})
	if err != nil {
		log.Fatalln("Error when calling check balance...", err)
		return 0, err
	}
	fmt.Println(result)
	var bal struct {
		LopeiId int
		Balance float32
	}
	err = json.Unmarshal([]byte(result.Result), &bal)
	if err != nil {
		log.Fatalln("Error when unmarshalling result...", err)
		return 0, err
	}
	return bal.Balance, err
}

func (c *lopeiRepo) DoPayment(lopeiId int32, amount float32) error {
	payment, err := c.client.DoPayment(context.Background(), &service.PaymentMessage{
		LopeiId: lopeiId,
		Amount:  amount,
	})
	if err != nil {
		log.Fatalln("Error when calling do payment...", err)
	}
	fmt.Println(payment)
	return nil
}

func NewLopeiRepo(client service.LopeiPaymentClient) LopeiRepo {
	repo := new(lopeiRepo)
	repo.client = client
	return repo
}

package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kardianos/service"
	"gorm.io/gorm"
	"ks-aggregator-rates/internal/pkg/client"
	"ks-aggregator-rates/internal/pkg/client/requests"
	"ks-aggregator-rates/internal/pkg/database"
	"math/big"
	"sync"
	"time"
)

func NewAndStartClient(req requests.ClientRequest) client.Client {
	return client.Client{
		req,
	}
}

type Program struct {
	service service.Service
}

func (p Program) Start(s service.Service) error {
	fmt.Println(s.String() + " started")
	p.service = s
	p.Run()
	return nil
}

func (p Program) Stop(s service.Service) error {
	fmt.Println(s.String() + " stopped")
	return nil
}

func setupClients() []client.Client {
	amountInETH := big.Int{}
	amountInETH.SetString("42000000000000000000", 10)
	amountInAPE := big.Int{}
	amountInAPE.SetString("11111000000000000000000", 10)
	amountInUSDC := big.Int{}
	amountInUSDC.SetString("100000000000", 10)
	amountInUSDT := big.Int{}
	amountInUSDT.SetString("100000000000", 10)

	/// KyberSwap
	ksRequestEthUsdc := requests.DefaultKyberSwapRequest(amountInETH)
	ksRequestEthUsdc.SetPairs(1)
	ksRequestUsdcUsdt := requests.DefaultKyberSwapRequest(amountInUSDC)
	ksRequestUsdcUsdt.SetPairs(2)
	ksRequestApeEth := requests.DefaultKyberSwapRequest(amountInAPE)
	ksRequestApeEth.SetPairs(3)

	ksRequestUsdcEth := requests.DefaultKyberSwapRequest(amountInUSDC)
	ksRequestUsdcEth.SetPairs(-1)
	ksRequestUsdtUsdc := requests.DefaultKyberSwapRequest(amountInUSDT)
	ksRequestUsdtUsdc.SetPairs(-2)
	ksRequestEthApe := requests.DefaultKyberSwapRequest(amountInETH)
	ksRequestEthApe.SetPairs(-3)

	/// OneInch
	oneInchRequestEthUsdc := requests.DefaultOneInchRequest(amountInETH)
	oneInchRequestEthUsdc.SetPairs(1)
	oneInchRequestUsdcUsdt := requests.DefaultOneInchRequest(amountInUSDC)
	oneInchRequestUsdcUsdt.SetPairs(2)
	oneInchRequestApeEth := requests.DefaultOneInchRequest(amountInAPE)
	oneInchRequestApeEth.SetPairs(3)

	oneInchRequestUsdcEth := requests.DefaultOneInchRequest(amountInUSDC)
	oneInchRequestUsdcEth.SetPairs(-1)
	oneInchRequestUsdtUsdc := requests.DefaultOneInchRequest(amountInUSDT)
	oneInchRequestUsdtUsdc.SetPairs(-2)
	oneInchRequestEthApe := requests.DefaultOneInchRequest(amountInETH)
	oneInchRequestEthApe.SetPairs(-3)

	/// ParaSwap
	paraRequestEthUsdc := requests.DefaultParaSwapRequest(amountInETH)
	paraRequestEthUsdc.SetPairs(1)
	paraRequestUsdcUsdt := requests.DefaultParaSwapRequest(amountInUSDC)
	paraRequestUsdcUsdt.SetPairs(2)
	paraRequestApeEth := requests.DefaultParaSwapRequest(amountInAPE)
	paraRequestApeEth.SetPairs(3)

	paraRequestUsdcEth := requests.DefaultParaSwapRequest(amountInUSDC)
	paraRequestUsdcEth.SetPairs(-1)
	paraRequestUsdtUsdc := requests.DefaultParaSwapRequest(amountInUSDT)
	paraRequestUsdtUsdc.SetPairs(-2)
	paraRequestEthApe := requests.DefaultParaSwapRequest(amountInETH)
	paraRequestEthApe.SetPairs(-3)

	/// 0x
	zeroXRequestEthUsdc := requests.DefaultZeroXRequest(amountInETH)
	zeroXRequestEthUsdc.SetPairs(1)
	zeroXRequestUsdcUsdt := requests.DefaultZeroXRequest(amountInUSDC)
	zeroXRequestUsdcUsdt.SetPairs(2)
	zeroXRequestApeEth := requests.DefaultZeroXRequest(amountInAPE)
	zeroXRequestApeEth.SetPairs(3)

	zeroXRequestUsdcEth := requests.DefaultZeroXRequest(amountInUSDC)
	zeroXRequestUsdcEth.SetPairs(-1)
	zeroXRequestUsdtUsdc := requests.DefaultZeroXRequest(amountInUSDT)
	zeroXRequestUsdtUsdc.SetPairs(-2)
	zeroXRequestEthApe := requests.DefaultZeroXRequest(amountInETH)
	zeroXRequestEthApe.SetPairs(-3)

	clients := []client.Client{
		NewAndStartClient(ksRequestEthUsdc),
		NewAndStartClient(ksRequestUsdcUsdt),
		NewAndStartClient(ksRequestApeEth),
		NewAndStartClient(ksRequestUsdcEth),
		NewAndStartClient(ksRequestUsdtUsdc),
		NewAndStartClient(ksRequestEthApe),

		NewAndStartClient(oneInchRequestEthUsdc),
		NewAndStartClient(oneInchRequestUsdcUsdt),
		NewAndStartClient(oneInchRequestApeEth),
		NewAndStartClient(oneInchRequestUsdcEth),
		NewAndStartClient(oneInchRequestUsdtUsdc),
		NewAndStartClient(oneInchRequestEthApe),

		NewAndStartClient(paraRequestEthUsdc),
		NewAndStartClient(paraRequestUsdcUsdt),
		NewAndStartClient(paraRequestApeEth),
		NewAndStartClient(paraRequestUsdcEth),
		NewAndStartClient(paraRequestUsdtUsdc),
		NewAndStartClient(paraRequestEthApe),

		NewAndStartClient(zeroXRequestEthUsdc),
		NewAndStartClient(zeroXRequestUsdcUsdt),
		NewAndStartClient(zeroXRequestApeEth),
		NewAndStartClient(zeroXRequestUsdcEth),
		NewAndStartClient(zeroXRequestUsdtUsdc),
		NewAndStartClient(zeroXRequestEthApe),
	}

	return clients
}

func (p Program) Run() {
	start := time.Now()

	var db = database.SetupDb()
	defer database.CloseDb(db)

	var clients = setupClients()

	for true {
		fetchTime := time.Now().Unix()
		var wg sync.WaitGroup
		for i, client_ := range clients {
			wg.Add(1)
			go func(wg *sync.WaitGroup, id int, c client.Client, db *gorm.DB) {
				defer wg.Done()
				err := c.Run(db, fetchTime)
				if err != nil {
					panic(err)
				}
			}(&wg, i, client_, db)
		}

		wg.Wait()
		time.Sleep(60 * 15 * time.Second)
		fmt.Println("Waited for 15 minutes, starting a new round")
	}

	/*	for {
			fmt.Println("Service is running")
			time.Sleep(1 * time.Second)
		}
	*/

	// p.Stop(p.service)
	elapsed := time.Since(start)
	fmt.Printf("Time elapsed %s\n", elapsed)
	return
}

func main() {
	go StartServer()
	serviceConfig := &service.Config{
		Name:        "KyberSwap Aggregator Rates",
		DisplayName: "KyberSwap Aggregator Rates",
		Description: "Description",
	}
	program := &Program{}
	s, err := service.New(program, serviceConfig)
	if err != nil {
		fmt.Println("Cannot create the service: " + err.Error())
	}

	err = s.Run()
	if err != nil {
		fmt.Println("Cannot start the service: " + err.Error())
	}
}

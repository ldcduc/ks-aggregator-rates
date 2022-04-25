package main

import (
	"fmt"
	"github.com/kardianos/service"
	"ks-aggregator-rates/internal/pkg/client"
	"ks-aggregator-rates/internal/pkg/client/requests"
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
	go p.Run()
	return nil
}

func (p Program) Stop(s service.Service) error {
	fmt.Println(s.String() + " stopped")
	return nil
}

func (p Program) Run() {
	start := time.Now()

	amountInETH := big.Int{}
	amountInETH.SetString("1000000000000000000", 10)
	amountInUSDC := big.Int{}
	amountInUSDC.SetString("10000000000", 10)
	amountInAPE := big.Int{}
	amountInAPE.SetString("1000000000000000000", 10)

	ksRequestEthUsdc := requests.DefaultKyberSwapRequest(amountInETH)
	ksRequestEthUsdc.SetPairs(1)
	ksRequestUsdcUsdt := requests.DefaultKyberSwapRequest(amountInUSDC)
	ksRequestUsdcUsdt.SetPairs(2)
	ksRequestApeEth := requests.DefaultKyberSwapRequest(amountInAPE)
	ksRequestApeEth.SetPairs(3)

	oneInchRequestEthUsdc := requests.DefaultOneInchRequest(amountInETH)
	oneInchRequestEthUsdc.SetPairs(1)
	oneInchRequestUsdcUsdt := requests.DefaultOneInchRequest(amountInUSDC)
	oneInchRequestUsdcUsdt.SetPairs(2)
	oneInchRequestApeEth := requests.DefaultOneInchRequest(amountInAPE)
	oneInchRequestApeEth.SetPairs(3)

	paraRequestEthUsdc := requests.DefaultParaSwapRequest(amountInETH)
	paraRequestEthUsdc.SetPairs(1)
	paraRequestUsdcUsdt := requests.DefaultParaSwapRequest(amountInUSDC)
	paraRequestUsdcUsdt.SetPairs(2)
	paraRequestApeEth := requests.DefaultParaSwapRequest(amountInAPE)
	paraRequestApeEth.SetPairs(3)

	zeroXRequestEthUsdc := requests.DefaultZeroXRequest(amountInETH)
	zeroXRequestEthUsdc.SetPairs(1)
	zeroXRequestUsdcUsdt := requests.DefaultZeroXRequest(amountInETH)
	zeroXRequestUsdcUsdt.SetPairs(2)
	zeroXRequestApeEth := requests.DefaultZeroXRequest(amountInETH)
	zeroXRequestApeEth.SetPairs(3)

	clients := []client.Client{
		NewAndStartClient(ksRequestEthUsdc),
		NewAndStartClient(ksRequestUsdcUsdt),
		NewAndStartClient(ksRequestApeEth),
		NewAndStartClient(oneInchRequestEthUsdc),
		NewAndStartClient(oneInchRequestUsdcUsdt),
		NewAndStartClient(oneInchRequestApeEth),
		NewAndStartClient(paraRequestEthUsdc),
		NewAndStartClient(paraRequestUsdcUsdt),
		NewAndStartClient(paraRequestApeEth),
		NewAndStartClient(zeroXRequestEthUsdc),
		NewAndStartClient(zeroXRequestUsdcUsdt),
		NewAndStartClient(zeroXRequestApeEth),
	}
	// messages := make(chan int)
	var wg sync.WaitGroup

	// wg.Add(len(clients))

	for i, client_ := range clients {
		wg.Add(1)
		go func(wg *sync.WaitGroup, id int, c client.Client) {
			defer wg.Done()
			err := c.Run()
			if err != nil {
				panic(err)
			}
			// messages <- id
		}(&wg, i, client_)
	}
	wg.Wait()

	elapsed := time.Since(start)
	fmt.Printf("Time elapsed %s\n", elapsed)
	/*	for {
			fmt.Println("Service is running")
			time.Sleep(1 * time.Second)
		}
	*/
	p.Stop(p.service)
}

func main() {
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

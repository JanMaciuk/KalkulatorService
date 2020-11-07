package server

import (
	"github/JahoPL/CalculatorService/files"
	"math"
)

//CalculatorService a
type CalculatorService struct {
	files.UnimplementedGreetServiceServer
}

//NewCaluclatorService a
func NewCaluclatorService() *CalculatorService {
	return &CalculatorService{}
}

//Calculator calculates stuff probably
func (s *CalculatorService) Calculator(req *files.ManyRequest, stream files.GreetService_CalculatorServer) error {

	number := req.GetNumber()
	powers := req.GetPowers()

	var i int32
	for i = 0; i < powers; i++ {

		res := &files.ManyResponse{
			Result: int32(math.Pow(float64(number), float64(i))),
		}
		err := stream.Send(res)

		if err != nil {
			return err
		}
	}
	return nil
}

package rpcdemo

import "errors"

//定义一个rpc的服务
type DemoService struct{}

type Args struct {
	A, B int
}

func (DemoService) Div(
	args Args, result *float64) error {
	if args.B == 0 {
		return errors.New("division by zero")
	}

	*result = float64(args.A) / float64(args.B)
	return nil
}

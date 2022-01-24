package handler

import (
	diction "grpccar/pb/diction"
)

type Serviceserver struct {
	diction.UnimplementedFinderServer
}
package service

import sz "acgfate/serializer"

type CateDetailService struct{}

func (c *CateDetailService) CateDetail() sz.Response {
	return sz.Success()
}

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	pb "github.com/GoogleCloudPlatform/microservices-demo/src/checkoutservice/genproto/hipstershop"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"io/ioutil"
)

func (cs *checkoutService) getUserCartByHttp(ctx context.Context, userID string) ([]*pb.CartItem, error) {
	var cartRequest = CartRequest{
		UserId: userID}

	cartRequestJson, err := json.Marshal(cartRequest)

	resp, err := otelhttp.Post(ctx, fmt.Sprintf("http://%s/Cart/GetCart", cs.cartSvcHttpAddr), "application/json", bytes.NewReader(cartRequestJson))
	if err != nil {
		log.Error("Error post GetCart request.")
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("Error read body from request")
		return nil, err
	}

	var c Cart

	err = json.Unmarshal([]byte(body), &c)
	if err != nil {
		log.Error("Error unmarshaling data from request.")
		return nil, err
	}

	var ci []*pb.CartItem

	for _, item := range c.Items {
		ci = append(ci, &pb.CartItem{
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
		})
	}
	return ci, nil
}

func (cs *checkoutService) emptyUserCartByHttp(ctx context.Context, userID string) error {
	var cartRequest = CartRequest{
		UserId: userID}

	cartRequestJson, err := json.Marshal(cartRequest)

	resp, err := otelhttp.Post(ctx, fmt.Sprintf("http://%s/Cart/EmptyCart", cs.cartSvcHttpAddr), "application/json", bytes.NewReader(cartRequestJson))
	if err != nil {
		log.Error("Error post EmptyCart request.")
		return err
	}
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("Error read body from request")
		return err
	}
	return nil
}

func (cs *checkoutService) quoteShippingByHttp(ctx context.Context, address *pb.Address, items []*pb.CartItem) (*pb.Money, error) {
	var shipOrderRequest = ShipOrderRequest{
		Address: &Address{
			StreetAddress: address.StreetAddress,
			City:          address.City,
			State:         address.State,
			Country:       address.Country,
			ZipCode:       address.ZipCode,
		},
		Items: mapItems(items)}

	shipOrderRequestJson, err := json.Marshal(shipOrderRequest)

	resp, err := otelhttp.Post(ctx, fmt.Sprintf("http://%s/getquote", cs.shippingSvcHttpAddr), "application/json", bytes.NewReader(shipOrderRequestJson))
	if err != nil {
		log.Error("Error post shipOrder request.")
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("Error read body from request")
		return nil, err
	}

	var m Money

	err = json.Unmarshal([]byte(body), &m)
	if err != nil {
		log.Error("Error unmarshaling data from request.")
		return nil, err
	}
	return &pb.Money{
		CurrencyCode: m.CurrencyCode,
		Units:        m.Units,
		Nanos:        m.Nanos,
	}, nil
}

func (cs *checkoutService) shipOrderByHttp(ctx context.Context, address *pb.Address, items []*pb.CartItem) (string, error) {
	var shipOrderRequest = ShipOrderRequest{
		Address: &Address{
			StreetAddress: address.StreetAddress,
			City:          address.City,
			State:         address.State,
			Country:       address.Country,
			ZipCode:       address.ZipCode,
		},
		Items: mapItems(items)}

	shipOrderRequestJson, err := json.Marshal(shipOrderRequest)

	resp, err := otelhttp.Post(ctx, fmt.Sprintf("http://%s/shiporder", cs.shippingSvcHttpAddr), "application/json", bytes.NewReader(shipOrderRequestJson))
	if err != nil {
		log.Error("Error post shipOrder request.")
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("Error read body from request")
		return "", err
	}

	var sor ShipOrderResponse

	err = json.Unmarshal([]byte(body), &sor)
	if err != nil {
		log.Error("Error unmarshaling data from request.")
		return "", err
	}
	return sor.TrackingId, nil
}

func (cs *checkoutService) convertCurrencyByHttp(ctx context.Context, from *pb.Money, toCurrency string) (*pb.Money, error) {
	var currencyConversionRequest = CurrencyConversionRequest{
		From: &Money{
			CurrencyCode: from.CurrencyCode,
			Units:        from.Units,
			Nanos:        from.Nanos,
		},
		ToCode: toCurrency,
	}

	currencyConversionRequestJson, err := json.Marshal(currencyConversionRequest)

	resp, err := otelhttp.Post(ctx, fmt.Sprintf("http://%s/convert", cs.currencySvcHttpAddr), "application/json", bytes.NewReader(currencyConversionRequestJson))
	if err != nil {
		log.Error("Error post convert request.")
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("Error read body from request")
		return nil, err
	}

	var m Money

	err = json.Unmarshal([]byte(body), &m)
	if err != nil {
		log.Error("Error unmarshaling data from request.")
		return nil, err
	}
	return &pb.Money{
		CurrencyCode: m.CurrencyCode,
		Units:        m.Units,
		Nanos:        m.Nanos,
	}, nil
}

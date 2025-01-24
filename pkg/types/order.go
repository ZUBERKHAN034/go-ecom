package types

type OrderItemPayload struct {
	ProductID uint `json:"productId"`
	Quantity  int  `json:"quantity"`
}

type OrderPayload struct {
	Items []OrderItemPayload `json:"orderItems"`
}

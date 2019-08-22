package orderbook

type Orderbook struct {
	Bids   []*Order
	Asks   []*Order
	trades []*Trade
	Id     int
}

func New() *Orderbook {
	orderbook := &Orderbook{}
	orderbook.Bids = []*Order{}
	orderbook.Asks = []*Order{}
	orderbook.trades = []*Trade{}

	return orderbook
}

func (orderbook *Orderbook) Match(order *Order) ([]*Trade, *Order) {
	switch order.Side.String() {
	case "BID":
		if order.Kind.String() == "MARKET" {
			return orderbook.BidMarket(order)
		}
		if order.Kind.String() == "LIMIT" {
			return orderbook.BidLimit(order)
		}
	case "ASK":
		if order.Kind.String() == "MARKET" {
			return orderbook.AskMarket(order)
		}
		if order.Kind.String() == "LIMIT" {
			return orderbook.AskLimit(order)
		}
	}
	return nil, nil
}

func (orderbook *Orderbook) AskMarket(order *Order) ([]*Trade, *Order) {
	orderbook.trades = nil
	for i := 0; i < len(orderbook.Bids); i++ {
		RtOrder := orderbook.Bids[i]
		trade := &Trade{}
		trade.Bid = RtOrder
		trade.Ask = order
		trade.Price = RtOrder.Price
		if RtOrder.Volume > order.Volume {
			trade.Volume = order.Volume
			RtOrder.Volume -= order.Volume
			order.Volume = 0
		} else {
			trade.Volume = RtOrder.Volume
			order.Volume -= RtOrder.Volume
			orderbook.Bids = append(orderbook.Bids[:i], orderbook.Bids[i+1:]...)
			i--
		}
		orderbook.trades = append(orderbook.trades, trade)
		if order.Volume == 0 {
			break
		}
	}
	if order.Volume > 1 {
		return orderbook.trades, order
	} else {
		return orderbook.trades, nil
	}
}

func (orderbook *Orderbook) BidMarket(order *Order) ([]*Trade, *Order) {
	orderbook.trades = nil
	for i := 0; i < len(orderbook.Asks); i++ {
		RtOrder := orderbook.Asks[i]
		trade := &Trade{}
		trade.Bid = order
		trade.Ask = RtOrder
		trade.Price = RtOrder.Price
		if RtOrder.Volume > order.Volume {
			trade.Volume = order.Volume
			RtOrder.Volume -= order.Volume
			order.Volume = 0
		} else {
			trade.Volume = RtOrder.Volume
			order.Volume -= RtOrder.Volume
			orderbook.Asks = append(orderbook.Asks[:i], orderbook.Asks[i+1:]...)
			i--
		}
		orderbook.trades = append(orderbook.trades, trade)
		if order.Volume == 0 {
			break
		}
	}
	if order.Volume > 1 {
		return orderbook.trades, order
	} else {
		return orderbook.trades, nil
	}
}

func (orderbook *Orderbook) AskLimit(order *Order) ([]*Trade, *Order) {
	orderbook.trades = nil
	for i := 0; i < len(orderbook.Bids); i++ {
		RtOrder := orderbook.Bids[i]
		if order.Price <= RtOrder.Price {
			trade := &Trade{}
			trade.Bid = RtOrder
			trade.Ask = order
			trade.Price = RtOrder.Price
			if RtOrder.Volume > order.Volume {
				trade.Volume = order.Volume
				RtOrder.Volume -= order.Volume
				order.Volume = 0
			} else {
				trade.Volume = RtOrder.Volume
				order.Volume -= RtOrder.Volume
				orderbook.Bids = append(orderbook.Bids[:i], orderbook.Bids[i+1:]...)
				i--
			}
			orderbook.trades = append(orderbook.trades, trade)
			if order.Volume == 0 {
				break
			}
		} else {
			break
		}
	}
	if order.Volume > 1 {
		orderbook.Asks = append(orderbook.Asks, order)
		for i := 0; i < len(orderbook.Asks)-1; i++ {
			if orderbook.Asks[i].Price > orderbook.Asks[i+1].Price {
				orderbook.Asks[i], orderbook.Asks[i+1] = orderbook.Asks[i+1], orderbook.Asks[i]
			} else {
				break
			}
		}
	}
	return orderbook.trades, nil
}

func (orderbook *Orderbook) BidLimit(order *Order) ([]*Trade, *Order) {
	orderbook.trades = nil
	for i := 0; i < len(orderbook.Asks); i++ {
		RtOrder := orderbook.Asks[i]
		if order.Price >= RtOrder.Price {
			trade := &Trade{}
			trade.Bid = order
			trade.Ask = RtOrder
			trade.Price = RtOrder.Price
			if RtOrder.Volume > order.Volume {
				trade.Volume = order.Volume
				RtOrder.Volume -= order.Volume
				order.Volume = 0
			} else {
				trade.Volume = RtOrder.Volume
				order.Volume -= RtOrder.Volume
				orderbook.Asks = append(orderbook.Asks[:i], orderbook.Asks[i+1:]...)
				i--
			}
			orderbook.trades = append(orderbook.trades, trade)
			if order.Volume == 0 {
				break
			}
		} else {
			break
		}
	}
	if order.Volume > 1 {
		orderbook.Bids = append(orderbook.Bids, order)
		for i := len(orderbook.Bids) - 1; i > 0; i-- {
			if orderbook.Bids[i].Price > orderbook.Bids[i-1].Price {
				orderbook.Bids[i], orderbook.Bids[i-1] = orderbook.Bids[i-1], orderbook.Bids[i]
			}
		}
	}
	return orderbook.trades, nil
}

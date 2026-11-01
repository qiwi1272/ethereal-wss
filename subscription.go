package meridian

import (
	"encoding/json"
	"errors"

	meridianv1 "github.com/qiwi1272/protos/gen/dex/meridian/v1"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// EventType represents the type of event for subscriptions.
type EventType = meridianv1.EventType

// Event type constants for easy access without importing protos.
const (
	EventTypeL2Book                EventType = meridianv1.EventType_EVENT_TYPE_L2_BOOK
	EventTypeTicker                EventType = meridianv1.EventType_EVENT_TYPE_TICKER
	EventTypeTradeFill             EventType = meridianv1.EventType_EVENT_TYPE_TRADE_FILL
	EventTypeSubaccountLiquidation EventType = meridianv1.EventType_EVENT_TYPE_SUBACCOUNT_LIQUIDATION
	EventTypePositionUpdate        EventType = meridianv1.EventType_EVENT_TYPE_POSITION_UPDATE
	EventTypeOrderUpdate           EventType = meridianv1.EventType_EVENT_TYPE_ORDER_UPDATE
	EventTypeOrderFill             EventType = meridianv1.EventType_EVENT_TYPE_ORDER_FILL
	EventTypeTokenTransfer         EventType = meridianv1.EventType_EVENT_TYPE_TOKEN_TRANSFER
)

var ErrUnknownEvent = errors.New("unknown event")

type Intent string

const (
	Sub   Intent = "subscribe"
	Unsub Intent = "unsubscribe"
)

type WebsocketRequest struct {
	I Intent      `json:"event"`
	D interface{} `json:"data"`
}

func (i Intent) MarshalEventData(event EventType, to string) ([]byte, error) {
	var data interface{}
	switch event {
	case meridianv1.EventType_EVENT_TYPE_L2_BOOK,
		meridianv1.EventType_EVENT_TYPE_TICKER,
		meridianv1.EventType_EVENT_TYPE_TRADE_FILL:
		data = struct {
			T string `json:"type"`
			S string `json:"symbol"`
		}{meridianv1.EventType_json_name[event], to}
	case meridianv1.EventType_EVENT_TYPE_SUBACCOUNT_LIQUIDATION,
		meridianv1.EventType_EVENT_TYPE_POSITION_UPDATE,
		meridianv1.EventType_EVENT_TYPE_ORDER_UPDATE,
		meridianv1.EventType_EVENT_TYPE_ORDER_FILL,
		meridianv1.EventType_EVENT_TYPE_TOKEN_TRANSFER:
		data = struct {
			T string `json:"type"`
			S string `json:"subaccountId"`
		}{meridianv1.EventType_json_name[event], to}
	default:
		return nil, ErrUnknownEvent
	}
	return json.Marshal(WebsocketRequest{i, data})
}

func UnmarshalEvent(event EventType, data []byte) (proto.Message, error) {
	var m proto.Message
	switch event {
	case meridianv1.EventType_EVENT_TYPE_L2_BOOK:
		m = new(meridianv1.L2Book)
	case meridianv1.EventType_EVENT_TYPE_TICKER:
		m = new(meridianv1.Ticker)
	case meridianv1.EventType_EVENT_TYPE_TRADE_FILL:
		m = new(meridianv1.TradeFill)
	case meridianv1.EventType_EVENT_TYPE_SUBACCOUNT_LIQUIDATION:
		m = new(meridianv1.SubaccountLiquidation)
	case meridianv1.EventType_EVENT_TYPE_POSITION_UPDATE:
		m = new(meridianv1.PositionUpdate)
	case meridianv1.EventType_EVENT_TYPE_ORDER_UPDATE:
		m = new(meridianv1.OrderUpdate)
	case meridianv1.EventType_EVENT_TYPE_ORDER_FILL:
		m = new(meridianv1.OrderFill)
	case meridianv1.EventType_EVENT_TYPE_TOKEN_TRANSFER:
		m = new(meridianv1.TokenTransfer)
	default:
		return nil, ErrUnknownEvent
	}
	if err := protojson.Unmarshal(data, m); err != nil {
		return nil, err
	}
	return m, nil
}

// type Subscription[Proto proto.Message, Data EventData] struct {
// 	*websocket.Conn
// 	eventName string
// 	data      Proto
// 	Callback  func(Proto)

// 	eventType EventData
// }

// func NewSubscription[P proto.Message, D EventData](callback func(P)) *Subscription[P, D] {
// 	return &Subscription[P, D]{}
// }

// func (c *Subscription[_, EventData]) Subscribe(ctx context.Context) error {
// 	if bytes, err := json.Marshal(&SubscriptionIntent[EventData]{I: Sub, D: c.eventType}); err != nil {
// 		return err
// 	} else {
// 		return c.Write(ctx, websocket.MessageBinary, bytes)
// 	}
// }

// func (c *Subscription[_, EventData]) Unsubscribe(ctx context.Context) error {
// 	if bytes, err := json.Marshal(&SubscriptionIntent[EventData]{I: Unsub, D: c.eventType}); err != nil {
// 		return err
// 	} else {
// 		return c.Write(ctx, websocket.MessageBinary, bytes)
// 	}
// }

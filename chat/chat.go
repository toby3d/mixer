package chat

import (
	"encoding/json"

	ws "github.com/gorilla/websocket"
	ffjson "github.com/pquerna/ffjson/ffjson"
)

const method = "method"

type (
	Method struct {
		Type      string        `json:"type"`
		Method    string        `json:"method"`
		Arguments []interface{} `json:"arguments"`
		ID        uint          `json:"id"`
	}

	Reply struct {
		Type  string          `json:"type"`
		Error string          `json:"error"`
		Data  json.RawMessage `json:"data"`
		ID    uint            `json:"id"`
	}

	Event struct {
		Type  string          `json:"type"`
		Event string          `json:"event"`
		Data  json.RawMessage `json:"data"`
	}

	Connection struct {
		*ws.Conn
	}
)

func Connect(endpoint string) (*Connection, error) {
	var conn Connection
	dial, _, err := ws.DefaultDialer.Dial(endpoint, nil)
	conn.Conn = dial
	return &conn, err
}

// Auth authenticating as a User successfully.
func (c *Connection) Auth(channelID, userID int, key string) error {
	var args []interface{}

	args = append(args, channelID)

	if userID != 0 {
		args = append(args, userID)
	}
	if key != "" {
		args = append(args, key)
	}

	mtd := &Method{
		Type:      method,
		Method:    "auth",
		Arguments: args,
		ID:        0,
	}

	msg, err := ffjson.Marshal(mtd)
	if err != nil {
		return err
	}

	return c.Conn.WriteMessage(ws.TextMessage, msg)
}

func (c *Connection) Msg(message string) error {
	mtd := &Method{
		Type:   method,
		Method: "msg",
		Arguments: []interface{}{
			message,
		},
		ID: 2,
	}

	msg, err := ffjson.Marshal(mtd)
	if err != nil {
		return err
	}

	return c.Conn.WriteMessage(ws.TextMessage, msg)
}

func (c *Connection) Whisper(targetUsername, message string) error {
	mtd := &Method{
		Type:   method,
		Method: "whisper",
		Arguments: []interface{}{
			targetUsername,
			message,
		},
		ID: 2,
	}

	msg, err := ffjson.Marshal(mtd)
	if err != nil {
		return err
	}

	return c.Conn.WriteMessage(ws.TextMessage, msg)
}

func (c *Connection) VoteChoose(voteIndex int) error {
	mtd := &Method{
		Type:   method,
		Method: "vote",
		Arguments: []interface{}{
			voteIndex,
		},
		ID: 3,
	}

	msg, err := ffjson.Marshal(mtd)
	if err != nil {
		return err
	}

	return c.Conn.WriteMessage(ws.TextMessage, msg)
}

func (c *Connection) VoteStart(question string, duration int, options ...string) error {
	mtd := &Method{
		Type:   method,
		Method: "vote:start",
		Arguments: []interface{}{
			question,
			options,
			duration,
		},
		ID: 3,
	}

	msg, err := ffjson.Marshal(mtd)
	if err != nil {
		return err
	}

	return c.Conn.WriteMessage(ws.TextMessage, msg)
}

func (c *Connection) Timeout(username string, duration int) error {
	mtd := &Method{
		Type:   method,
		Method: "timeout",
		Arguments: []interface{}{
			username,
			duration,
		},
		ID: 4,
	}

	msg, err := ffjson.Marshal(mtd)
	if err != nil {
		return err
	}

	return c.Conn.WriteMessage(ws.TextMessage, msg)
}

func (c *Connection) Purge(username string) error {
	mtd := &Method{
		Type:   method,
		Method: "purge",
		Arguments: []interface{}{
			username,
		},
		ID: 5,
	}

	msg, err := ffjson.Marshal(mtd)
	if err != nil {
		return err
	}

	return c.Conn.WriteMessage(ws.TextMessage, msg)
}

func (c *Connection) DeleteMessage(messageID string) error {
	mtd := &Method{
		Type:   method,
		Method: "deleteMessage",
		Arguments: []interface{}{
			messageID,
		},
		ID: 10,
	}

	msg, err := ffjson.Marshal(mtd)
	if err != nil {
		return err
	}

	return c.Conn.WriteMessage(ws.TextMessage, msg)
}

func (c *Connection) ClearMessages() error {
	mtd := &Method{
		Type:   method,
		Method: "clearMessages",
		ID:     11,
	}

	msg, err := ffjson.Marshal(mtd)
	if err != nil {
		return err
	}

	return c.Conn.WriteMessage(ws.TextMessage, msg)
}

func (c *Connection) History(limit int) error {
	mtd := &Method{
		Type:   method,
		Method: "history",
		Arguments: []interface{}{
			limit,
		},
		ID: 1,
	}

	msg, err := ffjson.Marshal(mtd)
	if err != nil {
		return err
	}

	return c.Conn.WriteMessage(ws.TextMessage, msg)
}

func (c *Connection) Giveaway() error {
	mtd := &Method{
		Type:   method,
		Method: "giveaway:start",
		ID:     11,
	}

	msg, err := ffjson.Marshal(mtd)
	if err != nil {
		return err
	}

	return c.Conn.WriteMessage(ws.TextMessage, msg)
}

func (c *Connection) Ping() error {
	mtd := &Method{
		Type:   method,
		Method: "ping",
		ID:     12,
	}

	msg, err := ffjson.Marshal(mtd)
	if err != nil {
		return err
	}

	return c.Conn.WriteMessage(ws.TextMessage, msg)
}

func (c *Connection) AttachEmotes() error {
	mtd := &Method{
		Type:   method,
		Method: "attachEmotes",
		ID:     12,
	}

	msg, err := ffjson.Marshal(mtd)
	if err != nil {
		return err
	}

	return c.Conn.WriteMessage(ws.TextMessage, msg)
}

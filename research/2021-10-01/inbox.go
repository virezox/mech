package goinsta

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Conversation is the representation of an instagram already established conversation through direct messages.
type Conversation struct {
	insta     *Instagram
	err       error
	firstRun  bool
	isPending bool

	ID   string `json:"thread_id"`
	V2ID string `json:"thread_v2_id"`
	// Items can be of many types.
	Items                      []*InboxItem          `json:"items"`
	Title                      string                `json:"thread_title"`
	Users                      []*User               `json:"users"`
	LeftUsers                  []*User               `json:"left_users"`
	AdminUserIDs               []int64               `json:"admin_user_ids"`
	ApprovalRequiredNewMembers bool                  `json:"approval_required_for_new_members"`
	Pending                    bool                  `json:"pending"`
	PendingScore               int64                 `json:"pending_score"`
	ReshareReceiveCount        int                   `json:"reshare_receive_count"`
	ReshareSendCount           int                   `json:"reshare_send_count"`
	ViewerID                   int64                 `json:"viewer_id"`
	ValuedRequest              bool                  `json:"valued_request"`
	LastActivityAt             int64                 `json:"last_activity_at"`
	Named                      bool                  `json:"named"`
	Muted                      bool                  `json:"muted"`
	Spam                       bool                  `json:"spam"`
	ShhModeEnabled             bool                  `json:"shh_mode_enabled"`
	ShhReplayEnabled           bool                  `json:"shh_replay_enabled"`
	IsPin                      bool                  `json:"is_pin"`
	IsGroup                    bool                  `json:"is_group"`
	IsVerifiedThread           bool                  `json:"is_verified_thread"`
	IsCloseFriendThread        bool                  `json:"is_close_friend_thread"`
	ThreadType                 string                `json:"thread_type"`
	ExpiringMediaSendCount     int                   `json:"expiring_media_send_count"`
	ExpiringMediaReceiveCount  int                   `json:"expiring_media_receive_count"`
	Inviter                    *User                 `json:"inviter"`
	HasOlder                   bool                  `json:"has_older"`
	HasNewer                   bool                  `json:"has_newer"`
	HasRestrictedUser          bool                  `json:"has_restricted_user"`
	Archived                   bool                  `json:"archived"`
	LastSeenAt                 map[string]lastSeenAt `json:"last_seen_at"`
	NewestCursor               string                `json:"newest_cursor"`
	OldestCursor               string                `json:"oldest_cursor"`

	LastPermanentItem InboxItem `json:"last_permanent_item"`
}

// InboxItem is any conversation message.
type InboxItem struct {
	ID            string `json:"item_id"`
	UserID        int64  `json:"user_id"`
	Timestamp     int64  `json:"timestamp"`
	ClientContext string `json:"client_context"`
	IsShhMode     bool   `json:"is_shh_mode"`
	TqSeqID       int    `json:"tq_seq_id"`

	// Type there are a few types:
	// text, like, raven_media, action_log, media_share, reel_share, link
	Type string `json:"item_type"`

	// Text is message text.
	Text string `json:"text"`

	// InboxItemLike is the heart that your girlfriend send to you.
	// (or in my case: the heart that my fans sends to me hehe)

	Like string `json:"like"`

	Reel      *reelShare `json:"reel_share"`
	Media     *Item      `json:"media_share"`
	ActionLog *actionLog `json:"action_log"`
	Link      struct {
		Text    string `json:"text"`
		Context struct {
			Url      string `json:"link_url"`
			Title    string `json:"link_title"`
			Summary  string `json:"link_summary"`
			ImageUrl string `json:"link_image_url"`
		} `json:"link_context"`
	} `json:"link"`
}

type inboxResp struct {
	isPending bool

	MostRecentInviter     *User  `json:"most_recent_inviter"`
	SeqID                 int64  `json:"seq_id"`
	PendingRequestsTotal  int    `json:"pending_requests_total"`
	SnapshotAtMs          int64  `json:"snapshot_at_ms"`
	Status                string `json:"status"`
	HasPendingTopRequests bool   `json:"has_pending_top_requests"`
}

type threadResp struct {
	Conversation *Conversation `json:"thread"`
	Status       string        `json:"status"`
}

type msgResp struct {
	Action  string `json:"action"`
	Payload struct {
		ClientContext string `json:"client_context"`
		ItemID        string `json:"item_id"`
		ThreadID      string `json:"thread_id"`
		Timestamp     string `json:"timestamp"`
	} `json:"payload"`
	Status     string `json:"status"`
	StatusCode string `json:"status_code"`
}

type reelShare struct {
	Text        string `json:"text"`
	IsPersisted bool   `json:"is_reel_persisted"`
	OwnderID    int64  `json:"reel_owner_id"`
	Type        string `json:"type"`
	ReelType    string `json:"reel_type"`
	Media       Item   `json:"media"`
}

type actionLog struct {
	Description string `json:"description"`
}

type lastSeenAt struct {
	Timestamp string `json:"timestamp"`
	ItemID    string `json:"item_id"`
}

func (c *Conversation) send(query map[string]string) error {
	body, _, err := c.insta.sendRequest(
		&reqOptions{
			Endpoint: urlInboxSend,
			IsPost:   true,
			Query:    query,
		},
	)
	if err != nil {
		return err
	}

	var resp msgResp
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return err
	}
	c.ID = resp.Payload.ThreadID

	ts, _ := strconv.ParseInt(resp.Payload.Timestamp, 10, 64)
	msg := &InboxItem{
		ID:            resp.Payload.ItemID,
		ClientContext: resp.Payload.ClientContext,
		Timestamp:     ts,
		Type:          "text",
	}
	c.addMessage(msg)
	return nil
}

func (conv *Conversation) approve() error {
	insta := conv.insta
	if !conv.isPending {
		return ErrConvNotPending
	}

	body, _, err := insta.sendRequest(
		&reqOptions{
			Endpoint: fmt.Sprintf(urlInboxApprove, conv.ID),
			IsPost:   true,
			Query: map[string]string{
				"filter": "DEFAULT",
				"_uuid":  insta.uuid,
			},
		},
	)
	if err != nil {
		return err
	}

	var resp struct {
		Status string `json:"status"`
	}
	err = json.Unmarshal(body, &resp)
	if resp.Status != "ok" {
		return fmt.Errorf("Failed to approve conversation with status: %s", resp.Status)
	}
	return nil
}

func (conv *Conversation) Hide() error {
	insta := conv.insta
	if !conv.isPending {
		return ErrConvNotPending
	}

	body, _, err := insta.sendRequest(
		&reqOptions{
			Endpoint: fmt.Sprintf(urlInboxHide, conv.ID),
			IsPost:   true,
			Query: map[string]string{
				"_uuid": insta.uuid,
			},
		},
	)
	if err != nil {
		return err
	}

	var resp struct {
		Status string `json:"status"`
	}
	err = json.Unmarshal(body, &resp)
	if resp.Status != "ok" {
		return fmt.Errorf("Failed to hide conversation with status: %s", resp.Status)
	}
	return nil
}

// Error will return Conversation.err
func (c *Conversation) Error() error {
	return c.err
}

func (c Conversation) lastItemID() string {
	n := len(c.Items)
	if n == 0 {
		return ""
	}
	return c.Items[n-1].ID
}

// DEPRICATED - doesn't work anymore
// Like sends heart to the conversation
//
// See example: examples/media/likeAll.go
// func (c *Conversation) Like() error {
// 	insta := c.insta
// 	to, err := prepareRecipients(c)
// 	if err != nil {
// 		return err
// 	}
//
// 	thread, err := json.Marshal([]string{c.ID})
// 	if err != nil {
// 		return err
// 	}
//
// 	data := insta.prepareDataQuery(
// 		map[string]interface{}{
// 			"recipient_users": to,
// 			"client_context":  generateUUID(),
// 			"thread_ids":      string(thread),
// 			"action":          "send_item",
// 		},
// 	)
// 	_, _, err = insta.sendRequest(
// 		&reqOptions{
// 			Endpoint: urlInboxSendLike,
// 			Query:    data,
// 			IsPost:   true,
// 		},
// 	)
// 	return err
// }

// Send sends message in conversation
func (c *Conversation) Send(text string) error {
	insta := c.insta
	// I DON'T KNOW WHY BUT INSTAGRAM WANTS A DOUBLE SLICE OF INTS FOR ONE ID. << lol
	to, err := prepareRecipients(c)
	if err != nil {
		return err
	}

	// I DONT KNOW WHY BUT INSTAGRAM WANTS SLICE OF STRINGS FOR ONE ID. << lol
	thread, err := json.Marshal([]string{c.ID})
	if err != nil {
		return err
	}
	query := map[string]string{
		"recipient_users": to,
		"client_context":  generateUUID(),
		"thread_ids":      string(thread),
		"action":          "send_item",
		"text":            text,
		"_uuid":           insta.uuid,
		"device_id":       insta.dID,
	}

	err = c.send(query)
	return err
}

// Write is like Send but being compatible with io.Writer.
func (c *Conversation) Write(b []byte) (int, error) {
	n := len(b)
	return n, c.Send(string(b))
}

// MarkAsSeen will marks a message as seen.
func (c *Conversation) MarkAsSeen(msg InboxItem) error {
	insta := c.insta
	token := "68" + randNum(17)
	body, _, err := insta.sendRequest(
		&reqOptions{
			Endpoint: fmt.Sprintf(urlInboxMsgSeen, c.ID, msg.ID),
			IsPost:   true,
			Query: map[string]string{
				"thread_id":            c.ID,
				"action":               "mark_seen",
				"client_context":       token,
				"_uuid":                insta.uuid,
				"offline_threading_id": token,
			},
		},
	)
	if err != nil {
		return err
	}
	var resp struct {
		Status string `json:"status"`
	}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return err
	}
	if resp.Status != "ok" {
		return fmt.Errorf("Status not ok while calling msg seen, '%s'", resp.Status)
	}
	return nil
}

// GetItems is an alternative way to get conversation messages, e.g. refresh.
// The app calls this when approving a DM request, for example.
func (c *Conversation) GetItems() error {
	insta := c.insta

	var ctxs []string
	var itemIDs []string
	for _, v := range c.Items {
		ctxs = append(ctxs, v.ClientContext)
		itemIDs = append(itemIDs, v.ID)
	}

	origCtxs, err := json.Marshal(ctxs)
	if err != nil {
		return err
	}

	x := itemIDs[0]
	for _, id := range itemIDs[1:] {
		x += "," + id
	}

	body, _, err := insta.sendRequest(
		&reqOptions{
			Endpoint: fmt.Sprintf(urlInboxThread, c.ID),
			Query: map[string]string{
				"visual_message_return_type":       "unseen",
				"_uuid":                            insta.uuid,
				"original_message_client_contexts": string(origCtxs),
				"item_ids":                         fmt.Sprintf("[%s]", x),
			},
		},
	)
	if err != nil {
		return err
	}

	var resp struct {
		Items  []*InboxItem `json:"items"`
		Status string       `json:"status"`
	}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return err
	}

	for _, msg := range resp.Items {
		c.addMessage(msg)
	}

	return nil
}

func (c *Conversation) update(newConv *Conversation) {
	insta := c.insta
	oldItems := c.Items
	newConv.setValues(insta)

	*c = *newConv
	c.Items = oldItems

	for _, msg := range newConv.Items {
		c.addMessage(msg)
	}
}

func (c *Conversation) addMessage(msg *InboxItem) {
	msg.setValues(c.insta)
	for _, m := range c.Items {
		// return if msg already present
		if msg.ID == m.ID {
			*m = *msg
			return
		}
	}
	if len(c.Items) == 0 {
		c.Items = []*InboxItem{msg}
		return
	} else if msg.Timestamp > c.Items[0].Timestamp {
		// If newer than newest
		c.Items = append([]*InboxItem{msg}, c.Items...)
		return
	} else if msg.Timestamp < c.Items[len(c.Items)-1].Timestamp {
		// if older than oldest
		c.Items = append(c.Items, msg)
		return
	}
	// if somewhere in between
	for i, m := range c.Items {
		if msg.Timestamp > m.Timestamp {
			l := append([]*InboxItem{msg}, c.Items[i:]...)
			c.Items = append(c.Items[:i], l...)
		}
	}
}

func (c *Conversation) setValues(insta *Instagram) {
	c.insta = insta

	for _, msg := range c.Items {
		msg.setValues(insta)
	}

	if c.Inviter != nil {
		c.Inviter.insta = insta
	}
	for _, u := range c.Users {
		u.insta = insta
	}
	for _, u := range c.LeftUsers {
		u.insta = insta
	}
}

func (msg *InboxItem) setValues(insta *Instagram) {
	if msg.Reel != nil {
		msg.Reel.Media.insta = insta
		msg.Reel.Media.User.insta = insta
	}
	if msg.Media != nil {
		msg.Media.insta = insta
		msg.Media.User.insta = insta
	}
}

package types

import (
	"bytes"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	crypto "github.com/tendermint/go-crypto"
)

func TestDepositMsg_ValidateBasic(t *testing.T) {
	coin := sdk.Coin{Amount: 100, Denom: "ATM"}
	coinNegative := sdk.Coin{Amount: -100, Denom: "ATM"}
	short := crypto.Address("foo")
	long := crypto.Address("hefkuhwqekufghwqekufgwqekufgkwuqgfkugfkuwgek")
	addr := crypto.GenPrivKeyEd25519().PubKey().Address()
	addr2 := crypto.GenPrivKeyEd25519().PubKey().Address()
	addr3 := crypto.GenPrivKeyEd25519().PubKey().Address()

	type fields struct {
		Operator  crypto.Address
		Sender    crypto.Address
		Recipient crypto.Address
		Amount    sdk.Coin
	}
	tests := []struct {
		name      string
		fields    fields
		errorCode sdk.CodeType
	}{
		{
			"empty msg",
			fields{},
			CodeInvalidAmount,
		},
		{
			"no denom",
			fields{Amount: sdk.Coin{Amount: 100}},
			CodeInvalidAmount,
		},
		{
			"no amount",
			fields{Amount: sdk.Coin{Denom: "Foo"}},
			CodeInvalidAmount,
		},
		{
			"missing operator address",
			fields{Amount: coin},
			CodeInvalidAddress,
		},
		{
			"missing sender address",
			fields{Amount: coin, Operator: addr},
			CodeInvalidAddress,
		},
		{
			"short address",
			fields{Amount: coin, Operator: short},
			CodeInvalidAddress,
		},
		{
			"long sender address",
			fields{Amount: coin, Operator: addr, Sender: long},
			CodeInvalidAddress,
		},
		{
			"long recipient address",
			fields{Amount: coin, Operator: addr, Sender: addr, Recipient: long},
			CodeInvalidAddress,
		},
		{
			"same address",
			fields{Amount: coin, Operator: addr2, Sender: addr, Recipient: addr},
			CodeInvalidAddress,
		},
		{
			"proper addresses",
			fields{Amount: coin, Operator: addr, Sender: addr2, Recipient: addr3},
			sdk.CodeOK,
		},
		{
			"negative amount",
			fields{Amount: coinNegative, Operator: addr, Sender: addr2, Recipient: addr3},
			CodeInvalidAmount,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := DepositMsg{
				Operator:  tt.fields.Operator,
				Sender:    tt.fields.Sender,
				Recipient: tt.fields.Recipient,
				Amount:    tt.fields.Amount,
			}
			got := d.ValidateBasic()
			if got == nil {
				assert.True(t, tt.errorCode.IsOK())
			} else {
				assert.Equal(t, tt.errorCode, got.ABCICode())
			}
		})
	}
}
func TestSettleMsg_ValidateBasic(t *testing.T) {
	coin := sdk.Coin{Amount: 100, Denom: "ATM"}
	coinNegative := sdk.Coin{Amount: -100, Denom: "ATM"}
	short := crypto.Address("foo")
	long := crypto.Address("hefkuhwqekufghwqekufgwqekufgkwuqgfkugfkuwgek")
	addr := crypto.GenPrivKeyEd25519().PubKey().Address()
	addr2 := crypto.GenPrivKeyEd25519().PubKey().Address()
	addr3 := crypto.GenPrivKeyEd25519().PubKey().Address()

	type fields struct {
		Operator  crypto.Address
		Sender    crypto.Address
		Recipient crypto.Address
		Amount    sdk.Coin
	}
	tests := []struct {
		name      string
		fields    fields
		errorCode sdk.CodeType
	}{
		{
			"empty msg",
			fields{},
			CodeInvalidAmount,
		},
		{
			"no denom",
			fields{Amount: sdk.Coin{Amount: 100}},
			CodeInvalidAmount,
		},
		{
			"no amount",
			fields{Amount: sdk.Coin{Denom: "Foo"}},
			CodeInvalidAmount,
		},
		{
			"missing operator address",
			fields{Amount: coin},
			CodeInvalidAddress,
		},
		{
			"short address",
			fields{Amount: coin, Operator: short, Sender: short, Recipient: short},
			CodeInvalidAddress,
		},
		{
			"long address",
			fields{Amount: coin, Operator: long, Sender: short, Recipient: long},
			CodeInvalidAddress,
		},
		{
			"long address2",
			fields{Amount: coin, Operator: addr, Sender: addr2, Recipient: long},
			CodeInvalidAddress,
		},
		{
			"sender and recipient got same address",
			fields{Amount: coin, Operator: addr, Sender: addr2, Recipient: addr2},
			CodeInvalidAddress,
		},
		{
			"proper address",
			fields{Amount: coin, Operator: addr, Sender: addr2, Recipient: addr3},
			sdk.CodeOK,
		},
		{
			"proper negative amount",
			fields{Amount: coinNegative, Operator: addr3, Sender: addr, Recipient: addr2},
			sdk.CodeOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := SettleMsg{
				Operator:  tt.fields.Operator,
				Sender:    tt.fields.Sender,
				Recipient: tt.fields.Recipient,
				Amount:    tt.fields.Amount,
			}
			got := d.ValidateBasic()
			if got == nil {
				assert.True(t, tt.errorCode.IsOK())
			} else {
				assert.Equal(t, tt.errorCode, got.ABCICode())
			}
		})
	}
}
func TestWithdrawMsg_ValidateBasic(t *testing.T) {
	coin := sdk.Coin{Amount: 100, Denom: "ATM"}
	coinNegative := sdk.Coin{Amount: -100, Denom: "ATM"}
	short := crypto.Address("foo")
	long := crypto.Address("hefkuhwqekufghwqekufgwqekufgkwuqgfkugfkuwgek")
	addr := crypto.GenPrivKeyEd25519().PubKey().Address()
	addr2 := crypto.GenPrivKeyEd25519().PubKey().Address()
	addr3 := crypto.GenPrivKeyEd25519().PubKey().Address()

	type fields struct {
		Sender    crypto.Address
		Recipient crypto.Address
		Operator  crypto.Address
		Amount    sdk.Coin
	}
	tests := []struct {
		name      string
		fields    fields
		errorCode sdk.CodeType
	}{
		{"empty msg", fields{}, CodeInvalidAmount},
		{"no denom", fields{Amount: sdk.Coin{Amount: 100}}, CodeInvalidAmount},
		{"no amount", fields{Amount: sdk.Coin{Denom: "Foo"}}, CodeInvalidAmount},
		{"missing address", fields{Amount: coin}, CodeInvalidAddress},
		{"short address", fields{Amount: coin, Sender: short, Recipient: short}, CodeInvalidAddress},
		{"long address", fields{Amount: coin, Sender: long, Recipient: long}, CodeInvalidAddress},
		{"long address2", fields{Amount: coin, Sender: addr, Recipient: long}, CodeInvalidAddress},
		{"same address", fields{Amount: coin, Sender: addr, Recipient: addr, Operator: addr3}, CodeInvalidAddress},
		{"missing proper address", fields{Amount: coin, Sender: addr, Recipient: addr2}, CodeInvalidAddress},
		{"negative amount", fields{Amount: coinNegative, Sender: addr, Recipient: addr2}, CodeInvalidAmount},
		{"proper address", fields{Amount: coin, Sender: addr, Recipient: addr2, Operator: addr3}, sdk.CodeOK},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := WithdrawMsg{
				Operator:  tt.fields.Operator,
				Sender:    tt.fields.Sender,
				Recipient: tt.fields.Recipient,
				Amount:    tt.fields.Amount,
			}
			got := w.ValidateBasic()
			if got == nil {
				assert.True(t, tt.errorCode.IsOK())
			} else {
				assert.Equal(t, tt.errorCode, got.ABCICode(), got.Error())
			}
		})
	}
}

func TestCreateAssetAccountMsg_ValidateBasic(t *testing.T) {
	creatorAddress := crypto.GenPrivKeyEd25519().PubKey().Address()
	newPubKey := crypto.GenPrivKeyEd25519().PubKey()
	type fields struct {
		Creator crypto.Address
		PubKey  crypto.PubKey
	}
	tests := []struct {
		name   string
		fields fields
		want   sdk.CodeType
	}{
		{"new CH acc ok", fields{creatorAddress, newPubKey}, sdk.CodeOK},
		{"new CUS acc ok", fields{creatorAddress, newPubKey}, sdk.CodeOK},
		{"new GCM acc ok", fields{creatorAddress, newPubKey}, sdk.CodeOK},
		{"new ICM acc ok", fields{creatorAddress, newPubKey}, sdk.CodeOK},
		{"creator is nil", fields{nil, newPubKey}, CodeInvalidAddress},
		{"invalid creator len", fields{crypto.Address("short"), newPubKey}, CodeInvalidAddress},
		{"new pubkey is nil", fields{creatorAddress, nil}, CodeInvalidPubKey},
		{"same creator and acct", fields{newPubKey.Address(), newPubKey}, CodeInvalidPubKey},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := CreateAssetAccountMsg{
				Creator: tt.fields.Creator,
				PubKey:  tt.fields.PubKey,
			}
			got := msg.ValidateBasic()
			if got == nil {
				assert.True(t, tt.want == sdk.CodeOK)
			} else {
				assert.Equal(t, tt.want, got.ABCICode(), got.Error())
			}
		})
	}
}

func TestBaseCreateUserMsg_ValidateBasic(t *testing.T) {
	pub := crypto.GenPrivKeyEd25519().PubKey()
	addr := crypto.GenPrivKeyEd25519().PubKey().Address()
	long := crypto.Address("hefkuhwqekufghwqekufgwqekufgkwuqgfkugfkuwgek")
	type fields struct {
		Creator crypto.Address
		PubKey  crypto.PubKey
	}
	tests := []struct {
		name   string
		fields fields
		want   sdk.CodeType
	}{
		{"nil pubkey", fields{Creator: pub.Address()}, CodeInvalidPubKey},
		{"nil address", fields{PubKey: pub}, CodeInvalidAddress},
		{"empty address", fields{Creator: crypto.Address(""), PubKey: pub}, CodeInvalidAddress},
		{"short address", fields{Creator: crypto.Address("foo"), PubKey: pub}, CodeInvalidAddress},
		{"long address", fields{Creator: long, PubKey: pub}, CodeInvalidAddress},
		{"self create", fields{Creator: pub.Address(), PubKey: pub}, CodeSelfCreate},
		{"good to go", fields{Creator: addr, PubKey: pub}, sdk.CodeOK},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := BaseCreateUserMsg{
				Creator: tt.fields.Creator,
				PubKey:  tt.fields.PubKey,
			}
			got := msg.ValidateBasic()
			if got != nil {
				assert.Equal(t, tt.want, got.ABCICode(), got.Error())
			} else {
				assert.Equal(t, tt.want, sdk.CodeOK)
			}
		})
	}
}

func TestCreateAdminMsg_ValidateBasic(t *testing.T) {
	addr := crypto.GenPrivKeyEd25519().PubKey().Address()
	pub := crypto.GenPrivKeyEd25519().PubKey()
	validEntity := BaseLegalEntity{
		EntityName: "CH",
		EntityType: EntityClearingHouse,
	}
	validCreateUser := BaseCreateUserMsg{
		Creator: addr,
		PubKey:  pub,
	}
	type fields struct {
		cm BaseCreateUserMsg
		le BaseLegalEntity
	}
	tests := []struct {
		name   string
		fields fields
		want   sdk.CodeType
	}{
		{"nil pubkey", fields{cm: BaseCreateUserMsg{nil, nil}, le: validEntity}, CodeInvalidPubKey},
		{"invalid entity type", fields{cm: validCreateUser, le: BaseLegalEntity{EntityName: "CH", EntityType: "invalid"}}, CodeInvalidEntity},
		{"empty entity name", fields{cm: validCreateUser, le: BaseLegalEntity{EntityName: "    ", EntityType: EntityClearingHouse}}, CodeInvalidEntity},
		{"self create", fields{cm: BaseCreateUserMsg{Creator: pub.Address(), PubKey: pub}, le: validEntity}, CodeSelfCreate},
		{"ok", fields{cm: validCreateUser, le: validEntity}, sdk.CodeOK},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := CreateAdminMsg{
				BaseCreateUserMsg: tt.fields.cm,
				BaseLegalEntity:   tt.fields.le,
			}
			got := msg.ValidateBasic()
			if got != nil {
				assert.Equal(t, tt.want, got.ABCICode(), got.ABCILog)
			} else {
				assert.Equal(t, tt.want, sdk.CodeOK)
			}
		})
	}
}

func TestBaseFreezeAccountMsg_ValidateBasic(t *testing.T) {
	addr1 := crypto.GenPrivKeyEd25519().PubKey().Address()
	addr2 := crypto.GenPrivKeyEd25519().PubKey().Address()
	type fields struct {
		a crypto.Address
		t crypto.Address
	}
	tests := []struct {
		name   string
		fields fields
		want   sdk.CodeType
	}{
		{"empty msg", fields{}, CodeInvalidAddress},
		{"empty target", fields{a: addr1}, CodeInvalidAddress},
		{"self freeze", fields{a: addr1, t: addr1}, CodeSelfFreeze},
		{"ok", fields{a: addr1, t: addr2}, sdk.CodeOK},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := BaseFreezeAccountMsg{
				Admin:  tt.fields.a,
				Target: tt.fields.t,
			}
			got := msg.ValidateBasic()
			if got != nil {
				assert.Equal(t, tt.want, got.ABCICode(), got.ABCILog)
			} else {
				assert.Equal(t, tt.want, sdk.CodeOK)
			}
		})
	}
}

func TestBaseFreezeAccountMsg_GetSigners(t *testing.T) {
	msg := BaseFreezeAccountMsg{
		Admin:  crypto.GenPrivKeyEd25519().PubKey().Address(),
		Target: crypto.GenPrivKeyEd25519().PubKey().Address(),
	}
	got := msg.GetSigners()
	assert.Equal(t, len(got), 1)
	assert.True(t, bytes.Equal(msg.Admin, got[0]))
}
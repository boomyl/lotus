package message_test

import (
	"bytes"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	datatransfer "github.com/filecoin-project/lotus/datatransfer"
	. "github.com/filecoin-project/lotus/datatransfer/message"
	"github.com/filecoin-project/lotus/datatransfer/testutil"
)

// TODO: get passing to complete
// https://github.com/filecoin-project/go-data-transfer/issues/37
func TestNewRequest(t *testing.T) {
	baseCid := testutil.GenerateCids(1)[0]
	selector := testutil.RandomBytes(100)
	isPull := false
	id := datatransfer.TransferID(rand.Int31())
	voucherIdentifier := "FakeVoucherType"
	voucher := testutil.RandomBytes(100)

	request := NewRequest(id, isPull, voucherIdentifier, voucher, baseCid, selector)
	assert.Equal(t, id, request.TransferID())
	assert.False(t, request.IsCancel())
	assert.False(t, request.IsPull())
	assert.True(t, request.IsRequest())
	assert.Equal(t, baseCid.String(), request.BaseCid().String())
	assert.Equal(t, voucherIdentifier, request.VoucherIdentifier())
	assert.Equal(t, voucher, request.Voucher())
	assert.Equal(t, selector, request.Selector())
}
func TestTransferRequest_MarshalCBOR(t *testing.T) {
	req := NewTestTransferRequest().(*TransferRequest)

	wbuf := new(bytes.Buffer)

	require.NoError(t, req.MarshalCBOR(wbuf))

	require.NotEmpty(t, wbuf)
	assert.Equal(t, 278, wbuf.Len())
}
func TestTransferRequest_UnmarshalCBOR(t *testing.T) {
	// deserializedRequest := .....
	//assert.Equal(t, request.TransferID(), deserializedRequest.TransferID())
	//assert.Equal(t, request.IsCancel(), deserializedRequest.IsCancel())
	//assert.Equal(t, request.IsPull(), deserializedRequest.IsPull())
	//assert.Equal(t, request.IsRequest(), deserializedRequest.IsRequest())
	//assert.Equal(t, request.BaseCid(), deserializedRequest.BaseCid())
	//assert.Equal(t, request.VoucherIdentifier(), deserializedRequest.VoucherIdentifier())
	//assert.Equal(t, request.Voucher(), deserializedRequest.Voucher())
	//assert.Equal(t, request.Selector(), deserializedRequest.Selector())
}

// TODO: get passing to complete
// https://github.com/filecoin-project/go-data-transfer/issues/37
func TestResponses(t *testing.T) {
	accepted := false
	id := datatransfer.TransferID(rand.Int31())
	response := NewResponse(id, accepted)
	assert.Equal(t, response.TransferID(), id)
	assert.False(t, response.Accepted())
	assert.False(t, response.IsRequest())

	// TODO: Uncomment once implemented
	//assert.False(t, cborMessage.IsRequest)
	//cborResponse := cborMessage.Response
	//require.Equal(t, cborResponse.TransferID, int32(id))
	//require.False(t, cborResponse.Accepted

	//deserialized, err := message.NewMessageFromProto(*cborMessage)
	//require.NoError(t, err)
	//
	//deserializedResponse, ok := deserialized.(message.DataTransferResponse)
	//require.True(t, ok)
	//
	//require.Equal(t, deserializedResponse.TransferID(), response.TransferID())
	//require.Equal(t, deserializedResponse.Accepted(), response.Accepted())
	//require.Equal(t, deserializedResponse.IsRequest(), response.IsRequest())
}

// TODO: get passing to complete
// https://github.com/filecoin-project/go-data-transfer/issues/37
func TestRequestCancel(t *testing.T) {
	id := datatransfer.TransferID(rand.Int31())
	request := CancelRequest(id)
	require.Equal(t, request.TransferID(), id)
	require.True(t, request.IsRequest())
	require.True(t, request.IsCancel())

	//cborMessage := request.MarshalCBOR()
	// TODO: Uncomment once implemented
	//require.True(t, cborMessage.IsRequest)
	//pbRequest := cborMessage.Request
	//require.Equal(t, pbRequest.TransferID, int32(id))
	//require.True(t, pbRequest.Cancel)

	//deserialized, err := message.NewMessageFromCBOR(*cborMessage)
	//require.NoError(t, err)
	//
	//deserializedRequest, ok := deserialized.(message.DataTransferRequest)
	//require.True(t, ok)
	//
	//require.Equal(t, deserializedRequest.TransferID(), request.TransferID())
	//require.Equal(t, deserializedRequest.IsCancel(), request.IsCancel())
	//require.Equal(t, deserializedRequest.IsRequest(), request.IsRequest())
}

// TODO: get passing to complete
// https://github.com/filecoin-project/go-data-transfer/issues/37
func TestToNetFromNetEquivalency(t *testing.T) {
	baseCid := testutil.GenerateCids(1)[0]
	selector := testutil.RandomBytes(100)
	isPull := false
	id := datatransfer.TransferID(rand.Int31())
	//accepted := false  // unused?
	voucherIdentifier := "FakeVoucherType"
	voucher := testutil.RandomBytes(100)
	request := NewRequest(id, isPull, voucherIdentifier, voucher, baseCid, selector)
	buf := new(bytes.Buffer)
	err := request.MarshalCBOR(buf)
	require.NoError(t, err)
	//deserialized, err := message.FromNet(buf)
	//require.NoError(t, err)
	//
	//deserializedRequest, ok := deserialized.(message.DataTransferRequest)
	//require.True(t, ok)
	//
	//require.Equal(t, deserializedRequest.TransferID(), request.TransferID())
	//require.Equal(t, deserializedRequest.IsCancel(), request.IsCancel())
	//require.Equal(t, deserializedRequest.IsPull(), request.IsPull())
	//require.Equal(t, deserializedRequest.IsRequest(), request.IsRequest())
	//require.Equal(t, deserializedRequest.BaseCid(), request.BaseCid())
	//require.Equal(t, deserializedRequest.VoucherIdentifier(), request.VoucherIdentifier())
	//require.Equal(t, deserializedRequest.Voucher(), request.Voucher())
	//require.Equal(t, deserializedRequest.Selector(), request.Selector())
	//
	//response := message.NewResponse(id, accepted)
	//err = request.ToNet(buf)
	//require.NoError(t, err)
	//deserialized, err = message.FromNet(buf)
	//require.NoError(t, err)
	//
	//deserializedResponse, ok := deserialized.(message.DataTransferResponse)
	//require.True(t, ok)
	//
	//require.Equal(t, deserializedResponse.TransferID(), response.TransferID())
	//require.Equal(t, deserializedResponse.Accepted(), response.Accepted())
	//require.Equal(t, deserializedResponse.IsRequest(), response.IsRequest())
	//
	//request = message.CancelRequest(id)
	//err = request.ToNet(buf)
	//require.NoError(t, err)
	//deserialized, err = message.FromNet(buf)
	//require.NoError(t, err)
	//
	//deserializedRequest, ok = deserialized.(message.DataTransferRequest)
	//require.True(t, ok)
	//
	//require.Equal(t, deserializedRequest.TransferID(), request.TransferID())
	//require.Equal(t, deserializedRequest.IsCancel(), request.IsCancel())
	//require.Equal(t, deserializedRequest.IsRequest(), request.IsRequest())
}

func NewTestTransferRequest() DataTransferRequest{
	baseCid := testutil.GenerateCids(1)[0]
	selector := testutil.RandomBytes(100)
	isPull := false
	id := datatransfer.TransferID(rand.Int31())
	voucherIdentifier := "FakeVoucherType"
	voucher := testutil.RandomBytes(100)
	return NewRequest(id, isPull, voucherIdentifier, voucher, baseCid, selector)
}
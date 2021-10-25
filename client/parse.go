package client

// import (
// 	"fmt"
// 	"github.com/golang/protobuf/proto"
// 	"github.com/hyperledger/fabric-protos-go/common"
// 	"github.com/hyperledger/fabric-protos-go/peer"
// 	"github.com/pkg/errors"
// 	"strings"

// 	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/rwsetutil"
// )

// // UnmarshalChannelHeader returns a ChannelHeader from bytes
// func GetChannelHeader(bytes []byte) (*common.ChannelHeader, error) {
// 	chdr := &common.ChannelHeader{}
// 	err := proto.Unmarshal(bytes, chdr)
// 	return chdr, errors.Wrap(err, "error unmarshaling ChannelHeader")
// }

// // GetSignatureHeader Get SignatureHeader from bytes
// func GetSignatureHeader(bytes []byte) (*common.SignatureHeader, error) {
// 	sh := &common.SignatureHeader{}
// 	err := proto.Unmarshal(bytes, sh)
// 	return sh, errors.Wrap(err, "error unmarshaling SignatureHeader")
// }

// // GetTransaction Get Transaction from bytes
// func GetTransaction(txBytes []byte) (*peer.Transaction, error) {
// 	tx := &peer.Transaction{}
// 	err := proto.Unmarshal(txBytes, tx)
// 	return tx, errors.Wrap(err, "error unmarshaling Transaction")
// }

// func parseTxAction(action *peer.TransactionAction, namespace string) ([]byte, error) {
// 	cc := peer.ChaincodeActionPayload{}
// 	if err := proto.Unmarshal(action.Payload, &cc); err != nil {
// 		return nil, err
// 	}

// 	respPayload := peer.ProposalResponsePayload{}
// 	if err := proto.Unmarshal(cc.Action.ProposalResponsePayload, &respPayload); err != nil {
// 		return nil, err
// 	}

// 	ccAction := peer.ChaincodeAction{}
// 	if err := proto.Unmarshal(respPayload.Extension, &ccAction); err != nil {
// 		return nil, err
// 	}

// 	if len(ccAction.Results) > 0 {
// 		txRWSet := &rwsetutil.TxRwSet{}
// 		if err := txRWSet.FromProtoBytes(ccAction.Results); err != nil {
// 			return nil, err
// 		}

// 		for _, nsRWSet := range txRWSet.NsRwSets {
// 			if nsRWSet.NameSpace == namespace {
// 				for _, wr := range nsRWSet.KvRwSet.Writes {
// 					if !wr.IsDelete && strings.HasPrefix(wr.Key, namespace+":") {
// 						return wr.Value, nil
// 					}
// 					fmt.Println(wr.Key)
// 				}
// 			}
// 		}
// 	}

// 	return nil, errors.New("Can't find valid write key.")
// }

// func ParseTxPayload(payload []byte, namespace string) ([]byte, error) {

// 	obj := common.Payload{}
// 	if err := proto.Unmarshal(payload, &obj); err != nil {
// 		return nil, err
// 	}

// 	// Parsing headers
// 	chlHeader, err := GetChannelHeader(obj.Header.ChannelHeader)
// 	if err != nil {
// 		return nil, err
// 	}

// 	//sigHeader, err := GetSignatureHeader(obj.Header.SignatureHeader)
// 	//if err != nil {
// 	//	return err
// 	//}
// 	//fmt.Println(sigHeader.Creator, sigHeader.Nonce)

// 	if common.HeaderType(chlHeader.Type) == common.HeaderType_ENDORSER_TRANSACTION {
// 		tx, err := GetTransaction(obj.Data)
// 		if err != nil {
// 			return nil, err
// 		}
// 		if len(tx.Actions) > 0 {
// 			return parseTxAction(tx.Actions[0], namespace)
// 		} else {
// 			return nil, errors.New("Not a valid transaction with empty actions")
// 		}
// 	} else {
// 		return nil, errors.New("Not a valid transaction")
// 	}
// }

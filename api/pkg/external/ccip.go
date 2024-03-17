package external

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
)

type CCIPQueryResponse struct {
	Data struct {
		MessageID struct {
			Nodes []interface{} `json:"nodes"`
		} `json:"messageId"`
		TransactionHash struct {
			Nodes []struct {
				TransactionHash     string      `json:"transactionHash"`
				DestTransactionHash *string     `json:"destTransactionHash"`
				OnrampAddress       string      `json:"onrampAddress"`
				OfframpAddress      string      `json:"offrampAddress"`
				CommitStoreAddress  string      `json:"commitStoreAddress"`
				State               interface{} `json:"state"`
				SourceChainID       string      `json:"sourceChainId"`
				SourceNetworkName   string      `json:"sourceNetworkName"`
				Sender              string      `json:"sender"`
				Receiver            string      `json:"receiver"`
				MessageID           string      `json:"messageId"`
				DestChainID         string      `json:"destChainId"`
				DestNetworkName     string      `json:"destNetworkName"`
				BlockTimestamp      string      `json:"blockTimestamp"`
				Data                string      `json:"data"`
				Strict              bool        `json:"strict"`
				Nonce               int         `json:"nonce"`
				GasLimit            string      `json:"gasLimit"`
				SequenceNumber      int         `json:"sequenceNumber"`
				FeeToken            string      `json:"feeToken"`
				FeeTokenAmount      string      `json:"feeTokenAmount"`
				TokenAmounts        []struct {
					Token  string `json:"token"`
					Amount string `json:"amount"`
				} `json:"tokenAmounts"`
			} `json:"nodes"`
		} `json:"transactionHash"`
		DestTransactionHash struct {
			Nodes []interface{} `json:"nodes"`
		} `json:"destTransactionHash"`
	} `json:"data"`
}

type CCIPQueryRequest struct {
	MsgIDOrTxnHash string `json:"msgIdOrTxnHash"`
}

func FetchCCIPMessage(sourceTxnHash string) (*CCIPQueryResponse, error) {
	baseUrl := "https://ccip.chain.link/api/query"

	request := CCIPQueryRequest{
		MsgIDOrTxnHash: sourceTxnHash,
	}
	vars, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	params := url.Values{
		"query":     {"TRANSACTION_SEARCH_QUERY"},
		"variables": {string(vars)},
	}
	reqUrl := baseUrl + "?" + params.Encode()

	resp, err := http.Get(reqUrl)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result CCIPQueryResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	log.Println("CCIP response:", string(body))

	return &result, nil
}

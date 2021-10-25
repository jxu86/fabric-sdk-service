package controller

import (
	// "encoding/json"
	"gas-fabric-service/common/gintool"
	"gas-fabric-service/common/json"
	"gas-fabric-service/common/util"
	"gas-fabric-service/model"
	"github.com/gin-gonic/gin"
	// "github.com/pkg/errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var (
	DefaultChaincodeId    = "gas"
	DefaultInvokeFuncName = "register"
	DefaultQueryFuncName  = "query"

	DefaultInvokePeers []string
	DefaultQueryPeers  []string
)

var BType map[string]string = map[string]string{
	"All":                      "全部业务",
	"LoginRecord":              "登陆记录",
	"RechargeRecord":           "充值记录",
	"IcBuyRecord":              "IC卡购气记录",
	"PayRecord":                "缴费记录",
	"MeterRecord":              "抄表业务",
	"HiddenDanger":             "安检隐患信息",
	"HiddenDangerFix":          "隐患整改",
	"InsuranceRecord":          "购买保险业务",
	"OpenAccount":              "开户业务",
	"ChangeAccount":            "过户业务",
	"BankWithhold":             "银行代扣",
	"CancelBankWithhold":       "取消银行代扣",
	"InsuranceWithhold":        "保险代扣",
	"CancelInsuranceWithhold":  "取消保险代扣",
	"InsuranceBillRecord":      "保险收费单记录",
	"GasCloseRecord":           "物联网关阀记录",
	"GasDiscountsRecord":       "用气优惠业务",
	"GasDiscountsLadderRecord": "用气优惠业务(阶梯调整)",
	"ContractSignAgain":        "合同补签业务",
	"BackupBillRecord":         "备款单业务",
	"FeeBillRecord":            "收费单(IC卡)业务",
	"OtherFeeBillRecord":       "其他收费业务",
	"BillRecord":               "发票业务",
	"WriteCardRecord":          "写卡状态记录业务",
	"ServiceCheckRecord":       "业务审核记录",
	"UserHandleRecord":         "用户操作记录",
	"WorkOrder":                "预约工单",
	"WorkOrderResult":          "工单结果",
	"ExtendInterface":			"扩展业务",
}

func init() {
	if peerStr := os.Getenv("DEFAULT_QUERY_PEERS"); peerStr != "" {
		DefaultQueryPeers = strings.Split(peerStr, ",")
	}

	if peerStr := os.Getenv("DEFAULT_INVOKE_PEERS"); peerStr != "" {
		DefaultInvokePeers = strings.Split(peerStr, ",")
	}

	if ccid := os.Getenv("DEFAULT_CHAINCODE_ID"); ccid != "" {
		DefaultChaincodeId = ccid
	}

	if funcName := os.Getenv("DEFAULT_INVOKE_FUNCNAME"); funcName != "" {
		DefaultInvokeFuncName = funcName
	}

	if funcName := os.Getenv("DEFAULT_QUERY_FUNCNAME"); funcName != "" {
		DefaultQueryFuncName = funcName
	}
}

func (a *ApiController) dispatch(funcName, key, m string) gin.H {
	disp := map[string]func(string, string) gin.H{
		"loginRecord":              a.loginRecord,
		"rechargeRecord":           a.rechargeRecord,
		"payRecord":                a.payRecord,
		"icBuyRecord":              a.icBuyRecord,
		"meterRecord":              a.meterRecord,
		"hiddenDanger":             a.hiddenDanger,
		"hiddenDangerFix":          a.hiddenDangerFix,
		"insuranceRecord":          a.insuranceRecord,
		"openAccount":              a.openAccount,
		"changeAccount":            a.changeAccount,
		"bankWithhold":             a.bankWithhold,
		"cancelBankWithhold":       a.cancelBankWithhold,
		"insuranceWithhold":        a.insuranceWithhold,
		"cancelInsuranceWithhold":  a.cancelInsuranceWithhold,
		"insuranceBillRecord":      a.insuranceBillRecord,
		"gasCloseRecord":           a.gasCloseRecord,
		"gasDiscountsRecord":       a.gasDiscountsRecord,
		"gasDiscountsLadderRecord": a.gasDiscountsLadderRecord,
		"contractSignAgain":        a.contractSignAgain,
		"backupBillRecord":         a.backupBillRecord,
		"feeBillRecord":            a.feeBillRecord,
		"otherFeeBillRecord":       a.otherFeeBillRecord,
		"billRecord":               a.billRecord,
		"writeCardRecord":          a.writeCardRecord,
		"userHandleRecord":         a.userHandleRecord,
		"serviceCheckRecord":       a.serviceCheckRecord,
		"workOrder":                a.workOrder,
		"workOrderResult":          a.workOrderResult,
		"extendInterface":   		a.extendInterface,
	}
	return disp[funcName](key, m)
}

func (a *ApiController) ChaincodeQuery(ctx *gin.Context) {

	var argJson model.ArgJson
	if err := ctx.ShouldBindJSON(&argJson); err != nil {
		gintool.ResultFail(ctx, err)
		return
	}
	logger.Info("argJson:", argJson)

	var ccReq model.ChaincodeCallRequest
	ccReq.ChaincodeId = DefaultChaincodeId
	ccReq.Peers = DefaultInvokePeers

	ccReq.FunctionName = util.FirstUpper(argJson.Event)
	ccReq.Args = append(ccReq.Args, argJson.Args)

	logger.Info("ccReq:", ccReq)
	resp, err := a.ChaincodeService.QueryChaincode(ccReq)
	if err != nil {
		gintool.ResultFail(ctx, "")
	} else {
		gintool.ResultOk(ctx, resp)
	}
}

func (a *ApiController) ChaincodeInvokeGas(ctx *gin.Context) {

	var argJson model.ArgJson
	if err := ctx.ShouldBindJSON(&argJson); err != nil {
		gintool.ResultFail(ctx, err.Error())
		return
	}
	logger.Info("argJson:", argJson)

	var ccReq model.ChaincodeCallRequest
	ccReq.ChaincodeId = DefaultChaincodeId
	ccReq.Peers = DefaultInvokePeers

	ccReq.FunctionName = util.FirstUpper(argJson.Event)
	ccReq.Args = append(ccReq.Args, argJson.Args)

	logger.Info("ccReq:", ccReq)
	resp, err := a.ChaincodeService.InvokeChaincode(ccReq)
	if err != nil {
		gintool.ResultFail(ctx, err.Error())
	} else {
		gintool.ResultOk(ctx, resp)
	}
}

func (a *ApiController) ChaincodeInvokeBatchGas(ctx *gin.Context) {
	type batchParam struct {
		Event string   `json:"event"`
		Args  []string `json:"args"`
	}
	type transaction struct {
		TranId string `json:"tranId"`
	}
	var bParam batchParam
	if err := ctx.ShouldBindJSON(&bParam); err != nil {
		gintool.ResultFail(ctx, err)
		return
	}

	logger.Info("###ChaincodeInvokeBatchGas=>bParam:", bParam)

	retData := make([]gin.H, 0)
	tranIdList := make([]string, 0)
	for _, arg := range bParam.Args {
		var tran transaction
		err := json.Unmarshal([]byte(arg), &tran)
		if err != nil || tran.TranId == "" {
			logger.Info("###no tranId:", err.Error())
			gintool.ResultFail(ctx, err.Error())
			return
		}
		tranIdList = append(tranIdList, tran.TranId)
	}

	event := util.FirstUpper(bParam.Event)
	Args, err := json.Marshal(bParam.Args)
	if err != nil {
		gintool.ResultFail(ctx, err)
		return
	}

	var ccReq model.ChaincodeCallRequest
	ccReq.ChaincodeId = DefaultChaincodeId
	ccReq.FunctionName = DefaultInvokeFuncName
	ccReq.Peers = DefaultInvokePeers

	ccReq.FunctionName = "WriteBatch"

	ccReq.Args = append(ccReq.Args, event)
	ccReq.Args = append(ccReq.Args, string(Args))

	logger.Info("###ccReq:", ccReq)

	resp, err := a.ChaincodeService.InvokeChaincode(ccReq)

	if err != nil {
		gintool.ResultFail(ctx, err.Error())
	} else {
		for _, tranId := range tranIdList {
			retData = append(retData, gin.H{"tranId": tranId, "txid": resp})
		}
		gintool.ResultOk(ctx, retData)
	}

}

func (a *ApiController) TransactionStatus(ctx *gin.Context) {

	txid := ctx.Query("id")
	logger.Info("txid=>", txid)
	err := a.ChaincodeService.QueryTransactionStatusByTxId(txid, DefaultInvokePeers)
	if err != nil {
		gintool.ResultFailData(ctx, "fail", err.Error())
	} else {
		gintool.ResultOk(ctx, "success")
	}
}

func (a *ApiController) TransactionStatusBatch(ctx *gin.Context) {
	logger.Info("##TransactionStatusBatchGas start")
	type batchParam struct {
		Args []string `json:"args"`
	}
	var bParam batchParam
	if err := ctx.ShouldBindJSON(&bParam); err != nil {
		gintool.ResultFail(ctx, err.Error())
		return
	}

	retData := make([]gin.H, 0)
	for _, txid := range bParam.Args {
		logger.Info("##TransactionStatusBatchGas=>txid:", txid)

		err := a.ChaincodeService.QueryTransactionStatusByTxId(txid, DefaultInvokePeers)
		if err != nil {
			logger.Info("err:", err.Error())
			retData = append(retData, gin.H{"txid": txid, "status": "fail"})
		} else {
			retData = append(retData, gin.H{"txid": txid, "status": "success"})
		}

	}
	gintool.ResultOk(ctx, retData)

}

func (a *ApiController) GetTypes(ctx *gin.Context) {

	retData := make([]gin.H, 0)

	// for k, v := range BType {
	// 	retData = append(retData, gin.H{"key": k, "value": v})
	// }

	retData = append(retData, gin.H{"key": "All", "value": "全部业务"})
	retData = append(retData, gin.H{"key": "LoginRecord", "value": "登陆记录"})
	retData = append(retData, gin.H{"key": "RechargeRecord", "value": "充值记录"})
	retData = append(retData, gin.H{"key": "IcBuyRecord", "value": "IC卡购气记录"})
	retData = append(retData, gin.H{"key": "PayRecord", "value": "缴费记录"})
	retData = append(retData, gin.H{"key": "MeterRecord", "value": "抄表业务"})
	retData = append(retData, gin.H{"key": "HiddenDanger", "value": "安检隐患信息"})
	retData = append(retData, gin.H{"key": "HiddenDangerFix", "value": "隐患整改"})
	retData = append(retData, gin.H{"key": "InsuranceRecord", "value": "购买保险业务"})
	retData = append(retData, gin.H{"key": "OpenAccount", "value": "开户业务"})
	retData = append(retData, gin.H{"key": "ChangeAccount", "value": "过户业务"})
	retData = append(retData, gin.H{"key": "BankWithhold", "value": "银行代扣"})
	retData = append(retData, gin.H{"key": "CancelBankWithhold", "value": "取消银行代扣"})
	retData = append(retData, gin.H{"key": "InsuranceWithhold", "value": "保险代扣"})
	retData = append(retData, gin.H{"key": "CancelInsuranceWithhold", "value": "取消保险代扣"})
	retData = append(retData, gin.H{"key": "InsuranceBillRecord", "value": "保险收费单记录"})
	retData = append(retData, gin.H{"key": "GasCloseRecord", "value": "物联网关阀记录"})
	retData = append(retData, gin.H{"key": "GasDiscountsRecord", "value": "用气优惠业务"})
	retData = append(retData, gin.H{"key": "GasDiscountsLadderRecord", "value": "用气优惠业务"})
	retData = append(retData, gin.H{"key": "ContractSignAgain", "value": "合同补签业务"})
	retData = append(retData, gin.H{"key": "BackupBillRecord", "value": "备款单业务"})
	retData = append(retData, gin.H{"key": "FeeBillRecord", "value": "收费单(IC卡)业务"})
	retData = append(retData, gin.H{"key": "OtherFeeBillRecord", "value": "其他收费业务"})
	retData = append(retData, gin.H{"key": "BillRecord", "value": "发票业务"})
	retData = append(retData, gin.H{"key": "WriteCardRecord", "value": "写卡状态记录业务"})
	retData = append(retData, gin.H{"key": "ServiceCheckRecord", "value": "业务审核记录"})
	retData = append(retData, gin.H{"key": "UserHandleRecord", "value": "用户操作记录"})
	retData = append(retData, gin.H{"key": "WorkOrder", "value": "预约工单"})
	retData = append(retData, gin.H{"key": "WorkOrderResult", "value": "工单结果"})
	retData = append(retData, gin.H{"key": "ExtendInterface", "value": "扩展业务"})

	gintool.ResultOk(ctx, retData)

}

func (a *ApiController) TransactionSummary(ctx *gin.Context) {
	totalTranCount := 0
	todayTranCount := 0
	resp, err := a.readFunc("TotalTxCount")
	if err == nil {
		totalTranCount, _ = strconv.Atoi(resp)
	}

	timeStr := time.Now().Format("20060102")
	dateKey := fmt.Sprintf("%s_%s", "DateCount", timeStr)
	resp, err = a.readFunc(dateKey)
	if err == nil {
		todayTranCount, _ = strconv.Atoi(resp)
	}

	gintool.ResultOk(ctx, gin.H{"totalTranCount": totalTranCount, "todayTranCount": todayTranCount})
}

func (a *ApiController) TransactionDetail(ctx *gin.Context) {

	// pageSize, pageNum, totalCount
	type DetailReq struct {
		PageSize int    `json:"pageSize"`
		PageNum  int    `json:"pageNum"`
		Type     string `json:"type"`
	}
	// var itemKey []string
	var detailReq DetailReq
	totalTranCount := 0
	var kKey string
	retData := make([]gin.H, 0)

	if err := ctx.ShouldBindJSON(&detailReq); err != nil {
		gintool.ResultFail(ctx, err)
		return
	}

	if detailReq.Type == "All" {
		resp, err := a.readFunc("TotalTxCount")
		if err != nil {
			gintool.ResultFail(ctx, err)
			return
		}
		totalTranCount, _ = strconv.Atoi(resp)
		// key := fmt.Sprintf("%s_%s", "TxRecord", strconv.Itoa(i))
		kKey = "TxRecord"
	} else {
		tKey := fmt.Sprintf("%s%s", detailReq.Type, "Count")
		logger.Info("tKey:", tKey)
		resp, err := a.readFunc(tKey)
		if err != nil {
			gintool.ResultFail(ctx, err)
			return
		}
		totalTranCount, _ = strconv.Atoi(resp)
		kKey = fmt.Sprintf("%s%s", detailReq.Type, "Key")
	}

	startItem := a.calcStartItem(detailReq.PageNum, detailReq.PageSize, totalTranCount)
	logger.Info("startItem:", startItem)
	if startItem != 0 {
		endItem := 0
		if startItem-detailReq.PageSize > 0 {
			endItem = startItem - detailReq.PageSize
		}
		for i := startItem; i > endItem; i-- {
			key := fmt.Sprintf("%s_%s", kKey, strconv.Itoa(i))
			logger.Info("key:", key)
			yKey, _ := a.readFunc(key)
			logger.Info("yKey:", yKey)
			resData, _ := a.readFunc(yKey)
			logger.Info("resData:", resData)
			if resData != "0" {
				r := a.analyseRes(yKey, resData)
				retData = append(retData, r)
			}
		}
	}

	gintool.ResultOk(ctx, gin.H{"totalCount": totalTranCount, "dataList": retData, "pageSize": detailReq.PageSize, "pageNum": detailReq.PageNum})
}

func (a *ApiController) analyseRes(key, value string) gin.H {
	k := strings.Split(key, "_")
	ykey := k[0]
	logger.Info("ykey:", ykey)
	logger.Info("value:", value)
	// bValue := []byte(value)
	// var lst model.OpenAccountST

	// if err := json.Unmarshal(bValue, &lst); err != nil {
	// 	logger.Info("err:", err)
	// } else {
	// 	logger.Info("lst:", lst)
	// }
	return a.dispatch(ykey, key, value)
}

func (a *ApiController) loginRecord(key, value string) gin.H {
	bValue := []byte(value)
	var lst model.LoginRecordST
	if err := json.Unmarshal(bValue, &lst); err != nil {
		logger.Info("err:", err)
	} else {
		logger.Info("lst:", lst)
	}
	// lst := l.(model.LoginRecordST)
	details := make([]gin.H, 0)
	a.rangeStruct(lst, func(k, v string) {
		details = append(details, gin.H{"key": k, "value": v})
	})
	k := strings.Split(key, "_")
	bKey := BType[util.FirstUpper(k[0])]
	summay := fmt.Sprintf("用户%s在%s登陆成功", lst.CsrId, lst.LoginTime)
	ret := gin.H{
		"tranType": bKey,
		"key":      key,
		"time":     lst.LoginTime,
		"summay":   summay,
		"details":  details,
	}
	logger.Info("ret:", ret)
	return ret
}

func (a *ApiController) rechargeRecord(key, value string) gin.H {
	bValue := []byte(value)
	var lst model.RechargeRecordST
	if err := json.Unmarshal(bValue, &lst); err != nil {
		logger.Info("err:", err)
	} else {
		logger.Info("lst:", lst)
	}
	details := make([]gin.H, 0)
	a.rangeStruct(lst, func(k, v string) {
		details = append(details, gin.H{"key": k, "value": v})
	})
	k := strings.Split(key, "_")
	bKey := BType[util.FirstUpper(k[0])]
	summay := fmt.Sprintf("用户%s在%s充值了%s元", lst.ZhmxYqzh, lst.ZhmxRq, lst.ZhmxJe)
	ret := gin.H{
		"tranType": bKey,
		"key":      key,
		"time":     lst.ZhmxRq,
		"summay":   summay,
		"details":  details,
	}
	logger.Info("ret:", ret)
	return ret
}

func (a *ApiController) payRecord(key, value string) gin.H {
	bValue := []byte(value)
	var lst model.PayRecordST
	if err := json.Unmarshal(bValue, &lst); err != nil {
		logger.Info("err:", err)
	} else {
		logger.Info("lst:", lst)
	}
	details := make([]gin.H, 0)
	a.rangeStruct(lst, func(k, v string) {
		details = append(details, gin.H{"key": k, "value": v})
	})
	k := strings.Split(key, "_")
	bKey := BType[util.FirstUpper(k[0])]
	summay := fmt.Sprintf("用户%s在%s缴费了%s元", lst.JfjlYqzh, lst.JfjlSjjfrq, lst.JfjlSjjfe)
	ret := gin.H{
		"tranType": bKey,
		"key":      key,
		"time":     lst.JfjlSjjfrq,
		"summay":   summay,
		"details":  details,
	}
	logger.Info("ret:", ret)
	return ret
}

func (a *ApiController) icBuyRecord(key, value string) gin.H {
	bValue := []byte(value)
	var lst model.IcBuyRecordST
	if err := json.Unmarshal(bValue, &lst); err != nil {
		logger.Info("err:", err)
	} else {
		logger.Info("lst:", lst)
	}
	details := make([]gin.H, 0)
	a.rangeStruct(lst, func(k, v string) {
		details = append(details, gin.H{"key": k, "value": v})
	})
	k := strings.Split(key, "_")
	bKey := BType[util.FirstUpper(k[0])]
	summay := fmt.Sprintf("用户%s在%s购买气量为:%s共%s元", lst.IcjlYqzh, lst.JfjlSjjfrq, lst.IcjlSl, lst.JfjlSjjfe)
	ret := gin.H{
		"tranType": bKey,
		"key":      key,
		"time":     lst.JfjlSjjfrq,
		"summay":   summay,
		"details":  details,
	}
	logger.Info("ret:", ret)
	return ret
}

func (a *ApiController) meterRecord(key, value string) gin.H {
	bValue := []byte(value)
	var lst model.MeterRecordST
	if err := json.Unmarshal(bValue, &lst); err != nil {
		logger.Info("err:", err)
	} else {
		logger.Info("lst:", lst)
	}
	details := make([]gin.H, 0)
	a.rangeStruct(lst, func(k, v string) {
		details = append(details, gin.H{"key": k, "value": v})
	})
	k := strings.Split(key, "_")
	bKey := BType[util.FirstUpper(k[0])]
	summay := fmt.Sprintf("用户%s在%s抄表度数为%s,抄表员:%s", lst.CbjlYqzh, lst.CbjlSjcbrq, lst.CbjlBcbd, lst.CbjlCbgId)
	ret := gin.H{
		"tranType": bKey,
		"key":      key,
		"time":     lst.CbjlSjcbrq,
		"summay":   summay,
		"details":  details,
	}
	logger.Info("ret:", ret)
	return ret
}

func (a *ApiController) hiddenDanger(key, value string) gin.H {
	bValue := []byte(value)
	var lst model.HiddenDangerST
	if err := json.Unmarshal(bValue, &lst); err != nil {
		logger.Info("err:", err)
	} else {
		logger.Info("lst:", lst)
	}
	details := make([]gin.H, 0)
	a.rangeStruct(lst, func(k, v string) {
		details = append(details, gin.H{"key": k, "value": v})
	})
	k := strings.Split(key, "_")
	bKey := BType[util.FirstUpper(k[0])]
	summay := fmt.Sprintf("用户%s在%s检测出安全隐患", lst.Ajjlyqzh, lst.Ajjllrsj)
	ret := gin.H{
		"tranType": bKey,
		"key":      key,
		"time":     lst.Ajjllrsj,
		"summay":   summay,
		"details":  details,
	}
	logger.Info("ret:", ret)
	return ret
}

func (a *ApiController) hiddenDangerFix(key, value string) gin.H {
	bValue := []byte(value)
	var lst model.HiddenDangerFixST
	if err := json.Unmarshal(bValue, &lst); err != nil {
		logger.Info("err:", err)
	} else {
		logger.Info("lst:", lst)
	}
	details := make([]gin.H, 0)
	a.rangeStruct(lst, func(k, v string) {
		details = append(details, gin.H{"key": k, "value": v})
	})
	k := strings.Split(key, "_")
	bKey := BType[util.FirstUpper(k[0])]
	summay := fmt.Sprintf("用户%s在%s进行安全整改", lst.Ajjlyqzh, lst.Anjianmxclrq)
	ret := gin.H{
		"tranType": bKey,
		"key":      key,
		"time":     lst.Anjianmxclrq,
		"summay":   summay,
		"details":  details,
	}
	logger.Info("ret:", ret)
	return ret
}

func (a *ApiController) insuranceRecord(key, value string) gin.H {
	bValue := []byte(value)
	var lst model.InsuranceRecordST
	if err := json.Unmarshal(bValue, &lst); err != nil {
		logger.Info("err:", err)
	} else {
		logger.Info("lst:", lst)
	}
	details := make([]gin.H, 0)
	a.rangeStruct(lst, func(k, v string) {
		details = append(details, gin.H{"key": k, "value": v})
	})
	k := strings.Split(key, "_")
	bKey := BType[util.FirstUpper(k[0])]
	summay := fmt.Sprintf("用户%s在%s购买了保险，保险金额为:%s,购买人为:%s", lst.Bxsfyqzh, lst.Bxsffsrq, lst.Bxsfje, lst.BxsfMc)
	ret := gin.H{
		"tranType": bKey,
		"key":      key,
		"time":     lst.Bxsffsrq,
		"summay":   summay,
		"details":  details,
	}
	logger.Info("ret:", ret)
	return ret
}

func (a *ApiController) openAccount(key, value string) gin.H {
	bValue := []byte(value)
	var lst model.OpenAccountST
	if err := json.Unmarshal(bValue, &lst); err != nil {
		logger.Info("err:", err)
	} else {
		logger.Info("lst:", lst)
	}
	details := make([]gin.H, 0)
	a.rangeStruct(lst, func(k, v string) {
		details = append(details, gin.H{"key": k, "value": v})
	})
	k := strings.Split(key, "_")
	bKey := BType[util.FirstUpper(k[0])]
	summay := fmt.Sprintf("用户%s在%s开户成功，用户号为:%s", lst.Username, lst.Applytime, lst.Yqzh)
	ret := gin.H{
		"tranType": bKey,
		"key":      key,
		"time":     lst.Applytime,
		"summay":   summay,
		"details":  details,
	}
	logger.Info("ret:", ret)
	return ret
}

func (a *ApiController) changeAccount(key, value string) gin.H {
	bValue := []byte(value)
	var lst model.ChangeAccountST
	if err := json.Unmarshal(bValue, &lst); err != nil {
		logger.Info("err:", err)
	} else {
		logger.Info("lst:", lst)
	}
	details := make([]gin.H, 0)
	a.rangeStruct(lst, func(k, v string) {
		details = append(details, gin.H{"key": k, "value": v})
	})
	k := strings.Split(key, "_")
	bKey := BType[util.FirstUpper(k[0])]
	summay := fmt.Sprintf("用户%s在%s进行过户申请", lst.Yqzh, lst.Applytime)
	ret := gin.H{
		"tranType": bKey,
		"key":      key,
		"time":     lst.Applytime,
		"summay":   summay,
		"details":  details,
	}
	logger.Info("ret:", ret)
	return ret
}

func (a *ApiController) bankWithhold(key, value string) gin.H {
	bValue := []byte(value)
	var lst model.BankWithholdST
	if err := json.Unmarshal(bValue, &lst); err != nil {
		logger.Info("err:", err)
	} else {
		logger.Info("lst:", lst)
	}
	details := make([]gin.H, 0)
	a.rangeStruct(lst, func(k, v string) {
		details = append(details, gin.H{"key": k, "value": v})
	})
	k := strings.Split(key, "_")
	bKey := BType[util.FirstUpper(k[0])]
	summay := fmt.Sprintf("用户%s在%s申请了银行代扣", lst.Yqzh, lst.Applytime)
	ret := gin.H{
		"tranType": bKey,
		"key":      key,
		"time":     lst.Applytime,
		"summay":   summay,
		"details":  details,
	}
	logger.Info("ret:", ret)
	return ret
}

func (a *ApiController) cancelBankWithhold(key, value string) gin.H {
	bValue := []byte(value)
	var lst model.CancelBankWithholdST
	if err := json.Unmarshal(bValue, &lst); err != nil {
		logger.Info("err:", err)
	} else {
		logger.Info("lst:", lst)
	}
	details := make([]gin.H, 0)
	a.rangeStruct(lst, func(k, v string) {
		details = append(details, gin.H{"key": k, "value": v})
	})
	k := strings.Split(key, "_")
	bKey := BType[util.FirstUpper(k[0])]
	summay := fmt.Sprintf("用户%s在%s取消了银行代扣", lst.Yqzh, lst.Checktime)
	ret := gin.H{
		"tranType": bKey,
		"key":      key,
		"time":     lst.Checktime,
		"summay":   summay,
		"details":  details,
	}
	logger.Info("ret:", ret)
	return ret
}

func (a *ApiController) insuranceWithhold(key, value string) gin.H {
	bValue := []byte(value)
	var lst model.InsuranceWithholdST
	if err := json.Unmarshal(bValue, &lst); err != nil {
		logger.Info("err:", err)
	} else {
		logger.Info("lst:", lst)
	}
	details := make([]gin.H, 0)
	a.rangeStruct(lst, func(k, v string) {
		details = append(details, gin.H{"key": k, "value": v})
	})
	k := strings.Split(key, "_")
	bKey := BType[util.FirstUpper(k[0])]
	summay := fmt.Sprintf("用户%s在%s申请了保险代扣", lst.Yqzh, lst.Applytime)
	ret := gin.H{
		"tranType": bKey,
		"key":      key,
		"time":     lst.Applytime,
		"summay":   summay,
		"details":  details,
	}
	logger.Info("ret:", ret)
	return ret
}

func (a *ApiController) cancelInsuranceWithhold(key, value string) gin.H {
	bValue := []byte(value)
	var lst model.CancelInsuranceWithholdST
	if err := json.Unmarshal(bValue, &lst); err != nil {
		logger.Info("err:", err)
	} else {
		logger.Info("lst:", lst)
	}
	details := make([]gin.H, 0)
	a.rangeStruct(lst, func(k, v string) {
		details = append(details, gin.H{"key": k, "value": v})
	})
	k := strings.Split(key, "_")
	bKey := BType[util.FirstUpper(k[0])]
	summay := fmt.Sprintf("用户%s在%s取消了保险代扣", lst.Yqzh, lst.Applytime)
	ret := gin.H{
		"tranType": bKey,
		"key":      key,
		"time":     lst.Applytime,
		"summay":   summay,
		"details":  details,
	}
	logger.Info("ret:", ret)
	return ret
}

func (a *ApiController) insuranceBillRecord(key, value string) gin.H {
	bValue := []byte(value)
	var lst model.InsuranceBillRecordST
	if err := json.Unmarshal(bValue, &lst); err != nil {
		logger.Info("err:", err)
	} else {
		logger.Info("lst:", lst)
	}
	details := make([]gin.H, 0)
	a.rangeStruct(lst, func(k, v string) {
		details = append(details, gin.H{"key": k, "value": v})
	})
	k := strings.Split(key, "_")
	bKey := BType[util.FirstUpper(k[0])]
	summay := fmt.Sprintf("用户%s在%s发送保险消费记录，金额为:%s", lst.Bxsfyqzh, lst.Bxsfjzrq, lst.Bxsfje)
	ret := gin.H{
		"tranType": bKey,
		"key":      key,
		"time":     lst.Bxsfjzrq,
		"summay":   summay,
		"details":  details,
	}
	logger.Info("ret:", ret)
	return ret
}

func (a *ApiController) gasCloseRecord(key, value string) gin.H {
	bValue := []byte(value)
	var lst model.GasCloseRecordST
	if err := json.Unmarshal(bValue, &lst); err != nil {
		logger.Info("err:", err)
	} else {
		logger.Info("lst:", lst)
	}
	details := make([]gin.H, 0)
	a.rangeStruct(lst, func(k, v string) {
		details = append(details, gin.H{"key": k, "value": v})
	})
	k := strings.Split(key, "_")
	bKey := BType[util.FirstUpper(k[0])]
	summay := fmt.Sprintf("关阀记录:用户号:%s,时间:%s,原因:%s", lst.Userid, lst.Shutofftime, lst.Shutoffreason)
	ret := gin.H{
		"tranType": bKey,
		"key":      key,
		"time":     lst.Shutofftime,
		"summay":   summay,
		"details":  details,
	}
	logger.Info("ret:", ret)
	return ret
}

func (a *ApiController) gasDiscountsRecord(key, value string) gin.H {
	bValue := []byte(value)
	var lst model.GasDiscountsRecordST
	if err := json.Unmarshal(bValue, &lst); err != nil {
		logger.Info("err:", err)
	} else {
		logger.Info("lst:", lst)
	}
	details := make([]gin.H, 0)
	a.rangeStruct(lst, func(k, v string) {
		details = append(details, gin.H{"key": k, "value": v})
	})
	k := strings.Split(key, "_")
	bKey := BType[util.FirstUpper(k[0])]
	summay := fmt.Sprintf("用气优惠:用户号:%s,时间:%s,办理人:%s", lst.Yqzh, lst.Applytime, lst.Username)
	ret := gin.H{
		"tranType": bKey,
		"key":      key,
		"time":     lst.Applytime,
		"summay":   summay,
		"details":  details,
	}
	logger.Info("ret:", ret)
	return ret
}

func (a *ApiController) gasDiscountsLadderRecord(key, value string) gin.H {
	bValue := []byte(value)
	var lst model.GasDiscountsLadderRecordST
	if err := json.Unmarshal(bValue, &lst); err != nil {
		logger.Info("err:", err)
	} else {
		logger.Info("lst:", lst)
	}
	details := make([]gin.H, 0)
	a.rangeStruct(lst, func(k, v string) {
		details = append(details, gin.H{"key": k, "value": v})
	})
	k := strings.Split(key, "_")
	bKey := BType[util.FirstUpper(k[0])]
	summay := fmt.Sprintf("用气优惠(阶梯调整):用户号:%s,时间:%s,办理人:%s", lst.Yqzh, lst.Applytime, lst.Username)
	ret := gin.H{
		"tranType": bKey,
		"key":      key,
		"time":     lst.Applytime,
		"summay":   summay,
		"details":  details,
	}
	logger.Info("ret:", ret)
	return ret
}

func (a *ApiController) contractSignAgain(key, value string) gin.H {
	bValue := []byte(value)
	var lst model.ContractSignAgainST
	if err := json.Unmarshal(bValue, &lst); err != nil {
		logger.Info("err:", err)
	} else {
		logger.Info("lst:", lst)
	}
	details := make([]gin.H, 0)
	a.rangeStruct(lst, func(k, v string) {
		details = append(details, gin.H{"key": k, "value": v})
	})
	k := strings.Split(key, "_")
	bKey := BType[util.FirstUpper(k[0])]
	summay := fmt.Sprintf("合同补签:用户号:%s,时间:%s,办理人:%s", lst.Yqzh, lst.Applytime, lst.Username)
	ret := gin.H{
		"tranType": bKey,
		"key":      key,
		"time":     lst.Applytime,
		"summay":   summay,
		"details":  details,
	}
	logger.Info("ret:", ret)
	return ret
}

func (a *ApiController) backupBillRecord(key, value string) gin.H {
	bValue := []byte(value)
	var lst model.BackupBillRecordST
	if err := json.Unmarshal(bValue, &lst); err != nil {
		logger.Info("err:", err)
	} else {
		logger.Info("lst:", lst)
	}
	details := make([]gin.H, 0)
	a.rangeStruct(lst, func(k, v string) {
		details = append(details, gin.H{"key": k, "value": v})
	})
	k := strings.Split(key, "_")
	bKey := BType[util.FirstUpper(k[0])]
	summay := fmt.Sprintf("备款单业务:用户号:%s,时间:%s", lst.Cbjlkhid, lst.Cbjlsjcbrq)
	ret := gin.H{
		"tranType": bKey,
		"key":      key,
		"time":     lst.Cbjlsjcbrq,
		"summay":   summay,
		"details":  details,
	}
	logger.Info("ret:", ret)
	return ret
}

func (a *ApiController) feeBillRecord(key, value string) gin.H {
	bValue := []byte(value)
	var lst model.FeeBillRecordST
	if err := json.Unmarshal(bValue, &lst); err != nil {
		logger.Info("err:", err)
	} else {
		logger.Info("lst:", lst)
	}
	details := make([]gin.H, 0)
	a.rangeStruct(lst, func(k, v string) {
		details = append(details, gin.H{"key": k, "value": v})
	})
	k := strings.Split(key, "_")
	bKey := BType[util.FirstUpper(k[0])]
	summay := fmt.Sprintf("收费单(IC卡)业务:用户号:%s,时间:%s", lst.Fyjlyqzh, lst.Fyjlsjjfrq)
	ret := gin.H{
		"tranType": bKey,
		"key":      key,
		"time":     lst.Fyjlsjjfrq,
		"summay":   summay,
		"details":  details,
	}
	logger.Info("ret:", ret)
	return ret
}

func (a *ApiController) otherFeeBillRecord(key, value string) gin.H {
	bValue := []byte(value)
	var lst model.OtherFeeBillRecordST
	if err := json.Unmarshal(bValue, &lst); err != nil {
		logger.Info("err:", err)
	} else {
		logger.Info("lst:", lst)
	}
	details := make([]gin.H, 0)
	a.rangeStruct(lst, func(k, v string) {
		details = append(details, gin.H{"key": k, "value": v})
	})
	k := strings.Split(key, "_")
	bKey := BType[util.FirstUpper(k[0])]
	summay := fmt.Sprintf("其他收费业务:用户号:%s,时间:%s", lst.Qtsfyqzh, lst.Qtsffsrq)
	ret := gin.H{
		"tranType": bKey,
		"key":      key,
		"time":     lst.Qtsffsrq,
		"summay":   summay,
		"details":  details,
	}
	logger.Info("ret:", ret)
	return ret
}

func (a *ApiController) billRecord(key, value string) gin.H {
	bValue := []byte(value)
	var lst model.BillRecordST
	if err := json.Unmarshal(bValue, &lst); err != nil {
		logger.Info("err:", err)
	} else {
		logger.Info("lst:", lst)
	}
	details := make([]gin.H, 0)
	a.rangeStruct(lst, func(k, v string) {
		details = append(details, gin.H{"key": k, "value": v})
	})
	k := strings.Split(key, "_")
	bKey := BType[util.FirstUpper(k[0])]
	summay := fmt.Sprintf("发票业务:用户号:%s,时间:%s", lst.Fpyqzh, lst.Fpsj)
	ret := gin.H{
		"tranType": bKey,
		"key":      key,
		"time":     lst.Fpsj,
		"summay":   summay,
		"details":  details,
	}
	logger.Info("ret:", ret)
	return ret
}

func (a *ApiController) writeCardRecord(key, value string) gin.H {
	bValue := []byte(value)
	var lst model.WriteCardRecordST
	if err := json.Unmarshal(bValue, &lst); err != nil {
		logger.Info("err:", err)
	} else {
		logger.Info("lst:", lst)
	}
	details := make([]gin.H, 0)
	a.rangeStruct(lst, func(k, v string) {
		details = append(details, gin.H{"key": k, "value": v})
	})
	k := strings.Split(key, "_")
	bKey := BType[util.FirstUpper(k[0])]
	summay := fmt.Sprintf("写卡状态记录业务:用户号:%s,时间:%s", lst.Userno, lst.Currenttime)
	ret := gin.H{
		"tranType": bKey,
		"key":      key,
		"time":     lst.Currenttime,
		"summay":   summay,
		"details":  details,
	}
	logger.Info("ret:", ret)
	return ret
}

func (a *ApiController) userHandleRecord(key, value string) gin.H {
	bValue := []byte(value)
	var lst model.UserHandleRecordST
	if err := json.Unmarshal(bValue, &lst); err != nil {
		logger.Info("err:", err)
	} else {
		logger.Info("lst:", lst)
	}
	details := make([]gin.H, 0)
	a.rangeStruct(lst, func(k, v string) {
		details = append(details, gin.H{"key": k, "value": v})
	})
	k := strings.Split(key, "_")
	bKey := BType[util.FirstUpper(k[0])]
	summay := fmt.Sprintf("用户操作记录:用户号:%s,时间:%s", lst.Yqzh, lst.Time)
	ret := gin.H{
		"tranType": bKey,
		"key":      key,
		"time":     lst.Time,
		"summay":   summay,
		"details":  details,
	}
	logger.Info("ret:", ret)
	return ret
}

func (a *ApiController) serviceCheckRecord(key, value string) gin.H {
	bValue := []byte(value)
	var lst model.ServiceCheckRecordST
	if err := json.Unmarshal(bValue, &lst); err != nil {
		logger.Info("err:", err)
	} else {
		logger.Info("lst:", lst)
	}
	details := make([]gin.H, 0)
	a.rangeStruct(lst, func(k, v string) {
		details = append(details, gin.H{"key": k, "value": v})
	})
	k := strings.Split(key, "_")
	bKey := BType[util.FirstUpper(k[0])]
	summay := fmt.Sprintf("业务审核记录(客服):审核员工号:%s,时间:%s,用户号:%s,业务类型:%s", lst.Checkman, lst.Checktime, lst.Userid, lst.Pbusinesstype)
	ret := gin.H{
		"tranType": bKey,
		"key":      key,
		"time":     lst.Checktime,
		"summay":   summay,
		"details":  details,
	}
	logger.Info("ret:", ret)
	return ret
}

func (a *ApiController) workOrder(key, value string) gin.H {
	bValue := []byte(value)
	var lst model.WorkOrderST
	if err := json.Unmarshal(bValue, &lst); err != nil {
		logger.Info("err:", err)
	} else {
		logger.Info("lst:", lst)
	}
	details := make([]gin.H, 0)
	a.rangeStruct(lst, func(k, v string) {
		details = append(details, gin.H{"key": k, "value": v})
	})
	k := strings.Split(key, "_")
	bKey := BType[util.FirstUpper(k[0])]
	summay := fmt.Sprintf("预约工单:工单编号%s,用户号:%s,", lst.Woid, lst.Userid)
	ret := gin.H{
		"tranType": bKey,
		"key":      key,
		"time":     lst.Anjiantime,
		"summay":   summay,
		"details":  details,
	}
	logger.Info("ret:", ret)
	return ret
}

func (a *ApiController) workOrderResult(key, value string) gin.H {
	bValue := []byte(value)
	var lst model.WorkOrderResultST
	if err := json.Unmarshal(bValue, &lst); err != nil {
		logger.Info("err:", err)
	} else {
		logger.Info("lst:", lst)
	}
	details := make([]gin.H, 0)
	a.rangeStruct(lst, func(k, v string) {
		details = append(details, gin.H{"key": k, "value": v})
	})
	k := strings.Split(key, "_")
	bKey := BType[util.FirstUpper(k[0])]
	summay := fmt.Sprintf("预约工单结果:工单编号%s,用户号:%s,", lst.Woid, lst.Userid)
	ret := gin.H{
		"tranType": bKey,
		"key":      key,
		"time":     lst.Completetime,
		"summay":   summay,
		"details":  details,
	}
	logger.Info("ret:", ret)
	return ret
}

func (a *ApiController) extendInterface(key, value string) gin.H {
	bValue := []byte(value)
	var lst model.ExtendInterfaceST
	if err := json.Unmarshal(bValue, &lst); err != nil {
		logger.Info("err:", err)
	} else {
		logger.Info("lst:", lst)
	}
	details := make([]gin.H, 0)
	a.rangeStruct(lst, func(k, v string) {
		details = append(details, gin.H{"key": k, "value": v})
	})
	k := strings.Split(key, "_")
	bKey := BType[util.FirstUpper(k[0])]
	summay := fmt.Sprintf("拓展业务:用户号:%s,", lst.Userid)
	ret := gin.H{
		"tranType": bKey,
		"key":      key,
		"time":     lst.Businesstime,
		"summay":   summay,
		"details":  details,
	}
	logger.Info("ret:", ret)
	return ret
}

func (a *ApiController) rangeStruct(in interface{}, h func(k, v string)) {
	rType, rVal := reflect.TypeOf(in), reflect.ValueOf(in)
	if rType.Kind() == reflect.Ptr { // 传入的in是指针,需要.Elem()取得指针指向的value
		rType, rVal = rType.Elem(), rVal.Elem()
	}
	if rType.Kind() != reflect.Struct {
		return
	}
	for i := 0; i < rType.NumField(); i++ { // 遍历结构体
		t, f := rType.Field(i), rVal.Field(i)
		// fmt.Printf("%v,%v\n", t, f)
		// 此处可以参照f.String(),f.Int(),f.Float()源码,处理不同类型,我这里统一转成string类型了
		if f.Kind() != reflect.Struct { // 不深入遍历结构体了,有需要自己实现吧
			h(t.Name, fmt.Sprint(f))
		}
	}
}

func (a *ApiController) calcStartItem(pageNum, pageSize, totalCount int) int {
	totalPage := totalCount/pageSize + 1
	if totalPage < pageNum {
		return 0
	} else {
		startCount := totalCount - (pageNum-1)*pageSize
		return startCount
	}

}

func (a *ApiController) readFunc(value string) (string, error) {
	var argJson model.ArgJson

	argJson.Event = "Read"
	argJson.Args = value
	logger.Info("argJson:", argJson)

	var ccReq model.ChaincodeCallRequest
	ccReq.ChaincodeId = DefaultChaincodeId
	ccReq.Peers = DefaultInvokePeers

	ccReq.FunctionName = util.FirstUpper(argJson.Event)
	ccReq.Args = append(ccReq.Args, argJson.Args)

	logger.Info("ccReq:", ccReq)
	resp, err := a.ChaincodeService.QueryChaincode(ccReq)
	logger.Info("resp:", resp)
	if err != nil {
		return "0", nil
	} else {
		var resStr string
		resStr = resp.(string)
		return resStr, nil
	}
}

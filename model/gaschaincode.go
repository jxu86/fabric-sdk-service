package model

type LoginRecordST struct {
	CsrId     string `json:"csrId"` // require
	LoginTime string `json:"loginTime"`
	LoginIp   string `json:"loginIp"`
	LoginType string `json:"loginType"`
}

// 充值记录
type RechargeRecordST struct {
	ZhmxYqzh  string `json:"zhmxYqzh"`  // 用户号
	ZhmxJe    int    `json:"zhmxJe"`    // 金额
	ZhmxRq    string `json:"zhmxRq"`    // 充值时间
	ZhmxFkfs  string `json:"zhmxFkfs"`  // 充值渠道
	TradeNo   string `json:"tradeNo"`   // 流水号
	ZhmxJylsh string `json:"zhmxJylsh"` // 支付交易流水号
	ZhId      string `json:"zhId"`      // 账户ID require
}

// 缴费记录
type PayRecordST struct {
	JfjlYqzh     string `json:"jfjlYqzh"` // require
	JfjlYje      int    `json:"jfjlYje"`
	JfjlZnj      int    `json:"jfjlZnj"`
	JfjlSjjfe    int    `json:"jfjlSjjfe"`
	JfjlSjjfrq   string `json:"jfjlSjjfrq"`
	JfjlSjjkfsId string `json:"jfjlSjjkfsId"`
	JfjlJylsh    string `json:"jfjlJylsh"`
	FyjlBiId     string `json:"fyjlBiId"`
	JfjlZt       string `json:"jfjlZt"`
	YwdxLx       string `json:"ywdxLx"`
	YwdxId       string `json:"ywdxId"`
}

// IC卡购气记录
type IcBuyRecordST struct {
	IcjlYqzh   string `json:"icjlYqzh"`   // 用户号require
	IcjlSl     int    `json:"icjlSl"`     // 气量
	JfjlSjjfe  int    `json:"jfjlSjjfe"`  // 金额
	JfjlSjjfrq string `json:"jfjlSjjfrq"` // 时间
	IcjlJyLsh  string `json:"icjlJyLsh"`  // 流水号
	IcjlKh     string `json:"icjlKh"`     // ic卡号
	IcjlGmCs   int    `json:"icjlGmCs"`   // 购买次数
	RqbId      string `json:"rqbId"`      // 燃气表ID
}

// 抄表记录
type MeterRecordST struct {
	CbjlCbgId   string `json:"cbjlCbgId"`   //抄表员
	CbjlSjcbrq  string `json:"cbjlSjcbrq"`  //抄表时间
	CbjlHzMc    string `json:"cbjlHzMc"`    //户主
	CbjlScbd    int    `json:"cbjlScbd"`    //上次抄表数
	CbjlBcbd    int    `json:"cbjlBcbd"`    //本次抄表数
	CbjlCbQk    string `json:"cbjlCbQk"`    //抄表情况
	CbjlRqbId   string `json:"cbjlRqbId"`   //抄表燃具
	CbjlBcyqlJs int    `json:"cbjlBcyqlJs"` //本次用气量
	CbjlYje     int    `json:"cbjlYje"`     //应收金额
	CbjlYqzh    string `json:"cbjlYqzh"`    //用户号
	CbjlYqdzMs  string `json:"cbjlYqdzMs"`  //用气地址
	CbjlId      string `json:"cbjlId"`      //抄表记录id require
}

// 隐患信息
type HiddenDangerST struct {
	Ajjlajrid      string `json:"ajjlAjrId"`      //	安检员
	Ajjllrsj       string `json:"ajjlLrsj"`       //	安检时间
	Ajjlkhmc       string `json:"ajjlKhMc"`       //	户主
	Ajjltzxx       string `json:"ajjlTzxx"`       //	隐患通知信息
	Ajjltzxxhash   string `json:"ajjlTzxxHash"`   //	隐患通知信息哈希值
	Ajjlzhenggaifs string `json:"ajjlZhenggaiFs"` //	隐患整改方式
	yhzpmc         string `json:"yhzpMc"`         //	隐患照片名称
	Qmzpmc         string `json:"qmzpMc"`      	  //	签名图片名称
	Ajjlyqzh       string `json:"ajjlYqzh"`       //	用户号
	Ajjlyqdzdes    string `json:"ajjlYqdzDes"`    //	用气地址
	Ajjlid         string `json:"ajjlId"`         //	安检记录id require
	Yhzphash       string `json:"yhzpHash"`       //	隐患照片哈希值
	Qmzphash       string `json:"qmzpHash"`       //	签名照片哈希值
}

// 隐患整改
type HiddenDangerFixST struct {
	Ajjlmxid 		string `json:"ajjlMxId"` 		//	安检记录明细id
	Ajjlzgxx 		string `json:"ajjlZgxx"` 		//	隐患整改信息
	AjjlZgxxhash 	string `json:"ajjlZgxxHash"` 	//	隐患整改信息
	Ajjlmxclrq 		string `json:"ajjlMxClRq"` 		//	处理时间
	Ajjlmxyhbm 		string `json:"ajjlMxYhbm"` 		//	隐患编码
	Ajjlmxclrid  	string `json:"ajjlMxClrId"`  	//	处理人
	Ajjlmxgjqk  	string `json:"ajjlMxGjqk"`  	//	跟进情况
	Ajjlmxgjjg  	string `json:"ajjlMxGjjg"`  	//	跟踪结果
	UserId      	string `json:"userId"`      	//	用户号
	Zgzpmc         	string `json:"zgzpMc"`         	//	整改照片名称
	Zgzphash       	string `json:"zgzpHash"`       	//	整改照片哈希值
}

// 购买保险业务
type InsuranceRecordST struct {
	Bxsfje     string `json:"bxsfJe"`     //	保险金额
	Bxsffsrq   string `json:"bxsfFsrq"`   //	购买时间
	Bxsfjzrq   string `json:"bxsfJzrq"`   //	保期
	BxsfMc     string `json:"bxsfMc"`     //	购买人姓名
	Bxsfluruyg string `json:"bxsfLuruYg"` //	购买渠道
	BxsfId     string `json:"bxsfId"`     //	流水号 require
	Bxsfyqzh   string `json:"bxsfYqzh"`   //	用户号
	Bxsfdzms   string `json:"bxsfDzMs"`   //	用气地址
}

// 开户业务
type OpenAccountST struct {
	Yqzh            	string `json:"yqzh"`            //	用户号	require
	Yqdzyqdzmms     	string `json:"yqdzYqdzmMs"`     //	用户地址
	Yqdzkhlx        	string `json:"yqdzKhLx"`        //	客户类型
	Username        	string `json:"username"`        //	联系人
	Phone           	string `json:"phone"`           //	联系电话
	Idcardfronthash 	string `json:"idcardfrontHash"` //	户主身份证正面
	Idcardbackhash  	string `json:"idcardbackHash"`  //	户主身份证反面
	Applytime       	string `json:"applyTime"`       //	申请时间
	Contractfilehash 	string `json:"contractFileHash"` //	合同文件唯一标识
	Houseprovepointhash string `json:"houseprovepointHash"` //	房产证材料哈希值
	Houseprovenamehash 	string `json:"houseprovenameHash"` //	房产证材料哈希值
	Houseprovestamphash string `json:"houseprovestampHash"` //	房产证材料哈希值
	Attachmenthash 		string `json:"attachmentHash"` //	其他证明材料哈希值
	Signaturehash 		string `json:"signatureHash"` //	签名哈希值
}

type ChangeAccountST struct {
	Yqzh            		string `json:"yqzh"`            //	用户号	require
	Yqdzyqdzmms     		string `json:"yqdzYqdzmMs"`     //	用户地址
	Yqdzkhlx        		string `json:"yqdzKhLx"`        //	客户类型
	Username        		string `json:"username"`        //	联系人
	Phone           		string `json:"phone"`           //	联系电话
	Idcardfronthash 		string `json:"idcardfrontHash"` //	新户主身份证正面
	Idcardbackhash  		string `json:"idcardbackHash"`  //	新户主身份证反面
	Applytime       		string `json:"applyTime"`       //	申请时间
	Contractfilehash 		string `json:"contractFileHash"` //	合同文件唯一标识
	Houseprovepointhash 	string `json:"houseprovepointHash"` //	房产证材料哈希值
	Houseprovenamehash 		string `json:"houseprovenameHash"` //	房产证材料哈希值
	Houseprovestamphash 	string `json:"houseprovestampHash"` //	房产证材料哈希值
	Attachmenthash 			string `json:"attachmentHash"` //	其他证明材料哈希值
	Signaturehash 			string `json:"signatureHash"` //	签名哈希值
	Meterreadingprovehash	string `json:"meterreadingproveHash"` //	燃气表读数资料路径和哈希值
}

// 银行代扣
type BankWithholdST struct {
	Yqzh            	string `json:"yqzh"`            //	用户号 require
	Yqdzyqdzmms     	string `json:"yqdzYqdzmMs"`     //	用气地址
	Username        	string `json:"username"`        //	持卡人姓名
	Idcard          	string `json:"idCard"`          //	持卡人身份证
	Phone           	string `json:"phone"`           //	手机号
	Bankcardtype    	string `json:"bankCardType"`    //	银行卡类型
	Bankcardnumber  	string `json:"bankCardNumber"`  //	银行卡号
	Applytime       	string `json:"applyTime"`       //	申请时间
	Idcardfronthash 	string `json:"idcardfrontHash"` //	户主身份证正面
	Idcardbackhash  	string `json:"idcardbackHash"`  //	户主身份证反面
	Contractfilehash 	string `json:"contractFileHash"` //	合同文件唯一标识
	Bankcardhash 		string `json:"bankcardHash"` //	合同文件唯一标识
	Signaturehash 		string `json:"signatureHash"` //	签名哈希值
	Attachmenthash 		string `json:"attachmentHash"` //	其他证明材料哈希值
}

type CancelBankWithholdST struct {
	Yqzh            string `json:"yqzh"`            //	用户号
	Yqdzyqdzmms     string `json:"yqdzYqdzmMs"`     //	用气地址
	Bankcardnumber  string `json:"bankCardNumber"`  //	银行卡号
	Checktime       string `json:"checkTime"`       //	取消时间
	Cancelyhdkreason    string `json:"cancelYhdkReason"`    //	取消原因
	Idcardfronthash 	string `json:"idcardfrontHash"` //	户主身份证正面
	Idcardbackhash  	string `json:"idcardbackHash"`  //	户主身份证反面
	Contractfilehash 	string `json:"contractFileHash"` //	合同文件唯一标识
	Bankcardhash 		string `json:"bankcardHash"` //	合同文件唯一标识
	Signaturehash 		string `json:"signatureHash"` //	签名哈希值
	Attachmenthash 		string `json:"attachmentHash"` //	其他证明材料哈希值
}

// 保险代扣
type InsuranceWithholdST struct {
	Yqzh            string `json:"yqzh"`            //	用户号	require
	Yqdzyqdzmms     string `json:"yqdzYqdzmMs"`     //	用气地址
	Username        string `json:"username"`        //	持卡人姓名
	Idcard          string `json:"idCard"`          //	持卡人身份证
	Phone           string `json:"phone"`           //	手机号
	Bankcardtype    string `json:"bankCardType"`    //	银行卡类型
	Bankcardnumber  string `json:"bankCardNumber"`  //	银行卡号
	Applytime       string `json:"applyTime"`       //	签约时间
	Idcardfront     string `json:"idCardFront"`     //	身份证正面
	Idcardback      string `json:"idCardBack"`      //	身份证反面
	Bxdkprice       string `json:"bxdkPrice"`       //	保险代扣金额
	Idcardfronthash 	string `json:"idcardfrontHash"` //	户主身份证正面
	Idcardbackhash  	string `json:"idcardbackHash"`  //	户主身份证反面
	Contractfilehash 	string `json:"contractFileHash"` //	合同文件唯一标识
	Bankcardhash 		string `json:"bankcardHash"` //	合同文件唯一标识
	Signaturehash 		string `json:"signatureHash"` //	签名哈希值
	Attachmenthash 		string `json:"attachmentHash"` //	其他证明材料哈希值
}

// 保险代扣
type CancelInsuranceWithholdST struct {
	Yqzh            	string `json:"yqzh"`            //	用户号	require
	Yqdzyqdzmms     	string `json:"yqdzYqdzmMs"`     //	用气地址
	Username        	string `json:"username"`        //	持卡人姓名
	Idcard          	string `json:"idCard"`          //	持卡人身份证
	Phone           	string `json:"phone"`           //	手机号
	Bankcardtype    	string `json:"bankCardType"`    //	银行卡类型
	Bankcardnumber  	string `json:"bankCardNumber"`  //	银行卡号
	Applytime       	string `json:"applyTime"`       //	申请时间
	Bxdkprice       	string `json:"bxdkPrice"`       //	保险代扣金额
	CheckTime       	string `json:"checkTime"`       //	取消时间
	Cancelbxdkreason    string `json:"cancelBxdkReason"`    //	取消原因
	Idcardfronthash 	string `json:"idcardfrontHash"` //	户主身份证正面
	Idcardbackhash  	string `json:"idcardbackHash"`  //	户主身份证反面
	Contractfilehash 	string `json:"contractFileHash"` //	合同文件唯一标识
	Bankcardhash 		string `json:"bankcardHash"` //	合同文件唯一标识
	Signaturehash 		string `json:"signatureHash"` //	签名哈希值
	Attachmenthash 		string `json:"attachmentHash"` //	其他证明材料哈希值
}

// 保险消费单记录
type InsuranceBillRecordST struct {
	Bxsfyqzh string `json:"bxsfYqzh"` //	用户号
	Bxsfdzms string `json:"bxsfDzMs"` //	用气地址
	Khlx     string `json:"khLx"`     //	客户类型
	BxsfMc   string `json:"bxsfMc"`   //	客户名称
	Bxsfjzrq string `json:"bxsfJzrq"` //	截止日期
	Bxsffsrq string `json:"bxsfFsrq"` //	费用发生日期
	Bxsfje   int    `json:"bxsfJe"`   //	保额
	Bxsfzt   string `json:"bxsfZt"`   //	状态
	Bxsfid   string `json:"bxsfId"`   //	保险收费记录id require

}

// 物联网关阀记录
type GasCloseRecordST struct {
	Userid        string `json:"userId"`        //	用户号
	Shutofftime   string `json:"shutOffTime"`   //	关阀时间
	Shutoffreason string `json:"shutOffReason"` //	关阀原因
	Shutoffsmsmsg string `json:"shutOffSmsMsg"` //	关阀短信内容
	Yqdzyqdzmms   string `json:"yqdzYqdzmMs"`   //	用气地址
	Csrid         string `json:"csrId"`         //	员工id
	Csrname       string `json:"csrName"`       //	员工姓名

}

// 用气优惠业务
type GasDiscountsRecordST struct {
	Yqzh             	string `json:"yqzh"`             //	用户号
	Yqdzyqdzmms      	string `json:"yqdzYqdzmMs"`      //	用气地址
	Population       	string `json:"population"`       //	家庭人口数
	Username         	string `json:"username"`         //	办理人
	Phone            	string `json:"phone"`            //	手机号
	Idcardfronthash 	string `json:"idcardfrontHash"` //	户主身份证正面
	Idcardbackhash  	string `json:"idcardbackHash"`  //	户主身份证反面
	Liveprovehash      	string `json:"liveproveHash"`      //	居住证材料
	Discountprovehash 	string `json:"discountproveHash"` //	优惠证明材料
	Signaturehash 		string `json:"signatureHash"` //	签名哈希值
	Applytime        	string `json:"applyTime"`        //	申请时间

}

// 用气优惠业务(阶梯调整)
type GasDiscountsLadderRecordST struct {
	Yqzh             string `json:"yqzh"`             //	用户号
	Yqdzyqdzmms      string `json:"yqdzYqdzmMs"`      //	用气地址
	Population       string `json:"population"`       //	家庭人口数
	Username         string `json:"username"`         //	办理人
	Phone            string `json:"phone"`            //	手机号
	Discountsfilecid string `json:"discountsFileCid"` //	居住证材料
	Familyfilecid    string `json:"familyFileCid"`    //	家庭成员身份证证明材料
	Applytime        string `json:"applyTime"`        //	办理时间
	Idcardfronthash 	string `json:"idcardfrontHash"` //	户主身份证正面
	Idcardbackhash  	string `json:"idcardbackHash"`  //	户主身份证反面
	Liveprovehash      	string `json:"liveproveHash"`      //	居住证材料
	Discountprovehash 	string `json:"discountproveHash"` //	优惠证明材料
	Signaturehash 		string `json:"signatureHash"` //	签名哈希值
	Contractfilehash 	string `json:"contractFileHash"` //	合同文件唯一标识

}

// 合同补签业务
type ContractSignAgainST struct {
	Yqzh         string `json:"yqzh"`         //	用户号
	Yqdzyqdzmms  string `json:"yqdzYqdzmMs"`  //	用气地址
	Username     string `json:"username"`     //	户主姓名
	Idcard       string `json:"idCard"`       //	身份证号
	Sex          string `json:"sex"`          //	性别
	Phone        string `json:"phone"`        //	户主手机号
	Applytime    string `json:"applyTime"`    //	办理时间
	Idcardfronthash 	string `json:"idcardfrontHash"` //	户主身份证正面
	Idcardbackhash  	string `json:"idcardbackHash"`  //	户主身份证反面
	Houseprovepointhash string `json:"houseprovepointHash"` //	房产证材料哈希值
	Signaturehash 		string `json:"signatureHash"` //	签名哈希值
	Contractfilehash 	string `json:"contractFileHash"` //	合同文件唯一标识

}

// 备款单业务
type BackupBillRecordST struct {
	Rhrhlxid   string `json:"rhRhlxId"`   //	计划类型
	Cbjlsjcbrq string `json:"cbjlSjcbrq"` //	抄表时间
	Cbjlcbgid  string `json:"cbjlCbgId"`  //	抄表工
	Rhid       string `json:"rhId"`       //	计划编号
	Cbjlzt     string `json:"cbjlZt"`     //	状态
	Cbjlkhid   string `json:"cbjlKhId"`   //	客户号
	Cbjlyqdzms string `json:"cbjlYqdzMs"` //	用气地址
	Cbjlbcbd   int    `json:"cbjlBcbd"`   //	本期读数
	Fyjldj     int    `json:"fyjlDj"`     //	燃气单价
	Cbjlscbd   int    `json:"cbjlScbd"`   //	上期读数
	Bxsfjzrq   string `json:"bxsfJzrq"`   //	保险截止日期
	Fyjlsl     int    `json:"fyjlSl"`     //	实用气量
	Fyjlznj    int    `json:"fyjlZnj"`    //	气费滞纳金
	Fyjlyje    int    `json:"fyjlYje"`    //	本期燃气费
	Fyjljfjzrq string `json:"fyjlJfjzrq"` //	缴费期限
	Fyjlyjehj  int    `json:"fyjlYjeHj"`  //	应缴费用合计
	Cbjlid     string `json:"cbjlId"`     //	抄表记录id
}

// 收费单(IC卡)业务
type FeeBillRecordST struct {
	Rqbkh        string `json:"rqbKh"`        //	卡号
	Fyjlyqzh     string `json:"fyjlYqzh"`     //	用户号
	fyjlyqdzms   string `json:"fyjlYqdzMs"`   //	用气地址
	Yqdzfgsid    string `json:"yqdzFgsId"`    //	所属分公司
	Orderamount  int    `json:"orderAmount"`  //	最后购气量
	Ordertime    string `json:"orderTime"`    //	最后购气时间
	Rqbljgql     int    `json:"rqbLjgql"`     //	累积购气量
	Fyjlhzmc     string `json:"fyjlHzMc"`     //	用户姓名
	Gmsl         int    `json:"gmsl"`         //	购气量
	Fyjlfylxid   string `json:"fyjlFylxId"`   //	费用类别
	Fyjldj       int    `json:"fyjlDj"`       //	单价
	Fyjlsl       int    `json:"fyjlSl"`       //	数量
	Fyje         int    `json:"fyje"`         //	金额
	Fyjlyhje     int    `json:"fyjlYhje"`     //	优惠金额
	Fyjlyje      int    `json:"fyjlYje"`      //	应缴金额
	Gqje         int    `json:"gqje"`         //	购气金额
	Fyjlsjjkfsid string `json:"fyjlSjjkfsId"` //	付款方式
	Fyjlsfyg     string `json:"fyjlSfyg"`     //	员工号
	Fyjlsjjfrq   string `json:"fyjlSjjfrq"`   //	充值时间
	Fyjljylsh    string `json:"fyjlJylsh"`    //	交易流水号
}

// 其他收费业务
type OtherFeeBillRecordST struct {
	Qtsfyqzh string `json:"qtsfYqzh"` //	用户号
	Qtsfdzms string `json:"qtsfDzMs"` //	用气地址
	Qtsflx   string `json:"qtsfLx"`   //	收费项目
	Qtsfdj   int    `json:"qtsfDj"`   //	单价
	Qtsfsl   int    `json:"qtsfSl"`   //	数量
	Qtsfje   int    `json:"qtsfJe"`   //	金额
	Qtsffsrq string `json:"qtsfFsrq"` //	发生日期
	Qtsfzt   string `json:"qtsfZt"`   //	状态
	Qtsfid   string `json:"qtsfId"`   //	其他收费记录id

}

// 发票业务
type BillRecordST struct {
	Fpyqzh     string `json:"fpYqzh"`     //	用户号
	Fqyhyqdz   string `json:"fqYhYqdz"`   //	用气地址
	Fpjylsh    string `json:"fpJyLsh"`    //	交易流水号
	Fpid       string `json:"fpId"`       //	发票号
	Fpsj       string `json:"fpSj"`       //	开票时间
	Fpjflxid   string `json:"fpJflxId"`   //	费用类型
	Fpsjjkfsid string `json:"fpSjjkfsId"` //	付款方式
	Fpjyjfrq   string `json:"fpJyJfrq"`   //	缴费日期
	Fpssje     int    `json:"fpSsJe"`     //	应缴金额
	Fphjjehj   int    `json:"fpHjJehj"`   //	滞纳金
	Fphjznjhj  int    `json:"fpHjZnjhj"`  //	实缴金额

}

type WriteCardRecordST struct {
	Jystatus     string `json:"jyStatus"`     //	IC卡充值写卡状态
	Currenttime  string `json:"currentTime"`  //	IC卡充值时间
	Fyjlsl       int    `json:"fyjlSl"`       //	充值方数
	Userno       string `json:"userNo"`       //	用户号
	Pcno         string `json:"pcNo"`         //	终端号
	Fyjlsjjkfsid string `json:"fyjlSjjkfsId"` //	支付方式
	Fyjljylsh    string `json:"fyjlJylsh"`    //	流水号
	Banktype     string `json:"bankType"`     //	银行卡类型
	Bankno       string `json:"bankNo"`       //	银行卡号
}

// 用户操作记录
type UserHandleRecordST struct {
	Yqzh       string `json:"yqzh"`       //	用户号	 require
	Handletype string `json:"handleType"` //	审核业务类型，填写对应的业务event字段
	Time       string `json:"time"`       //	审核时间
	Status     string `json:"status"`     //	状态

}

// 业务审核记录(客服)
type ServiceCheckRecordST struct {
	Checktime     string `json:"checkTime"`     //	审核时间
	Csrname       string `json:"csrName"`       //	审核操作员姓名
	Checkman      string `json:"checkMan"`      //	审核员工号
	Pbusinesstype string `json:"pBusinessType"` //	审核业务类型
	Applystatus   string `json:"applyStatus"`   //	审核状态
	Reason        string `json:"reason"`        //	审核不通过原因
	Loginip       string `json:"loginIp"`       //	审核计算机ip
	Userid        string `json:"userId"`        //	用户号
	Address       string `json:"address"`       //	用气地址
	Serialnumber  string `json:"serialNumber"`  //	交易流水号
}

// 预约工单
type WorkOrderST struct {
	Woid         string `json:"woId"`         //	工单编号
	Userid       string `json:"userId"`       //	用户号
	Address      string `json:"address"`      //	用气地址
	Name         string `json:"name"`         //	联系人姓名
	Homephone    string `json:"homePhone"`    //	联系电话
	Secondtypeid string `json:"secondTypeId"` //	工单类型
	Thirdtypeid  string `json:"thirdTypeId"`  //	工单种类
	Beizhu       string `json:"beizhu"`       //	备注
	Anjiantime   string `json:"anJianTime"`   //	上次安检日期
	Anjianresult string `json:"anJianResult"` //	上次安检结果
	Pretime      string `json:"preTime"`      //	预约时间
}

type WorkOrderResultST struct {
	Userid       string `json:"userId"`       //	用户号
	Woid         string `json:"woId"`         //	工单编号
	Address      string `json:"address"`      //	用气地址
	Result       string `json:"result"`       //	完成情况
	Completetime string `json:"completeTime"` //	完成时间
	Content      string `json:"content"`      //	内容
}

//
type ExtendInterfaceST struct {
	TranId         string `json:"tranId"`         //	交易id
	Userid         string `json:"userId"`         //	用户号
	Address        string `json:"address"`        //	用气地址
	Username       string `json:"username"`       //	户主姓名
	Phone          string `json:"phone"`          //	联系电话
	Idcard         string `json:"idCard"`         //	户主身份证
	Bankcardnumber string `json:"bankCardNumber"` //	银行卡号
	Serialnumber   string `json:"serialNumber"`   //	流水号
	Businesstime   string `json:"businessTime"`   //	业务时间
	Businesstype   string `json:"businessType"`   //	业务类型
	Amount         int    `json:"amount"`         //	金额
	Status         string `json:"status"`         //	状态
	Filecid        string `json:"fileCid"`        //	文件列表
	Remarks        string `json:"remarks"`        //	备注
	Comefrom       string `json:"comeFrom"`       //	来源
	Csrid          string `json:"csrId"`          //	员工ID
	Csrname        string `json:"csrName"`        //	员工姓名
	Extend1        string `json:"extend1"`        //	预留字段1
	Extend2        string `json:"extend2"`        //	预留字段2
	Extend3        string `json:"extend3"`        //	预留字段3
	Extend4        string `json:"extend4"`        //	预留字段4
	Extend5        string `json:"extend5"`        //	预留字段5
	Extend6        string `json:"extend6"`        //	预留字段6
	Extend7        string `json:"extend7"`        //	预留字段7
	Extend8        string `json:"extend8"`        //	预留字段8
	Extend9        string `json:"extend9"`        //	预留字段9
	Extend10       string `json:"extend10"`       //	预留字段10
	Extend11       string `json:"extend11"`       //	预留字段11
	Extend12       string `json:"extend12"`       //	预留字段12
	Extend13       string `json:"extend13"`       //	预留字段13
	Extend14       string `json:"extend14"`       //	预留字段14
	Extend15       string `json:"extend15"`       //	预留字段15
	Extend16       string `json:"extend16"`       //	预留字段16
	Extend17       string `json:"extend17"`       //	预留字段17
	Extend18       string `json:"extend18"`       //	预留字段18
	Extend19       string `json:"extend19"`       //	预留字段19
	Extend20       string `json:"extend20"`       //	预留字段20
}
package args

type ContactArg struct {
	UserId int64 `json:"userid" form:"userid"`
	DstId  int64 `json:"dstid" form:"dstid"`
}

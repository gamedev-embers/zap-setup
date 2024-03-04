package tencentcls

import (
	"log"

	cls "github.com/tencentcloud/tencentcloud-cls-sdk-go"
)

type Callback struct{}

func (cb *Callback) Success(rs *cls.Result) {
	// attemptList := result.GetReservedAttempts()
	// for _, attempt := range attemptList {
	// 	fmt.Printf("%+v \n", attempt)
	// }
}

func (cb *Callback) Fail(rs *cls.Result) {
	if errCode := rs.GetErrorCode(); errCode != "" {
		log.Printf("auditlog pusher failed. errCode=%s", errCode)
	}
	// log.Println(result.IsSuccessful())
	// log.Println(result.GetErrorCode())
	// log.Println(result.GetErrorMessage())
	// log.Println(result.GetReservedAttempts())
	// log.Println(result.GetRequestId())
	// log.Println(result.GetTimeStampMs())
}

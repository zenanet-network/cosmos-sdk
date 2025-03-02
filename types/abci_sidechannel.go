package types

// tendermint에서 추가해야하는것

// RequestBeginSideBlock는 사이드 체인 블록 시작 요청을 나타냅니다.
type RequestBeginSideBlock struct {
	Hash []byte
	// 기타 필요한 필드
}

// ResponseBeginSideBlock는 사이드 체인 블록 시작 응답을 나타냅니다.
type ResponseBeginSideBlock struct {
	// 필요한 필드
}

// SideTxResultType은 사이드 트랜잭션 결과 유형을 나타냅니다.
type SideTxResultType int32

const (
	// SideTxResultType_Unknown은 알 수 없는 결과를 나타냅니다.
	SideTxResultType_Unknown SideTxResultType = 0
	// SideTxResultType_Skip은 제안을 건너뛰어야 함을 나타냅니다.
	SideTxResultType_Skip SideTxResultType = 1
	// SideTxResultType_Accept는 제안을 수락해야 함을 나타냅니다.
	SideTxResultType_Accept SideTxResultType = 2
)

// RequestDeliverSideTx는 사이드 트랜잭션 전달 요청을 나타냅니다.
type RequestDeliverSideTx struct {
	Tx []byte
	// 기타 필요한 필드
}

// ResponseDeliverSideTx는 사이드 트랜잭션 전달 응답을 나타냅니다.
type ResponseDeliverSideTx struct {
	Result    SideTxResultType
	Code      uint32
	Codespace string
	// 기타 필요한 필드
}

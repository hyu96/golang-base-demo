package errormsg

import (
	comm "github.com/huydq/proto/gen-go/common/v2"
	"golang.org/x/text/language"
)

type MsgError map[string]string // language:error_msg

const DEFAULT_MESSAGE = "Lỗi hệ thống, xin hãy thử lại!"

var (
	VI           = language.Vietnamese.String()
	EN           = language.English.String()
	DEFAULT_LANG = VI
)

var mapMessError = map[comm.Code]MsgError{
	comm.Code_ERR_DATA_NOT_FOUND: {
		VI: "Không tìm thấy kết quả",
		EN: "Not found",
	},
	comm.Code_ERR_AUTH_TOKEN_INVALID: {
		VI: "Lỗi xác thực",
		EN: "Unauthenticated",
	},
	comm.Code_ERR_PHONE_INVALID: {
		VI: "Số điện thoại không hợp lệ",
		EN: "Phone number invalid",
	},
	comm.Code_ERR_BIRTHDAY_INVALID: {
		VI: "Ngày sinh không hợp lệ",
		EN: "Birthday invalid",
	},
	comm.Code_ERR_FULLNAME_EMPTY: {
		VI: "Họ tên không hợp lệ",
		EN: "Full name is empty",
	},
	comm.Code_ERR_GENDER_EMPTY: {
		VI: "Giới tính không hợp lệ",
		EN: "Gender is empty",
	},
	comm.Code_ERR_INTERNAL: {
		VI: "Lỗi server",
		EN: "server error",
	},
	comm.Code_ERR_OTP_RESEND_LOCKED: {
		VI: "Bạn cần đợi hết thời gian hiệu lực để gửi lại OTP",
		EN: "Cannot resend OTP",
	},
	comm.Code_ERR_OTP_LOCKED: {
		VI: "Tài khoản của bạn bị tạm khóa vì nhập sai mã OTP quá số lần quy định. Vui lòng thử lại sau 24h.",
		EN: "OTP feature is locked",
	},
	comm.Code_ERR_OTP_EXPIRED: {
		VI: "OTP hết hạn sử dụng",
		EN: "OTP feature is locked",
	},
}

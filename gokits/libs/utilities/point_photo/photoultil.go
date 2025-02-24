package pointphoto

import (
	"context"
	"strings"

	"github.com/spf13/viper"

	iconst "github.com/huydq/gokits/constants"
	aes "github.com/huydq/gokits/libs/crypto/mysql_aes"
)

func GetPhoto(ctx context.Context, imageID, updateTimestamp string) string {
	return strings.Join(
		[]string{
			viper.GetString("DOMAIN_ROOT"),
			ctx.Value(iconst.KAuthWorkingSiteId).(string),
			"gup2start/rest/photoReader/1.0.0",
			aes.AESEncryptWithHex(imageID, viper.GetString("AES_KEY_1_0_0")),
			updateTimestamp,
		},
		"/")
}

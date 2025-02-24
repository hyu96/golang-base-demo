package imagereader

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/viper"

	aes "github.com/huydq/gokits/libs/crypto/mysql_aes"
)

type ImageReader struct {
	DOMAIN_ROOT   string
	AES_KEY_1_0_0 string
	// APP_ROOT       string
	// DATA_ROOT_PATH string
}

var imageReader *ImageReader
var allowedImageExtensions = []string{".jpg", ".jpeg", ".png", ".gif"}

func NewImageReader() *ImageReader {
	if imageReader != nil {
		return imageReader
	}

	imageReader = &ImageReader{
		DOMAIN_ROOT:   viper.GetString("DOMAIN_ROOT"),
		AES_KEY_1_0_0: viper.GetString("AES_KEY_1_0_0"),
		// APP_ROOT:       viper.GetString("APP_ROOT"),
		// DATA_ROOT_PATH: viper.GetString("DATA_ROOT_PATH"),
	}

	return imageReader
}

func GetImageReader() *ImageReader {
	return imageReader
}

func (r *ImageReader) SqlReader(workingSiteID, alias string) string {
	d := r.DOMAIN_ROOT
	k := r.AES_KEY_1_0_0
	return "IF(" + alias + ".id,CONCAT('" + d + "/" + workingSiteID + "/gup2start/rest/photoReader/1.0.0/',HEX(AES_ENCRYPT(" + alias + ".id,'" + k + "')),'/',UNIX_TIMESTAMP(COALESCE(" + alias + ".update_time," + alias + ".create_time))),'')"
}

func (r *ImageReader) SqlPath(alias string, field string, hasTime bool, workingSiteId int) string {
	d := r.DOMAIN_ROOT
	k := r.AES_KEY_1_0_0

	// Check if an alias is provided and add a dot if it is
	if alias != "" {
		alias += "."
	}

	// Determine the timestamp part of the query based on the hasTime parameter
	var ts string
	if hasTime {
		ts = ",'/',UNIX_TIMESTAMP(COALESCE(" + alias + "update_time," + alias + "create_time))"
	} else {
		ts = ""
	}

	// Construct the SQL query string
	result := fmt.Sprintf("IF(LENGTH(%s%s),CONCAT('%s/%d/gup2start/rest/photoReader/1.0.0/',HEX(AES_ENCRYPT(%s%s,'%s'))%s),'')",
		alias, field, d, workingSiteId, alias, field, k, ts)

	return result
}

func (r *ImageReader) PlatformImageReaderEncode(path, prefixPath string, createTime, updateTime time.Time) string {
	if len(path) > 4 && path[0:4] == "http" {
		return path
	}
	d := r.DOMAIN_ROOT
	k := r.AES_KEY_1_0_0
	hex := aes.AESEncryptWithHex(path, k)
	encryptedTimestamp := createTime.Unix() + updateTime.Unix()
	return fmt.Sprintf("%s/%s/%s/%d", d, prefixPath, hex, encryptedTimestamp)
}

func (r *ImageReader) PlatformImageReaderDecode(encryptedStr string) (string, bool) {
	path := aes.AESDecryptWithHex(encryptedStr, r.AES_KEY_1_0_0)
	if path == "" {
		return path, false
	}
	return path, true
}

// func (r *ImageReader) SqlPath(alias string, field string, hasTime bool) string {
// 	d := r.DOMAIN_ROOT
// 	k := r.AES_KEY_1_0_0
// 	if alias != "" {
// 		alias += "."
// 	}
// 	ts := ""
// 	if hasTime {
// 		ts = ",'/',UNIX_TIMESTAMP(COALESCE(" + alias + "update_time," + alias + "create_time))"
// 	}
// 	return "IF(LENGTH(" + alias + field + "),CONCAT('" + d + "/{" + "$workingSite['id']}/gup2start/rest/photoReader/1.0.0/',HEX(AES_ENCRYPT(" + alias + field + ",'" + k + "'))" + ts + "),'')"
// }

// func (r *ImageReader) EncodeReader(table string, key string, field string, id int, time int64) string {
// 	data := []interface{}{table, key, field, id}
// 	b, err := json.Marshal(data)
// 	if err != nil {
// 		panic(err)
// 	}
// 	block, err := aes.NewCipher([]byte(r.AES_KEY_1_0_0))
// 	if err != nil {
// 		panic(err)
// 	}
// 	ciphertext := make([]byte, len(b))
// 	aesEncrypter := cipher.NewCBCEncrypter(block, make([]byte, aes.BlockSize))
// 	aesEncrypter.CryptBlocks(ciphertext, b)
// 	encodedCiphertext := base64.StdEncoding.EncodeToString(ciphertext)
// 	return r.DOMAIN_ROOT + "/{" + "$workingSite['id']}/gup2start/rest/photoReader/1.0.0/0/" + encodedCiphertext + "/" + strconv.FormatInt(time, 10)
// }

// func (r *ImageReader) decodeReader(data string, showNoImage bool) string {
// 	var table, keyname, fieldname string
// 	var id int
// 	var path string
// 	b, err := base64.StdEncoding.DecodeString(data)
// 	if err != nil {
// 		panic(err)
// 	}
// 	block, err := aes.NewCipher([]byte(r.AES_KEY_1_0_0))
// 	if err != nil {
// 		panic(err)
// 	}
// 	aesDecrypter := cipher.NewCBCDecrypter(block, make([]byte, aes.BlockSize))
// 	aesDecrypter.CryptBlocks(b, b)
// 	err = json.Unmarshal(b, &[]interface{}{&table, &keyname, &fieldname, &id})
// 	if err != nil {
// 		if showNoImage {
// 			path = r.APP_ROOT + "/resource/common/images/no-image.png"
// 		}
// 		return path
// 	}
// 	if !strings.ContainsAny(table+keyname+fieldname, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_") {
// 		if showNoImage {
// 			path = r.APP_ROOT + "/resource/common/images/no-image.png"
// 		}
// 		return path
// 	}
// 	sql := "SELECT " + fieldname + " FROM " + table + " WHERE " + keyname + "=?"
// 	row := database.QueryRow(sql, id)
// 	var val string
// 	err = row.Scan(&val)
// 	if err != nil {
// 		if showNoImage {
// 			path = r.APP_ROOT + "/resource/common/images/no-image.png"
// 		}
// 		return path
// 	}
// 	if strings.HasPrefix(val, "http") {
// 		// Redirect to external URL
// 		// Note: this requires access to the ResponseWriter object
// 		// w.Header().Set("Location", val)
// 		// w.WriteHeader(http.StatusFound)
// 		return ""
// 	}
// 	path = r.DATA_ROOT_PATH + val
// 	if !fileExists(path) {
// 		if showNoImage {
// 			path = r.APP_ROOT + "/resource/common/images/no-image.png"
// 		} else {
// 			path = ""
// 		}
// 	}
// 	return path
// }

func ValidateImageExt(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	for _, allowedExt := range allowedImageExtensions {
		if ext == allowedExt {
			return true
		}
	}
	return false
}

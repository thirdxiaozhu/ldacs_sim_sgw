package util

/*
#cgo CFLAGS: -I/usr/local/include/ldacs -I/usr/local/include/km
#cgo LDFLAGS:  -lsqlite3 -lsdf -luuid -lkm_src -L/usr/local/lib/ldacs -lldacsutils
#include "key_manage.h"
#include "kmdb.h"
#include "km_field.h"
#include <utils/ld_log.h>
#include <utils/ld_buffer.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

/*
@author:wencheng
@date:20240723
@biref key manage API
*/

/**************************************************
*                   根密钥预置                     *
**************************************************/

/*
*
GenerateRootKey 生成根密钥并导出到指定文件中。
参数说明：
- asName: 应用服务器名称，用于生成密钥的标识。
- sgwName: 网关名称，用于生成密钥的标识。
- keyLen: 所需生成密钥的长度，以字节为单位。
- validityPeriod: 密钥的有效期，以天数表示。
- dbname: 密钥存储的数据库名称。
- tablename: 密钥存储的表名称。
- exportFilePath: 导出生成的根密钥的文件路径。
返回值：
生成根密钥的结果，返回一个整数表示操作的状态或错误码。
*/
func GenerateRootKey(asName, sgwName string, keyLen, validityPeriod int,
	dbname, tablename, exportFilePath string) error {
	// 将字符串转换为 C 字符串
	cAsName := C.CString(asName)
	cSgwName := C.CString(sgwName)
	cDbname := C.CString(dbname)
	cTablename := C.CString(tablename)
	cFilePath := C.CString(exportFilePath)

	defer C.free(unsafe.Pointer(cAsName))
	defer C.free(unsafe.Pointer(cSgwName))
	defer C.free(unsafe.Pointer(cDbname))
	defer C.free(unsafe.Pointer(cTablename))
	defer C.free(unsafe.Pointer(cFilePath))

	// 将 int 类型的 keyLen 和 validityPeriod 转换为 C.uint32_t
	cKeyLen := C.uint32_t(keyLen)
	cValidityPeriod := C.uint32_t(validityPeriod)

	// 调用 C 函数
	result := int(C.km_rkey_gen_export(cAsName, cSgwName, cKeyLen, cValidityPeriod, cDbname, cTablename, cFilePath))
	if result != C.LD_KM_OK {
		return fmt.Errorf("failed to write file to cryptocard, error code: %d", int(result))
	}

	// 返回封装后的结果类型
	return nil
}

/*
KmWriteFileToCryptocard 将文件存入密码卡文件区。
参数说明：
- filepath: 指定输入文件的路径。应为以 null 结尾的字符串，表示文件在文件系统中的完整路径。
- filename: 存入密码卡时的文件名。应为以 null 结尾的字符串，表示文件在密码卡中的存储名称。
返回值：
返回一个 error 类型的值，如果操作成功则返回 nil，否则返回一个描述错误的消息。
*/
func KmWriteFileToCryptocard(filepath, filename string) error {
	cFilePath := C.CString(filepath)
	cFileName := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilePath))
	defer C.free(unsafe.Pointer(cFileName))

	result := C.km_writefile_to_cryptocard(
		(*C.uint8_t)(unsafe.Pointer(cFilePath)),
		(*C.uint8_t)(unsafe.Pointer(cFileName)),
	)

	if result != C.LD_KM_OK {
		return fmt.Errorf("failed to write file to cryptocard, error code: %d", int(result))
	}

	return nil
}

// RootKeyImport 将根密钥导入并保存到数据库中。
//
// 参数说明：
// - dbname: 数据库名称，用于存储根密钥。应为以 null 结尾的字符串，表示数据库的名称。
// - tablename: 数据库中的表名，用于存储根密钥。应为以 null 结尾的字符串，表示表的名称。
// - rkeyFilenameInCcard: 在密码卡中根密钥文件名。应为以 null 结尾的字符串，表示密码卡中的文件名。
//
// 返回值：
// 返回一个 l_km_err 类型的错误码。如果操作成功，则返回 LD_KM_OK（0），否则返回其他错误码。
func RootKeyImport(dbname, tablename, rkeyFilenameInCcard string) error {
	cDbname := C.CString(dbname)
	cTablename := C.CString(tablename)
	cRkeyFilenameInCcard := C.CString(rkeyFilenameInCcard)
	defer C.free(unsafe.Pointer(cDbname))
	defer C.free(unsafe.Pointer(cTablename))
	defer C.free(unsafe.Pointer(cRkeyFilenameInCcard))

	result := C.km_rkey_import(
		cDbname,
		cTablename,
		cRkeyFilenameInCcard,
	)

	if result != C.LD_KM_OK {
		fmt.Printf("Error importing root key: %d\n", int(result))
	}

	return nil
}

/*
// Convert Go string to C string
func goStringToCString(s string) *C.char {
    return C.CString(s)
}

// Convert C string to Go string
func cStringToGoString(s *C.char) string {
    return C.GoString(s)
}
*/

/*
// QueryResult struct to wrap C struct
type QueryIDResult struct {
    queryId int
    // other fields
}

//go:generate CGO_LDFLAGS_ALLOW=.-lcgo_example

// NewQueryResult creates a new QueryResult from C struct
func NewQueryResult(cRes *C.QueryResult_for_queryid) *QueryResult {
    return &QueryResult{
        queryId: int(cRes.query_id),
        // initialize other fields
    }
}

// QueryID 封装了 C 语言的 query_id 函数，用于根据给定的数据库名、表名、所有者、密钥类型和状态查询密钥标识。
// 参数：
// - dbName: 数据库名
// - tableName: 表名
// - owner1: 第一个所有者
// - owner2: 第二个所有者
// - keyType: 密钥类型，表示为 int 类型，需转换为 C 的 KEY_TYPE 类型
// - state: 状态，表示为 int 类型，需转换为 C 的 STATE 类型
// 返回：
// - 字符串类型的密钥标识，如果未找到结果返回空字符串
// - 错误信息，若调用过程中出现错误则返回该错误

func QueryID(dbName, tableName, owner1, owner2 string, keyType KeyType, state State) (*QueryResult, error) {
    cDBName := C.CString(dbName)
    cTableName := C.CString(tableName)
    cOwner1 := C.CString(owner1)
    cOwner2 := C.CString(owner2)

    defer C.free(unsafe.Pointer(cDBName))
    defer C.free(unsafe.Pointer(cTableName))
    defer C.free(unsafe.Pointer(cOwner1))
    defer C.free(unsafe.Pointer(cOwner2))

    cResult := C.query_id(cDBName, cTableName, cOwner1, cOwner2, C.enum_KEY_TYPE(keyType), C.enum_STATE(state))

    if cResult == nil {
        return nil, fmt.Errorf("query_id returned nil")
    }

    result := NewQueryResult(cResult)
    return result, nil
}
*/
// EnableKey 调用 C 语言的 enable_key 函数
func EnableKey(dbName, tableName, id string) error {
	cDbName := C.CString(dbName)
	cTableName := C.CString(tableName)
	cId := C.CString(id)
	defer C.free(unsafe.Pointer(cDbName))
	defer C.free(unsafe.Pointer(cTableName))
	defer C.free(unsafe.Pointer(cId))

	result := (C.enable_key(cDbName, cTableName, cId))
	if result != C.LD_KM_OK {
		fmt.Printf("Error enable key: %d\n", int(result))
	}

	return nil
}

/**************************************************
*                   密钥派生                      *
**************************************************/

// DeriveKey 调用 C 语言的 km_derive_key 函数
func DeriveKey(dbName, tableName, id, gsName string, keyLen uint32, rand []byte) error {
	// 将 Go 字符串转换为 C 字符串
	cDbName := C.CString(dbName)
	cTableName := C.CString(tableName)
	cId := C.CString(id)
	cGsName := C.CString(gsName)

	// 将 Go 字节切片转换为 C 字节数组
	cRand := C.CBytes(rand)

	// 确保在函数退出时释放分配的内存
	defer C.free(unsafe.Pointer(cDbName))
	defer C.free(unsafe.Pointer(cTableName))
	defer C.free(unsafe.Pointer(cId))
	defer C.free(unsafe.Pointer(cGsName))
	defer C.free(cRand) // 直接使用 cRand，不需要 unsafe.Pointer 的转换

	// 调用 C 函数
	result := C.km_derive_key(
		(*C.uint8_t)(unsafe.Pointer(cDbName)),
		(*C.uint8_t)(unsafe.Pointer(cTableName)),
		(*C.uint8_t)(unsafe.Pointer(cId)),
		C.uint32_t(keyLen),
		(*C.uint8_t)(unsafe.Pointer(cGsName)),
		(*C.uint8_t)(unsafe.Pointer(cRand)),
		C.uint32_t(len(rand)),
	)
	fmt.Println()

	// 检查 C 函数的返回结果
	if result != C.LD_KM_OK { // 确保与实际错误码匹配
		return fmt.Errorf("error deriving key: %d", int(result))
	}

	return nil
}

/**************************************************
*           密钥查询，用于分发主密钥给GS             *
**************************************************/
// QueryResult holds the result of the key query.
type QueryResult struct {
	KeyLen int
	Key    []byte
}

// QueryKeyValue queries a key value from the external C function.
func QueryKeyValue(dbName, tableName, id string) (*QueryResult, error) {
	cDbName := C.CString(dbName)
	defer C.free(unsafe.Pointer(cDbName))
	cTableName := C.CString(tableName)
	defer C.free(unsafe.Pointer(cTableName))
	cID := C.CString(id)
	defer C.free(unsafe.Pointer(cID))

	// Convert *C.char to *C.uchar
	result := C.query_keyvalue((*C.uchar)(unsafe.Pointer(cDbName)),
		(*C.uchar)(unsafe.Pointer(cTableName)),
		(*C.uchar)(unsafe.Pointer(cID)))
	if result == nil {
		return nil, fmt.Errorf("no result found")
	}
	defer C.free(unsafe.Pointer(result.key))
	defer C.free(unsafe.Pointer(result))

	keyLen := int(result.key_len)
	keySlice := C.GoBytes(unsafe.Pointer(result.key), C.int(keyLen))

	return &QueryResult{
		KeyLen: keyLen,
		Key:    keySlice,
	}, nil
}

/**************************************************
*                   密钥存储                      *
**************************************************/

// 主密钥安装
func InstallKey(dbname, tablename string, key []byte, sacAs, sacGs, nonce []byte) error {
	// 将 Go 字符串转换为 C 字符串
	cDbname := C.CString(dbname)
	cTablename := C.CString(tablename)
	cKey := C.CBytes(key)
	cSacAs := C.CBytes(sacAs)
	cSacGs := C.CBytes(sacGs)
	cNonce := C.CBytes(nonce)
	defer C.free(unsafe.Pointer(cDbname))
	defer C.free(unsafe.Pointer(cTablename))
	defer C.free(unsafe.Pointer(cKey))
	defer C.free(unsafe.Pointer(cSacAs))
	defer C.free(unsafe.Pointer(cSacGs))
	defer C.free(unsafe.Pointer(cNonce))

	// 获取字节切片的长度
	keyLen := C.uint32_t(len(key))
	nonceLen := C.uint32_t(len(nonce))

	// 调用 C 语言函数
	errCode := C.km_install_key(
		(*C.uint8_t)(unsafe.Pointer(cDbname)),
		(*C.uint8_t)(unsafe.Pointer(cTablename)),
		keyLen,
		(*C.uint8_t)(cKey),
		(*C.uint8_t)(cSacAs),
		(*C.uint8_t)(cSacGs),
		nonceLen,
		(*C.uint8_t)(cNonce),
	)

	// 根据返回的错误码返回 Go 语言的错误
	if errCode != C.LD_KM_OK {
		return fmt.Errorf("failed with error code %d", errCode)
	}
	return nil
}

// 查询密钥id

// 查询密钥值

/**************************************************
*                   密钥更新                      *
**************************************************/

/**
 * @brief 外部接口：网关端更新主密钥 KAS-GS 使用id进行定位
 * @param[in] dbname   密钥库名
 * @param[in] tablename 密钥表名
 * @param[in] keyID 密钥id标识
 * @param[in] sgwName 网关名字
 * @param[in] gsName 目的GS名字
 * @param[in] nonce 随机数
 */
func SGWUpdateMasterKey(dbname, tablename, keyID, sgwName, gsName string, nonce []byte) int {
	cDbName := C.CString(dbname)
	cTableName := C.CString(tablename)
	cKeyID := C.CString(keyID)
	cSgwName := C.CString(sgwName)
	cGsName := C.CString(gsName)
	cNonce := (*C.uint8_t)(unsafe.Pointer(&nonce[0]))
	cLenNonce := C.uint16_t(len(nonce))

	ret := C.sgw_update_master_key(
		(*C.uint8_t)(unsafe.Pointer(cDbName)),
		(*C.uint8_t)(unsafe.Pointer(cTableName)),
		(*C.uint8_t)(unsafe.Pointer(cKeyID)),
		(*C.uint8_t)(unsafe.Pointer(cSgwName)),
		(*C.uint8_t)(unsafe.Pointer(cGsName)),
		cLenNonce,
		cNonce,
	)

	C.free(unsafe.Pointer(cDbName))
	C.free(unsafe.Pointer(cTableName))
	C.free(unsafe.Pointer(cKeyID))
	C.free(unsafe.Pointer(cSgwName))
	C.free(unsafe.Pointer(cGsName))

	return int(ret)
}

/*
// UpdateMasterKey 更新主密钥
func UpdateMasterKey(
    dbName, tableName, sacSgw, sacGsS, sacGsT, sacAs []byte,
    lenNonce uint16, nonce []byte,
) error {
    // 转换Go语言的byte切片为C语言的指针
    cDbName := C.CBytes(dbName)
    defer C.free(cDbName)
    cTableName := C.CBytes(tableName)
    defer C.free(cTableName)
    cSacSgw := C.CBytes(sacSgw)
    defer C.free(cSacSgw)
    cSacGsS := C.CBytes(sacGsS)
    defer C.free(cSacGsS)
    cSacGsT := C.CBytes(sacGsT)
    defer C.free(cSacGsT)
    cSacAs := C.CBytes(sacAs)
    defer C.free(cSacAs)
    cNonce := C.CBytes(nonce)
    defer C.free(cNonce)

    // 调用C语言函数
    result := C.km_update_masterkey(
        (*C.uint8_t)(cDbName),
        (*C.uint8_t)(cTableName),
        (*C.uint8_t)(cSacSgw),
        (*C.uint8_t)(cSacGsS),
        (*C.uint8_t)(cSacGsT),
        (*C.uint8_t)(cSacAs),
        C.uint16_t(lenNonce),
        (*C.uint8_t)(cNonce),
    )

    // 检查返回值
    if result != C.LD_KM_OK {
		return fmt.Errorf("error Update key: %d", int(result))
    }

    return nil
}
*/

// 密钥撤销
func RevokeKey(dbname, tablename, id string) error {
	cDbname := C.CString(dbname)
	cTablename := C.CString(tablename)
	cId := C.CString(id)
	defer C.free(unsafe.Pointer(cDbname))
	defer C.free(unsafe.Pointer(cTablename))
	defer C.free(unsafe.Pointer(cId))

	// 调用 C 语言函数
	errCode := C.km_revoke_key(
		(*C.uint8_t)(unsafe.Pointer(cDbname)),
		(*C.uint8_t)(unsafe.Pointer(cTablename)),
		(*C.uint8_t)(unsafe.Pointer(cId)),
	)

	// 根据返回的错误码返回 Go 语言的错误
	if errCode != 0 {
		return fmt.Errorf("failed with error code %d", errCode)
	}
	return nil
}

func GetKeyHandle(dbname, tablename, id string) (unsafe.Pointer, error) {
	//cPtr := (*unsafe.Pointer)(unsafe.Pointer(handler))
	handler := unsafe.Pointer(nil)

	cDbname := C.CString(dbname)
	cTablename := C.CString(tablename)
	cId := C.CString(id)
	defer C.free(unsafe.Pointer(cDbname))
	defer C.free(unsafe.Pointer(cTablename))
	defer C.free(unsafe.Pointer(cId))

	errCode := C.get_handle_from_db(
		(*C.uint8_t)(unsafe.Pointer(cDbname)),
		(*C.uint8_t)(unsafe.Pointer(cTablename)),
		(*C.uint8_t)(unsafe.Pointer(cId)),
		&handler)

	if errCode != 0 {
		return nil, fmt.Errorf("GetKeyHandle failed with error code %d", errCode)
	}
	return handler, nil
}

func CalcHMAC(handler unsafe.Pointer, data []byte, limit uint32) ([]byte, error) {
	hmacResult := make([]byte, limit)

	//cData := C.CBytes(data)
	//defer C.free(cData)
	//cResult := C.CBytes(hmacResult)
	//defer C.free(cResult)
	//cResultLen := C.uint32_t(hmacLen)
	cHmacLen := C.uint32_t(0)
	//logger.Warn("??????????")

	errCode := C.km_hmac_with_keyhandle(
		handler,
		(*C.uint8_t)(unsafe.Pointer(&data[0])),
		C.uint32_t(len(data)),
		(*C.uint8_t)(unsafe.Pointer(&hmacResult[0])),
		&cHmacLen)
	if errCode != 0 {
		return nil, fmt.Errorf("GetKeyHandle failed with error code %d", errCode)
	}

	return hmacResult, nil
}

func Encrypt(handler unsafe.Pointer) {
	iv := make([]byte, 16)
	data := make([]byte, 16)
	res := make([]byte, 256)
	cResLen := C.uint32_t(0)
	errCode := C.km_encrypt(
		handler,
		C.uint32_t(0x00000402),
		(*C.uint8_t)(unsafe.Pointer(&iv[0])),
		(*C.uint8_t)(unsafe.Pointer(&data[0])),
		C.uint32_t(16),
		(*C.uint8_t)(unsafe.Pointer(&res[0])),
		&cResLen)

	for i := range res {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Printf("%02x", res[i])
	}
	fmt.Println()
	if errCode != 0 {
		fmt.Printf("GetKeyHandle failed with error code %d", errCode)
	}
}

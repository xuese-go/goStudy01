package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

// 指定加密密钥
var jwtSecret = []byte("xueSeToken")

//签名
var Autograph = "xue-se"

//Claim是一些实体（通常指的用户）的状态和额外的元数据
type Claims struct {
	Uuid string `json:"uuid"`
	jwt.StandardClaims
}

// 根据用户的uuid产生token
func GenerateToken(uuid string) (string, error) {
	//设置token有效时间
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		Uuid: uuid,
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: expireTime.Unix(),
			// 指定token发行人
			Issuer: Autograph,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//该方法内部生成签名字符串，再用于获取完整、已签名的token
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// 根据传入的token值获取到Claims对象信息
func ParseToken(token string) (*Claims, error) {

	//用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		// 从tokenClaims中获取到Claims对象，并使用断言，将该对象转换为我们自己定义的Claims
		// 要传入指针，项目中结构体都是用指针传递，节省空间。
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err

}

//令牌合法性判断
func IsToken(token string) bool {
	//解析
	if t, err := ParseToken(token); err != nil {
		log.Println("令牌解析错误")
		return false
	} else {
		//签名
		autograph := t.StandardClaims.Issuer
		if autograph == Autograph {
			//判断令牌过期时间
			cl := t.ExpiresAt
			//当前时间
			t := time.Now().Unix()
			if cl-t < 0 {
				return true
			} else {
				return false
			}
		} else {
			return false
		}
	}
}

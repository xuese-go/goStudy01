package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/xuese-go/goStudy01/util/md5"
	"log"
	"time"
)

// 指定加密密钥
var jwtSecret = []byte("www.xueSeToken.com")

type Claims struct {
	jwt.StandardClaims
}

/**
生成令牌
*/
func GenerateToken(str string, ip string) (string, error) {
	claims := Claims{
		StandardClaims: jwt.StandardClaims{
			//签发时间
			IssuedAt: time.Now().Unix(),
			//过期时间
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
			//	发放者
			Issuer: str,
			//	主题
			Subject: md5.Enc(ip, "逗你玩!!!"),
		},
	}
	//声明新的token
	token := jwt.NewWithClaims(jwt.SigningMethodHS384, claims)
	//	获取完整的签名令牌
	t, err := token.SignedString(jwtSecret)
	return t, err
}

/**
验证令牌
*/
func ParseToken(tokenStr string) (bool, *Claims) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if token != nil && token.Valid {
		//	令牌有效
		return true, token.Claims.(*Claims)
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		//令牌格式不正确
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			log.Println("That's not even a token")
			return false, nil
		} else if ve.Errors&jwt.ValidationErrorUnverifiable != 0 {
			log.Println("由于签名问题，无法验证令牌")
			return false, nil
		} else if ve.Errors&jwt.ValidationErrorSignatureInvalid != 0 {
			log.Println("签名验证失败")
			return false, nil
		} else if ve.Errors&jwt.ValidationErrorAudience != 0 {
			log.Println("AUD验证失败")
			return false, nil
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// 令牌已过期或尚未激活
			log.Println("EXP验证失败")
			return false, nil
		} else if ve.Errors&jwt.ValidationErrorIssuedAt != 0 {
			log.Println("IAT验证失败")
			return false, nil
		} else if ve.Errors&jwt.ValidationErrorIssuer != 0 {
			log.Println("ISS验证失败")
			return false, nil
		} else if ve.Errors&jwt.ValidationErrorId != 0 {
			log.Println("JTI验证失败")
			return false, nil
		} else if ve.Errors&jwt.ValidationErrorClaimsInvalid != 0 {
			log.Println("通用声明验证错误")
			return false, nil
		} else {
			//无法处理此令牌
			log.Println("Couldn't handle this token:", err)
			return false, nil
		}
	} else {
		log.Println("Couldn't handle this token:", err)
		return false, nil
	}
}

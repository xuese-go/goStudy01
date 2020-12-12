/**
注意此文件log所采用的输出方法
*/
package boltdb

import (
	"github.com/boltdb/bolt"
	"log"
	"time"
)

func init() {
	//打开我的数据库当前目录中的数据文件。
	//如果它不存在，它将被创建。
	//同时只能打开一次，所以防止无限等待则设置超时时间
	var err error
	db, err := bolt.Open(DbName, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		//注意此处
		log.Fatal(err)
	}
	defer db.Close()

	//	读写事务
	err = db.Update(func(tx *bolt.Tx) error {

		//判断要创建的桶是否存在
		b := tx.Bucket([]byte(BucketName))
		if b == nil {

			//创建叫"token_table"的桶
			_, err := tx.CreateBucket([]byte(BucketName))
			if err != nil {
				log.Fatal("桶token_table创建失败", err.Error())
			}
		}

		//一定要返回nil
		return nil
	})

	//更新数据库失败
	if err != nil {
		log.Fatal(err)
	}
}

package boltdb

import (
	"github.com/boltdb/bolt"
	"log"
	"time"
)

/**
保存token
*/
func Save(token string) error {
	var err error
	db, err := bolt.Open(DbName, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		//注意此处
		log.Fatal(err)
	}
	//插入数据
	err = db.Update(func(tx *bolt.Tx) error {

		//取出桶
		b := tx.Bucket([]byte(BucketName))

		//往表里面存储数据
		if b != nil {
			//插入的键值对数据类型必须是字节数组
			err := b.Put([]byte(token), []byte(token))
			if err != nil {
				log.Println(err.Error())
				return err
			}
		}

		//一定要返回nil
		return nil
	})

	//更新数据库失败
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	return nil
}

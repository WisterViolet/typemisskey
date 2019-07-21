// Package util : 仮に他のSNSで使うときになっても使用できる関数郡
package util

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
	"path"
)

// LoadJSON : .jsonを読み込んで、変数に保存する関数
func LoadJSON(filename string, v interface{}) error {
	/*
		filename: .jsonファイルのパス
		v: 保存先の変数のアドレス
	*/
	text, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(text, v)
	if err != nil {
		return err
	}
	return nil
}

// WrapJSONString : 任意の型をjsonの[]byteに変換する関数
func WrapJSONString(v interface{}) ([]byte, error) {
	/*
		v: 任意の型
	*/
	body, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// JoinURLPath : baseURLにエンドポイントを加える関数
func JoinURLPath(base, endpoint string) (string, error) {
	/*
		base: ベースになるURL
		endpoint: エンドポイントへのパス
	*/
	baseURL, err := url.Parse(base)
	if err != nil {
		return "", err
	}
	tmp := path.Join(baseURL.Path, endpoint)
	baseURL.Path = tmp
	return baseURL.String(), nil
}

// Copyright 2017~2022 The Bottos Authors
// This file is part of the Bottos Chain library.
// Created by Rocket Core Team of Bottos.

//This program is free software: you can distribute it and/or modify
//it under the terms of the GNU General Public License as published by
//the Free Software Foundation, either version 3 of the License, or
//(at your option) any later version.

//This program is distributed in the hope that it will be useful,
//but WITHOUT ANY WARRANTY; without even the implied warranty of
//MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//GNU General Public License for more details.

//You should have received a copy of the GNU General Public License
// along with bottos.  If not, see <http://www.gnu.org/licenses/>.

/*
 * file description:  provide a interface error definition for all modules
 * @Author: zl
 * @Date: 2018-4-25
 * @Last Modified by:
 * @Last Modified time:
*/
package error

import (
	"io/ioutil"
	"encoding/json"
	log "github.com/cihub/seelog"
)

type ErrorCode struct {
	Code    int64 `json:"code"`
	Msg     struct {
		Cn string `json:"cn"`
		En string `json:"en"`
	} `json:"msg"`
	Details string  `json:"details"`
}

type Ret struct {
	Code    int64 		`json:"code"`
	Data 	interface{} `json:"data"`
	Msg     string		`json:"msg"`
}

func GetErrorInfo(code int64) ErrorCode {
	d := GetAllErrorInfos()
	for _, v := range d {
		if code == v.Code {
			return v
		}
	}
	return ErrorCode{}
}

func Return(b interface{}) string {
	buf, err:= json.Marshal(b)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	var ret Ret
	json.Unmarshal(buf, &ret)
	if(ret.Code == 0 || ret.Code == 1){
		ret.Code = 1
		ret.Msg = "ok"

		body, err:= json.Marshal(ret)
		if err != nil {
			log.Error(err)
			panic(err)
		}
		return string(body)
	}


	d := GetAllErrorInfos()
	for _, v := range d {
		if ret.Code == v.Code {
			v.Details = ret.Msg
			json, err := json.Marshal(v)
			if err != nil {
				log.Error(err)
				panic(err)
			}
			return string(json)
		}
	}

	json, err := json.Marshal(ErrorCode{})
	if err != nil {
		log.Error(err)
		panic(err)
	}
	return string(json)
}

func ReturnError(code int64, e ...error) string {
	d := GetAllErrorInfos()
	for _, v := range d {
		if code == v.Code {
			if len(e) > 0 {
				v.Details = e[0].Error()
			}
			json, err := json.Marshal(v)
			if err != nil {
				log.Error(err)
				panic(err)
			}
			return string(json)
		}
	}
	json, err := json.Marshal(ErrorCode{})
	if err != nil {
		log.Error(err)
		panic(err)
	}
	return string(json)
}

func GetAllErrorInfos() []ErrorCode {
	fr, err := ioutil.ReadFile("./error/err-code.json")
	if err != nil {
		log.Error(err)
		panic(err)
	}

	var d []ErrorCode
	err = json.Unmarshal(fr, &d)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	return d
}



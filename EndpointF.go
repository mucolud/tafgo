// **********************************************************************
// This file was generated by a TAF parser!
// TAF version 3.2.2.2 by WSRD Tencent.
// Generated from `EndpointF.jce'
// **********************************************************************

package tafgo

import (
	"bytes"
)

type EndpointF struct {
	Host          string `tag:"0"  required:"true"  json:"host"`
	Port          int32  `tag:"1"  required:"true"  json:"port"`
	Timeout       int32  `tag:"2"  required:"true"  json:"timeout"`
	Istcp         int32  `tag:"3"  required:"true"  json:"istcp"`
	Grid          int32  `tag:"4"  required:"true"  json:"grid"`
	Groupworkid   int32  `tag:"5"  required:"false"  json:"groupworkid"`
	Grouprealid   int32  `tag:"6"  required:"false"  json:"grouprealid"`
	SetId         string `tag:"7"  required:"false"  json:"setId"`
	Qos           int32  `tag:"8"  required:"false"  json:"qos"`
	BakFlag       int32  `tag:"9"  required:"false"  json:"bakFlag"`
	GridFlag      int32  `tag:"10"  required:"false"  json:"gridFlag"`
	Weight        int32  `tag:"11"  required:"false"  json:"weight"`
	WeightType    int32  `tag:"12"  required:"false"  json:"weightType"`
	Cpuload       int32  `tag:"13"  required:"false"  json:"cpuload"`
	Sampletime    int64  `tag:"14"  required:"false"  json:"sampletime"`
	ContainerName string `tag:"15"  required:"false"  json:"containerName"`
}

func (p *EndpointF) ClassName() string {
	return "tafgo.EndpointF"
}
func (p *EndpointF) MD5() string {
	return "a2ddb3f27d41b6b47cb4186b06d8ff8b"
}
func (p *EndpointF) ResetDefautlt() {
	var empty EndpointF
	*p = empty
}
func (p *EndpointF) Encode(buf *bytes.Buffer) error {
	var err error
	err = EncodeTagStringValue(buf, p.Host, 0)
	if nil != err {
		return err
	}
	err = EncodeTagInt32Value(buf, p.Port, 1)
	if nil != err {
		return err
	}
	err = EncodeTagInt32Value(buf, p.Timeout, 2)
	if nil != err {
		return err
	}
	err = EncodeTagInt32Value(buf, p.Istcp, 3)
	if nil != err {
		return err
	}
	err = EncodeTagInt32Value(buf, p.Grid, 4)
	if nil != err {
		return err
	}
	err = EncodeTagInt32Value(buf, p.Groupworkid, 5)
	if nil != err {
		return err
	}
	err = EncodeTagInt32Value(buf, p.Grouprealid, 6)
	if nil != err {
		return err
	}
	err = EncodeTagStringValue(buf, p.SetId, 7)
	if nil != err {
		return err
	}
	err = EncodeTagInt32Value(buf, p.Qos, 8)
	if nil != err {
		return err
	}
	err = EncodeTagInt32Value(buf, p.BakFlag, 9)
	if nil != err {
		return err
	}
	err = EncodeTagInt32Value(buf, p.GridFlag, 10)
	if nil != err {
		return err
	}
	err = EncodeTagInt32Value(buf, p.Weight, 11)
	if nil != err {
		return err
	}
	err = EncodeTagInt32Value(buf, p.WeightType, 12)
	if nil != err {
		return err
	}
	err = EncodeTagInt32Value(buf, p.Cpuload, 13)
	if nil != err {
		return err
	}
	err = EncodeTagInt64Value(buf, p.Sampletime, 14)
	if nil != err {
		return err
	}
	err = EncodeTagStringValue(buf, p.ContainerName, 15)
	if nil != err {
		return err
	}
	return nil
}
func (p *EndpointF) Decode(buf *bytes.Buffer) error {
	var err error
	err = DecodeTagStringValue(buf, &p.Host, 0, true)
	if nil != err {
		return err
	}
	err = DecodeTagInt32Value(buf, &p.Port, 1, true)
	if nil != err {
		return err
	}
	err = DecodeTagInt32Value(buf, &p.Timeout, 2, true)
	if nil != err {
		return err
	}
	err = DecodeTagInt32Value(buf, &p.Istcp, 3, true)
	if nil != err {
		return err
	}
	err = DecodeTagInt32Value(buf, &p.Grid, 4, true)
	if nil != err {
		return err
	}
	err = DecodeTagInt32Value(buf, &p.Groupworkid, 5, false)
	if nil != err {
		return err
	}
	err = DecodeTagInt32Value(buf, &p.Grouprealid, 6, false)
	if nil != err {
		return err
	}
	err = DecodeTagStringValue(buf, &p.SetId, 7, false)
	if nil != err {
		return err
	}
	err = DecodeTagInt32Value(buf, &p.Qos, 8, false)
	if nil != err {
		return err
	}
	err = DecodeTagInt32Value(buf, &p.BakFlag, 9, false)
	if nil != err {
		return err
	}
	err = DecodeTagInt32Value(buf, &p.GridFlag, 10, false)
	if nil != err {
		return err
	}
	err = DecodeTagInt32Value(buf, &p.Weight, 11, false)
	if nil != err {
		return err
	}
	err = DecodeTagInt32Value(buf, &p.WeightType, 12, false)
	if nil != err {
		return err
	}
	err = DecodeTagInt32Value(buf, &p.Cpuload, 13, false)
	if nil != err {
		return err
	}
	err = DecodeTagInt64Value(buf, &p.Sampletime, 14, false)
	if nil != err {
		return err
	}
	err = DecodeTagStringValue(buf, &p.ContainerName, 15, false)
	if nil != err {
		return err
	}
	return err
}

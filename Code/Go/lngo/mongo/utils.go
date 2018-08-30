package mongo

import (
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
)

// TODO: 不是很懂UnmarshalBson 合Marshal
type M map[string]interface{}

func (m M) UnmarshalBSON(payload []byte) error {
	var out bson.Document
	err := bson.Unmarshal(payload, &out)
	if err != nil {
		return err
	}
	for k, v := range convertDocument(&out) {
		m[k] = v
	}
	return nil
}

func (m M) ToDocument() *bson.Document {
	d := bson.NewDocument()
	for k, v := range m {
		switch vv := v.(type) {
		case M:
			d.Append(bson.EC.SubDocument(k, vv.ToDocument()))
		case []M:
			a := bson.NewArray()
			for _, av := range vv {
				a.Append(bson.VC.Document(av.ToDocument()))
			}
			d.Append(bson.EC.Array(k, a))
		case time.Time:
			d.Append(bson.EC.Time(k, vv))
		case *time.Time:
			if vv != nil {
				d.Append(bson.EC.Time(k, *vv))
			}
		default:
			d.Append(bson.EC.Interface(k, v))
		}
	}
	return d
}

func (m M) MarshalBSON() (ret []byte, err error) {
	return m.ToDocument().MarshalBSON()
}

func convertArray(arr *bson.Array) []interface{} {
	ret := []interface{}{}
	i, _ := arr.Iterator()
	for i.Next() {
		ret = append(ret, convertValue(i.Value()))
	}
	return ret
}

func convertDocument(doc *bson.Document) M {
	ret := M{}
	i := doc.Iterator()
	for i.Next() {
		elem := i.Element()
		ret[elem.Key()] = convertValue(elem.Value())
	}
	return ret
}

func convertValue(v *bson.Value) interface{} {
	switch v.Type() {
	case bson.TypeArray:
		return convertArray(v.MutableArray())
	case bson.TypeEmbeddedDocument:
		return convertDocument(v.MutableDocument())
	default:
		return v.Interface()
	}
}

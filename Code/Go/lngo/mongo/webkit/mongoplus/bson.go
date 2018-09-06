package mongoplus

import (
	"log"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
)

// Convert using bson to transform a struct to another struct
func Convert(from, to interface{}) error {
	buf, err := bson.Marshal(from)
	if err != nil {
		return err
	}
	return bson.Unmarshal(buf, to)
}

// PrintJSON print struct with bson
func PrintJSON(input interface{}) {
	v, err := bson.Marshal(input)
	if err != nil {
		log.Println("error: ", err)
		return
	}
	ret, err := bson.ToExtJSON(false, v)
	if err != nil {
		log.Println("error: ", err)
		return
	}
	log.Println(ret)
}

// MergeM merge multiple M, duplicate key are overwrited
func MergeM(args ...M) (ret M) {
	ret = make(M)
	for _, m := range args {
		for k, v := range m {
			ret[k] = v
		}
	}
	return
}

// OrQuery helps build mongo query with $or
func OrQuery(queries ...M) M {
	return M{
		"$or": queries,
	}
}

func andQuery(a, b M) M {
	if a == nil && b == nil {
		return M{}
	}
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	aOR, aOr := a["$or"].([]M)
	bOR, bOr := b["$or"].([]M)
	if aOr && bOr {
		delete(a, "$or")
		delete(b, "$or")
		return MergeM(a, b, M{"$and": []M{M{"$or": aOR}, M{"$or": bOR}}})
	}
	return MergeM(a, b)
}

// AndQuery helps build mongo query with $and
func AndQuery(queries ...M) M {
	if len(queries) == 1 {
		return queries[0]
	}
	ret := andQuery(queries[0], queries[1])
	for _, q := range queries[2:] {
		ret = andQuery(ret, q)
	}
	return ret
}

// RealTimeRangeQuery return query select item overlap with specific period of time.
func RealTimeRangeQuery(fromKey, toKey string, from, to time.Time) M {
	return OrQuery(
		M{fromKey: TimeRangeQuery(&from, &to)},
		M{toKey: TimeRangeQuery(&from, &to)},
		M{fromKey: M{"$lte": from}, toKey: M{"$gte": to}},
	)
}

// TimeRangeQuery return query select item between time
func TimeRangeQuery(from *time.Time, to *time.Time) M {
	ret := M{}
	if to != nil {
		ret["$lte"] = *to
	}
	if from != nil {
		ret["$gte"] = *from
	}
	if from == nil && to == nil {
		// FIXME: find a better match all query
		ret["$exists"] = true
	}
	return ret
}

// M just likes mgo/bson.M
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

// ToDocument ransform M to bson.Document
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
		case bson.Document:
			d.Append(bson.EC.SubDocument(k, &vv))
		case *bson.Document:
			d.Append(bson.EC.SubDocument(k, vv))
		default:
			d.Append(bson.EC.Interface(k, v))
		}
	}
	return d
}

func (m M) MarshalBSON() (ret []byte, err error) {
	return m.ToDocument().MarshalBSON()
}

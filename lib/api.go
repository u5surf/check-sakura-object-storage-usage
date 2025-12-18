package checkusage

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Usage struct {
	quota  float64
	amount float64
}

type AmountGibPerBucket struct {
	IsApplied bool    `json:"is_applied"`
	Quota     float64 `json:"quota"`
	Val       float64 `json:"val"`
}

type NumObjectsPerBucket struct {
	IsApplied bool    `json:"is_applied"`
	Quota     float64 `json:"quota"`
	Val       float64 `json:"val"`
}

type Data struct {
	AmountGibPerBkt AmountGibPerBucket  `json:"amount_gib_per_bucket"`
	NumObjPerBkt    NumObjectsPerBucket `json:"num_objects_per_bucket"`
}

type Response struct {
	Data Data `json:"data"`
}

type APIClient interface {
	GetUsage() (*Usage, error)
}

type ObjectStorageAPI struct {
	url    string
	key    string
	secret string
}

func NewObjectStorageAPI(site, bucket, key, secret string) *ObjectStorageAPI {
	url := fmt.Sprintf("https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/%s/v2/buckets/%s/penalty", site, bucket)
	return &ObjectStorageAPI{
		url:    url,
		key:    key,
		secret: secret,
	}
}

func (a *ObjectStorageAPI) GetUsage() (*Usage, error) {
	req, err := http.NewRequest("GET", a.url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(a.key, a.secret)
	client := &http.Client{}
	var res Response
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(body, &res); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("bad response status code %d", resp.StatusCode)
	}
	return &Usage{
		quota:  res.Data.AmountGibPerBkt.Quota,
		amount: res.Data.AmountGibPerBkt.Val,
	}, nil
}

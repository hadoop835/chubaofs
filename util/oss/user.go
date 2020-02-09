package oss

import "sync"

type UserAK struct {
	UserID    string `json:"user_id"`
	AccessKey string `json:"access_key"`
}

type AKPolicy struct {
	AccessKey string      `json:"access_key"`
	SecretKey string      `json:"secret_key"`
	Policy    *UserPolicy `json:"policy"`
	UserID    string      `json:"user_id"`
}

type UserPolicy struct {
	OwnVol     []string
	NoneOwnVol map[string][]string
	Mutex      sync.RWMutex
}

type VolAK struct {
	Vol       string   `json:"vol"`
	AKAndAPIs []string `json:"user_aks"`
}

func (policy *UserPolicy) Add(addPolicy *UserPolicy) {
	policy.Mutex.Lock()
	defer policy.Mutex.Unlock()
	policy.OwnVol = append(policy.OwnVol, addPolicy.OwnVol...)
	for k, v := range addPolicy.NoneOwnVol {
		if apis, ok := policy.NoneOwnVol[k]; ok {
			policy.NoneOwnVol[k] = append(apis, addPolicy.NoneOwnVol[k]...)
		} else {
			policy.NoneOwnVol[k] = v
		}
	}
}

func (policy *UserPolicy) Delete(addPolicy *UserPolicy) {
	policy.Mutex.Lock()
	defer policy.Mutex.Unlock()
	policy.OwnVol = append(policy.OwnVol, addPolicy.OwnVol...)
	for k, v := range addPolicy.NoneOwnVol {
		if apis, ok := policy.NoneOwnVol[k]; ok {
			policy.NoneOwnVol[k] = append(apis, addPolicy.NoneOwnVol[k]...)
		} else {
			policy.NoneOwnVol[k] = v
		}
	}
}
package models


type Quota struct {

    /* 配额项的名称 (Optional) */
    Name string `json:"name"`

    /* 配额 (Optional) */
    Max int `json:"max"`

    /* 已使用的数目 (Optional) */
    Used int `json:"used"`
}

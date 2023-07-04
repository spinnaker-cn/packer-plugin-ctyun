package models


type Sort struct {

    /* 排序条件的名称 (Optional) */
    Name *string `json:"name"`

    /* 排序条件的方向 (Optional) */
    Direction *string `json:"direction"`
}

package models


type TagFilter struct {

    /* Tag键  */
    Key string `json:"key"`

    /* Tag值  */
    Values []string `json:"values"`
}
